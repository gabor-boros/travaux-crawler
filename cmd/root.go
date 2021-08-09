package cmd

import (
	"errors"
	"fmt"
	"gabor-boros/travaux-crawler/internal/pkg/crawler"
	"gabor-boros/travaux-crawler/internal/pkg/ocim"
	"gabor-boros/travaux-crawler/internal/pkg/travaux"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgFile is globally set by the argument parser
var cfgFile string

var version string
var commit string

// rootCmd is the main command used by the tool
var rootCmd = &cobra.Command{
	Use:   "travaux-crawler",
	Short: "Crawl the travaux.ovh.net site for incidents affecting Ocim instances.",
	Long: `
Travaux crawler is a CLI tool that crawls Travaux, OVH's incident report site
to find matching cloud instance IDs in the incident description.

The crawler is built for OpenCraft Instance Manager (Ocim) operators to make it
easier finding out if an ongoing outage affects an Open edX installation.`,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, err := cmd.Flags().GetBool("version")
		cobra.CheckErr(err)

		if showVersion {
			fmt.Printf("travaux-crawler version %s, commit %s\n", version, commit[:8])
			os.Exit(0)
		}

		rawProject, err := cmd.Flags().GetString("project")
		cobra.CheckErr(err)

		project := travaux.Project(rawProject)

		rawStatus, err := cmd.Flags().GetString("status")
		cobra.CheckErr(err)

		status := travaux.TaskStatus(rawStatus)

		crawlAllPages, err := cmd.Flags().GetBool("all-pages")
		cobra.CheckErr(err)

		verbose, err := cmd.Flags().GetBool("verbose")
		cobra.CheckErr(err)

		travauxURL, err := travaux.URL(project, status)
		cobra.CheckErr(err)

		appServerIDs, err := cmd.Flags().GetIntSlice("app-server")
		cobra.CheckErr(err)

		if len(appServerIDs) == 0 {
			cobra.CheckErr(errors.New("no app servers were provided"))
		}

		ocimURL, err := url.Parse(viper.GetString("ocim_url"))
		cobra.CheckErr(err)

		ocimURL.User = url.UserPassword(
			viper.GetString("ocim_username"),
			viper.GetString("ocim_password"),
		)

		if crawlAllPages && (status == travaux.TaskStatusAll || status == travaux.TaskStatusClosed) {
			cobra.CheckErr("cannot crawl all pages for all statuses, please set a status instead")
		}

		var appServers []ocim.AppServer
		for _, appServerID := range appServerIDs {
			fmt.Println("Getting app server info for", appServerID)
			appServer, err := ocim.GetAppServer(http.DefaultClient, ocimURL, appServerID)
			cobra.CheckErr(err)

			appServers = append(appServers, *appServer)
		}

		results, err := crawler.Crawl(travauxURL, status, appServers, crawlAllPages, verbose)
		cobra.CheckErr(err)

		if len(results) == 0 {
			fmt.Println("No app servers are affected")
		} else {
			for appServer, incidents := range results {
				fmt.Printf(
					"App server #%d (%s) (%s) is potentically affected by:\n",
					appServer.ID,
					appServer.Server.Name,
					appServer.Server.OpenstackId.String(),
				)

				for _, incident := range incidents {
					fmt.Printf("\t- %s\n", incident)
				}
			}
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.travauxCrawler-crawler.yaml)")

	rootCmd.Flags().StringP("project", "p", string(travaux.ProjectAll), "set travaux project name")
	rootCmd.Flags().StringP("status", "s", string(travaux.TaskStatusAll), "set travaux project name")

	rootCmd.Flags().IntSliceP("app-server", "a", []int{}, "potentially affected app server ID (required)")

	rootCmd.Flags().BoolP("all-pages", "", false, "crawl all pages using the paginator")
	rootCmd.Flags().BoolP("verbose", "", false, "print the visited page URLs")
	rootCmd.Flags().BoolP("version", "", false, "show command version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".travaux-crawler")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if _, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed()); err != nil {
			cobra.CheckErr(err)
		}
	}
}

// Execute called by main program and executes the root cmd
func Execute(cmdVersion string, cmdCommit string) {
	version = cmdVersion
	commit = cmdCommit

	cobra.CheckErr(rootCmd.Execute())
}
