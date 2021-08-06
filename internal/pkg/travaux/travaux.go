package travaux

import (
	"net/url"
	"strconv"
	"strings"
)

type Project string
type TaskStatus string

const (
	ProjectAll                    Project = "All"
	ProjectDomainNames                    = "Domain names"
	ProjectFreeServices                   = "Free services"
	ProjectEmails                         = "E-mails"
	ProjectWebHostingCloudDB              = "Web Hosting / CloudDB"
	ProjectDedicated                      = "Dedicated"
	ProjectEBackup                        = "E-Backup"
	ProjectTel2Pay                        = "Tel2Pay"
	ProjectNetworkAndRacks                = "Network and racks"
	ProjectManager                        = "Manager"
	ProjectDatacenters                    = "Datacenters"
	ProjectSupport                        = "Support"
	ProjectDemo1g                         = "Demo1g"
	ProjectRPS                            = "RPS"
	ProjectVoIP                           = "VoIP"
	ProjectDistributionsEtOS              = "Distributions et OS"
	ProjectCloud                          = "Cloud"
	ProjectVRackInfrastructure            = "vRack infrastructure"
	ProjectXDSLFTTH                       = "xDSL / FTTH"
	ProjectDedicatedCloud                 = "Dedicated Cloud"
	ProjectVPS                            = "VPS"
	ProjectCDN                            = "CDN"
	ProjectAntiDDoS                       = "Anti-DDoS"
	ProjectHubiC                          = "hubiC"
	ProjectCorporateWebsite               = "Corporate Website"
	ProjectTest                           = "test"
	ProjectDBaaS                          = "DBaaS"
	ProjectOverTheBox                     = "OverTheBox"
	ProjectMicrosoft                      = "Microsoft"
	ProjectDocker                         = "Docker"
	ProjectVirtualDesktop                 = "Virtual Desktop"
	ProjectNasHA                          = "Nas-HA"
	ProjectKubernetes                     = "Kubernetes"
	ProjectManagedPrivateRegistry         = "Managed Private Registry"
	ProjectOVHCloudConnect                = "OVHcloud Connect"
	ProjectAISolutions                    = "AI Solutions"
	ProjectDataAndAnalytics               = "Data & Analytics"

	TaskStatusAll        TaskStatus = "All"
	TaskStatusInProgress            = "In progress"
	TaskStatusPlanned               = "Planned"
	TaskStatusClosed                = "Closed"
)

var projectMapping = map[Project]int{
	ProjectAll:                    0,
	ProjectDomainNames:            1,
	ProjectFreeServices:           2,
	ProjectEmails:                 3,
	ProjectWebHostingCloudDB:      4,
	ProjectDedicated:              5,
	ProjectEBackup:                7,
	ProjectTel2Pay:                8,
	ProjectNetworkAndRacks:        9,
	ProjectManager:                10,
	ProjectDatacenters:            11,
	ProjectSupport:                12,
	ProjectDemo1g:                 14,
	ProjectRPS:                    15,
	ProjectVoIP:                   16,
	ProjectDistributionsEtOS:      17,
	ProjectCloud:                  18,
	ProjectVRackInfrastructure:    19,
	ProjectXDSLFTTH:               20,
	ProjectDedicatedCloud:         21,
	ProjectVPS:                    22,
	ProjectCDN:                    23,
	ProjectAntiDDoS:               24,
	ProjectHubiC:                  26,
	ProjectCorporateWebsite:       27,
	ProjectTest:                   28,
	ProjectDBaaS:                  29,
	ProjectOverTheBox:             30,
	ProjectMicrosoft:              31,
	ProjectDocker:                 32,
	ProjectVirtualDesktop:         33,
	ProjectNasHA:                  34,
	ProjectKubernetes:             35,
	ProjectManagedPrivateRegistry: 36,
	ProjectOVHCloudConnect:        37,
	ProjectAISolutions:            38,
	ProjectDataAndAnalytics:       39,
}

var incidentStatusMapping = map[TaskStatus]int{
	TaskStatusInProgress: 1,
	TaskStatusPlanned:    2,
}

func (p Project) ProjectID() int {
	return projectMapping[p]
}

func (s TaskStatus) StatusID() string {
	switch s {
	case TaskStatusAll, TaskStatusClosed:
		return strings.ToLower(string(s))
	default:
		return strconv.Itoa(incidentStatusMapping[s])
	}
}

// URL builds a Travaux site URL to crawl.
func URL(project Project, status TaskStatus) (*url.URL, error) {
	travauxURL, err := url.Parse("http://travaux.ovh.net/")
	if err != nil {
		return nil, err
	}

	query := travauxURL.Query()
	query.Set("project", strconv.Itoa(project.ProjectID()))
	query.Set("status", status.StatusID())

	travauxURL.RawQuery = query.Encode()

	return travauxURL, nil
}
