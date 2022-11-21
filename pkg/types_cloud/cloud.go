package types_cloud

import (
	"time"
	"github.com/lalmeras/clio/pkg/types"
)


type Cloud_AccessTypeEnum string
const (
	Cloud_AccessTypeEnum_full Cloud_AccessTypeEnum = "full"
	Cloud_AccessTypeEnum_restricted  = "restricted"
)

type Cloud_Acl struct {
	AccountId string            `json:"accountId"`
	Type      Cloud_AclTypeEnum `json:"type"`
}

type Cloud_AclTypeEnum string
const (
	Cloud_AclTypeEnum_readOnly Cloud_AclTypeEnum = "readOnly"
	Cloud_AclTypeEnum_readWrite  = "readWrite"
)

type Cloud_Alerting struct {
	CreationDate              time.Time               `json:"creationDate"`
	Delay                     Cloud_AlertingDelayEnum `json:"delay"`
	Email                     string                  `json:"email"`
	FormattedMonthlyThreshold Order_Price             `json:"formattedMonthlyThreshold"`
	Id                        string                  `json:"id"`
	MonthlyThreshold          int64                   `json:"monthlyThreshold"`
}

type Cloud_AlertingAlert struct {
	AlertDate time.Time `json:"alertDate"`
	AlertId   int64     `json:"alertId"`
	Emails    []string  `json:"emails"`
}

type Cloud_AlertingDelayEnum int64
const (
	Cloud_AlertingDelayEnum__10800 Cloud_AlertingDelayEnum = 10800
	Cloud_AlertingDelayEnum__172800  = 172800
	Cloud_AlertingDelayEnum__21600  = 21600
	Cloud_AlertingDelayEnum__259200  = 259200
	Cloud_AlertingDelayEnum__3600  = 3600
	Cloud_AlertingDelayEnum__43200  = 43200
	Cloud_AlertingDelayEnum__604800  = 604800
	Cloud_AlertingDelayEnum__86400  = 86400
)

type Cloud_ArchiveStoragePrice struct {
	MonthlyPrice Order_Price `json:"monthlyPrice"`
	Region       string      `json:"region"`
}

type Cloud_AvailableRegion struct {
	ContinentCode      Cloud_RegionContinentEnum `json:"continentCode"`
	DatacenterLocation string                    `json:"datacenterLocation"`
	Name               string                    `json:"name"`
}

type Cloud_Backup struct {
	BackupName string            `json:"backupName"`
	CreatedAt  time.Time         `json:"createdAt"`
	Cron       string            `json:"cron"`
	Executions []Cloud_Execution `json:"executions"`
	Id         string            `json:"id"`
	InstanceId string            `json:"instanceId"`
	Name       string            `json:"name"`
}

type Cloud_BandwidthStoragePrice struct {
	Price  Order_Price `json:"price"`
	Region string      `json:"region"`
}

type Cloud_ColdArchiveContainer struct {
	CreatedAt    time.Time                            `json:"createdAt"`
	Name         string                               `json:"name"`
	Objects      []Cloud_StorageObject                `json:"objects"`
	ObjectsCount int64                                `json:"objectsCount"`
	ObjectsSize  int64                                `json:"objectsSize"`
	OwnerId      int64                                `json:"ownerId"`
	Status       Cloud_ColdArchiveContainerStatusEnum `json:"status"`
	VirtualHost  string                               `json:"virtualHost"`
}

type Cloud_ColdArchiveContainerStatusEnum string
const (
	Cloud_ColdArchiveContainerStatusEnum_archived Cloud_ColdArchiveContainerStatusEnum = "archived"
	Cloud_ColdArchiveContainerStatusEnum_archiving  = "archiving"
	Cloud_ColdArchiveContainerStatusEnum_deleting  = "deleting"
	Cloud_ColdArchiveContainerStatusEnum_draining  = "draining"
	Cloud_ColdArchiveContainerStatusEnum_flushed  = "flushed"
	Cloud_ColdArchiveContainerStatusEnum_locked  = "locked"
	Cloud_ColdArchiveContainerStatusEnum_none  = "none"
	Cloud_ColdArchiveContainerStatusEnum_restored  = "restored"
	Cloud_ColdArchiveContainerStatusEnum_restoring  = "restoring"
)

type Cloud_Component struct {
	Endpoint string                  `json:"endpoint"`
	Name     string                  `json:"name"`
	Status   Cloud_ServiceStatusEnum `json:"status"`
}

type Cloud_Credit struct {
	Available_credit Order_Price                 `json:"available_credit"`
	Bill             string                      `json:"bill"`
	Description      string                      `json:"description"`
	Id               int64                       `json:"id"`
	Products         []string                    `json:"products"`
	Total_credit     Order_Price                 `json:"total_credit"`
	Used_credit      Order_Price                 `json:"used_credit"`
	Validity         Cloudcommon_VoucherValidity `json:"validity"`
	Voucher          string                      `json:"voucher"`
}

type Cloud_Execution struct {
	ExecutedAt time.Time                `json:"executedAt"`
	State      Cloud_ExecutionStateEnum `json:"state"`
	StateInfo  string                   `json:"stateInfo"`
}

type Cloud_ExecutionState string
const (
	Cloud_ExecutionState_IDLE Cloud_ExecutionState = "IDLE"
	Cloud_ExecutionState_RUNNING  = "RUNNING"
	Cloud_ExecutionState_SUCCESS  = "SUCCESS"
	Cloud_ExecutionState_ERROR  = "ERROR"
	Cloud_ExecutionState_PAUSED  = "PAUSED"
)

type Cloud_ExecutionStateEnum string
const (
	Cloud_ExecutionStateEnum_CANCELED Cloud_ExecutionStateEnum = "CANCELED"
	Cloud_ExecutionStateEnum_ERROR  = "ERROR"
	Cloud_ExecutionStateEnum_IDLE  = "IDLE"
	Cloud_ExecutionStateEnum_PAUSED  = "PAUSED"
	Cloud_ExecutionStateEnum_RUNNING  = "RUNNING"
	Cloud_ExecutionStateEnum_SUCCESS  = "SUCCESS"
)

type Cloud_FlavorPrice struct {
	FlavorId     string      `json:"flavorId"`
	FlavorName   string      `json:"flavorName"`
	MonthlyPrice Order_Price `json:"monthlyPrice"`
	Price        Order_Price `json:"price"`
	Region       string      `json:"region"`
}

type Cloud_IpCountryEnum string
const (
	Cloud_IpCountryEnum_au Cloud_IpCountryEnum = "au"
	Cloud_IpCountryEnum_be  = "be"
	Cloud_IpCountryEnum_ca  = "ca"
	Cloud_IpCountryEnum_cz  = "cz"
	Cloud_IpCountryEnum_de  = "de"
	Cloud_IpCountryEnum_es  = "es"
	Cloud_IpCountryEnum_fi  = "fi"
	Cloud_IpCountryEnum_fr  = "fr"
	Cloud_IpCountryEnum_ie  = "ie"
	Cloud_IpCountryEnum_it  = "it"
	Cloud_IpCountryEnum_lt  = "lt"
	Cloud_IpCountryEnum_nl  = "nl"
	Cloud_IpCountryEnum_pl  = "pl"
	Cloud_IpCountryEnum_pt  = "pt"
	Cloud_IpCountryEnum_sg  = "sg"
	Cloud_IpCountryEnum_uk  = "uk"
	Cloud_IpCountryEnum_us  = "us"
)

type Cloud_Lab struct {
	Id     string              `json:"id"`
	Name   string              `json:"name"`
	Status Cloud_LabStatusEnum `json:"status"`
}

type Cloud_LabAgreements struct {
	Accepted []int64 `json:"accepted"`
	ToAccept []int64 `json:"toAccept"`
}

type Cloud_LabStatus string
const (
	Cloud_LabStatus_open Cloud_LabStatus = "open"
	Cloud_LabStatus_activating  = "activating"
	Cloud_LabStatus_activated  = "activated"
	Cloud_LabStatus_closed  = "closed"
)

type Cloud_LabStatusEnum string
const (
	Cloud_LabStatusEnum_activated Cloud_LabStatusEnum = "activated"
	Cloud_LabStatusEnum_activating  = "activating"
	Cloud_LabStatusEnum_closed  = "closed"
	Cloud_LabStatusEnum_open  = "open"
)

type Cloud_Operation struct {
	Action      string                    `json:"action"`
	CompletedAt time.Time                 `json:"completedAt"`
	CreatedAt   time.Time                 `json:"createdAt"`
	Id          string                    `json:"id"`
	Progress    int64                     `json:"progress"`
	Regions     []string                  `json:"regions"`
	ResourceId  string                    `json:"resourceId"`
	StartedAt   time.Time                 `json:"startedAt"`
	Status      Cloud_OperationStatusEnum `json:"status"`
}

type Cloud_OperationStatus string
const (
	Cloud_OperationStatus_created Cloud_OperationStatus = "created"
	Cloud_OperationStatus_inprogress  = "in-progress"
	Cloud_OperationStatus_completed  = "completed"
	Cloud_OperationStatus_inerror  = "in-error"
	Cloud_OperationStatus_unknown  = "unknown"
)

type Cloud_OperationStatusEnum string
const (
	Cloud_OperationStatusEnum_completed Cloud_OperationStatusEnum = "completed"
	Cloud_OperationStatusEnum_created  = "created"
	Cloud_OperationStatusEnum_inerror  = "in-error"
	Cloud_OperationStatusEnum_inprogress  = "in-progress"
	Cloud_OperationStatusEnum_unknown  = "unknown"
)

type Cloud_Price struct {
	Archive             []Cloud_ArchiveStoragePrice   `json:"archive"`
	BandwidthArchiveIn  []Cloud_BandwidthStoragePrice `json:"bandwidthArchiveIn"`
	BandwidthArchiveOut []Cloud_BandwidthStoragePrice `json:"bandwidthArchiveOut"`
	BandwidthStorage    []Cloud_BandwidthStoragePrice `json:"bandwidthStorage"`
	Instances           []Cloud_FlavorPrice           `json:"instances"`
	ProjectCreation     Order_Price                   `json:"projectCreation"`
	Snapshots           []Cloud_SnapshotPrice         `json:"snapshots"`
	Storage             []Cloud_StoragePrice          `json:"storage"`
	Volumes             []Cloud_VolumePrice           `json:"volumes"`
}

type Cloud_Project struct {
	Access       Cloud_AccessTypeEnum           `json:"access"`
	CreationDate time.Time                      `json:"creationDate"`
	Description  string                         `json:"description"`
	Expiration   time.Time                      `json:"expiration"`
	ManualQuota  bool                           `json:"manualQuota"`
	OrderId      int64                          `json:"orderId"`
	PlanCode     string                         `json:"planCode"`
	ProjectName  string                         `json:"projectName"`
	Project_id   string                         `json:"project_id"`
	Status       Cloudproject_ProjectStatusEnum `json:"status"`
	Unleash      bool                           `json:"unleash"`
}

type Cloud_ProjectActivateMonthlyBillingCreation struct {
	Instances []Cloudinstance_MonthlyInstanceBulkParams `json:"instances"`
}

type Cloud_ProjectContainerRegistryCreation struct {
	Name   string `json:"name"`
	PlanID string `json:"planID"`
	Region string `json:"region"`
}

type Cloud_ProjectContainerRegistryUpdate struct {
	Name string `json:"name"`
}

type Cloud_ProjectContainerRegistryUsersCreation struct {
	Email string `json:"email"`
	Login string `json:"login"`
}

type Cloud_ProjectInstanceBulkCreation struct {
	Autobackup     Cloudinstance_AutoBackup          `json:"autobackup"`
	FlavorId       string                            `json:"flavorId"`
	GroupId        string                            `json:"groupId"`
	ImageId        string                            `json:"imageId"`
	MonthlyBilling bool                              `json:"monthlyBilling"`
	Name           string                            `json:"name"`
	Networks       []Cloudinstance_NetworkBulkParams `json:"networks"`
	Number         int64                             `json:"number"`
	Region         string                            `json:"region"`
	SshKeyId       string                            `json:"sshKeyId"`
	UserData       string                            `json:"userData"`
	VolumeId       string                            `json:"volumeId"`
}

type Cloud_ProjectInstanceCreation struct {
	Autobackup     Cloudinstance_AutoBackup      `json:"autobackup"`
	FlavorId       string                        `json:"flavorId"`
	GroupId        string                        `json:"groupId"`
	ImageId        string                        `json:"imageId"`
	MonthlyBilling bool                          `json:"monthlyBilling"`
	Name           string                        `json:"name"`
	Networks       []Cloudinstance_NetworkParams `json:"networks"`
	Region         string                        `json:"region"`
	SshKeyId       string                        `json:"sshKeyId"`
	UserData       string                        `json:"userData"`
	VolumeId       string                        `json:"volumeId"`
}

type Cloud_ProjectInstanceGroupCreation struct {
	Name   string                                   `json:"name"`
	Region string                                   `json:"region"`
	Type   Cloudinstancegroup_InstanceGroupTypeEnum `json:"type"`
}

type Cloud_ProjectInstanceInterfaceCreation struct {
	Ip        string `json:"ip"`
	NetworkId string `json:"networkId"`
}

type Cloud_ProjectInstanceRebootCreation struct {
	Type Cloudinstance_RebootTypeEnum `json:"type"`
}

type Cloud_ProjectInstanceReinstallCreation struct {
	ImageId string `json:"imageId"`
}

type Cloud_ProjectInstanceRescueModeCreation struct {
	ImageId string `json:"imageId"`
	Rescue  bool   `json:"rescue"`
}

type Cloud_ProjectInstanceResizeCreation struct {
	FlavorId string `json:"flavorId"`
}

type Cloud_ProjectInstanceSnapshotCreation struct {
	SnapshotName string `json:"snapshotName"`
}

type Cloud_ProjectInstanceUpdate struct {
	InstanceName string `json:"instanceName"`
}

type Cloud_ProjectIpFailoverAttachCreation struct {
	InstanceId string `json:"instanceId"`
}

type Cloud_ProjectKubeCreation struct {
	Customization               Cloud_ProjectKubeCustomization        `json:"customization"`
	Name                        string                                `json:"name"`
	Nodepool                    Cloud_ProjectKubeCreationNodePool     `json:"nodepool"`
	PrivateNetworkConfiguration Cloudkube_PrivateNetworkConfiguration `json:"privateNetworkConfiguration"`
	PrivateNetworkId            string                                `json:"privateNetworkId"`
	Region                      string                                `json:"region"`
	UpdatePolicy                Cloudkube_UpdatePolicyEnum            `json:"updatePolicy"`
	Version                     Cloudkube_VersionEnum                 `json:"version"`
}

type Cloud_ProjectKubeCreationNodePool struct {
	AntiAffinity  bool                       `json:"antiAffinity"`
	Autoscale     bool                       `json:"autoscale"`
	DesiredNodes  int64                      `json:"desiredNodes"`
	FlavorName    string                     `json:"flavorName"`
	MaxNodes      int64                      `json:"maxNodes"`
	MinNodes      int64                      `json:"minNodes"`
	MonthlyBilled bool                       `json:"monthlyBilled"`
	Name          string                     `json:"name"`
	Template      Cloudkube_NodePoolTemplate `json:"template"`
}

type Cloud_ProjectKubeCustomization struct {
	ApiServer Cloud_ProjectKubeCustomizationAPIServer `json:"apiServer"`
}

type Cloud_ProjectKubeCustomizationAPIServer struct {
	AdmissionPlugins Cloud_ProjectKubeCustomizationAPIServerAdmissionPlugins `json:"admissionPlugins"`
}

type Cloud_ProjectKubeCustomizationAPIServerAdmissionPlugins struct {
	Disabled []Cloud_ProjectKubeCustomizationAPIServerAdmissionPluginsEnum `json:"disabled"`
	Enabled  []Cloud_ProjectKubeCustomizationAPIServerAdmissionPluginsEnum `json:"enabled"`
}

type Cloud_ProjectKubeCustomizationAPIServerAdmissionPluginsEnum string
const (
	Cloud_ProjectKubeCustomizationAPIServerAdmissionPluginsEnum_AlwaysPullImages Cloud_ProjectKubeCustomizationAPIServerAdmissionPluginsEnum = "AlwaysPullImages"
	Cloud_ProjectKubeCustomizationAPIServerAdmissionPluginsEnum_NodeRestriction  = "NodeRestriction"
)

type Cloud_ProjectKubeIpRestrictionUpsert struct {
	Ips []string `json:"ips"`
}

type Cloud_ProjectKubeNodeCreation struct {
	FlavorName string `json:"flavorName"`
	Name       string `json:"name"`
}

type Cloud_ProjectKubeNodePoolAutoscalingParams struct {
	ScaleDownUnneededTimeSeconds  int64   `json:"scaleDownUnneededTimeSeconds"`
	ScaleDownUnreadyTimeSeconds   int64   `json:"scaleDownUnreadyTimeSeconds"`
	ScaleDownUtilizationThreshold float64 `json:"scaleDownUtilizationThreshold"`
}

type Cloud_ProjectKubeNodePoolCreation struct {
	AntiAffinity  bool                                       `json:"antiAffinity"`
	Autoscale     bool                                       `json:"autoscale"`
	Autoscaling   Cloud_ProjectKubeNodePoolAutoscalingParams `json:"autoscaling"`
	DesiredNodes  int64                                      `json:"desiredNodes"`
	FlavorName    string                                     `json:"flavorName"`
	MaxNodes      int64                                      `json:"maxNodes"`
	MinNodes      int64                                      `json:"minNodes"`
	MonthlyBilled bool                                       `json:"monthlyBilled"`
	Name          string                                     `json:"name"`
	Template      Cloudkube_NodePoolTemplate                 `json:"template"`
}

type Cloud_ProjectKubeNodePoolUpdate struct {
	Autoscale     bool                                       `json:"autoscale"`
	Autoscaling   Cloud_ProjectKubeNodePoolAutoscalingParams `json:"autoscaling"`
	DesiredNodes  int64                                      `json:"desiredNodes"`
	MaxNodes      int64                                      `json:"maxNodes"`
	MinNodes      int64                                      `json:"minNodes"`
	NodesToRemove []string                                   `json:"nodesToRemove"`
	Template      Cloudkube_NodePoolTemplate                 `json:"template"`
}

type Cloud_ProjectKubeOpenIdConnectCreation struct {
	CaContent         string                                         `json:"caContent"`
	ClientId          string                                         `json:"clientId"`
	GroupsClaim       []string                                       `json:"groupsClaim"`
	GroupsPrefix      string                                         `json:"groupsPrefix"`
	IssuerUrl         string                                         `json:"issuerUrl"`
	RequiredClaim     []string                                       `json:"requiredClaim"`
	SigningAlgorithms []Cloudkube_OpenIdConnectSigningAlgorithmsEnum `json:"signingAlgorithms"`
	UsernameClaim     string                                         `json:"usernameClaim"`
	UsernamePrefix    string                                         `json:"usernamePrefix"`
}

type Cloud_ProjectKubeOpenIdConnectUpdate struct {
	CaContent         string                                         `json:"caContent"`
	ClientId          string                                         `json:"clientId"`
	GroupsClaim       []string                                       `json:"groupsClaim"`
	GroupsPrefix      string                                         `json:"groupsPrefix"`
	IssuerUrl         string                                         `json:"issuerUrl"`
	RequiredClaim     []string                                       `json:"requiredClaim"`
	SigningAlgorithms []Cloudkube_OpenIdConnectSigningAlgorithmsEnum `json:"signingAlgorithms"`
	UsernameClaim     string                                         `json:"usernameClaim"`
	UsernamePrefix    string                                         `json:"usernamePrefix"`
}

type Cloud_ProjectKubeResetCreation struct {
	Customization               Cloud_ProjectKubeCustomization        `json:"customization"`
	Name                        string                                `json:"name"`
	PrivateNetworkConfiguration Cloudkube_PrivateNetworkConfiguration `json:"privateNetworkConfiguration"`
	PrivateNetworkId            string                                `json:"privateNetworkId"`
	UpdatePolicy                Cloudkube_UpdatePolicyEnum            `json:"updatePolicy"`
	Version                     Cloudkube_VersionEnum                 `json:"version"`
	WorkerNodesPolicy           Cloudkube_ResetWorkerNodesPolicyEnum  `json:"workerNodesPolicy"`
}

type Cloud_ProjectKubeRestart struct {
	Force bool `json:"force"`
}

type Cloud_ProjectKubeUpdate struct {
	Name         string                     `json:"name"`
	UpdatePolicy Cloudkube_UpdatePolicyEnum `json:"updatePolicy"`
}

type Cloud_ProjectKubeUpdateCreation struct {
	Force    bool                         `json:"force"`
	Strategy Cloudkube_UpdateStrategyEnum `json:"strategy"`
}

type Cloud_ProjectKubeUpdatePolicyUpdate struct {
	UpdatePolicy Cloudkube_UpdatePolicyEnum `json:"updatePolicy"`
}

type Cloud_ProjectMigrationUpdate struct {
	Date time.Time `json:"date"`
}

type Cloud_ProjectNetworkPrivateCreation struct {
	Name    string   `json:"name"`
	Regions []string `json:"regions"`
	VlanId  int64    `json:"vlanId"`
}

type Cloud_ProjectNetworkPrivateRegionCreation struct {
	Region string `json:"region"`
}

type Cloud_ProjectNetworkPrivateSubnetCreation struct {
	Dhcp      bool   `json:"dhcp"`
	End       string `json:"end"`
	Network   string `json:"network"`
	NoGateway bool   `json:"noGateway"`
	Region    string `json:"region"`
	Start     string `json:"start"`
}

type Cloud_ProjectNetworkPrivateSubnetUpdate struct {
	Dhcp           bool   `json:"dhcp"`
	DisableGateway bool   `json:"disableGateway"`
	GatewayIp      string `json:"gatewayIp"`
}

type Cloud_ProjectNetworkPrivateUpdate struct {
	Name string `json:"name"`
}

type Cloud_ProjectRegionCreation struct {
	Region string `json:"region"`
}

type Cloud_ProjectRegionQuotaCreation struct {
	Name string `json:"name"`
}

type Cloud_ProjectRegionWorkflowBackupCreation struct {
	Cron              string `json:"cron"`
	InstanceId        string `json:"instanceId"`
	MaxExecutionCount int64  `json:"maxExecutionCount"`
	Name              string `json:"name"`
	Rotation          int64  `json:"rotation"`
}

type Cloud_ProjectSshkeyCreation struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	Region    string `json:"region"`
}

type Cloud_ProjectStorageCorsCreation struct {
	Origin string `json:"origin"`
}

type Cloud_ProjectStorageCreation struct {
	Archive       bool   `json:"archive"`
	ContainerName string `json:"containerName"`
	Region        string `json:"region"`
}

type Cloud_ProjectStoragePublicUrlCreation struct {
	ExpirationDate time.Time `json:"expirationDate"`
	ObjectName     string    `json:"objectName"`
}

type Cloud_ProjectStorageUpdate struct {
	ContainerType Cloudstorage_TypeEnum `json:"containerType"`
}

type Cloud_ProjectStorageUserCreation struct {
	Description string                 `json:"description"`
	Right       Cloudstorage_RightEnum `json:"right"`
}

type Cloud_ProjectUserCreation struct {
	Description string               `json:"description"`
	Role        Clouduser_RoleEnum   `json:"role"`
	Roles       []Clouduser_RoleEnum `json:"roles"`
}

type Cloud_ProjectUserRoleCreation struct {
	RoleId string `json:"roleId"`
}

type Cloud_ProjectUserRoleUpdate struct {
	RolesIds []string `json:"rolesIds"`
}

type Cloud_ProjectUserTokenCreation struct {
	Password string `json:"password"`
}

type Cloud_ProjectVolumeAttachCreation struct {
	InstanceId string `json:"instanceId"`
}

type Cloud_ProjectVolumeCreation struct {
	Description string                     `json:"description"`
	ImageId     string                     `json:"imageId"`
	Name        string                     `json:"name"`
	Region      string                     `json:"region"`
	Size        int64                      `json:"size"`
	SnapshotId  string                     `json:"snapshotId"`
	Type        Cloudvolume_VolumeTypeEnum `json:"type"`
}

type Cloud_ProjectVolumeDetachCreation struct {
	InstanceId string `json:"instanceId"`
}

type Cloud_ProjectVolumeSnapshotCreation struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Cloud_ProjectVolumeUpdate struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Cloud_ProjectVolumeUpsizeCreation struct {
	Size int64 `json:"size"`
}

type Cloud_Region struct {
	ContinentCode      Cloud_RegionContinentEnum `json:"continentCode"`
	DatacenterLocation string                    `json:"datacenterLocation"`
	IpCountries        []Cloud_IpCountryEnum     `json:"ipCountries"`
	Name               string                    `json:"name"`
	Services           []Cloud_Component         `json:"services"`
	Status             Cloud_RegionStatusEnum    `json:"status"`
}

type Cloud_RegionContinent string
const (
	Cloud_RegionContinent_EU Cloud_RegionContinent = "EU"
	Cloud_RegionContinent_NA  = "NA"
	Cloud_RegionContinent_US  = "US"
	Cloud_RegionContinent_ASIA  = "ASIA"
)

type Cloud_RegionContinentEnum string
const (
	Cloud_RegionContinentEnum_ASIA Cloud_RegionContinentEnum = "ASIA"
	Cloud_RegionContinentEnum_EU  = "EU"
	Cloud_RegionContinentEnum_NA  = "NA"
	Cloud_RegionContinentEnum_US  = "US"
)

type Cloud_RegionStatus string
const (
	Cloud_RegionStatus_UP Cloud_RegionStatus = "UP"
	Cloud_RegionStatus_DOWN  = "DOWN"
	Cloud_RegionStatus_MAINTENANCE  = "MAINTENANCE"
)

type Cloud_RegionStatusEnum string
const (
	Cloud_RegionStatusEnum_DOWN Cloud_RegionStatusEnum = "DOWN"
	Cloud_RegionStatusEnum_MAINTENANCE  = "MAINTENANCE"
	Cloud_RegionStatusEnum_UP  = "UP"
)

type Cloud_ServiceStatus string
const (
	Cloud_ServiceStatus_UP Cloud_ServiceStatus = "UP"
	Cloud_ServiceStatus_DOWN  = "DOWN"
)

type Cloud_ServiceStatusEnum string
const (
	Cloud_ServiceStatusEnum_DOWN Cloud_ServiceStatusEnum = "DOWN"
	Cloud_ServiceStatusEnum_UP  = "UP"
)

type Cloud_SnapshotPrice struct {
	MonthlyPrice Order_Price `json:"monthlyPrice"`
	Price        Order_Price `json:"price"`
	Region       string      `json:"region"`
}

type Cloud_StorageContainer struct {
	CreatedAt    time.Time             `json:"createdAt"`
	Name         string                `json:"name"`
	Objects      []Cloud_StorageObject `json:"objects"`
	ObjectsCount int64                 `json:"objectsCount"`
	ObjectsSize  int64                 `json:"objectsSize"`
	OwnerId      int64                 `json:"ownerId"`
	Region       string                `json:"region"`
	VirtualHost  string                `json:"virtualHost"`
}

type Cloud_StorageContainerCreation struct {
	Name    string `json:"name"`
	OwnerId int64  `json:"ownerId"`
}

type Cloud_StorageObject struct {
	Etag         string    `json:"etag"`
	Key          string    `json:"key"`
	LastModified time.Time `json:"lastModified"`
	Size         int64     `json:"size"`
}

type Cloud_StoragePrice struct {
	MonthlyPrice Order_Price `json:"monthlyPrice"`
	Price        Order_Price `json:"price"`
	Region       string      `json:"region"`
}

type Cloud_VolumePrice struct {
	MonthlyPrice Order_Price `json:"monthlyPrice"`
	Price        Order_Price `json:"price"`
	Region       string      `json:"region"`
	VolumeName   string      `json:"volumeName"`
}

type Cloud_Vrack struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Name        string `json:"name"`
}

type Cloudauthentication_Catalog struct {
	Endpoints []Cloudauthentication_Endpoint `json:"endpoints"`
	Id        string                         `json:"id"`
	Type      string                         `json:"type"`
}

type Cloudauthentication_Domain struct {
	Name string `json:"name"`
}

type Cloudauthentication_Endpoint struct {
	Id                 string `json:"id"`
	Interface          string `json:"interface"`
	Legacy_endpoint_id string `json:"legacy_endpoint_id"`
	Region_id          string `json:"region_id"`
	Service_id         string `json:"service_id"`
	Url                string `json:"url"`
}

type Cloudauthentication_OpenstackToken struct {
	Catalog    []Cloudauthentication_Catalog    `json:"catalog"`
	Expires_at time.Time                        `json:"expires_at"`
	Issued_at  time.Time                        `json:"issued_at"`
	Methods    []string                         `json:"methods"`
	Project    Cloudauthentication_TokenProject `json:"project"`
	Roles      []Cloudauthentication_Role       `json:"roles"`
	User       Cloudauthentication_UserToken    `json:"user"`
}

type Cloudauthentication_Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cloudauthentication_Token struct {
	XAuthToken string                             `json:"X-Auth-Token"`
	Token      Cloudauthentication_OpenstackToken `json:"token"`
}

type Cloudauthentication_TokenProject struct {
	Domain Cloudauthentication_Domain `json:"domain"`
	Id     string                     `json:"id"`
	Name   string                     `json:"name"`
}

type Cloudauthentication_UserToken struct {
	Domain Cloudauthentication_Domain `json:"domain"`
	Id     string                     `json:"id"`
	Name   string                     `json:"name"`
}

type CloudbillingView_BandwidthInstance struct {
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_BandwidthStorage struct {
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_Component struct {
	Name       string                    `json:"name"`
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_HourlyInstance struct {
	Details    []CloudbillingView_HourlyInstanceDetail `json:"details"`
	Quantity   CloudbillingView_Quantity               `json:"quantity"`
	Reference  string                                  `json:"reference"`
	Region     string                                  `json:"region"`
	TotalPrice float64                                 `json:"totalPrice"`
}

type CloudbillingView_HourlyInstanceBandwidth struct {
	IncomingBandwidth CloudbillingView_BandwidthInstance `json:"incomingBandwidth"`
	OutgoingBandwidth CloudbillingView_BandwidthInstance `json:"outgoingBandwidth"`
	Region            string                             `json:"region"`
	TotalPrice        float64                            `json:"totalPrice"`
}

type CloudbillingView_HourlyInstanceDetail struct {
	InstanceId string                    `json:"instanceId"`
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_HourlyInstanceOption struct {
	Details    []CloudbillingView_HourlyInstanceOptionDetail `json:"details"`
	Quantity   CloudbillingView_Quantity                     `json:"quantity"`
	Reference  string                                        `json:"reference"`
	Region     string                                        `json:"region"`
	TotalPrice float64                                       `json:"totalPrice"`
}

type CloudbillingView_HourlyInstanceOptionDetail struct {
	InstanceId string                    `json:"instanceId"`
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_HourlyResources struct {
	Instance          []CloudbillingView_HourlyInstance          `json:"instance"`
	InstanceBandwidth []CloudbillingView_HourlyInstanceBandwidth `json:"instanceBandwidth"`
	InstanceOption    []CloudbillingView_HourlyInstanceOption    `json:"instanceOption"`
	Snapshot          []CloudbillingView_HourlySnapshot          `json:"snapshot"`
	Storage           []CloudbillingView_HourlyStorage           `json:"storage"`
	Volume            []CloudbillingView_HourlyVolume            `json:"volume"`
}

type CloudbillingView_HourlySnapshot struct {
	Instance   CloudbillingView_InstanceSnapshot `json:"instance"`
	Region     string                            `json:"region"`
	TotalPrice float64                           `json:"totalPrice"`
	Volume     CloudbillingView_VolumeSnapshot   `json:"volume"`
}

type CloudbillingView_HourlyStorage struct {
	BucketName                string                            `json:"bucketName"`
	IncomingBandwidth         CloudbillingView_BandwidthStorage `json:"incomingBandwidth"`
	IncomingInternalBandwidth CloudbillingView_BandwidthStorage `json:"incomingInternalBandwidth"`
	OutgoingBandwidth         CloudbillingView_BandwidthStorage `json:"outgoingBandwidth"`
	OutgoingInternalBandwidth CloudbillingView_BandwidthStorage `json:"outgoingInternalBandwidth"`
	Region                    string                            `json:"region"`
	Stored                    CloudbillingView_StoredStorage    `json:"stored"`
	TotalPrice                float64                           `json:"totalPrice"`
	Type                      CloudbillingView_StorageTypeEnum  `json:"type"`
}

type CloudbillingView_HourlyVolume struct {
	Details    []CloudbillingView_HourlyVolumeDetail `json:"details"`
	Quantity   CloudbillingView_Quantity             `json:"quantity"`
	Region     string                                `json:"region"`
	TotalPrice float64                               `json:"totalPrice"`
	Type       string                                `json:"type"`
}

type CloudbillingView_HourlyVolumeDetail struct {
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
	VolumeId   string                    `json:"volumeId"`
}

type CloudbillingView_InstanceSnapshot struct {
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_MonthlyCertification struct {
	Details    []CloudbillingView_MonthlyCertificationDetail `json:"details"`
	Reference  string                                        `json:"reference"`
	TotalPrice float64                                       `json:"totalPrice"`
}

type CloudbillingView_MonthlyCertificationDetail struct {
	Activation time.Time `json:"activation"`
	TotalPrice float64   `json:"totalPrice"`
}

type CloudbillingView_MonthlyInstance struct {
	Details    []CloudbillingView_MonthlyInstanceDetail `json:"details"`
	Reference  string                                   `json:"reference"`
	Region     string                                   `json:"region"`
	TotalPrice float64                                  `json:"totalPrice"`
}

type CloudbillingView_MonthlyInstanceDetail struct {
	Activation time.Time `json:"activation"`
	InstanceId string    `json:"instanceId"`
	TotalPrice float64   `json:"totalPrice"`
}

type CloudbillingView_MonthlyInstanceOption struct {
	Details    []CloudbillingView_MonthlyInstanceOptionDetail `json:"details"`
	Reference  string                                         `json:"reference"`
	Region     string                                         `json:"region"`
	TotalPrice float64                                        `json:"totalPrice"`
}

type CloudbillingView_MonthlyInstanceOptionDetail struct {
	InstanceId string  `json:"instanceId"`
	TotalPrice float64 `json:"totalPrice"`
}

type CloudbillingView_MonthlyResources struct {
	Certification  []CloudbillingView_MonthlyCertification  `json:"certification"`
	Instance       []CloudbillingView_MonthlyInstance       `json:"instance"`
	InstanceOption []CloudbillingView_MonthlyInstanceOption `json:"instanceOption"`
}

type CloudbillingView_Quantity struct {
	Unit  CloudbillingView_UnitQuantityEnum `json:"unit"`
	Value float64                           `json:"value"`
}

type CloudbillingView_RegionalizedResource struct {
	Components []CloudbillingView_Component `json:"components"`
	Region     string                       `json:"region"`
}

type CloudbillingView_StorageTypeEnum string
const (
	CloudbillingView_StorageTypeEnum_pca CloudbillingView_StorageTypeEnum = "pca"
	CloudbillingView_StorageTypeEnum_pcs  = "pcs"
	CloudbillingView_StorageTypeEnum_storagecoldarchive  = "storage-coldarchive"
	CloudbillingView_StorageTypeEnum_storagehighperf  = "storage-high-perf"
	CloudbillingView_StorageTypeEnum_storagestandard  = "storage-standard"
)

type CloudbillingView_StoredStorage struct {
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type CloudbillingView_TypedResources struct {
	Resources  []CloudbillingView_RegionalizedResource `json:"resources"`
	TotalPrice float64                                 `json:"totalPrice"`
	Type       string                                  `json:"type"`
}

type CloudbillingView_UnitQuantity string
const (
	CloudbillingView_UnitQuantity_GiB CloudbillingView_UnitQuantity = "GiB"
	CloudbillingView_UnitQuantity_GiBh  = "GiBh"
	CloudbillingView_UnitQuantity_Hour  = "Hour"
)

type CloudbillingView_UnitQuantityEnum string
const (
	CloudbillingView_UnitQuantityEnum_GiB CloudbillingView_UnitQuantityEnum = "GiB"
	CloudbillingView_UnitQuantityEnum_GiBh  = "GiBh"
	CloudbillingView_UnitQuantityEnum_Hour  = "Hour"
	CloudbillingView_UnitQuantityEnum_Minute  = "Minute"
	CloudbillingView_UnitQuantityEnum_Second  = "Second"
	CloudbillingView_UnitQuantityEnum_Unit  = "Unit"
)

type CloudbillingView_UsedCredit struct {
	Description string  `json:"description"`
	Id          int64   `json:"id"`
	UsedAmount  float64 `json:"usedAmount"`
}

type CloudbillingView_UsedCredits struct {
	Details     []CloudbillingView_UsedCredit `json:"details"`
	TotalCredit float64                       `json:"totalCredit"`
}

type CloudbillingView_VolumeSnapshot struct {
	Quantity   CloudbillingView_Quantity `json:"quantity"`
	TotalPrice float64                   `json:"totalPrice"`
}

type Cloudcapabilities_Availability struct {
	Plans    []Cloudcapabilities_AvailabilityPlan    `json:"plans"`
	Products []Cloudcapabilities_AvailabilityProduct `json:"products"`
}

type Cloudcapabilities_AvailabilityPlan struct {
	Code    string                                 `json:"code"`
	Regions []Cloudcapabilities_AvailabilityRegion `json:"regions"`
}

type Cloudcapabilities_AvailabilityProduct struct {
	Name    string                                 `json:"name"`
	Regions []Cloudcapabilities_AvailabilityRegion `json:"regions"`
}

type Cloudcapabilities_AvailabilityRegion struct {
	ContinentCode Cloud_RegionContinentEnum `json:"continentCode"`
	Datacenter    string                    `json:"datacenter"`
	Enabled       bool                      `json:"enabled"`
	Name          string                    `json:"name"`
}

type Cloudcapabilities_Capability struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

type Cloudcommon_VoucherValidity struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type CloudcontainerRegistry_Capability struct {
	Plans      []CloudcontainerRegistry_Plan `json:"plans"`
	RegionName string                        `json:"regionName"`
}

type CloudcontainerRegistry_Features struct {
	Vulnerability bool `json:"vulnerability"`
}

type CloudcontainerRegistry_Limits struct {
	ImageStorage    int64 `json:"imageStorage"`
	ParallelRequest int64 `json:"parallelRequest"`
}

type CloudcontainerRegistry_Plan struct {
	Code           string                          `json:"code"`
	CreatedAt      time.Time                       `json:"createdAt"`
	Features       CloudcontainerRegistry_Features `json:"features"`
	Id             string                          `json:"id"`
	Name           string                          `json:"name"`
	RegistryLimits CloudcontainerRegistry_Limits   `json:"registryLimits"`
	UpdatedAt      time.Time                       `json:"updatedAt"`
}

type CloudcontainerRegistry_PlanUpdate struct {
	PlanID string `json:"planID"`
}

type CloudcontainerRegistry_Registry struct {
	CreatedAt time.Time                         `json:"createdAt"`
	Id        string                            `json:"id"`
	Name      string                            `json:"name"`
	ProjectID string                            `json:"projectID"`
	Region    string                            `json:"region"`
	Size      int64                             `json:"size"`
	Status    CloudcontainerRegistry_StatusEnum `json:"status"`
	UpdatedAt time.Time                         `json:"updatedAt"`
	Url       string                            `json:"url"`
	Version   string                            `json:"version"`
}

type CloudcontainerRegistry_StatusEnum string
const (
	CloudcontainerRegistry_StatusEnum_DELETED CloudcontainerRegistry_StatusEnum = "DELETED"
	CloudcontainerRegistry_StatusEnum_DELETING  = "DELETING"
	CloudcontainerRegistry_StatusEnum_ERROR  = "ERROR"
	CloudcontainerRegistry_StatusEnum_INSTALLING  = "INSTALLING"
	CloudcontainerRegistry_StatusEnum_READY  = "READY"
	CloudcontainerRegistry_StatusEnum_RESTORING  = "RESTORING"
	CloudcontainerRegistry_StatusEnum_SUSPENDED  = "SUSPENDED"
	CloudcontainerRegistry_StatusEnum_SUSPENDING  = "SUSPENDING"
	CloudcontainerRegistry_StatusEnum_UPDATING  = "UPDATING"
)

type CloudcontainerRegistry_User struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Password string `json:"password"`
	User     string `json:"user"`
}

type CloudcontainerRegistryregistry_RegionEnum string
const (
	CloudcontainerRegistryregistry_RegionEnum_GRA7 CloudcontainerRegistryregistry_RegionEnum = "GRA7"
)

type CloudcontainerRegistryregistry_Registry struct {
	CreatedAt time.Time                                 `json:"createdAt"`
	Id        string                                    `json:"id"`
	Name      string                                    `json:"name"`
	ProjectID string                                    `json:"projectID"`
	Region    CloudcontainerRegistryregistry_RegionEnum `json:"region"`
	Status    CloudcontainerRegistryregistry_StatusEnum `json:"status"`
	UpdatedAt time.Time                                 `json:"updatedAt"`
	Url       string                                    `json:"url"`
	Version   string                                    `json:"version"`
}

type CloudcontainerRegistryregistry_StatusEnum string
const (
	CloudcontainerRegistryregistry_StatusEnum_ERROR CloudcontainerRegistryregistry_StatusEnum = "ERROR"
	CloudcontainerRegistryregistry_StatusEnum_READY  = "READY"
	CloudcontainerRegistryregistry_StatusEnum_DELETED  = "DELETED"
	CloudcontainerRegistryregistry_StatusEnum_SUSPENDED  = "SUSPENDED"
	CloudcontainerRegistryregistry_StatusEnum_INSTALLING  = "INSTALLING"
	CloudcontainerRegistryregistry_StatusEnum_UPDATING  = "UPDATING"
	CloudcontainerRegistryregistry_StatusEnum_RESTORING  = "RESTORING"
	CloudcontainerRegistryregistry_StatusEnum_SUSPENDING  = "SUSPENDING"
	CloudcontainerRegistryregistry_StatusEnum_DELETING  = "DELETING"
)

type CloudcontainerRegistryuser_User struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Password string `json:"password"`
	User     string `json:"user"`
}

type Cloudflavor_Capability struct {
	Enabled bool                           `json:"enabled"`
	Name    Cloudflavor_CapabilityNameEnum `json:"name"`
}

type Cloudflavor_CapabilityNameEnum string
const (
	Cloudflavor_CapabilityNameEnum_failoverip Cloudflavor_CapabilityNameEnum = "failoverip"
	Cloudflavor_CapabilityNameEnum_resize  = "resize"
	Cloudflavor_CapabilityNameEnum_snapshot  = "snapshot"
	Cloudflavor_CapabilityNameEnum_volume  = "volume"
)

type Cloudflavor_Flavor struct {
	Available         bool                        `json:"available"`
	Capabilities      []Cloudflavor_Capability    `json:"capabilities"`
	Disk              int64                       `json:"disk"`
	Id                string                      `json:"id"`
	InboundBandwidth  int64                       `json:"inboundBandwidth"`
	Name              string                      `json:"name"`
	OsType            string                      `json:"osType"`
	OutboundBandwidth int64                       `json:"outboundBandwidth"`
	PlanCodes         Cloudflavor_FlavorPlanCodes `json:"planCodes"`
	Quota             int64                       `json:"quota"`
	Ram               int64                       `json:"ram"`
	Region            string                      `json:"region"`
	Type              string                      `json:"type"`
	Vcpus             int64                       `json:"vcpus"`
}

type Cloudflavor_FlavorPlanCodes struct {
	Hourly  string `json:"hourly"`
	Monthly string `json:"monthly"`
}

type Cloudforecast_ProjectForecast struct {
	LastMetric      time.Time   `json:"lastMetric"`
	ProjectForecast Order_Price `json:"projectForecast"`
}

type Cloudimage_Image struct {
	CreationDate time.Time `json:"creationDate"`
	FlavorType   string    `json:"flavorType"`
	Id           string    `json:"id"`
	MinDisk      int64     `json:"minDisk"`
	MinRam       int64     `json:"minRam"`
	Name         string    `json:"name"`
	PlanCode     string    `json:"planCode"`
	Region       string    `json:"region"`
	Size         float64   `json:"size"`
	Status       string    `json:"status"`
	Tags         []string  `json:"tags"`
	Type         string    `json:"type"`
	User         string    `json:"user"`
	Visibility   string    `json:"visibility"`
}

type Cloudimage_OSTypeEnum string
const (
	Cloudimage_OSTypeEnum_baremetallinux Cloudimage_OSTypeEnum = "baremetal-linux"
	Cloudimage_OSTypeEnum_bsd  = "bsd"
	Cloudimage_OSTypeEnum_linux  = "linux"
	Cloudimage_OSTypeEnum_windows  = "windows"
)

type Cloudinstance_Access struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}

type Cloudinstance_ApplicationAccess struct {
	Accesses []Cloudinstance_Access                   `json:"accesses"`
	Status   Cloudinstance_ApplicationAccessStateEnum `json:"status"`
}

type Cloudinstance_ApplicationAccessStateEnum string
const (
	Cloudinstance_ApplicationAccessStateEnum_installing Cloudinstance_ApplicationAccessStateEnum = "installing"
	Cloudinstance_ApplicationAccessStateEnum_ok  = "ok"
)

type Cloudinstance_AssociateFloatingIp struct {
	FloatingIpId string                            `json:"floatingIpId"`
	Gateway      Cloudnetwork_CreateGatewaySummary `json:"gateway"`
	Ip           string                            `json:"ip"`
}

type Cloudinstance_AutoBackup struct {
	Cron     string `json:"cron"`
	Rotation int64  `json:"rotation"`
}

type Cloudinstance_CreateFloatingIp struct {
	Gateway Cloudnetwork_CreateGatewaySummary `json:"gateway"`
	Ip      string                            `json:"ip"`
}

type Cloudinstance_Instance struct {
	Created                     time.Time                        `json:"created"`
	CurrentMonthOutgoingTraffic int64                            `json:"currentMonthOutgoingTraffic"`
	FlavorId                    string                           `json:"flavorId"`
	Id                          string                           `json:"id"`
	ImageId                     string                           `json:"imageId"`
	IpAddresses                 []Cloudinstance_IpAddress        `json:"ipAddresses"`
	MonthlyBilling              Cloudinstance_MonthlyBilling     `json:"monthlyBilling"`
	Name                        string                           `json:"name"`
	OperationIds                []string                         `json:"operationIds"`
	PlanCode                    string                           `json:"planCode"`
	Region                      string                           `json:"region"`
	SshKeyId                    string                           `json:"sshKeyId"`
	Status                      Cloudinstance_InstanceStatusEnum `json:"status"`
}

type Cloudinstance_InstanceDetail struct {
	Created                     time.Time                        `json:"created"`
	CurrentMonthOutgoingTraffic int64                            `json:"currentMonthOutgoingTraffic"`
	Flavor                      Cloudflavor_Flavor               `json:"flavor"`
	Id                          string                           `json:"id"`
	Image                       Cloudimage_Image                 `json:"image"`
	IpAddresses                 []Cloudinstance_IpAddress        `json:"ipAddresses"`
	MonthlyBilling              Cloudinstance_MonthlyBilling     `json:"monthlyBilling"`
	Name                        string                           `json:"name"`
	OperationIds                []string                         `json:"operationIds"`
	PlanCode                    string                           `json:"planCode"`
	Region                      string                           `json:"region"`
	RescuePassword              string                           `json:"rescuePassword"`
	SshKey                      Cloudsshkey_SshKeyDetail         `json:"sshKey"`
	Status                      Cloudinstance_InstanceStatusEnum `json:"status"`
}

type Cloudinstance_InstanceMetrics struct {
	Unit   string                               `json:"unit"`
	Values []Cloudinstance_InstanceMetricsValue `json:"values"`
}

type Cloudinstance_InstanceMetricsValue struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type Cloudinstance_InstanceStatusEnum string
const (
	Cloudinstance_InstanceStatusEnum_ACTIVE Cloudinstance_InstanceStatusEnum = "ACTIVE"
	Cloudinstance_InstanceStatusEnum_BUILD  = "BUILD"
	Cloudinstance_InstanceStatusEnum_BUILDING  = "BUILDING"
	Cloudinstance_InstanceStatusEnum_DELETED  = "DELETED"
	Cloudinstance_InstanceStatusEnum_DELETING  = "DELETING"
	Cloudinstance_InstanceStatusEnum_ERROR  = "ERROR"
	Cloudinstance_InstanceStatusEnum_HARD_REBOOT  = "HARD_REBOOT"
	Cloudinstance_InstanceStatusEnum_MIGRATING  = "MIGRATING"
	Cloudinstance_InstanceStatusEnum_PASSWORD  = "PASSWORD"
	Cloudinstance_InstanceStatusEnum_PAUSED  = "PAUSED"
	Cloudinstance_InstanceStatusEnum_REBOOT  = "REBOOT"
	Cloudinstance_InstanceStatusEnum_REBUILD  = "REBUILD"
	Cloudinstance_InstanceStatusEnum_RESCUE  = "RESCUE"
	Cloudinstance_InstanceStatusEnum_RESCUED  = "RESCUED"
	Cloudinstance_InstanceStatusEnum_RESCUING  = "RESCUING"
	Cloudinstance_InstanceStatusEnum_RESIZE  = "RESIZE"
	Cloudinstance_InstanceStatusEnum_RESIZED  = "RESIZED"
	Cloudinstance_InstanceStatusEnum_RESUMING  = "RESUMING"
	Cloudinstance_InstanceStatusEnum_REVERT_RESIZE  = "REVERT_RESIZE"
	Cloudinstance_InstanceStatusEnum_SHELVED  = "SHELVED"
	Cloudinstance_InstanceStatusEnum_SHELVED_OFFLOADED  = "SHELVED_OFFLOADED"
	Cloudinstance_InstanceStatusEnum_SHELVING  = "SHELVING"
	Cloudinstance_InstanceStatusEnum_SHUTOFF  = "SHUTOFF"
	Cloudinstance_InstanceStatusEnum_SNAPSHOTTING  = "SNAPSHOTTING"
	Cloudinstance_InstanceStatusEnum_SOFT_DELETED  = "SOFT_DELETED"
	Cloudinstance_InstanceStatusEnum_STOPPED  = "STOPPED"
	Cloudinstance_InstanceStatusEnum_SUSPENDED  = "SUSPENDED"
	Cloudinstance_InstanceStatusEnum_UNKNOWN  = "UNKNOWN"
	Cloudinstance_InstanceStatusEnum_UNRESCUING  = "UNRESCUING"
	Cloudinstance_InstanceStatusEnum_UNSHELVING  = "UNSHELVING"
	Cloudinstance_InstanceStatusEnum_VERIFY_RESIZE  = "VERIFY_RESIZE"
)

type Cloudinstance_InstanceSummary struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cloudinstance_InstanceVnc struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Cloudinstance_IpAddress struct {
	GatewayIp string `json:"gatewayIp"`
	Ip        string `json:"ip"`
	NetworkId string `json:"networkId"`
	Type      string `json:"type"`
	Version   int64  `json:"version"`
}

type Cloudinstance_MetricsPeriod string
const (
	Cloudinstance_MetricsPeriod_lastday Cloudinstance_MetricsPeriod = "lastday"
	Cloudinstance_MetricsPeriod_lastmonth  = "lastmonth"
	Cloudinstance_MetricsPeriod_lastweek  = "lastweek"
	Cloudinstance_MetricsPeriod_lastyear  = "lastyear"
	Cloudinstance_MetricsPeriod_today  = "today"
)

type Cloudinstance_MetricsPeriodEnum string
const (
	Cloudinstance_MetricsPeriodEnum_lastday Cloudinstance_MetricsPeriodEnum = "lastday"
	Cloudinstance_MetricsPeriodEnum_lastmonth  = "lastmonth"
	Cloudinstance_MetricsPeriodEnum_lastweek  = "lastweek"
	Cloudinstance_MetricsPeriodEnum_lastyear  = "lastyear"
	Cloudinstance_MetricsPeriodEnum_today  = "today"
)

type Cloudinstance_MetricsType string
const (
	Cloudinstance_MetricsType_memused Cloudinstance_MetricsType = "mem:used"
	Cloudinstance_MetricsType_memmax  = "mem:max"
	Cloudinstance_MetricsType_cpuused  = "cpu:used"
	Cloudinstance_MetricsType_cpumax  = "cpu:max"
	Cloudinstance_MetricsType_nettx  = "net:tx"
	Cloudinstance_MetricsType_netrx  = "net:rx"
)

type Cloudinstance_MetricsTypeEnum string
const (
	Cloudinstance_MetricsTypeEnum_cpumax Cloudinstance_MetricsTypeEnum = "cpu:max"
	Cloudinstance_MetricsTypeEnum_cpuused  = "cpu:used"
	Cloudinstance_MetricsTypeEnum_memmax  = "mem:max"
	Cloudinstance_MetricsTypeEnum_memused  = "mem:used"
	Cloudinstance_MetricsTypeEnum_netrx  = "net:rx"
	Cloudinstance_MetricsTypeEnum_nettx  = "net:tx"
)

type Cloudinstance_MonthlyBilling struct {
	Since  time.Time                              `json:"since"`
	Status Cloudinstance_MonthlyBillingStatusEnum `json:"status"`
}

type Cloudinstance_MonthlyBillingStatusEnum string
const (
	Cloudinstance_MonthlyBillingStatusEnum_activationPending Cloudinstance_MonthlyBillingStatusEnum = "activationPending"
	Cloudinstance_MonthlyBillingStatusEnum_ok  = "ok"
)

type Cloudinstance_MonthlyInstanceBulkParams struct {
	InstanceId string `json:"instanceId"`
	Region     string `json:"region"`
}

type Cloudinstance_NetworkBulkParams struct {
	NetworkId string `json:"networkId"`
}

type Cloudinstance_NetworkParams struct {
	Ip        string `json:"ip"`
	NetworkId string `json:"networkId"`
}

type Cloudinstance_RebootTypeEnum string
const (
	Cloudinstance_RebootTypeEnum_hard Cloudinstance_RebootTypeEnum = "hard"
	Cloudinstance_RebootTypeEnum_soft  = "soft"
)

type Cloudinstance_RescueAdminPassword struct {
	AdminPassword string `json:"adminPassword"`
}

type CloudinstanceInterface_FixedIp struct {
	Ip       string `json:"ip"`
	SubnetId string `json:"subnetId"`
}

type CloudinstanceInterface_Interface struct {
	FixedIps   []CloudinstanceInterface_FixedIp `json:"fixedIps"`
	Id         string                           `json:"id"`
	MacAddress string                           `json:"macAddress"`
	NetworkId  string                           `json:"networkId"`
	State      string                           `json:"state"`
	Type       string                           `json:"type"`
}

type Cloudinstancegroup_InstanceGroup struct {
	Id           string                                   `json:"id"`
	Instance_ids []string                                 `json:"instance_ids"`
	Name         string                                   `json:"name"`
	Region       string                                   `json:"region"`
	Type         Cloudinstancegroup_InstanceGroupTypeEnum `json:"type"`
}

type Cloudinstancegroup_InstanceGroupTypeEnum string
const (
	Cloudinstancegroup_InstanceGroupTypeEnum_affinity Cloudinstancegroup_InstanceGroupTypeEnum = "affinity"
	Cloudinstancegroup_InstanceGroupTypeEnum_antiaffinity  = "anti-affinity"
)

type Cloudip_CloudIp struct {
	Id     string               `json:"id"`
	Ip     string               `json:"ip"`
	Status Cloudip_IpStatusEnum `json:"status"`
	Type   string               `json:"type"`
}

type Cloudip_FailoverIp struct {
	Block         string                `json:"block"`
	ContinentCode string                `json:"continentCode"`
	Geoloc        string                `json:"geoloc"`
	Id            string                `json:"id"`
	Ip            string                `json:"ip"`
	Progress      int64                 `json:"progress"`
	RoutedTo      string                `json:"routedTo"`
	Status        Cloudip_IpStatusEnum  `json:"status"`
	SubType       Cloudip_IpSubTypeEnum `json:"subType"`
}

type Cloudip_IpStatusEnum string
const (
	Cloudip_IpStatusEnum_ok Cloudip_IpStatusEnum = "ok"
	Cloudip_IpStatusEnum_operationPending  = "operationPending"
)

type Cloudip_IpSubTypeEnum string
const (
	Cloudip_IpSubTypeEnum_cloud Cloudip_IpSubTypeEnum = "cloud"
	Cloudip_IpSubTypeEnum_ovh  = "ovh"
)

type Cloudkeymanager_Certificate struct {
	Id      string                      `json:"id"`
	Name    string                      `json:"name"`
	Region  string                      `json:"region"`
	Secrets []Cloudkeymanager_SecretRef `json:"secrets"`
}

type Cloudkeymanager_CertificateCreate struct {
	Certificate          string `json:"certificate"`
	Intermediates        string `json:"intermediates"`
	Name                 string `json:"name"`
	PrivateKey           string `json:"privateKey"`
	PrivateKeyPassphrase string `json:"privateKeyPassphrase"`
}

type Cloudkeymanager_Secret struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}

type Cloudkeymanager_SecretRef struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cloudkube_AuditLogs struct {
	ExpirationDate time.Time `json:"expirationDate"`
	Url            string    `json:"url"`
}

type Cloudkube_Cluster struct {
	ControlPlaneIsUpToDate      bool                                  `json:"controlPlaneIsUpToDate"`
	CreatedAt                   time.Time                             `json:"createdAt"`
	Customization               Cloud_ProjectKubeCustomization        `json:"customization"`
	Id                          string                                `json:"id"`
	IsUpToDate                  bool                                  `json:"isUpToDate"`
	Name                        string                                `json:"name"`
	NextUpgradeVersions         []Cloudkube_VersionEnum               `json:"nextUpgradeVersions"`
	NodesUrl                    string                                `json:"nodesUrl"`
	PrivateNetworkConfiguration Cloudkube_PrivateNetworkConfiguration `json:"privateNetworkConfiguration"`
	PrivateNetworkId            string                                `json:"privateNetworkId"`
	Region                      Cloudkube_RegionEnum                  `json:"region"`
	Status                      Cloudkube_ClusterStatusEnum           `json:"status"`
	UpdatePolicy                string                                `json:"updatePolicy"`
	UpdatedAt                   time.Time                             `json:"updatedAt"`
	Url                         string                                `json:"url"`
	Version                     string                                `json:"version"`
}

type Cloudkube_ClusterStatus string
const (
	Cloudkube_ClusterStatus_INSTALLING Cloudkube_ClusterStatus = "INSTALLING"
	Cloudkube_ClusterStatus_UPDATING  = "UPDATING"
	Cloudkube_ClusterStatus_RESETTING  = "RESETTING"
	Cloudkube_ClusterStatus_SUSPENDING  = "SUSPENDING"
	Cloudkube_ClusterStatus_REOPENING  = "REOPENING"
	Cloudkube_ClusterStatus_DELETING  = "DELETING"
	Cloudkube_ClusterStatus_SUSPENDED  = "SUSPENDED"
	Cloudkube_ClusterStatus_ERROR  = "ERROR"
	Cloudkube_ClusterStatus_USER_ERROR  = "USER_ERROR"
	Cloudkube_ClusterStatus_USER_QUOTA_ERROR  = "USER_QUOTA_ERROR"
	Cloudkube_ClusterStatus_READY  = "READY"
)

type Cloudkube_ClusterStatusEnum string
const (
	Cloudkube_ClusterStatusEnum_DELETED Cloudkube_ClusterStatusEnum = "DELETED"
	Cloudkube_ClusterStatusEnum_DELETING  = "DELETING"
	Cloudkube_ClusterStatusEnum_DOWNSCALING  = "DOWNSCALING"
	Cloudkube_ClusterStatusEnum_ERROR  = "ERROR"
	Cloudkube_ClusterStatusEnum_INSTALLING  = "INSTALLING"
	Cloudkube_ClusterStatusEnum_MAINTENANCE  = "MAINTENANCE"
	Cloudkube_ClusterStatusEnum_READY  = "READY"
	Cloudkube_ClusterStatusEnum_REDEPLOYING  = "REDEPLOYING"
	Cloudkube_ClusterStatusEnum_REOPENING  = "REOPENING"
	Cloudkube_ClusterStatusEnum_RESETTING  = "RESETTING"
	Cloudkube_ClusterStatusEnum_RESIZING  = "RESIZING"
	Cloudkube_ClusterStatusEnum_SUSPENDED  = "SUSPENDED"
	Cloudkube_ClusterStatusEnum_SUSPENDING  = "SUSPENDING"
	Cloudkube_ClusterStatusEnum_UPDATING  = "UPDATING"
	Cloudkube_ClusterStatusEnum_UPSCALING  = "UPSCALING"
	Cloudkube_ClusterStatusEnum_USER_ERROR  = "USER_ERROR"
	Cloudkube_ClusterStatusEnum_USER_NODE_NOT_FOUND_ERROR  = "USER_NODE_NOT_FOUND_ERROR"
	Cloudkube_ClusterStatusEnum_USER_NODE_SUSPENDED_SERVICE  = "USER_NODE_SUSPENDED_SERVICE"
	Cloudkube_ClusterStatusEnum_USER_QUOTA_ERROR  = "USER_QUOTA_ERROR"
)

type Cloudkube_Flavor struct {
	Category Cloudkube_FlavorCategoryEnum `json:"category"`
	Gpus     int64                        `json:"gpus"`
	Name     string                       `json:"name"`
	Ram      int64                        `json:"ram"`
	State    Cloudkube_FlavorStateEnum    `json:"state"`
	VCPUs    int64                        `json:"vCPUs"`
}

type Cloudkube_FlavorCategory string
const (
	Cloudkube_FlavorCategory_c Cloudkube_FlavorCategory = "c"
	Cloudkube_FlavorCategory_g  = "g"
	Cloudkube_FlavorCategory_t  = "t"
	Cloudkube_FlavorCategory_b  = "b"
	Cloudkube_FlavorCategory_r  = "r"
)

type Cloudkube_FlavorCategoryEnum string
const (
	Cloudkube_FlavorCategoryEnum_b Cloudkube_FlavorCategoryEnum = "b"
	Cloudkube_FlavorCategoryEnum_c  = "c"
	Cloudkube_FlavorCategoryEnum_d  = "d"
	Cloudkube_FlavorCategoryEnum_g  = "g"
	Cloudkube_FlavorCategoryEnum_i  = "i"
	Cloudkube_FlavorCategoryEnum_r  = "r"
	Cloudkube_FlavorCategoryEnum_t  = "t"
)

type Cloudkube_FlavorState string
const (
	Cloudkube_FlavorState_available Cloudkube_FlavorState = "available"
	Cloudkube_FlavorState_unavailable  = "unavailable"
)

type Cloudkube_FlavorStateEnum string
const (
	Cloudkube_FlavorStateEnum_available Cloudkube_FlavorStateEnum = "available"
	Cloudkube_FlavorStateEnum_unavailable  = "unavailable"
)

type Cloudkube_Kubeconfig struct {
	Content string `json:"content"`
}

type Cloudkube_Node struct {
	CreatedAt  time.Time                `json:"createdAt"`
	DeployedAt time.Time                `json:"deployedAt"`
	Flavor     string                   `json:"flavor"`
	Id         string                   `json:"id"`
	InstanceId string                   `json:"instanceId"`
	IsUpToDate bool                     `json:"isUpToDate"`
	Name       string                   `json:"name"`
	NodePoolId string                   `json:"nodePoolId"`
	ProjectId  string                   `json:"projectId"`
	Status     Cloudkube_NodeStatusEnum `json:"status"`
	UpdatedAt  time.Time                `json:"updatedAt"`
	Version    string                   `json:"version"`
}

type Cloudkube_NodePool struct {
	AntiAffinity   bool                             `json:"antiAffinity"`
	Autoscale      bool                             `json:"autoscale"`
	Autoscaling    Cloudkube_NodePoolAutoscaling    `json:"autoscaling"`
	AvailableNodes int64                            `json:"availableNodes"`
	CreatedAt      time.Time                        `json:"createdAt"`
	CurrentNodes   int64                            `json:"currentNodes"`
	DesiredNodes   int64                            `json:"desiredNodes"`
	Flavor         string                           `json:"flavor"`
	Id             string                           `json:"id"`
	MaxNodes       int64                            `json:"maxNodes"`
	MinNodes       int64                            `json:"minNodes"`
	MonthlyBilled  bool                             `json:"monthlyBilled"`
	Name           string                           `json:"name"`
	ProjectId      string                           `json:"projectId"`
	SizeStatus     Cloudkube_NodePoolSizeStatusEnum `json:"sizeStatus"`
	Status         Cloudkube_NodePoolStatusEnum     `json:"status"`
	Template       Cloudkube_NodePoolTemplate       `json:"template"`
	UpToDateNodes  int64                            `json:"upToDateNodes"`
	UpdatedAt      time.Time                        `json:"updatedAt"`
}

type Cloudkube_NodePoolAutoscaling struct {
	ScaleDownUnneededTimeSeconds  int64   `json:"scaleDownUnneededTimeSeconds"`
	ScaleDownUnreadyTimeSeconds   int64   `json:"scaleDownUnreadyTimeSeconds"`
	ScaleDownUtilizationThreshold float64 `json:"scaleDownUtilizationThreshold"`
}

type Cloudkube_NodePoolSizeStatusEnum string
const (
	Cloudkube_NodePoolSizeStatusEnum_CAPACITY_OK Cloudkube_NodePoolSizeStatusEnum = "CAPACITY_OK"
	Cloudkube_NodePoolSizeStatusEnum_OVER_CAPACITY  = "OVER_CAPACITY"
	Cloudkube_NodePoolSizeStatusEnum_UNDER_CAPACITY  = "UNDER_CAPACITY"
)

type Cloudkube_NodePoolStatusEnum string
const (
	Cloudkube_NodePoolStatusEnum_DELETED Cloudkube_NodePoolStatusEnum = "DELETED"
	Cloudkube_NodePoolStatusEnum_DELETING  = "DELETING"
	Cloudkube_NodePoolStatusEnum_DOWNSCALING  = "DOWNSCALING"
	Cloudkube_NodePoolStatusEnum_ERROR  = "ERROR"
	Cloudkube_NodePoolStatusEnum_INSTALLING  = "INSTALLING"
	Cloudkube_NodePoolStatusEnum_MAINTENANCE  = "MAINTENANCE"
	Cloudkube_NodePoolStatusEnum_READY  = "READY"
	Cloudkube_NodePoolStatusEnum_REDEPLOYING  = "REDEPLOYING"
	Cloudkube_NodePoolStatusEnum_REOPENING  = "REOPENING"
	Cloudkube_NodePoolStatusEnum_RESETTING  = "RESETTING"
	Cloudkube_NodePoolStatusEnum_RESIZING  = "RESIZING"
	Cloudkube_NodePoolStatusEnum_SUSPENDED  = "SUSPENDED"
	Cloudkube_NodePoolStatusEnum_SUSPENDING  = "SUSPENDING"
	Cloudkube_NodePoolStatusEnum_UPDATING  = "UPDATING"
	Cloudkube_NodePoolStatusEnum_UPSCALING  = "UPSCALING"
	Cloudkube_NodePoolStatusEnum_USER_ERROR  = "USER_ERROR"
	Cloudkube_NodePoolStatusEnum_USER_NODE_NOT_FOUND_ERROR  = "USER_NODE_NOT_FOUND_ERROR"
	Cloudkube_NodePoolStatusEnum_USER_NODE_SUSPENDED_SERVICE  = "USER_NODE_SUSPENDED_SERVICE"
	Cloudkube_NodePoolStatusEnum_USER_QUOTA_ERROR  = "USER_QUOTA_ERROR"
)

type Cloudkube_NodePoolTemplate struct {
	Metadata Cloudkube_NodePoolTemplateMetadata `json:"metadata"`
	Spec     Cloudkube_NodePoolTemplateSpec     `json:"spec"`
}

type Cloudkube_NodePoolTemplateMetadata struct {
	Annotations map[string]string `json:"annotations"`
	Finalizers  []string          `json:"finalizers"`
	Labels      map[string]string `json:"labels"`
}

type Cloudkube_NodePoolTemplateSpec struct {
	Taints        []Cloudkube_Taint `json:"taints"`
	Unschedulable bool              `json:"unschedulable"`
}

type Cloudkube_NodeStatus string
const (
	Cloudkube_NodeStatus_INSTALLING Cloudkube_NodeStatus = "INSTALLING"
	Cloudkube_NodeStatus_UPDATING  = "UPDATING"
	Cloudkube_NodeStatus_RESETTING  = "RESETTING"
	Cloudkube_NodeStatus_SUSPENDING  = "SUSPENDING"
	Cloudkube_NodeStatus_REOPENING  = "REOPENING"
	Cloudkube_NodeStatus_DELETING  = "DELETING"
	Cloudkube_NodeStatus_SUSPENDED  = "SUSPENDED"
	Cloudkube_NodeStatus_ERROR  = "ERROR"
	Cloudkube_NodeStatus_USER_ERROR  = "USER_ERROR"
	Cloudkube_NodeStatus_USER_QUOTA_ERROR  = "USER_QUOTA_ERROR"
	Cloudkube_NodeStatus_USER_NODE_NOT_FOUND_ERROR  = "USER_NODE_NOT_FOUND_ERROR"
	Cloudkube_NodeStatus_USER_NODE_SUSPENDED_SERVICE  = "USER_NODE_SUSPENDED_SERVICE"
	Cloudkube_NodeStatus_READY  = "READY"
)

type Cloudkube_NodeStatusEnum string
const (
	Cloudkube_NodeStatusEnum_DELETED Cloudkube_NodeStatusEnum = "DELETED"
	Cloudkube_NodeStatusEnum_DELETING  = "DELETING"
	Cloudkube_NodeStatusEnum_DOWNSCALING  = "DOWNSCALING"
	Cloudkube_NodeStatusEnum_ERROR  = "ERROR"
	Cloudkube_NodeStatusEnum_INSTALLING  = "INSTALLING"
	Cloudkube_NodeStatusEnum_MAINTENANCE  = "MAINTENANCE"
	Cloudkube_NodeStatusEnum_READY  = "READY"
	Cloudkube_NodeStatusEnum_REDEPLOYING  = "REDEPLOYING"
	Cloudkube_NodeStatusEnum_REOPENING  = "REOPENING"
	Cloudkube_NodeStatusEnum_RESETTING  = "RESETTING"
	Cloudkube_NodeStatusEnum_RESIZING  = "RESIZING"
	Cloudkube_NodeStatusEnum_SUSPENDED  = "SUSPENDED"
	Cloudkube_NodeStatusEnum_SUSPENDING  = "SUSPENDING"
	Cloudkube_NodeStatusEnum_UPDATING  = "UPDATING"
	Cloudkube_NodeStatusEnum_UPSCALING  = "UPSCALING"
	Cloudkube_NodeStatusEnum_USER_ERROR  = "USER_ERROR"
	Cloudkube_NodeStatusEnum_USER_NODE_NOT_FOUND_ERROR  = "USER_NODE_NOT_FOUND_ERROR"
	Cloudkube_NodeStatusEnum_USER_NODE_SUSPENDED_SERVICE  = "USER_NODE_SUSPENDED_SERVICE"
	Cloudkube_NodeStatusEnum_USER_QUOTA_ERROR  = "USER_QUOTA_ERROR"
)

type Cloudkube_OpenIdConnect struct {
	CaContent         string                                         `json:"caContent"`
	ClientId          string                                         `json:"clientId"`
	GroupsClaim       []string                                       `json:"groupsClaim"`
	GroupsPrefix      string                                         `json:"groupsPrefix"`
	IssuerUrl         string                                         `json:"issuerUrl"`
	RequiredClaim     []string                                       `json:"requiredClaim"`
	SigningAlgorithms []Cloudkube_OpenIdConnectSigningAlgorithmsEnum `json:"signingAlgorithms"`
	UsernameClaim     string                                         `json:"usernameClaim"`
	UsernamePrefix    string                                         `json:"usernamePrefix"`
}

type Cloudkube_OpenIdConnectSigningAlgorithmsEnum string
const (
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_ES256 Cloudkube_OpenIdConnectSigningAlgorithmsEnum = "ES256"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_ES384  = "ES384"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_ES512  = "ES512"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_PS256  = "PS256"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_PS384  = "PS384"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_PS512  = "PS512"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_RS256  = "RS256"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_RS384  = "RS384"
	Cloudkube_OpenIdConnectSigningAlgorithmsEnum_RS512  = "RS512"
)

type Cloudkube_PrivateNetworkConfiguration struct {
	DefaultVrackGateway            string `json:"defaultVrackGateway"`
	PrivateNetworkRoutingAsDefault bool   `json:"privateNetworkRoutingAsDefault"`
}

type Cloudkube_Region string
const (
	Cloudkube_Region_GRA5 Cloudkube_Region = "GRA5"
	Cloudkube_Region_GRA7  = "GRA7"
	Cloudkube_Region_BHS5  = "BHS5"
)

type Cloudkube_RegionEnum string
const (
	Cloudkube_RegionEnum_BHS5 Cloudkube_RegionEnum = "BHS5"
	Cloudkube_RegionEnum_DE1  = "DE1"
	Cloudkube_RegionEnum_GRA5  = "GRA5"
	Cloudkube_RegionEnum_GRA7  = "GRA7"
	Cloudkube_RegionEnum_GRA9  = "GRA9"
	Cloudkube_RegionEnum_SBG5  = "SBG5"
	Cloudkube_RegionEnum_SGP1  = "SGP1"
	Cloudkube_RegionEnum_SYD1  = "SYD1"
	Cloudkube_RegionEnum_UK1  = "UK1"
	Cloudkube_RegionEnum_USEASTVA1  = "US-EAST-VA-1"
	Cloudkube_RegionEnum_USWESTOR1  = "US-WEST-OR-1"
	Cloudkube_RegionEnum_WAW1  = "WAW1"
)

type Cloudkube_ResetWorkerNodesPolicy string
const (
	Cloudkube_ResetWorkerNodesPolicy_reinstall Cloudkube_ResetWorkerNodesPolicy = "reinstall"
	Cloudkube_ResetWorkerNodesPolicy_delete  = "delete"
)

type Cloudkube_ResetWorkerNodesPolicyEnum string
const (
	Cloudkube_ResetWorkerNodesPolicyEnum_delete Cloudkube_ResetWorkerNodesPolicyEnum = "delete"
	Cloudkube_ResetWorkerNodesPolicyEnum_reinstall  = "reinstall"
)

type Cloudkube_ResponseMessage struct {
	Message string `json:"message"`
}

type Cloudkube_Taint struct {
	Effect Cloudkube_TaintEffectEnum `json:"effect"`
	Key    string                    `json:"key"`
	Value  string                    `json:"value"`
}

type Cloudkube_TaintEffectEnum string
const (
	Cloudkube_TaintEffectEnum_NoExecute Cloudkube_TaintEffectEnum = "NoExecute"
	Cloudkube_TaintEffectEnum_NoSchedule  = "NoSchedule"
	Cloudkube_TaintEffectEnum_PreferNoSchedule  = "PreferNoSchedule"
)

type Cloudkube_UpdatePolicy string
const (
	Cloudkube_UpdatePolicy_ALWAYS_UPDATE Cloudkube_UpdatePolicy = "ALWAYS_UPDATE"
	Cloudkube_UpdatePolicy_MINIMAL_DOWNTIME  = "MINIMAL_DOWNTIME"
	Cloudkube_UpdatePolicy_NEVER_UPDATE  = "NEVER_UPDATE"
)

type Cloudkube_UpdatePolicyEnum string
const (
	Cloudkube_UpdatePolicyEnum_ALWAYS_UPDATE Cloudkube_UpdatePolicyEnum = "ALWAYS_UPDATE"
	Cloudkube_UpdatePolicyEnum_MINIMAL_DOWNTIME  = "MINIMAL_DOWNTIME"
	Cloudkube_UpdatePolicyEnum_NEVER_UPDATE  = "NEVER_UPDATE"
)

type Cloudkube_UpdateStrategy string
const (
	Cloudkube_UpdateStrategy_LATEST_PATCH Cloudkube_UpdateStrategy = "LATEST_PATCH"
	Cloudkube_UpdateStrategy_NEXT_MINOR  = "NEXT_MINOR"
)

type Cloudkube_UpdateStrategyEnum string
const (
	Cloudkube_UpdateStrategyEnum_LATEST_PATCH Cloudkube_UpdateStrategyEnum = "LATEST_PATCH"
	Cloudkube_UpdateStrategyEnum_NEXT_MINOR  = "NEXT_MINOR"
)

type Cloudkube_UpgradeVersion string
const (
	Cloudkube_UpgradeVersion__112 Cloudkube_UpgradeVersion = "1.12"
	Cloudkube_UpgradeVersion__113  = "1.13"
	Cloudkube_UpgradeVersion__114  = "1.14"
	Cloudkube_UpgradeVersion__115  = "1.15"
	Cloudkube_UpgradeVersion__116  = "1.16"
)

type Cloudkube_Version string
const (
	Cloudkube_Version__113 Cloudkube_Version = "1.13"
	Cloudkube_Version__114  = "1.14"
	Cloudkube_Version__115  = "1.15"
)

type Cloudkube_VersionEnum string
const (
	Cloudkube_VersionEnum__120 Cloudkube_VersionEnum = "1.20"
	Cloudkube_VersionEnum__121  = "1.21"
	Cloudkube_VersionEnum__122  = "1.22"
	Cloudkube_VersionEnum__123  = "1.23"
	Cloudkube_VersionEnum__124  = "1.24"
)

type Cloudloadbalancing_AssociateFloatingIp struct {
	FloatingIpId string                            `json:"floatingIpId"`
	Gateway      Cloudnetwork_CreateGatewaySummary `json:"gateway"`
	Ip           string                            `json:"ip"`
}

type Cloudloadbalancing_CreateFloatingIp struct {
	Gateway Cloudnetwork_CreateGatewaySummary `json:"gateway"`
	Ip      string                            `json:"ip"`
}

type Cloudloadbalancing_CreateListener struct {
	CertificateId  string                                  `json:"certificateId"`
	DefaultPoolId  string                                  `json:"defaultPoolId"`
	LoadbalancerId string                                  `json:"loadbalancerId"`
	Name           string                                  `json:"name"`
	Port           int64                                   `json:"port"`
	Protocol       Cloudloadbalancing_ListenerProtocolEnum `json:"protocol"`
}

type Cloudloadbalancing_Flavor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cloudloadbalancing_HealthMonitor struct {
	HttpConfiguration  Cloudloadbalancing_HealthMonitorHTTPConfiguration     `json:"httpConfiguration"`
	Id                 string                                                `json:"id"`
	MaxRetries         int64                                                 `json:"maxRetries"`
	MaxRetriesDown     int64                                                 `json:"maxRetriesDown"`
	MonitorType        Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum  `json:"monitorType"`
	Name               string                                                `json:"name"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	Periodicity        int64                                                 `json:"periodicity"`
	PoolId             string                                                `json:"poolId"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
	Timeout            int64                                                 `json:"timeout"`
}

type Cloudloadbalancing_HealthMonitorHTTPConfiguration struct {
	DomainName    string                                                      `json:"domainName"`
	ExpectedCodes []int64                                                     `json:"expectedCodes"`
	HttpMethod    Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum  `json:"httpMethod"`
	HttpVersion   Cloudloadbalancing_LoadBalancerHealthMonitorHTTPVersionEnum `json:"httpVersion"`
	UrlPath       string                                                      `json:"urlPath"`
}

type Cloudloadbalancing_L7Policy struct {
	Action             Cloudloadbalancing_L7PolicyActionEnum                 `json:"action"`
	Description        string                                                `json:"description"`
	Id                 string                                                `json:"id"`
	ListenerId         string                                                `json:"listenerId"`
	Name               string                                                `json:"name"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	Position           int64                                                 `json:"position"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
	RedirectHttpCode   int64                                                 `json:"redirectHttpCode"`
	RedirectPoolId     string                                                `json:"redirectPoolId"`
	RedirectPrefix     string                                                `json:"redirectPrefix"`
	RedirectUrl        string                                                `json:"redirectUrl"`
}

type Cloudloadbalancing_L7PolicyActionEnum string
const (
	Cloudloadbalancing_L7PolicyActionEnum_redirectPrefix Cloudloadbalancing_L7PolicyActionEnum = "redirectPrefix"
	Cloudloadbalancing_L7PolicyActionEnum_redirectToPool  = "redirectToPool"
	Cloudloadbalancing_L7PolicyActionEnum_redirectToURL  = "redirectToURL"
	Cloudloadbalancing_L7PolicyActionEnum_reject  = "reject"
)

type Cloudloadbalancing_L7PolicyUpdate struct {
	Action           Cloudloadbalancing_L7PolicyActionEnum `json:"action"`
	Description      string                                `json:"description"`
	ListenerId       string                                `json:"listenerId"`
	Name             string                                `json:"name"`
	Position         int64                                 `json:"position"`
	RedirectHttpCode int64                                 `json:"redirectHttpCode"`
	RedirectPoolId   string                                `json:"redirectPoolId"`
	RedirectPrefix   string                                `json:"redirectPrefix"`
	RedirectUrl      string                                `json:"redirectUrl"`
}

type Cloudloadbalancing_L7Rule struct {
	CompareType        Cloudloadbalancing_L7RuleCompareTypeEnum              `json:"compareType"`
	Id                 string                                                `json:"id"`
	Invert             bool                                                  `json:"invert"`
	Key                string                                                `json:"key"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
	RuleType           Cloudloadbalancing_L7RuleTypeEnum                     `json:"ruleType"`
	Value              string                                                `json:"value"`
}

type Cloudloadbalancing_L7RuleCompareTypeEnum string
const (
	Cloudloadbalancing_L7RuleCompareTypeEnum_contains Cloudloadbalancing_L7RuleCompareTypeEnum = "contains"
	Cloudloadbalancing_L7RuleCompareTypeEnum_endsWith  = "endsWith"
	Cloudloadbalancing_L7RuleCompareTypeEnum_equalTo  = "equalTo"
	Cloudloadbalancing_L7RuleCompareTypeEnum_regex  = "regex"
	Cloudloadbalancing_L7RuleCompareTypeEnum_startsWith  = "startsWith"
)

type Cloudloadbalancing_L7RuleTypeEnum string
const (
	Cloudloadbalancing_L7RuleTypeEnum_cookie Cloudloadbalancing_L7RuleTypeEnum = "cookie"
	Cloudloadbalancing_L7RuleTypeEnum_fileType  = "fileType"
	Cloudloadbalancing_L7RuleTypeEnum_header  = "header"
	Cloudloadbalancing_L7RuleTypeEnum_hostName  = "hostName"
	Cloudloadbalancing_L7RuleTypeEnum_path  = "path"
	Cloudloadbalancing_L7RuleTypeEnum_sslConnHasCert  = "sslConnHasCert"
	Cloudloadbalancing_L7RuleTypeEnum_sslDNField  = "sslDNField"
	Cloudloadbalancing_L7RuleTypeEnum_sslVerifyResult  = "sslVerifyResult"
)

type Cloudloadbalancing_Listener struct {
	CertificateId      string                                                `json:"certificateId"`
	DefaultPoolId      string                                                `json:"defaultPoolId"`
	Id                 string                                                `json:"id"`
	LoadBalancerIds    []string                                              `json:"loadBalancerIds"`
	Name               string                                                `json:"name"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	Port               int64                                                 `json:"port"`
	Protocol           Cloudloadbalancing_ListenerProtocolEnum               `json:"protocol"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
}

type Cloudloadbalancing_ListenerProtocolEnum string
const (
	Cloudloadbalancing_ListenerProtocolEnum_http Cloudloadbalancing_ListenerProtocolEnum = "http"
	Cloudloadbalancing_ListenerProtocolEnum_https  = "https"
	Cloudloadbalancing_ListenerProtocolEnum_sctp  = "sctp"
	Cloudloadbalancing_ListenerProtocolEnum_tcp  = "tcp"
	Cloudloadbalancing_ListenerProtocolEnum_terminatedHTTPS  = "terminatedHTTPS"
	Cloudloadbalancing_ListenerProtocolEnum_udp  = "udp"
)

type Cloudloadbalancing_LoadBalancer struct {
	CreatedAt          time.Time                                             `json:"createdAt"`
	FlavorId           string                                                `json:"flavorId"`
	Id                 string                                                `json:"id"`
	Name               string                                                `json:"name"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
	UpdatedAt          time.Time                                             `json:"updatedAt"`
	VipAddress         string                                                `json:"vipAddress"`
	VipNetworkId       string                                                `json:"vipNetworkId"`
	VipSubnetId        string                                                `json:"vipSubnetId"`
}

type Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum string
const (
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_CONNECT Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum = "CONNECT"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_DELETE  = "DELETE"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_GET  = "GET"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_HEAD  = "HEAD"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_OPTIONS  = "OPTIONS"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_PATCH  = "PATCH"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_POST  = "POST"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_PUT  = "PUT"
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPMethodEnum_TRACE  = "TRACE"
)

type Cloudloadbalancing_LoadBalancerHealthMonitorHTTPVersionEnum float64
const (
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPVersionEnum__10 Cloudloadbalancing_LoadBalancerHealthMonitorHTTPVersionEnum = 1.0
	Cloudloadbalancing_LoadBalancerHealthMonitorHTTPVersionEnum__11  = 1.1
)

type Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum string
const (
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_http Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum = "http"
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_https  = "https"
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_ping  = "ping"
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_sctp  = "sctp"
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_tcp  = "tcp"
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_tlshello  = "tls-hello"
	Cloudloadbalancing_LoadBalancerHealthMonitorTypeEnum_udpconnect  = "udp-connect"
)

type Cloudloadbalancing_LoadBalancerOperatingStatusEnum string
const (
	Cloudloadbalancing_LoadBalancerOperatingStatusEnum_degraded Cloudloadbalancing_LoadBalancerOperatingStatusEnum = "degraded"
	Cloudloadbalancing_LoadBalancerOperatingStatusEnum_draining  = "draining"
	Cloudloadbalancing_LoadBalancerOperatingStatusEnum_error  = "error"
	Cloudloadbalancing_LoadBalancerOperatingStatusEnum_noMonitor  = "noMonitor"
	Cloudloadbalancing_LoadBalancerOperatingStatusEnum_offline  = "offline"
	Cloudloadbalancing_LoadBalancerOperatingStatusEnum_online  = "online"
)

type Cloudloadbalancing_LoadBalancerProvisioningStatusEnum string
const (
	Cloudloadbalancing_LoadBalancerProvisioningStatusEnum_active Cloudloadbalancing_LoadBalancerProvisioningStatusEnum = "active"
	Cloudloadbalancing_LoadBalancerProvisioningStatusEnum_creating  = "creating"
	Cloudloadbalancing_LoadBalancerProvisioningStatusEnum_deleted  = "deleted"
	Cloudloadbalancing_LoadBalancerProvisioningStatusEnum_deleting  = "deleting"
	Cloudloadbalancing_LoadBalancerProvisioningStatusEnum_error  = "error"
	Cloudloadbalancing_LoadBalancerProvisioningStatusEnum_updating  = "updating"
)

type Cloudloadbalancing_LoadbalancerCreate struct {
	FlavorId           string                                                  `json:"flavorId"`
	Listeners          []Cloudloadbalancingloadbalancer_ListenerCreate         `json:"listeners"`
	Name               string                                                  `json:"name"`
	NetworkInformation Cloudloadbalancingloadbalancer_NetworkInformationCreate `json:"networkInformation"`
}

type Cloudloadbalancing_Pool struct {
	Algorithm          Cloudloadbalancing_PoolAlgorithmEnum                  `json:"algorithm"`
	Id                 string                                                `json:"id"`
	ListenerId         string                                                `json:"listenerId"`
	LoadbalancerId     string                                                `json:"loadbalancerId"`
	Name               string                                                `json:"name"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	Protocol           Cloudloadbalancing_PoolProtocolEnum                   `json:"protocol"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
	SessionPersistence Cloudloadbalancing_PoolSessionPersistence             `json:"sessionPersistence"`
}

type Cloudloadbalancing_PoolAlgorithmEnum string
const (
	Cloudloadbalancing_PoolAlgorithmEnum_leastConnections Cloudloadbalancing_PoolAlgorithmEnum = "leastConnections"
	Cloudloadbalancing_PoolAlgorithmEnum_roundRobin  = "roundRobin"
	Cloudloadbalancing_PoolAlgorithmEnum_sourceIP  = "sourceIP"
)

type Cloudloadbalancing_PoolCreate struct {
	Algorithm          Cloudloadbalancing_PoolAlgorithmEnum      `json:"algorithm"`
	ListenerId         string                                    `json:"listenerId"`
	LoadbalancerId     string                                    `json:"loadbalancerId"`
	Name               string                                    `json:"name"`
	Protocol           Cloudloadbalancing_PoolProtocolEnum       `json:"protocol"`
	SessionPersistence Cloudloadbalancing_PoolSessionPersistence `json:"sessionPersistence"`
}

type Cloudloadbalancing_PoolProtocolEnum string
const (
	Cloudloadbalancing_PoolProtocolEnum_http Cloudloadbalancing_PoolProtocolEnum = "http"
	Cloudloadbalancing_PoolProtocolEnum_https  = "https"
	Cloudloadbalancing_PoolProtocolEnum_proxy  = "proxy"
	Cloudloadbalancing_PoolProtocolEnum_sctp  = "sctp"
	Cloudloadbalancing_PoolProtocolEnum_tcp  = "tcp"
	Cloudloadbalancing_PoolProtocolEnum_udp  = "udp"
)

type Cloudloadbalancing_PoolSessionPersistence struct {
	CookieName string                                            `json:"cookieName"`
	Type       Cloudloadbalancing_PoolSessionPersistenceTypeEnum `json:"type"`
}

type Cloudloadbalancing_PoolSessionPersistenceTypeEnum string
const (
	Cloudloadbalancing_PoolSessionPersistenceTypeEnum_appCookie Cloudloadbalancing_PoolSessionPersistenceTypeEnum = "appCookie"
	Cloudloadbalancing_PoolSessionPersistenceTypeEnum_disabled  = "disabled"
	Cloudloadbalancing_PoolSessionPersistenceTypeEnum_httpCookie  = "httpCookie"
	Cloudloadbalancing_PoolSessionPersistenceTypeEnum_sourceIP  = "sourceIP"
)

type Cloudloadbalancing_PoolUpdate struct {
	Algorithm          Cloudloadbalancing_PoolAlgorithmEnum      `json:"algorithm"`
	Name               string                                    `json:"name"`
	SessionPersistence Cloudloadbalancing_PoolSessionPersistence `json:"sessionPersistence"`
}

type Cloudloadbalancing_UpdateHealthMonitor struct {
	HttpConfiguration Cloudloadbalancing_HealthMonitorHTTPConfiguration `json:"httpConfiguration"`
	MaxRetries        int64                                             `json:"maxRetries"`
	MaxRetriesDown    int64                                             `json:"maxRetriesDown"`
	Name              string                                            `json:"name"`
	Periodicity       int64                                             `json:"periodicity"`
	Timeout           int64                                             `json:"timeout"`
}

type Cloudloadbalancingloadbalancer_ListenerCreate struct {
	Name     string                                      `json:"name"`
	Pools    []Cloudloadbalancingloadbalancer_PoolCreate `json:"pools"`
	Port     int64                                       `json:"port"`
	Protocol Cloudloadbalancing_ListenerProtocolEnum     `json:"protocol"`
	SecretId string                                      `json:"secretId"`
}

type Cloudloadbalancingloadbalancer_MemberCreate struct {
	Address      string `json:"address"`
	Name         string `json:"name"`
	ProtocolPort int64  `json:"protocolPort"`
	Weight       int64  `json:"weight"`
}

type Cloudloadbalancingloadbalancer_NetworkInformationCreate struct {
	Gateway   Cloudnetwork_CreateGatewaySummary `json:"gateway"`
	NetworkId string                            `json:"networkId"`
	SubnetId  string                            `json:"subnetId"`
}

type Cloudloadbalancingloadbalancer_PoolCreate struct {
	Algorithm          Cloudloadbalancing_PoolAlgorithmEnum          `json:"algorithm"`
	Default            bool                                          `json:"default"`
	Members            []Cloudloadbalancingloadbalancer_MemberCreate `json:"members"`
	Name               string                                        `json:"name"`
	Protocol           Cloudloadbalancing_PoolProtocolEnum           `json:"protocol"`
	SessionPersistence Cloudloadbalancing_PoolSessionPersistence     `json:"sessionPersistence"`
}

type Cloudloadbalancingpool_Member struct {
	Address            string                                                `json:"address"`
	Id                 string                                                `json:"id"`
	Name               string                                                `json:"name"`
	OperatingStatus    Cloudloadbalancing_LoadBalancerOperatingStatusEnum    `json:"operatingStatus"`
	ProtocolPort       int64                                                 `json:"protocolPort"`
	ProvisioningStatus Cloudloadbalancing_LoadBalancerProvisioningStatusEnum `json:"provisioningStatus"`
	Weight             int64                                                 `json:"weight"`
}

type Cloudloadbalancingpool_MemberUpdate struct {
	Name   string `json:"name"`
	Weight int64  `json:"weight"`
}

type Cloudmigration_Migration struct {
	Date         time.Time                       `json:"date"`
	MigrationId  string                          `json:"migrationId"`
	ResourceId   string                          `json:"resourceId"`
	ResourceType Cloudmigration_ResourceTypeEnum `json:"resourceType"`
}

type Cloudmigration_ResourceTypeEnum string
const (
	Cloudmigration_ResourceTypeEnum_instance Cloudmigration_ResourceTypeEnum = "instance"
)

type Cloudnetwork_CreateGateway struct {
	Model   Cloudnetwork_GatewayModelEnum               `json:"model"`
	Name    string                                      `json:"name"`
	Network Cloudnetwork_CreateNetworkForGatewaySummary `json:"network"`
}

type Cloudnetwork_CreateGatewaySummary struct {
	Model Cloudnetwork_GatewayModelEnum `json:"model"`
	Name  string                        `json:"name"`
}

type Cloudnetwork_CreateNetwork struct {
	Gateway Cloudnetwork_CreateGatewaySummary `json:"gateway"`
	Name    string                            `json:"name"`
	Subnet  Cloudnetwork_CreateSubnetSummary  `json:"subnet"`
	VlanId  int64                             `json:"vlanId"`
}

type Cloudnetwork_CreateNetworkForGatewaySummary struct {
	Name   string                                     `json:"name"`
	Subnet Cloudnetwork_CreateSubnetForGatewaySummary `json:"subnet"`
	VlanId int64                                      `json:"vlanId"`
}

type Cloudnetwork_CreateSubnetForGatewaySummary struct {
	Cidr       string `json:"cidr"`
	EnableDhcp bool   `json:"enableDhcp"`
	IpVersion  int64  `json:"ipVersion"`
}

type Cloudnetwork_CreateSubnetSummary struct {
	Cidr            string `json:"cidr"`
	EnableDhcp      bool   `json:"enableDhcp"`
	EnableGatewayIp bool   `json:"enableGatewayIp"`
	IpVersion       int64  `json:"ipVersion"`
}

type Cloudnetwork_Gateway struct {
	ExternalInformation Cloudnetworkgateway_ExternalInformation `json:"externalInformation"`
	Id                  string                                  `json:"id"`
	Interfaces          []Cloudnetworkgateway_Interface         `json:"interfaces"`
	Model               Cloudnetwork_GatewayModelEnum           `json:"model"`
	Name                string                                  `json:"name"`
	Region              string                                  `json:"region"`
	Status              Cloudnetwork_GatewayStatusEnum          `json:"status"`
}

type Cloudnetwork_GatewayModelEnum string
const (
	Cloudnetwork_GatewayModelEnum_l Cloudnetwork_GatewayModelEnum = "l"
	Cloudnetwork_GatewayModelEnum_m  = "m"
	Cloudnetwork_GatewayModelEnum_s  = "s"
)

type Cloudnetwork_GatewayStatusEnum string
const (
	Cloudnetwork_GatewayStatusEnum_active Cloudnetwork_GatewayStatusEnum = "active"
	Cloudnetwork_GatewayStatusEnum_building  = "building"
	Cloudnetwork_GatewayStatusEnum_down  = "down"
	Cloudnetwork_GatewayStatusEnum_error  = "error"
)

type Cloudnetwork_IPPool struct {
	Dhcp    bool   `json:"dhcp"`
	End     string `json:"end"`
	Network string `json:"network"`
	Region  string `json:"region"`
	Start   string `json:"start"`
}

type Cloudnetwork_Network struct {
	Id      string                         `json:"id"`
	Name    string                         `json:"name"`
	Regions []Cloudnetwork_NetworkRegion   `json:"regions"`
	Status  Cloudnetwork_NetworkStatusEnum `json:"status"`
	Type    Cloudnetwork_NetworkTypeEnum   `json:"type"`
	VlanId  int64                          `json:"vlanId"`
}

type Cloudnetwork_NetworkRegion struct {
	OpenstackId string                               `json:"openstackId"`
	Region      string                               `json:"region"`
	Status      Cloudnetwork_NetworkRegionStatusEnum `json:"status"`
}

type Cloudnetwork_NetworkRegionStatusEnum string
const (
	Cloudnetwork_NetworkRegionStatusEnum_ACTIVE Cloudnetwork_NetworkRegionStatusEnum = "ACTIVE"
	Cloudnetwork_NetworkRegionStatusEnum_BUILDING  = "BUILDING"
)

type Cloudnetwork_NetworkStatusEnum string
const (
	Cloudnetwork_NetworkStatusEnum_ACTIVE Cloudnetwork_NetworkStatusEnum = "ACTIVE"
	Cloudnetwork_NetworkStatusEnum_BUILDING  = "BUILDING"
	Cloudnetwork_NetworkStatusEnum_DELETING  = "DELETING"
)

type Cloudnetwork_NetworkTypeEnum string
const (
	Cloudnetwork_NetworkTypeEnum_private Cloudnetwork_NetworkTypeEnum = "private"
	Cloudnetwork_NetworkTypeEnum_public  = "public"
)

type Cloudnetwork_NetworkVisibilityEnum string
const (
	Cloudnetwork_NetworkVisibilityEnum_private Cloudnetwork_NetworkVisibilityEnum = "private"
	Cloudnetwork_NetworkVisibilityEnum_public  = "public"
)

type Cloudnetwork_Subnet struct {
	Cidr        string                `json:"cidr"`
	DhcpEnabled bool                  `json:"dhcpEnabled"`
	GatewayIp   string                `json:"gatewayIp"`
	Id          string                `json:"id"`
	IpPools     []Cloudnetwork_IPPool `json:"ipPools"`
}

type Cloudnetwork_UpdateGateway struct {
	Model Cloudnetwork_GatewayModelEnum `json:"model"`
	Name  string                        `json:"name"`
}

type Cloudnetworkgateway_CreateInterface struct {
	SubnetId string `json:"subnetId"`
}

type Cloudnetworkgateway_ExternalInformation struct {
	Ips       []Cloudnetworkgateway_IpSubnet `json:"ips"`
	NetworkId string                         `json:"networkId"`
}

type Cloudnetworkgateway_Interface struct {
	Id        string `json:"id"`
	Ip        string `json:"ip"`
	NetworkId string `json:"networkId"`
	SubnetId  string `json:"subnetId"`
}

type Cloudnetworkgateway_IpSubnet struct {
	Ip       string `json:"ip"`
	SubnetId string `json:"subnetId"`
}

type Cloudorder_Order struct {
	Date        time.Time             `json:"date"`
	OrderId     int64                 `json:"orderId"`
	PlanCode    string                `json:"planCode"`
	ServiceName string                `json:"serviceName"`
	Status      Cloudorder_StatusEnum `json:"status"`
}

type Cloudorder_StatusEnum string
const (
	Cloudorder_StatusEnum_delivered Cloudorder_StatusEnum = "delivered"
	Cloudorder_StatusEnum_delivering  = "delivering"
	Cloudorder_StatusEnum_unknown  = "unknown"
	Cloudorder_StatusEnum_unpaid  = "unpaid"
)

type Cloudorderrule_Availability struct {
	Plans    []Cloudorderrule_AvailabilityPlan    `json:"plans"`
	Products []Cloudorderrule_AvailabilityProduct `json:"products"`
}

type Cloudorderrule_AvailabilityPlan struct {
	Code    string   `json:"code"`
	Regions []string `json:"regions"`
}

type Cloudorderrule_AvailabilityProduct struct {
	Name    string   `json:"name"`
	Regions []string `json:"regions"`
}

type Cloudorderrule_InstanceCategories struct {
	Categories      []Cloudorderrule_InstanceCategory `json:"categories"`
	DefaultCategory string                            `json:"defaultCategory"`
}

type Cloudorderrule_InstanceCategory struct {
	Category Cloudorderrule_InstanceCategoryTypeEnum `json:"category"`
	IsNew    bool                                    `json:"isNew"`
	Kinds    []string                                `json:"kinds"`
	Title    string                                  `json:"title"`
}

type Cloudorderrule_InstanceCategoryTypeEnum string
const (
	Cloudorderrule_InstanceCategoryTypeEnum_accelerated Cloudorderrule_InstanceCategoryTypeEnum = "accelerated"
	Cloudorderrule_InstanceCategoryTypeEnum_balanced  = "balanced"
	Cloudorderrule_InstanceCategoryTypeEnum_baremetal  = "baremetal"
	Cloudorderrule_InstanceCategoryTypeEnum_discovery  = "discovery"
	Cloudorderrule_InstanceCategoryTypeEnum_iops  = "iops"
	Cloudorderrule_InstanceCategoryTypeEnum_ram  = "ram"
	Cloudorderrule_InstanceCategoryTypeEnum_vps  = "vps"
)

type Cloudproject_AllocationPool struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

type Cloudproject_ApplicationLoadBalancer struct {
	Address         Cloudprojectloadbalancer_Address              `json:"address"`
	Configuration   Cloudprojectloadbalancer_ConfigurationVersion `json:"configuration"`
	CreatedAt       time.Time                                     `json:"createdAt"`
	Description     string                                        `json:"description"`
	EgressAddress   Cloudprojectloadbalancer_Addresses            `json:"egressAddress"`
	Id              string                                        `json:"id"`
	Name            string                                        `json:"name"`
	OpenstackRegion string                                        `json:"openstackRegion"`
	Region          string                                        `json:"region"`
	Size            Cloudprojectloadbalancer_SizeEnum             `json:"size"`
	Status          Cloudprojectloadbalancer_StatusEnum           `json:"status"`
}

type Cloudproject_ApplicationLoadBalancerCreation struct {
	Description     string                                                `json:"description"`
	Id              string                                                `json:"id"`
	Name            string                                                `json:"name"`
	Networking      Cloudprojectloadbalancernetworking_NetworkingCreation `json:"networking"`
	OpenstackRegion string                                                `json:"openstackRegion"`
	Region          string                                                `json:"region"`
	Size            Cloudprojectloadbalancer_SizeEnum                     `json:"size"`
}

type Cloudproject_BandwidthStorageUsage struct {
	DownloadedBytes int64       `json:"downloadedBytes"`
	Region          string      `json:"region"`
	Total           Order_Price `json:"total"`
}

type Cloudproject_Bill struct {
	BillId string                    `json:"billId"`
	Type   Cloudproject_BillTypeEnum `json:"type"`
}

type Cloudproject_BillTypeEnum string
const (
	Cloudproject_BillTypeEnum_creditPurchased Cloudproject_BillTypeEnum = "creditPurchased"
	Cloudproject_BillTypeEnum_monthlyConsumption  = "monthlyConsumption"
	Cloudproject_BillTypeEnum_monthlyInstanceActivation  = "monthlyInstanceActivation"
)

type Cloudproject_Certificate struct {
	ExpireAt               time.Time                                       `json:"expireAt"`
	Fingerprint            string                                          `json:"fingerprint"`
	Id                     string                                          `json:"id"`
	Issuer                 string                                          `json:"issuer"`
	Kind                   Cloudproject_CertificateKindEnum                `json:"kind"`
	Name                   string                                          `json:"name"`
	SerialNumber           string                                          `json:"serialNumber"`
	ServerAlternativeNames []Cloudprojectcertificate_ServerAlternativeName `json:"serverAlternativeNames"`
	Status                 Cloudproject_CertificateStatusEnum              `json:"status"`
	Subject                string                                          `json:"subject"`
	ValidAt                time.Time                                       `json:"validAt"`
	Version                int64                                           `json:"version"`
}

type Cloudproject_CertificateAdd struct {
	Import Cloudprojectcertificate_Import `json:"import"`
	Name   string                         `json:"name"`
}

type Cloudproject_CertificateKindEnum string
const (
	Cloudproject_CertificateKindEnum_IMPORTED Cloudproject_CertificateKindEnum = "IMPORTED"
)

type Cloudproject_CertificateStatusEnum string
const (
	Cloudproject_CertificateStatusEnum_EXPIRED Cloudproject_CertificateStatusEnum = "EXPIRED"
	Cloudproject_CertificateStatusEnum_NOT_YET_VALID  = "NOT_YET_VALID"
	Cloudproject_CertificateStatusEnum_OK  = "OK"
	Cloudproject_CertificateStatusEnum_REVOKED  = "REVOKED"
)

type Cloudproject_CurrentUsage struct {
	Instances       Cloudproject_InstancesUsage `json:"instances"`
	Snapshots       Cloudproject_SnapshotsUsage `json:"snapshots"`
	Storage         Cloudproject_StorageUsage   `json:"storage"`
	Total           Order_Price                 `json:"total"`
	VolumeSnapshots Cloudproject_SnapshotsUsage `json:"volumeSnapshots"`
	Volumes         Cloudproject_VolumesUsage   `json:"volumes"`
}

type Cloudproject_EligibilityAction string
const (
	Cloudproject_EligibilityAction_addPaymentMethod Cloudproject_EligibilityAction = "addPaymentMethod"
	Cloudproject_EligibilityAction_askIncreaseProjectsQuota  = "askIncreaseProjectsQuota"
	Cloudproject_EligibilityAction_challengePaymentMethod  = "challengePaymentMethod"
	Cloudproject_EligibilityAction_verifyPaypal  = "verifyPaypal"
)

type Cloudproject_EligibilityInfo struct {
	ActionsRequired          []Cloudproject_EligibilityAction       `json:"actionsRequired"`
	MinimumCredit            Order_Price                            `json:"minimumCredit"`
	PaymentMethodsAuthorized []Cloudproject_PaymentMethodAuthorized `json:"paymentMethodsAuthorized"`
	Voucher                  Cloudproject_NewProjectInfoVoucher     `json:"voucher"`
}

type Cloudproject_FloatingIp struct {
	AssociatedEntity CloudprojectfloatingIp_AssociatedEntity `json:"associatedEntity"`
	Id               string                                  `json:"id"`
	Ip               string                                  `json:"ip"`
	NetworkId        string                                  `json:"networkId"`
	Region           string                                  `json:"region"`
	Status           CloudprojectfloatingIp_StatusEnum       `json:"status"`
}

type Cloudproject_InstanceMonthlyBilling struct {
	ActivatedOn time.Time   `json:"activatedOn"`
	Cost        Order_Price `json:"cost"`
}

type Cloudproject_InstanceUsageDetail struct {
	Hourly         Order_Price                         `json:"hourly"`
	InstanceId     string                              `json:"instanceId"`
	Monthly        Cloudproject_InstanceMonthlyBilling `json:"monthly"`
	MonthlyBilling bool                                `json:"monthlyBilling"`
	Reference      string                              `json:"reference"`
}

type Cloudproject_InstancesUsage struct {
	Detail []Cloudproject_InstanceUsageDetail `json:"detail"`
	Total  Order_Price                        `json:"total"`
}

type Cloudproject_LoadBalancer struct {
	Address         Cloudprojectloadbalancer_Address              `json:"address"`
	Configuration   Cloudprojectloadbalancer_ConfigurationVersion `json:"configuration"`
	CreatedAt       time.Time                                     `json:"createdAt"`
	Description     string                                        `json:"description"`
	EgressAddress   Cloudprojectloadbalancer_Addresses            `json:"egressAddress"`
	Id              string                                        `json:"id"`
	Name            string                                        `json:"name"`
	Networking      Cloudprojectloadbalancernetworking_Networking `json:"networking"`
	OpenstackRegion string                                        `json:"openstackRegion"`
	Region          string                                        `json:"region"`
	Size            Cloudprojectloadbalancer_SizeEnum             `json:"size"`
	Status          Cloudprojectloadbalancer_StatusEnum           `json:"status"`
}

type Cloudproject_LoadBalancerCreation struct {
	Description     string                                                `json:"description"`
	Id              string                                                `json:"id"`
	Name            string                                                `json:"name"`
	Networking      Cloudprojectloadbalancernetworking_NetworkingCreation `json:"networking"`
	OpenstackRegion string                                                `json:"openstackRegion"`
	Region          string                                                `json:"region"`
	Size            Cloudprojectloadbalancer_SizeEnum                     `json:"size"`
}

type Cloudproject_Network struct {
	Id         string                             `json:"id"`
	Name       string                             `json:"name"`
	Region     string                             `json:"region"`
	Visibility Cloudnetwork_NetworkVisibilityEnum `json:"visibility"`
	VlanId     int64                              `json:"vlanId"`
}

type Cloudproject_NetworkLoadBalancer struct {
	Address         Cloudprojectloadbalancer_Address              `json:"address"`
	Configuration   Cloudprojectloadbalancer_ConfigurationVersion `json:"configuration"`
	CreatedAt       time.Time                                     `json:"createdAt"`
	Description     string                                        `json:"description"`
	EgressAddress   Cloudprojectloadbalancer_Addresses            `json:"egressAddress"`
	Id              string                                        `json:"id"`
	Name            string                                        `json:"name"`
	OpenstackRegion string                                        `json:"openstackRegion"`
	Region          string                                        `json:"region"`
	Size            Cloudprojectloadbalancer_SizeEnum             `json:"size"`
	Status          Cloudprojectloadbalancer_StatusEnum           `json:"status"`
}

type Cloudproject_NetworkLoadBalancerCreation struct {
	Description     string                                                `json:"description"`
	Id              string                                                `json:"id"`
	Name            string                                                `json:"name"`
	Networking      Cloudprojectloadbalancernetworking_NetworkingCreation `json:"networking"`
	OpenstackRegion string                                                `json:"openstackRegion"`
	Region          string                                                `json:"region"`
	Size            Cloudprojectloadbalancer_SizeEnum                     `json:"size"`
}

type Cloudproject_NewProject struct {
	Agreements  []int64                           `json:"agreements"`
	Credit      Cloudproject_NewProjectCredit     `json:"credit"`
	Description string                            `json:"description"`
	OrderId     int64                             `json:"orderId"`
	Project     string                            `json:"project"`
	Status      Cloudproject_NewProjectStatusEnum `json:"status"`
}

type Cloudproject_NewProjectCredit struct {
	Description  string                      `json:"description"`
	Id           int64                       `json:"id"`
	Products     []string                    `json:"products"`
	Total_credit Order_Price                 `json:"total_credit"`
	Validity     Cloudcommon_VoucherValidity `json:"validity"`
}

type Cloudproject_NewProjectInfo struct {
	Agreements []int64                            `json:"agreements"`
	Error      Cloudproject_NewProjectInfoError   `json:"error"`
	Order      Order_Price                        `json:"order"`
	Voucher    Cloudproject_NewProjectInfoVoucher `json:"voucher"`
}

type Cloudproject_NewProjectInfoError struct {
	Code    Cloudproject_NewProjectInfoErrorCodeEnum `json:"code"`
	Message string                                   `json:"message"`
}

type Cloudproject_NewProjectInfoErrorCodeEnum string
const (
	Cloudproject_NewProjectInfoErrorCodeEnum_accountNotEligible Cloudproject_NewProjectInfoErrorCodeEnum = "accountNotEligible"
	Cloudproject_NewProjectInfoErrorCodeEnum_challengePaymentMethodRequested  = "challengePaymentMethodRequested"
	Cloudproject_NewProjectInfoErrorCodeEnum_invalidPaymentMean  = "invalidPaymentMean"
	Cloudproject_NewProjectInfoErrorCodeEnum_maxProjectsLimitReached  = "maxProjectsLimitReached"
	Cloudproject_NewProjectInfoErrorCodeEnum_paypalAccountNotVerified  = "paypalAccountNotVerified"
	Cloudproject_NewProjectInfoErrorCodeEnum_unpaidDebts  = "unpaidDebts"
)

type Cloudproject_NewProjectInfoVoucher struct {
	Credit                Order_Price `json:"credit"`
	PaymentMethodRequired bool        `json:"paymentMethodRequired"`
}

type Cloudproject_NewProjectStatusEnum string
const (
	Cloudproject_NewProjectStatusEnum_creating Cloudproject_NewProjectStatusEnum = "creating"
	Cloudproject_NewProjectStatusEnum_ok  = "ok"
	Cloudproject_NewProjectStatusEnum_validationPending  = "validationPending"
	Cloudproject_NewProjectStatusEnum_waitingAgreementsValidation  = "waitingAgreementsValidation"
)

type Cloudproject_PaymentMethodAuthorized string
const (
	Cloudproject_PaymentMethodAuthorized_bankAccount Cloudproject_PaymentMethodAuthorized = "bankAccount"
	Cloudproject_PaymentMethodAuthorized_credit  = "credit"
	Cloudproject_PaymentMethodAuthorized_creditCard  = "creditCard"
	Cloudproject_PaymentMethodAuthorized_paypal  = "paypal"
	Cloudproject_PaymentMethodAuthorized_sepaDirectDebit  = "sepaDirectDebit"
)

type Cloudproject_ProductAgreements struct {
	AgreementsToValidate []int64 `json:"agreementsToValidate"`
	AgreementsValidated  []int64 `json:"agreementsValidated"`
}

type Cloudproject_ProductNameEnum string
const (
	Cloudproject_ProductNameEnum_registry Cloudproject_ProductNameEnum = "registry"
)

type Cloudproject_ProjectStatus string
const (
	Cloudproject_ProjectStatus_creating Cloudproject_ProjectStatus = "creating"
	Cloudproject_ProjectStatus_deleted  = "deleted"
	Cloudproject_ProjectStatus_deleting  = "deleting"
	Cloudproject_ProjectStatus_ok  = "ok"
	Cloudproject_ProjectStatus_suspended  = "suspended"
)

type Cloudproject_ProjectStatusEnum string
const (
	Cloudproject_ProjectStatusEnum_creating Cloudproject_ProjectStatusEnum = "creating"
	Cloudproject_ProjectStatusEnum_deleted  = "deleted"
	Cloudproject_ProjectStatusEnum_deleting  = "deleting"
	Cloudproject_ProjectStatusEnum_ok  = "ok"
	Cloudproject_ProjectStatusEnum_suspended  = "suspended"
)

type Cloudproject_ProjectUsage struct {
	Current Cloudproject_CurrentUsage `json:"current"`
}

type Cloudproject_SnapshotUsageDetail struct {
	Price      Order_Price               `json:"price"`
	Region     string                    `json:"region"`
	StoredSize types.UnitAndValueFloat64 `json:"storedSize"`
}

type Cloudproject_SnapshotsUsage struct {
	Detail []Cloudproject_SnapshotUsageDetail `json:"detail"`
	Total  Order_Price                        `json:"total"`
}

type Cloudproject_StorageUsage struct {
	Bandwidth []Cloudproject_BandwidthStorageUsage `json:"bandwidth"`
	Total     Order_Price                          `json:"total"`
	Volume    []Cloudproject_StorageVolumeUsage    `json:"volume"`
}

type Cloudproject_StorageVolumeUsage struct {
	Region      string      `json:"region"`
	StoredBytes int64       `json:"storedBytes"`
	Total       Order_Price `json:"total"`
}

type Cloudproject_Subnet struct {
	AllocationPools []Cloudproject_AllocationPool `json:"allocationPools"`
	Cidr            string                        `json:"cidr"`
	DhcpEnabled     bool                          `json:"dhcpEnabled"`
	GatewayIp       string                        `json:"gatewayIp"`
	Id              string                        `json:"id"`
	IpVersion       int64                         `json:"ipVersion"`
	Name            string                        `json:"name"`
}

type Cloudproject_VolumeType string
const (
	Cloudproject_VolumeType_classic Cloudproject_VolumeType = "classic"
	Cloudproject_VolumeType_highspeed  = "high-speed"
)

type Cloudproject_VolumeUsageDetail struct {
	Price          Order_Price             `json:"price"`
	VolumeCapacity types.UnitAndValueInt64 `json:"volumeCapacity"`
	VolumeId       string                  `json:"volumeId"`
	VolumeType     Cloudproject_VolumeType `json:"volumeType"`
}

type Cloudproject_VolumesUsage struct {
	Detail []Cloudproject_VolumeUsageDetail `json:"detail"`
	Total  Order_Price                      `json:"total"`
}

type Cloudprojectai_AuthorizationStatus struct {
	Authorized bool `json:"authorized"`
}

type Cloudprojectai_Command struct {
	Command string `json:"command"`
}

type Cloudprojectai_Info struct {
	Code    Cloudprojectai_InfoCodeEnum `json:"code"`
	Message string                      `json:"message"`
}

type Cloudprojectai_InfoCodeEnum string
const (
	Cloudprojectai_InfoCodeEnum_APP_CREATE_ERROR Cloudprojectai_InfoCodeEnum = "APP_CREATE_ERROR"
	Cloudprojectai_InfoCodeEnum_APP_ERROR  = "APP_ERROR"
	Cloudprojectai_InfoCodeEnum_APP_FAILED  = "APP_FAILED"
	Cloudprojectai_InfoCodeEnum_APP_INITIALIZING  = "APP_INITIALIZING"
	Cloudprojectai_InfoCodeEnum_APP_INTERRUPTED_BY_PLATFORM  = "APP_INTERRUPTED_BY_PLATFORM"
	Cloudprojectai_InfoCodeEnum_APP_QUEUED  = "APP_QUEUED"
	Cloudprojectai_InfoCodeEnum_APP_RUNNING  = "APP_RUNNING"
	Cloudprojectai_InfoCodeEnum_APP_SCALING  = "APP_SCALING"
	Cloudprojectai_InfoCodeEnum_APP_STOPPED  = "APP_STOPPED"
	Cloudprojectai_InfoCodeEnum_APP_STOPPING  = "APP_STOPPING"
	Cloudprojectai_InfoCodeEnum_COMPATIBILITY  = "COMPATIBILITY"
	Cloudprojectai_InfoCodeEnum_DATASYNC_AUTHENTICATE_FAILED  = "DATASYNC_AUTHENTICATE_FAILED"
	Cloudprojectai_InfoCodeEnum_DATASYNC_DONE  = "DATASYNC_DONE"
	Cloudprojectai_InfoCodeEnum_DATASYNC_ERROR  = "DATASYNC_ERROR"
	Cloudprojectai_InfoCodeEnum_DATASYNC_FAILED  = "DATASYNC_FAILED"
	Cloudprojectai_InfoCodeEnum_DATASYNC_INTERRUPTED  = "DATASYNC_INTERRUPTED"
	Cloudprojectai_InfoCodeEnum_DATASYNC_INVALID_CONTAINER  = "DATASYNC_INVALID_CONTAINER"
	Cloudprojectai_InfoCodeEnum_DATASYNC_QUEUED  = "DATASYNC_QUEUED"
	Cloudprojectai_InfoCodeEnum_DATASYNC_RETRY_ERROR  = "DATASYNC_RETRY_ERROR"
	Cloudprojectai_InfoCodeEnum_DATASYNC_RUNNING  = "DATASYNC_RUNNING"
	Cloudprojectai_InfoCodeEnum_JOB_CREATE_CONTAINER_CONFIG_ERROR  = "JOB_CREATE_CONTAINER_CONFIG_ERROR"
	Cloudprojectai_InfoCodeEnum_JOB_CREATE_CONTAINER_ERROR  = "JOB_CREATE_CONTAINER_ERROR"
	Cloudprojectai_InfoCodeEnum_JOB_DONE  = "JOB_DONE"
	Cloudprojectai_InfoCodeEnum_JOB_ERROR  = "JOB_ERROR"
	Cloudprojectai_InfoCodeEnum_JOB_EVICTED  = "JOB_EVICTED"
	Cloudprojectai_InfoCodeEnum_JOB_FAILED  = "JOB_FAILED"
	Cloudprojectai_InfoCodeEnum_JOB_FAILED_WITH_MESSAGE  = "JOB_FAILED_WITH_MESSAGE"
	Cloudprojectai_InfoCodeEnum_JOB_FINALIZING  = "JOB_FINALIZING"
	Cloudprojectai_InfoCodeEnum_JOB_IMAGE_INSPECT_ERROR  = "JOB_IMAGE_INSPECT_ERROR"
	Cloudprojectai_InfoCodeEnum_JOB_IMAGE_PULL  = "JOB_IMAGE_PULL"
	Cloudprojectai_InfoCodeEnum_JOB_IMAGE_PULL_BACKOFF  = "JOB_IMAGE_PULL_BACKOFF"
	Cloudprojectai_InfoCodeEnum_JOB_INITIALIZING  = "JOB_INITIALIZING"
	Cloudprojectai_InfoCodeEnum_JOB_INTERRUPTED  = "JOB_INTERRUPTED"
	Cloudprojectai_InfoCodeEnum_JOB_INTERRUPTED_BY_PLATFORM  = "JOB_INTERRUPTED_BY_PLATFORM"
	Cloudprojectai_InfoCodeEnum_JOB_INTERRUPTING  = "JOB_INTERRUPTING"
	Cloudprojectai_InfoCodeEnum_JOB_INVALID_IMAGE_NAME  = "JOB_INVALID_IMAGE_NAME"
	Cloudprojectai_InfoCodeEnum_JOB_PENDING  = "JOB_PENDING"
	Cloudprojectai_InfoCodeEnum_JOB_QUEUED  = "JOB_QUEUED"
	Cloudprojectai_InfoCodeEnum_JOB_REGISTRY_UNAVAILABLE  = "JOB_REGISTRY_UNAVAILABLE"
	Cloudprojectai_InfoCodeEnum_JOB_RUNNING  = "JOB_RUNNING"
	Cloudprojectai_InfoCodeEnum_JOB_TIMEOUT  = "JOB_TIMEOUT"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_FAILED  = "NOTEBOOK_FAILED"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_FAILED_WITH_MESSAGE  = "NOTEBOOK_FAILED_WITH_MESSAGE"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_FINALIZING  = "NOTEBOOK_FINALIZING"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_INITIALIZING  = "NOTEBOOK_INITIALIZING"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_PENDING  = "NOTEBOOK_PENDING"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_RUNNING  = "NOTEBOOK_RUNNING"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_STARTING  = "NOTEBOOK_STARTING"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_STOPPED  = "NOTEBOOK_STOPPED"
	Cloudprojectai_InfoCodeEnum_NOTEBOOK_STOPPING  = "NOTEBOOK_STOPPING"
)

type Cloudprojectai_Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Cloudprojectai_LogLine struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Cloudprojectai_Logs struct {
	LastActivity time.Time                `json:"lastActivity"`
	Logs         []Cloudprojectai_LogLine `json:"logs"`
}

type Cloudprojectai_OrderEnum string
const (
	Cloudprojectai_OrderEnum_asc Cloudprojectai_OrderEnum = "asc"
	Cloudprojectai_OrderEnum_desc  = "desc"
)

type Cloudprojectai_Resources struct {
	Cpu              int64  `json:"cpu"`
	EphemeralStorage int64  `json:"ephemeralStorage"`
	Flavor           string `json:"flavor"`
	Gpu              int64  `json:"gpu"`
	GpuBrand         string `json:"gpuBrand"`
	GpuMemory        int64  `json:"gpuMemory"`
	GpuModel         string `json:"gpuModel"`
	Memory           int64  `json:"memory"`
	PrivateNetwork   int64  `json:"privateNetwork"`
	PublicNetwork    int64  `json:"publicNetwork"`
}

type Cloudprojectai_ResourcesInput struct {
	Cpu              int64  `json:"cpu"`
	EphemeralStorage int64  `json:"ephemeralStorage"`
	Flavor           string `json:"flavor"`
	Gpu              int64  `json:"gpu"`
	GpuBrand         string `json:"gpuBrand"`
	GpuMemory        int64  `json:"gpuMemory"`
	GpuModel         string `json:"gpuModel"`
	Memory           int64  `json:"memory"`
	PrivateNetwork   int64  `json:"privateNetwork"`
	PublicNetwork    int64  `json:"publicNetwork"`
}

type Cloudprojectai_ShutdownStrategyEnum string
const (
	Cloudprojectai_ShutdownStrategyEnum_Stop Cloudprojectai_ShutdownStrategyEnum = "Stop"
)

type Cloudprojectai_TokenRoleEnum string
const (
	Cloudprojectai_TokenRoleEnum_ai_training_operator Cloudprojectai_TokenRoleEnum = "ai_training_operator"
	Cloudprojectai_TokenRoleEnum_ai_training_read  = "ai_training_read"
)

type Cloudprojectai_VolumePermissionEnum string
const (
	Cloudprojectai_VolumePermissionEnum_RO Cloudprojectai_VolumePermissionEnum = "RO"
	Cloudprojectai_VolumePermissionEnum_RW  = "RW"
	Cloudprojectai_VolumePermissionEnum_RWD  = "RWD"
)

type Cloudprojectaiapp_App struct {
	CreatedAt time.Time                   `json:"createdAt"`
	Id        string                      `json:"id"`
	Spec      Cloudprojectaiapp_AppSpec   `json:"spec"`
	Status    Cloudprojectaiapp_AppStatus `json:"status"`
	UpdatedAt time.Time                   `json:"updatedAt"`
	User      string                      `json:"user"`
}

type Cloudprojectaiapp_AppImageInput struct {
	Url string `json:"url"`
}

type Cloudprojectaiapp_AppSpec struct {
	Command         []string                          `json:"command"`
	DefaultHttpPort int64                             `json:"defaultHttpPort"`
	Env             []Cloudprojectaijob_JobEnv        `json:"env"`
	Image           string                            `json:"image"`
	Labels          map[string]string                 `json:"labels"`
	Name            string                            `json:"name"`
	PartnerId       string                            `json:"partnerId"`
	Probe           Cloudprojectaiapp_Probe           `json:"probe"`
	Region          string                            `json:"region"`
	Resources       Cloudprojectai_Resources          `json:"resources"`
	ScalingStrategy Cloudprojectaiapp_ScalingStrategy `json:"scalingStrategy"`
	UnsecureHttp    bool                              `json:"unsecureHttp"`
	Volumes         []Cloudprojectaivolume_Volume     `json:"volumes"`
}

type Cloudprojectaiapp_AppSpecInput struct {
	Command         []string                               `json:"command"`
	DefaultHttpPort int64                                  `json:"defaultHttpPort"`
	Env             []Cloudprojectaijob_JobEnv             `json:"env"`
	Image           string                                 `json:"image"`
	Labels          map[string]string                      `json:"labels"`
	Name            string                                 `json:"name"`
	PartnerId       string                                 `json:"partnerId"`
	Probe           Cloudprojectaiapp_ProbeInput           `json:"probe"`
	Region          string                                 `json:"region"`
	Resources       Cloudprojectai_ResourcesInput          `json:"resources"`
	ScalingStrategy Cloudprojectaiapp_ScalingStrategyInput `json:"scalingStrategy"`
	UnsecureHttp    bool                                   `json:"unsecureHttp"`
	Volumes         []Cloudprojectaivolume_Volume          `json:"volumes"`
}

type Cloudprojectaiapp_AppStateEnum string
const (
	Cloudprojectaiapp_AppStateEnum_DELETED Cloudprojectaiapp_AppStateEnum = "DELETED"
	Cloudprojectaiapp_AppStateEnum_DELETING  = "DELETING"
	Cloudprojectaiapp_AppStateEnum_ERROR  = "ERROR"
	Cloudprojectaiapp_AppStateEnum_FAILED  = "FAILED"
	Cloudprojectaiapp_AppStateEnum_INITIALIZING  = "INITIALIZING"
	Cloudprojectaiapp_AppStateEnum_QUEUED  = "QUEUED"
	Cloudprojectaiapp_AppStateEnum_RUNNING  = "RUNNING"
	Cloudprojectaiapp_AppStateEnum_SCALING  = "SCALING"
	Cloudprojectaiapp_AppStateEnum_STOPPED  = "STOPPED"
	Cloudprojectaiapp_AppStateEnum_STOPPING  = "STOPPING"
)

type Cloudprojectaiapp_AppStateHistory struct {
	Date  time.Time                      `json:"date"`
	State Cloudprojectaiapp_AppStateEnum `json:"state"`
}

type Cloudprojectaiapp_AppStatus struct {
	AvailableReplicas  int64                               `json:"availableReplicas"`
	DataSync           []Cloudprojectaivolume_DataSync     `json:"dataSync"`
	History            []Cloudprojectaiapp_AppStateHistory `json:"history"`
	Info               Cloudprojectai_Info                 `json:"info"`
	InfoUrl            string                              `json:"infoUrl"`
	LastTransitionDate time.Time                           `json:"lastTransitionDate"`
	MonitoringUrl      string                              `json:"monitoringUrl"`
	State              Cloudprojectaiapp_AppStateEnum      `json:"state"`
	Url                string                              `json:"url"`
}

type Cloudprojectaiapp_Probe struct {
	Path string `json:"path"`
	Port int64  `json:"port"`
}

type Cloudprojectaiapp_ProbeInput struct {
	Path string `json:"path"`
	Port int64  `json:"port"`
}

type Cloudprojectaiapp_ScalingAutomaticStrategy struct {
	AverageUsageTarget int64                                                      `json:"averageUsageTarget"`
	ReplicasMax        int64                                                      `json:"replicasMax"`
	ReplicasMin        int64                                                      `json:"replicasMin"`
	ResourceType       Cloudprojectaiapp_ScalingAutomaticStrategyResourceTypeEnum `json:"resourceType"`
}

type Cloudprojectaiapp_ScalingAutomaticStrategyInput struct {
	AverageUsageTarget int64                                                      `json:"averageUsageTarget"`
	ReplicasMax        int64                                                      `json:"replicasMax"`
	ReplicasMin        int64                                                      `json:"replicasMin"`
	ResourceType       Cloudprojectaiapp_ScalingAutomaticStrategyResourceTypeEnum `json:"resourceType"`
}

type Cloudprojectaiapp_ScalingAutomaticStrategyResourceTypeEnum string
const (
	Cloudprojectaiapp_ScalingAutomaticStrategyResourceTypeEnum_CPU Cloudprojectaiapp_ScalingAutomaticStrategyResourceTypeEnum = "CPU"
	Cloudprojectaiapp_ScalingAutomaticStrategyResourceTypeEnum_RAM  = "RAM"
)

type Cloudprojectaiapp_ScalingFixedStrategy struct {
	Replicas int64 `json:"replicas"`
}

type Cloudprojectaiapp_ScalingFixedStrategyInput struct {
	Replicas int64 `json:"replicas"`
}

type Cloudprojectaiapp_ScalingStrategy struct {
	Automatic Cloudprojectaiapp_ScalingAutomaticStrategy `json:"automatic"`
	Fixed     Cloudprojectaiapp_ScalingFixedStrategy     `json:"fixed"`
}

type Cloudprojectaiapp_ScalingStrategyInput struct {
	Automatic Cloudprojectaiapp_ScalingAutomaticStrategyInput `json:"automatic"`
	Fixed     Cloudprojectaiapp_ScalingFixedStrategyInput     `json:"fixed"`
}

type Cloudprojectaicapabilities_Features struct {
	Lab      bool `json:"lab"`
	Registry bool `json:"registry"`
}

type Cloudprojectaicapabilities_Flavor struct {
	Default          bool                                              `json:"default"`
	Description      string                                            `json:"description"`
	GpuInformation   Cloudprojectaicapabilitiesflavor_GpuInformation   `json:"gpuInformation"`
	Id               string                                            `json:"id"`
	Max              int64                                             `json:"max"`
	ResourcesPerUnit Cloudprojectaicapabilitiesflavor_ResourcesPerUnit `json:"resourcesPerUnit"`
	Type             Cloudprojectaicapabilities_FlavorTypeEnum         `json:"type"`
}

type Cloudprojectaicapabilities_FlavorTypeEnum string
const (
	Cloudprojectaicapabilities_FlavorTypeEnum_cpu Cloudprojectaicapabilities_FlavorTypeEnum = "cpu"
	Cloudprojectaicapabilities_FlavorTypeEnum_gpu  = "gpu"
)

type Cloudprojectaicapabilities_Preset struct {
	Capabilities Cloudprojectaicapabilities_PresetCapabilities       `json:"capabilities"`
	Descriptions []string                                            `json:"descriptions"`
	DocUrl       []Cloudprojectaicapabilities_PresetDocumentationUrl `json:"docUrl"`
	Id           string                                              `json:"id"`
	LogoUrl      string                                              `json:"logoUrl"`
	Name         string                                              `json:"name"`
	Partner      Cloudprojectaijob_Partner                           `json:"partner"`
	Snippet      string                                              `json:"snippet"`
	Type         Cloudprojectaicapabilities_PresetTypeEnum           `json:"type"`
}

type Cloudprojectaicapabilities_PresetCapabilities struct {
	Exec        bool                                        `json:"exec"`
	FlavorTypes []Cloudprojectaicapabilities_FlavorTypeEnum `json:"flavorTypes"`
	Log         bool                                        `json:"log"`
	Resources   Cloudprojectaicapabilities_PresetResources  `json:"resources"`
	Ssh         bool                                        `json:"ssh"`
	Volume      bool                                        `json:"volume"`
}

type Cloudprojectaicapabilities_PresetDocumentationUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Cloudprojectaicapabilities_PresetResources struct {
	MaxGpu int64 `json:"maxGpu"`
}

type Cloudprojectaicapabilities_PresetTypeEnum string
const (
	Cloudprojectaicapabilities_PresetTypeEnum_app Cloudprojectaicapabilities_PresetTypeEnum = "app"
	Cloudprojectaicapabilities_PresetTypeEnum_job  = "job"
	Cloudprojectaicapabilities_PresetTypeEnum_notebook  = "notebook"
)

type Cloudprojectaicapabilities_ProjectQuotas struct {
	Resources map[string]int64 `json:"resources"`
	Storage   int64            `json:"storage"`
}

type Cloudprojectaicapabilities_Region struct {
	CliInstallUrl    string `json:"cliInstallUrl"`
	DocumentationUrl string `json:"documentationUrl"`
	Id               string `json:"id"`
	RegistryUrl      string `json:"registryUrl"`
	Version          string `json:"version"`
}

type Cloudprojectaicapabilitiesflavor_GpuInformation struct {
	GpuBrand  string `json:"gpuBrand"`
	GpuMemory int64  `json:"gpuMemory"`
	GpuModel  string `json:"gpuModel"`
}

type Cloudprojectaicapabilitiesflavor_ResourcesPerUnit struct {
	Cpu              int64 `json:"cpu"`
	EphemeralStorage int64 `json:"ephemeralStorage"`
	Memory           int64 `json:"memory"`
	PrivateNetwork   int64 `json:"privateNetwork"`
	PublicNetwork    int64 `json:"publicNetwork"`
}

type Cloudprojectaijob_Job struct {
	CreatedAt time.Time                   `json:"createdAt"`
	Id        string                      `json:"id"`
	Spec      Cloudprojectaijob_JobSpec   `json:"spec"`
	Status    Cloudprojectaijob_JobStatus `json:"status"`
	UpdatedAt time.Time                   `json:"updatedAt"`
	User      string                      `json:"user"`
}

type Cloudprojectaijob_JobEnv struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Cloudprojectaijob_JobSpec struct {
	Command         []string                            `json:"command"`
	DefaultHttpPort int64                               `json:"defaultHttpPort"`
	Env             []Cloudprojectaijob_JobEnv          `json:"env"`
	Image           string                              `json:"image"`
	Labels          map[string]string                   `json:"labels"`
	Name            string                              `json:"name"`
	PartnerId       string                              `json:"partnerId"`
	ReadUser        string                              `json:"readUser"`
	Region          string                              `json:"region"`
	Resources       Cloudprojectai_Resources            `json:"resources"`
	Shutdown        Cloudprojectai_ShutdownStrategyEnum `json:"shutdown"`
	SshPublicKeys   []string                            `json:"sshPublicKeys"`
	Timeout         int64                               `json:"timeout"`
	UnsecureHttp    bool                                `json:"unsecureHttp"`
	Volumes         []Cloudprojectaivolume_Volume       `json:"volumes"`
}

type Cloudprojectaijob_JobSpecInput struct {
	Command         []string                            `json:"command"`
	DefaultHttpPort int64                               `json:"defaultHttpPort"`
	Env             []Cloudprojectaijob_JobEnv          `json:"env"`
	Image           string                              `json:"image"`
	Labels          map[string]string                   `json:"labels"`
	Name            string                              `json:"name"`
	PartnerId       string                              `json:"partnerId"`
	ReadUser        string                              `json:"readUser"`
	Region          string                              `json:"region"`
	Resources       Cloudprojectai_ResourcesInput       `json:"resources"`
	Shutdown        Cloudprojectai_ShutdownStrategyEnum `json:"shutdown"`
	SshPublicKeys   []string                            `json:"sshPublicKeys"`
	Timeout         int64                               `json:"timeout"`
	UnsecureHttp    bool                                `json:"unsecureHttp"`
	Volumes         []Cloudprojectaivolume_Volume       `json:"volumes"`
}

type Cloudprojectaijob_JobStateEnum string
const (
	Cloudprojectaijob_JobStateEnum_DONE Cloudprojectaijob_JobStateEnum = "DONE"
	Cloudprojectaijob_JobStateEnum_ERROR  = "ERROR"
	Cloudprojectaijob_JobStateEnum_FAILED  = "FAILED"
	Cloudprojectaijob_JobStateEnum_FINALIZING  = "FINALIZING"
	Cloudprojectaijob_JobStateEnum_INITIALIZING  = "INITIALIZING"
	Cloudprojectaijob_JobStateEnum_INTERRUPTED  = "INTERRUPTED"
	Cloudprojectaijob_JobStateEnum_INTERRUPTING  = "INTERRUPTING"
	Cloudprojectaijob_JobStateEnum_PENDING  = "PENDING"
	Cloudprojectaijob_JobStateEnum_QUEUED  = "QUEUED"
	Cloudprojectaijob_JobStateEnum_RUNNING  = "RUNNING"
	Cloudprojectaijob_JobStateEnum_TIMEOUT  = "TIMEOUT"
)

type Cloudprojectaijob_JobStatus struct {
	DataSync           []Cloudprojectaivolume_DataSync      `json:"dataSync"`
	Duration           int64                                `json:"duration"`
	ExitCode           int64                                `json:"exitCode"`
	ExternalIp         string                               `json:"externalIp"`
	FinalizedAt        time.Time                            `json:"finalizedAt"`
	History            []Cloudprojectaijob_JobStatusHistory `json:"history"`
	Info               Cloudprojectai_Info                  `json:"info"`
	InfoUrl            string                               `json:"infoUrl"`
	InitializingAt     time.Time                            `json:"initializingAt"`
	Ip                 string                               `json:"ip"`
	LastTransitionDate time.Time                            `json:"lastTransitionDate"`
	MonitoringUrl      string                               `json:"monitoringUrl"`
	QueuedAt           time.Time                            `json:"queuedAt"`
	SshUrl             string                               `json:"sshUrl"`
	StartedAt          time.Time                            `json:"startedAt"`
	State              Cloudprojectaijob_JobStateEnum       `json:"state"`
	StoppedAt          time.Time                            `json:"stoppedAt"`
	Url                string                               `json:"url"`
}

type Cloudprojectaijob_JobStatusHistory struct {
	Date  time.Time                      `json:"date"`
	State Cloudprojectaijob_JobStateEnum `json:"state"`
}

type Cloudprojectaijob_Partner struct {
	Flavor string `json:"flavor"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

type Cloudprojectaijob_PresetImage struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Link        string `json:"link"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
}

type Cloudprojectainotebook_Editor struct {
	Description string `json:"description"`
	DocUrl      string `json:"docUrl"`
	Id          string `json:"id"`
	LogoUrl     string `json:"logoUrl"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

type Cloudprojectainotebook_Framework struct {
	Description string   `json:"description"`
	DocUrl      string   `json:"docUrl"`
	Id          string   `json:"id"`
	LogoUrl     string   `json:"logoUrl"`
	Name        string   `json:"name"`
	Versions    []string `json:"versions"`
}

type Cloudprojectainotebook_Notebook struct {
	CreatedAt time.Time                             `json:"createdAt"`
	Id        string                                `json:"id"`
	Spec      Cloudprojectainotebook_NotebookSpec   `json:"spec"`
	Status    Cloudprojectainotebook_NotebookStatus `json:"status"`
	UpdatedAt time.Time                             `json:"updatedAt"`
	User      string                                `json:"user"`
}

type Cloudprojectainotebook_NotebookEnv struct {
	EditorId         string `json:"editorId"`
	FrameworkId      string `json:"frameworkId"`
	FrameworkVersion string `json:"frameworkVersion"`
}

type Cloudprojectainotebook_NotebookSpec struct {
	Env           Cloudprojectainotebook_NotebookEnv  `json:"env"`
	Flavor        string                              `json:"flavor"`
	Labels        map[string]string                   `json:"labels"`
	Name          string                              `json:"name"`
	Region        string                              `json:"region"`
	Resources     Cloudprojectai_Resources            `json:"resources"`
	Shutdown      Cloudprojectai_ShutdownStrategyEnum `json:"shutdown"`
	SshPublicKeys []string                            `json:"sshPublicKeys"`
	UnsecureHttp  bool                                `json:"unsecureHttp"`
	Volumes       []Cloudprojectaivolume_Volume       `json:"volumes"`
}

type Cloudprojectainotebook_NotebookSpecInput struct {
	Env           Cloudprojectainotebook_NotebookEnv  `json:"env"`
	Labels        map[string]string                   `json:"labels"`
	Name          string                              `json:"name"`
	Region        string                              `json:"region"`
	Resources     Cloudprojectai_ResourcesInput       `json:"resources"`
	Shutdown      Cloudprojectai_ShutdownStrategyEnum `json:"shutdown"`
	SshPublicKeys []string                            `json:"sshPublicKeys"`
	UnsecureHttp  bool                                `json:"unsecureHttp"`
	Volumes       []Cloudprojectaivolume_Volume       `json:"volumes"`
}

type Cloudprojectainotebook_NotebookStateEnum string
const (
	Cloudprojectainotebook_NotebookStateEnum_DELETING Cloudprojectainotebook_NotebookStateEnum = "DELETING"
	Cloudprojectainotebook_NotebookStateEnum_FAILED  = "FAILED"
	Cloudprojectainotebook_NotebookStateEnum_RUNNING  = "RUNNING"
	Cloudprojectainotebook_NotebookStateEnum_STARTING  = "STARTING"
	Cloudprojectainotebook_NotebookStateEnum_STOPPED  = "STOPPED"
	Cloudprojectainotebook_NotebookStateEnum_STOPPING  = "STOPPING"
)

type Cloudprojectainotebook_NotebookStatus struct {
	DataSync      []Cloudprojectaivolume_DataSync          `json:"dataSync"`
	Duration      int64                                    `json:"duration"`
	Info          Cloudprojectai_Info                      `json:"info"`
	InfoUrl       string                                   `json:"infoUrl"`
	LastJobStatus Cloudprojectaijob_JobStatus              `json:"lastJobStatus"`
	LastStartedAt time.Time                                `json:"lastStartedAt"`
	LastStoppedAt time.Time                                `json:"lastStoppedAt"`
	MonitoringUrl string                                   `json:"monitoringUrl"`
	SshUrl        string                                   `json:"sshUrl"`
	State         Cloudprojectainotebook_NotebookStateEnum `json:"state"`
	Url           string                                   `json:"url"`
	Workspace     Cloudprojectainotebook_NotebookWorkspace `json:"workspace"`
}

type Cloudprojectainotebook_NotebookUpdate struct {
	Labels        map[string]string             `json:"labels"`
	Resources     Cloudprojectai_ResourcesInput `json:"resources"`
	SshPublicKeys []string                      `json:"sshPublicKeys"`
	UnsecureHttp  bool                          `json:"unsecureHttp"`
	Volumes       []Cloudprojectaivolume_Volume `json:"volumes"`
}

type Cloudprojectainotebook_NotebookWorkspace struct {
	StorageFree int64 `json:"storageFree"`
	StorageUsed int64 `json:"storageUsed"`
}

type Cloudprojectairegistry_Registry struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Password  string    `json:"password"`
	Region    string    `json:"region"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url       string    `json:"url"`
	User      string    `json:"user"`
	Username  string    `json:"username"`
}

type Cloudprojectaiserving_APIStatusEnum string
const (
	Cloudprojectaiserving_APIStatusEnum_pending Cloudprojectaiserving_APIStatusEnum = "pending"
	Cloudprojectaiserving_APIStatusEnum_running  = "running"
	Cloudprojectaiserving_APIStatusEnum_scaling  = "scaling"
	Cloudprojectaiserving_APIStatusEnum_sleeping  = "sleeping"
	Cloudprojectaiserving_APIStatusEnum_starting  = "starting"
	Cloudprojectaiserving_APIStatusEnum_waking  = "waking"
)

type Cloudprojectaiserving_AutoscalingSpec struct {
	CpuAverageUtilization    int64 `json:"cpuAverageUtilization"`
	MaxReplicas              int64 `json:"maxReplicas"`
	MemoryAverageUtilization int64 `json:"memoryAverageUtilization"`
	MinReplicas              int64 `json:"minReplicas"`
}

type Cloudprojectaiserving_Backend struct {
	Id   Cloudprojectaiserving_BackendIdEnum `json:"id"`
	Link string                              `json:"link"`
	Name string                              `json:"name"`
}

type Cloudprojectaiserving_BackendIdEnum string
const (
	Cloudprojectaiserving_BackendIdEnum_bentoml Cloudprojectaiserving_BackendIdEnum = "bentoml"
	Cloudprojectaiserving_BackendIdEnum_servingruntime  = "serving-runtime"
)

type Cloudprojectaiserving_Features struct {
	ChooseBackend bool `json:"chooseBackend"`
}

type Cloudprojectaiserving_Flavor struct {
	CpuCore     int64  `json:"cpuCore"`
	Description string `json:"description"`
	Id          string `json:"id"`
	RamMB       int64  `json:"ramMB"`
}

type Cloudprojectaiserving_Framework struct {
	Backends           []Cloudprojectaiserving_BackendIdEnum `json:"backends"`
	DocPage            string                                `json:"docPage"`
	Id                 Cloudprojectaiserving_FrameworkIdEnum `json:"id"`
	Link               string                                `json:"link"`
	Logo               string                                `json:"logo"`
	Name               string                                `json:"name"`
	RecommendedBackend Cloudprojectaiserving_BackendIdEnum   `json:"recommendedBackend"`
	Version            string                                `json:"version"`
}

type Cloudprojectaiserving_FrameworkIdEnum string
const (
	Cloudprojectaiserving_FrameworkIdEnum_fastai Cloudprojectaiserving_FrameworkIdEnum = "fastai"
	Cloudprojectaiserving_FrameworkIdEnum_flow  = "flow"
	Cloudprojectaiserving_FrameworkIdEnum_huggingface  = "huggingface"
	Cloudprojectaiserving_FrameworkIdEnum_onnx  = "onnx"
	Cloudprojectaiserving_FrameworkIdEnum_pmml  = "pmml"
	Cloudprojectaiserving_FrameworkIdEnum_tensorflow_1  = "tensorflow_1"
	Cloudprojectaiserving_FrameworkIdEnum_torch  = "torch"
)

type Cloudprojectaiserving_Metrics struct {
	Endpoints []Cloudprojectaiserving_MetricsEndpoint `json:"endpoints"`
	Token     string                                  `json:"token"`
}

type Cloudprojectaiserving_MetricsEndpoint struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Cloudprojectaiserving_Model struct {
	ApiStatus                  Cloudprojectaiserving_APIStatusEnum                  `json:"apiStatus"`
	AutoscalingSpec            Cloudprojectaiserving_AutoscalingSpec                `json:"autoscalingSpec"`
	CreatedAt                  time.Time                                            `json:"createdAt"`
	Flavor                     string                                               `json:"flavor"`
	Id                         string                                               `json:"id"`
	Replicas                   int64                                                `json:"replicas"`
	Url                        string                                               `json:"url"`
	Version                    int64                                                `json:"version"`
	VersionStatus              Cloudprojectaiserving_VersionStatusEnum              `json:"versionStatus"`
	WorkflowTemplate           Cloudprojectaiserving_WorkflowTemplateEnum           `json:"workflowTemplate"`
	WorkflowTemplateParameters Cloudprojectaiserving_ModelWorkflowTemplateParameter `json:"workflowTemplateParameters"`
}

type Cloudprojectaiserving_ModelDefinition struct {
	AutoscalingSpec  Cloudprojectaiserving_AutoscalingSpec      `json:"autoscalingSpec"`
	Backend          Cloudprojectaiserving_BackendIdEnum        `json:"backend"`
	Flavor           string                                     `json:"flavor"`
	Framework        Cloudprojectaiserving_FrameworkIdEnum      `json:"framework"`
	Id               string                                     `json:"id"`
	ImageId          string                                     `json:"imageId"`
	StoragePath      string                                     `json:"storagePath"`
	WorkflowTemplate Cloudprojectaiserving_WorkflowTemplateEnum `json:"workflowTemplate"`
}

type Cloudprojectaiserving_ModelWorkflowTemplateParameter struct {
	Backend     Cloudprojectaiserving_BackendIdEnum   `json:"backend"`
	Framework   Cloudprojectaiserving_FrameworkIdEnum `json:"framework"`
	ImageId     string                                `json:"imageId"`
	StoragePath string                                `json:"storagePath"`
}

type Cloudprojectaiserving_Namespace struct {
	ClusterId   string    `json:"clusterId"`
	Container   string    `json:"container"`
	ContainerId string    `json:"containerId"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	HubUrl      string    `json:"hubUrl"`
	Id          string    `json:"id"`
	Region      string    `json:"region"`
	Url         string    `json:"url"`
}

type Cloudprojectaiserving_NamespaceCreation struct {
	Container   string `json:"container"`
	Description string `json:"description"`
	Region      string `json:"region"`
}

type Cloudprojectaiserving_PresetImage struct {
	Description      string `json:"description"`
	Id               string `json:"id"`
	Link             string `json:"link"`
	Name             string `json:"name"`
	RamRequirementMB int64  `json:"ramRequirementMB"`
}

type Cloudprojectaiserving_Registry struct {
	Custom   bool   `json:"custom"`
	Password string `json:"password"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type Cloudprojectaiserving_RegistryResponse struct {
	Message string `json:"message"`
}

type Cloudprojectaiserving_Token struct {
	CreatedAt time.Time                              `json:"createdAt"`
	Groups    []Cloudprojectaiserving_TokenGroupEnum `json:"groups"`
	Id        string                                 `json:"id"`
	Resource  string                                 `json:"resource"`
	Token     string                                 `json:"token"`
}

type Cloudprojectaiserving_TokenGroupEnum string
const (
	Cloudprojectaiserving_TokenGroupEnum_modelevaluation Cloudprojectaiserving_TokenGroupEnum = "model-evaluation"
	Cloudprojectaiserving_TokenGroupEnum_modelmanagement  = "model-management"
)

type Cloudprojectaiserving_VersionStatusEnum string
const (
	Cloudprojectaiserving_VersionStatusEnum_builderror Cloudprojectaiserving_VersionStatusEnum = "build-error"
	Cloudprojectaiserving_VersionStatusEnum_building  = "building"
	Cloudprojectaiserving_VersionStatusEnum_built  = "built"
	Cloudprojectaiserving_VersionStatusEnum_deployed  = "deployed"
	Cloudprojectaiserving_VersionStatusEnum_deploying  = "deploying"
	Cloudprojectaiserving_VersionStatusEnum_failed  = "failed"
	Cloudprojectaiserving_VersionStatusEnum_pending  = "pending"
	Cloudprojectaiserving_VersionStatusEnum_rollback  = "rollback"
)

type Cloudprojectaiserving_WorkflowTemplateEnum string
const (
	Cloudprojectaiserving_WorkflowTemplateEnum_buildimage Cloudprojectaiserving_WorkflowTemplateEnum = "build-image"
	Cloudprojectaiserving_WorkflowTemplateEnum_presetimage  = "preset-image"
)

type Cloudprojectaitoken_Token struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Id        string                          `json:"id"`
	Spec      Cloudprojectaitoken_TokenSpec   `json:"spec"`
	Status    Cloudprojectaitoken_TokenStatus `json:"status"`
	UpdatedAt time.Time                       `json:"updatedAt"`
}

type Cloudprojectaitoken_TokenSpec struct {
	LabelSelector string                       `json:"labelSelector"`
	Name          string                       `json:"name"`
	Region        string                       `json:"region"`
	Role          Cloudprojectai_TokenRoleEnum `json:"role"`
}

type Cloudprojectaitoken_TokenStatus struct {
	Value   string `json:"value"`
	Version int64  `json:"version"`
}

type Cloudprojectaivolume_DataSync struct {
	CreatedAt time.Time                           `json:"createdAt"`
	Id        string                              `json:"id"`
	Spec      Cloudprojectaivolume_DataSyncSpec   `json:"spec"`
	Status    Cloudprojectaivolume_DataSyncStatus `json:"status"`
	UpdatedAt time.Time                           `json:"updatedAt"`
}

type Cloudprojectaivolume_DataSyncEnum string
const (
	Cloudprojectaivolume_DataSyncEnum_pull Cloudprojectaivolume_DataSyncEnum = "pull"
	Cloudprojectaivolume_DataSyncEnum_push  = "push"
)

type Cloudprojectaivolume_DataSyncProgressStateEnum string
const (
	Cloudprojectaivolume_DataSyncProgressStateEnum_DONE Cloudprojectaivolume_DataSyncProgressStateEnum = "DONE"
	Cloudprojectaivolume_DataSyncProgressStateEnum_ERROR  = "ERROR"
	Cloudprojectaivolume_DataSyncProgressStateEnum_FAILED  = "FAILED"
	Cloudprojectaivolume_DataSyncProgressStateEnum_INTERRUPTED  = "INTERRUPTED"
	Cloudprojectaivolume_DataSyncProgressStateEnum_QUEUED  = "QUEUED"
	Cloudprojectaivolume_DataSyncProgressStateEnum_RUNNING  = "RUNNING"
)

type Cloudprojectaivolume_DataSyncSpec struct {
	Direction Cloudprojectaivolume_DataSyncEnum `json:"direction"`
	Manual    bool                              `json:"manual"`
}

type Cloudprojectaivolume_DataSyncStateEnum string
const (
	Cloudprojectaivolume_DataSyncStateEnum_DONE Cloudprojectaivolume_DataSyncStateEnum = "DONE"
	Cloudprojectaivolume_DataSyncStateEnum_ERROR  = "ERROR"
	Cloudprojectaivolume_DataSyncStateEnum_FAILED  = "FAILED"
	Cloudprojectaivolume_DataSyncStateEnum_INTERRUPTED  = "INTERRUPTED"
	Cloudprojectaivolume_DataSyncStateEnum_QUEUED  = "QUEUED"
	Cloudprojectaivolume_DataSyncStateEnum_RUNNING  = "RUNNING"
)

type Cloudprojectaivolume_DataSyncStatus struct {
	EndedAt   time.Time                              `json:"endedAt"`
	Info      Cloudprojectai_Info                    `json:"info"`
	Progress  []Cloudprojectaivolume_Progress        `json:"progress"`
	QueuedAt  time.Time                              `json:"queuedAt"`
	StartedAt time.Time                              `json:"startedAt"`
	State     Cloudprojectaivolume_DataSyncStateEnum `json:"state"`
}

type Cloudprojectaivolume_PrivateSwift struct {
	Archive   string `json:"archive"`
	Container string `json:"container"`
	Internal  bool   `json:"internal"`
	Prefix    string `json:"prefix"`
	Region    string `json:"region"`
}

type Cloudprojectaivolume_Progress struct {
	Completed        int64                                          `json:"completed"`
	CreatedAt        time.Time                                      `json:"createdAt"`
	Deleted          int64                                          `json:"deleted"`
	Direction        Cloudprojectaivolume_DataSyncEnum              `json:"direction"`
	Eta              int64                                          `json:"eta"`
	Failed           int64                                          `json:"failed"`
	Id               string                                         `json:"id"`
	Info             string                                         `json:"info"`
	Processed        int64                                          `json:"processed"`
	Skipped          int64                                          `json:"skipped"`
	State            Cloudprojectaivolume_DataSyncProgressStateEnum `json:"state"`
	Total            int64                                          `json:"total"`
	TransferredBytes int64                                          `json:"transferredBytes"`
	UpdatedAt        time.Time                                      `json:"updatedAt"`
}

type Cloudprojectaivolume_PublicGit struct {
	Url string `json:"url"`
}

type Cloudprojectaivolume_PublicSwift struct {
	Url string `json:"url"`
}

type Cloudprojectaivolume_Standalone struct {
	Name string `json:"name"`
}

type Cloudprojectaivolume_Volume struct {
	Cache              bool                                `json:"cache"`
	Container          string                              `json:"container"`
	MountPath          string                              `json:"mountPath"`
	Permission         Cloudprojectai_VolumePermissionEnum `json:"permission"`
	Prefix             string                              `json:"prefix"`
	PrivateSwift       Cloudprojectaivolume_PrivateSwift   `json:"privateSwift"`
	PublicGit          Cloudprojectaivolume_PublicGit      `json:"publicGit"`
	PublicSwift        Cloudprojectaivolume_PublicSwift    `json:"publicSwift"`
	Region             string                              `json:"region"`
	Standalone         Cloudprojectaivolume_Standalone     `json:"standalone"`
	TargetPrivateSwift Cloudprojectaivolume_PrivateSwift   `json:"targetPrivateSwift"`
}

type Cloudprojectcertificate_Import struct {
	Cert  string `json:"cert"`
	Chain string `json:"chain"`
	Key   string `json:"key"`
}

type Cloudprojectcertificate_ServerAlternativeName struct {
	Kind Cloudprojectcertificate_ServerAlternativeNameKindEnum `json:"kind"`
	Name string                                                `json:"name"`
}

type Cloudprojectcertificate_ServerAlternativeNameKindEnum string
const (
	Cloudprojectcertificate_ServerAlternativeNameKindEnum_DNS Cloudprojectcertificate_ServerAlternativeNameKindEnum = "DNS"
	Cloudprojectcertificate_ServerAlternativeNameKindEnum_EMAIL  = "EMAIL"
	Cloudprojectcertificate_ServerAlternativeNameKindEnum_IP  = "IP"
	Cloudprojectcertificate_ServerAlternativeNameKindEnum_URI  = "URI"
)

type CloudprojectdataProcessing_AuthorizationStatus struct {
	Authorized bool `json:"authorized"`
}

type CloudprojectdataProcessing_CapabilitiesEngineParameter struct {
	Default     string                                        `json:"default"`
	Description string                                        `json:"description"`
	Mandatory   bool                                          `json:"mandatory"`
	Name        string                                        `json:"name"`
	Type        string                                        `json:"type"`
	Validator   CloudprojectdataProcessing_ParameterValidator `json:"validator"`
}

type CloudprojectdataProcessing_CapabilitiesTemplate struct {
	Cores  int64 `json:"cores"`
	Id     int64 `json:"id"`
	Memory int64 `json:"memory"`
}

type CloudprojectdataProcessing_Capability struct {
	AvailableVersions []CloudprojectdataProcessing_EngineVersion               `json:"availableVersions"`
	Name              string                                                   `json:"name"`
	Parameters        []CloudprojectdataProcessing_CapabilitiesEngineParameter `json:"parameters"`
	Templates         []CloudprojectdataProcessing_CapabilitiesTemplate        `json:"templates"`
}

type CloudprojectdataProcessing_EngineParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CloudprojectdataProcessing_EngineVersion struct {
	AvailableRegions []string `json:"availableRegions"`
	Description      string   `json:"description"`
	Name             string   `json:"name"`
}

type CloudprojectdataProcessing_Job struct {
	ContainerName    string                                       `json:"containerName"`
	CreationDate     time.Time                                    `json:"creationDate"`
	EndDate          time.Time                                    `json:"endDate"`
	Engine           string                                       `json:"engine"`
	EngineParameters []CloudprojectdataProcessing_EngineParameter `json:"engineParameters"`
	EngineVersion    string                                       `json:"engineVersion"`
	Id               string                                       `json:"id"`
	Name             string                                       `json:"name"`
	Region           string                                       `json:"region"`
	ReturnCode       int64                                        `json:"returnCode"`
	StartDate        time.Time                                    `json:"startDate"`
	Status           CloudprojectdataProcessing_StatusEnum        `json:"status"`
	Ttl              time.Duration                                `json:"ttl"`
}

type CloudprojectdataProcessing_JobLogs struct {
	Logs        []CloudprojectdataProcessing_LogLine `json:"logs"`
	LogsAddress string                               `json:"logsAddress"`
	StartDate   time.Time                            `json:"startDate"`
}

type CloudprojectdataProcessing_LogLine struct {
	Content   string    `json:"content"`
	Id        int64     `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type CloudprojectdataProcessing_Metrics struct {
	Endpoints []CloudprojectdataProcessing_MetricsEndpoint `json:"endpoints"`
	Token     string                                       `json:"token"`
}

type CloudprojectdataProcessing_MetricsEndpoint struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CloudprojectdataProcessing_ParameterValidator struct {
	Max   float64 `json:"max"`
	Min   float64 `json:"min"`
	Regex string  `json:"regex"`
}

type CloudprojectdataProcessing_StatusEnum string
const (
	CloudprojectdataProcessing_StatusEnum_CANCELLING CloudprojectdataProcessing_StatusEnum = "CANCELLING"
	CloudprojectdataProcessing_StatusEnum_COMPLETED  = "COMPLETED"
	CloudprojectdataProcessing_StatusEnum_FAILED  = "FAILED"
	CloudprojectdataProcessing_StatusEnum_PENDING  = "PENDING"
	CloudprojectdataProcessing_StatusEnum_RUNNING  = "RUNNING"
	CloudprojectdataProcessing_StatusEnum_SUBMITTED  = "SUBMITTED"
	CloudprojectdataProcessing_StatusEnum_TERMINATED  = "TERMINATED"
	CloudprojectdataProcessing_StatusEnum_UNKNOWN  = "UNKNOWN"
)

type Cloudprojectdatabase_Availability struct {
	Backup              Cloudprojectdatabase_BackupTypeEnum         `json:"backup"`
	BackupRetentionDays int64                                       `json:"backupRetentionDays"`
	Default             bool                                        `json:"default"`
	EndOfLife           time.Time                                   `json:"endOfLife"`
	Engine              string                                      `json:"engine"`
	Flavor              string                                      `json:"flavor"`
	MaxDiskSize         int64                                       `json:"maxDiskSize"`
	MaxNodeNumber       int64                                       `json:"maxNodeNumber"`
	MinDiskSize         int64                                       `json:"minDiskSize"`
	MinNodeNumber       int64                                       `json:"minNodeNumber"`
	Network             Cloudprojectdatabase_NetworkTypeEnum        `json:"network"`
	Plan                string                                      `json:"plan"`
	Region              string                                      `json:"region"`
	StartDate           time.Time                                   `json:"startDate"`
	Status              Cloudprojectdatabaseavailability_StatusEnum `json:"status"`
	StepDiskSize        int64                                       `json:"stepDiskSize"`
	UpstreamEndOfLife   time.Time                                   `json:"upstreamEndOfLife"`
	Version             string                                      `json:"version"`
}

type Cloudprojectdatabase_BackupTypeEnum string
const (
	Cloudprojectdatabase_BackupTypeEnum_automatic Cloudprojectdatabase_BackupTypeEnum = "automatic"
	Cloudprojectdatabase_BackupTypeEnum_manual  = "manual"
)

type Cloudprojectdatabase_Capabilities struct {
	Disks   []string                                  `json:"disks"`
	Engines []Cloudprojectdatabasecapabilities_Engine `json:"engines"`
	Flavors []Cloudprojectdatabasecapabilities_Flavor `json:"flavors"`
	Options []Cloudprojectdatabasecapabilities_Option `json:"options"`
	Plans   []Cloudprojectdatabasecapabilities_Plan   `json:"plans"`
	Regions []string                                  `json:"regions"`
}

type Cloudprojectdatabase_EngineEnum string
const (
	Cloudprojectdatabase_EngineEnum_cassandra Cloudprojectdatabase_EngineEnum = "cassandra"
	Cloudprojectdatabase_EngineEnum_grafana  = "grafana"
	Cloudprojectdatabase_EngineEnum_kafka  = "kafka"
	Cloudprojectdatabase_EngineEnum_kafkaConnect  = "kafkaConnect"
	Cloudprojectdatabase_EngineEnum_kafkaMirrorMaker  = "kafkaMirrorMaker"
	Cloudprojectdatabase_EngineEnum_m3aggregator  = "m3aggregator"
	Cloudprojectdatabase_EngineEnum_m3db  = "m3db"
	Cloudprojectdatabase_EngineEnum_mongodb  = "mongodb"
	Cloudprojectdatabase_EngineEnum_mysql  = "mysql"
	Cloudprojectdatabase_EngineEnum_opensearch  = "opensearch"
	Cloudprojectdatabase_EngineEnum_postgresql  = "postgresql"
	Cloudprojectdatabase_EngineEnum_redis  = "redis"
)

type Cloudprojectdatabase_IpRestriction struct {
	Description string                          `json:"description"`
	Ip          string                          `json:"ip"`
	Status      Cloudprojectdatabase_StatusEnum `json:"status"`
}

type Cloudprojectdatabase_IpRestrictionCreation struct {
	Description string `json:"description"`
	Ip          string `json:"ip"`
}

type Cloudprojectdatabase_NetworkTypeEnum string
const (
	Cloudprojectdatabase_NetworkTypeEnum_private Cloudprojectdatabase_NetworkTypeEnum = "private"
	Cloudprojectdatabase_NetworkTypeEnum_public  = "public"
)

type Cloudprojectdatabase_Service struct {
	BackupTime        time.Time                                     `json:"backupTime"`
	CreatedAt         time.Time                                     `json:"createdAt"`
	Description       string                                        `json:"description"`
	Disk              Cloudprojectdatabaseservice_Disk              `json:"disk"`
	Domain            string                                        `json:"domain"`
	Endpoints         []Cloudprojectdatabaseservice_Endpoint        `json:"endpoints"`
	Engine            Cloudprojectdatabase_EngineEnum               `json:"engine"`
	Flavor            string                                        `json:"flavor"`
	Id                string                                        `json:"id"`
	MaintenanceTime   time.Time                                     `json:"maintenanceTime"`
	MaintenanceWindow Cloudprojectdatabaseservice_MaintenanceWindow `json:"maintenanceWindow"`
	NetworkId         string                                        `json:"networkId"`
	NetworkType       Cloudprojectdatabase_NetworkTypeEnum          `json:"networkType"`
	NodeNumber        int64                                         `json:"nodeNumber"`
	Plan              string                                        `json:"plan"`
	Port              int64                                         `json:"port"`
	SslMode           string                                        `json:"sslMode"`
	Status            Cloudprojectdatabase_StatusEnum               `json:"status"`
	SubnetId          string                                        `json:"subnetId"`
	Uri               string                                        `json:"uri"`
	Version           string                                        `json:"version"`
}

type Cloudprojectdatabase_ServiceCreation struct {
	Backup          Cloudprojectdatabaseservicecreation_BackupFork `json:"backup"`
	BackupTime      time.Time                                      `json:"backupTime"`
	Description     string                                         `json:"description"`
	Disk            Cloudprojectdatabaseservice_Disk               `json:"disk"`
	MaintenanceTime time.Time                                      `json:"maintenanceTime"`
	NetworkId       string                                         `json:"networkId"`
	NodesList       []Cloudprojectdatabaseservice_NodeCreation     `json:"nodesList"`
	NodesPattern    Cloudprojectdatabaseservice_NodePattern        `json:"nodesPattern"`
	Plan            string                                         `json:"plan"`
	SubnetId        string                                         `json:"subnetId"`
	Version         string                                         `json:"version"`
}

type Cloudprojectdatabase_StatusEnum string
const (
	Cloudprojectdatabase_StatusEnum_CREATING Cloudprojectdatabase_StatusEnum = "CREATING"
	Cloudprojectdatabase_StatusEnum_DELETING  = "DELETING"
	Cloudprojectdatabase_StatusEnum_ERROR  = "ERROR"
	Cloudprojectdatabase_StatusEnum_ERROR_INCONSISTENT_SPEC  = "ERROR_INCONSISTENT_SPEC"
	Cloudprojectdatabase_StatusEnum_LOCKED  = "LOCKED"
	Cloudprojectdatabase_StatusEnum_LOCKED_PENDING  = "LOCKED_PENDING"
	Cloudprojectdatabase_StatusEnum_LOCKED_UPDATING  = "LOCKED_UPDATING"
	Cloudprojectdatabase_StatusEnum_PENDING  = "PENDING"
	Cloudprojectdatabase_StatusEnum_READY  = "READY"
	Cloudprojectdatabase_StatusEnum_UPDATING  = "UPDATING"
)

type Cloudprojectdatabase_TemporaryWriteDeadline struct {
	Until time.Time `json:"until"`
}

type Cloudprojectdatabase_TypeEnum string
const (
	Cloudprojectdatabase_TypeEnum_boolean Cloudprojectdatabase_TypeEnum = "boolean"
	Cloudprojectdatabase_TypeEnum_double  = "double"
	Cloudprojectdatabase_TypeEnum_duration  = "duration"
	Cloudprojectdatabase_TypeEnum_long  = "long"
	Cloudprojectdatabase_TypeEnum_string  = "string"
)

type Cloudprojectdatabaseavailability_StatusEnum string
const (
	Cloudprojectdatabaseavailability_StatusEnum_BETA Cloudprojectdatabaseavailability_StatusEnum = "BETA"
	Cloudprojectdatabaseavailability_StatusEnum_DEPRECATED  = "DEPRECATED"
	Cloudprojectdatabaseavailability_StatusEnum_STABLE  = "STABLE"
)

type Cloudprojectdatabasecapabilities_Engine struct {
	DefaultVersion string   `json:"defaultVersion"`
	Description    string   `json:"description"`
	Name           string   `json:"name"`
	SslModes       []string `json:"sslModes"`
	Versions       []string `json:"versions"`
}

type Cloudprojectdatabasecapabilities_Flavor struct {
	Core    int64  `json:"core"`
	Memory  int64  `json:"memory"`
	Name    string `json:"name"`
	Storage int64  `json:"storage"`
}

type Cloudprojectdatabasecapabilities_Integration struct {
	DestinationEngine Cloudprojectdatabase_EngineEnum                         `json:"destinationEngine"`
	Parameters        []Cloudprojectdatabasecapabilitiesintegration_Parameter `json:"parameters"`
	SourceEngine      Cloudprojectdatabase_EngineEnum                         `json:"sourceEngine"`
	Type              Cloudprojectdatabaseserviceintegration_TypeEnum         `json:"type"`
}

type Cloudprojectdatabasecapabilities_Option struct {
	Name string                        `json:"name"`
	Type Cloudprojectdatabase_TypeEnum `json:"type"`
}

type Cloudprojectdatabasecapabilities_Plan struct {
	BackupRetention time.Duration `json:"backupRetention"`
	Description     string        `json:"description"`
	Name            string        `json:"name"`
}

type CloudprojectdatabasecapabilitiesadvancedConfiguration_Property struct {
	Description string                                                                 `json:"description"`
	Maximum     float64                                                                `json:"maximum"`
	Minimum     float64                                                                `json:"minimum"`
	Name        string                                                                 `json:"name"`
	Type        CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum `json:"type"`
	Values      []string                                                               `json:"values"`
}

type CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum string
const (
	CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum_boolean CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum = "boolean"
	CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum_double  = "double"
	CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum_long  = "long"
	CloudprojectdatabasecapabilitiesadvancedConfigurationproperty_TypeEnum_string  = "string"
)

type Cloudprojectdatabasecapabilitiesintegration_Parameter struct {
	Name string                                                        `json:"name"`
	Type Cloudprojectdatabasecapabilitiesintegrationparameter_TypeEnum `json:"type"`
}

type Cloudprojectdatabasecapabilitiesintegrationparameter_TypeEnum string
const (
	Cloudprojectdatabasecapabilitiesintegrationparameter_TypeEnum_integer Cloudprojectdatabasecapabilitiesintegrationparameter_TypeEnum = "integer"
	Cloudprojectdatabasecapabilitiesintegrationparameter_TypeEnum_string  = "string"
)

type Cloudprojectdatabasekafka_Acl struct {
	Id         string `json:"id"`
	Permission string `json:"permission"`
	Topic      string `json:"topic"`
	Username   string `json:"username"`
}

type Cloudprojectdatabasekafka_Permissions struct {
	Names []string `json:"names"`
}

type Cloudprojectdatabasekafka_Service struct {
	BackupTime        time.Time                                     `json:"backupTime"`
	CreatedAt         time.Time                                     `json:"createdAt"`
	Description       string                                        `json:"description"`
	Disk              Cloudprojectdatabaseservice_Disk              `json:"disk"`
	Domain            string                                        `json:"domain"`
	Endpoints         []Cloudprojectdatabaseservice_Endpoint        `json:"endpoints"`
	Engine            Cloudprojectdatabase_EngineEnum               `json:"engine"`
	Flavor            string                                        `json:"flavor"`
	Id                string                                        `json:"id"`
	MaintenanceTime   time.Time                                     `json:"maintenanceTime"`
	MaintenanceWindow Cloudprojectdatabaseservice_MaintenanceWindow `json:"maintenanceWindow"`
	NetworkId         string                                        `json:"networkId"`
	NetworkType       Cloudprojectdatabase_NetworkTypeEnum          `json:"networkType"`
	NodeNumber        int64                                         `json:"nodeNumber"`
	Plan              string                                        `json:"plan"`
	Port              int64                                         `json:"port"`
	RestApi           bool                                          `json:"restApi"`
	SslMode           string                                        `json:"sslMode"`
	Status            Cloudprojectdatabase_StatusEnum               `json:"status"`
	SubnetId          string                                        `json:"subnetId"`
	Uri               string                                        `json:"uri"`
	Version           string                                        `json:"version"`
}

type Cloudprojectdatabasekafka_Topic struct {
	Id                string `json:"id"`
	MinInsyncReplicas int64  `json:"minInsyncReplicas"`
	Name              string `json:"name"`
	Partitions        int64  `json:"partitions"`
	Replication       int64  `json:"replication"`
	RetentionBytes    int64  `json:"retentionBytes"`
	RetentionHours    int64  `json:"retentionHours"`
}

type Cloudprojectdatabasekafka_TopicCreation struct {
	Id                string `json:"id"`
	MinInsyncReplicas int64  `json:"minInsyncReplicas"`
	Name              string `json:"name"`
	Partitions        int64  `json:"partitions"`
	Replication       int64  `json:"replication"`
	RetentionBytes    int64  `json:"retentionBytes"`
	RetentionHours    int64  `json:"retentionHours"`
}

type Cloudprojectdatabasekafkauser_Access struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type CloudprojectdatabasekafkaConnect_Connector struct {
	Configuration map[string]string                                    `json:"configuration"`
	ConnectorId   string                                               `json:"connectorId"`
	Id            string                                               `json:"id"`
	Name          string                                               `json:"name"`
	Status        CloudprojectdatabasekafkaConnectconnector_StatusEnum `json:"status"`
}

type CloudprojectdatabasekafkaConnect_ConnectorCreation struct {
	Configuration map[string]string `json:"configuration"`
	ConnectorId   string            `json:"connectorId"`
	Name          string            `json:"name"`
}

type CloudprojectdatabasekafkaConnectcapabilities_Connector struct {
	Author           string                                                         `json:"author"`
	DocumentationUrl string                                                         `json:"documentationUrl"`
	Id               string                                                         `json:"id"`
	Latest           bool                                                           `json:"latest"`
	Name             string                                                         `json:"name"`
	Preview          bool                                                           `json:"preview"`
	Type             CloudprojectdatabasekafkaConnectcapabilitiesconnector_TypeEnum `json:"type"`
	Version          string                                                         `json:"version"`
}

type CloudprojectdatabasekafkaConnectcapabilitiesconnector_Transform struct {
	Description   string                                                     `json:"description"`
	DisplayName   string                                                     `json:"displayName"`
	Name          string                                                     `json:"name"`
	Required      bool                                                       `json:"required"`
	TransformType string                                                     `json:"transformType"`
	Type          CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum `json:"type"`
	Values        []string                                                   `json:"values"`
}

type CloudprojectdatabasekafkaConnectcapabilitiesconnector_TypeEnum string
const (
	CloudprojectdatabasekafkaConnectcapabilitiesconnector_TypeEnum_sink CloudprojectdatabasekafkaConnectcapabilitiesconnector_TypeEnum = "sink"
	CloudprojectdatabasekafkaConnectcapabilitiesconnector_TypeEnum_source  = "source"
)

type CloudprojectdatabasekafkaConnectcapabilitiesconnectorconfiguration_Property struct {
	DefaultValue string                                                                       `json:"defaultValue"`
	Description  string                                                                       `json:"description"`
	DisplayName  string                                                                       `json:"displayName"`
	Group        string                                                                       `json:"group"`
	Importance   CloudprojectdatabasekafkaConnectcapabilitiesconnectorproperty_ImportanceEnum `json:"importance"`
	Name         string                                                                       `json:"name"`
	Required     bool                                                                         `json:"required"`
	Type         CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum                   `json:"type"`
	Values       []string                                                                     `json:"values"`
}

type CloudprojectdatabasekafkaConnectcapabilitiesconnectorproperty_ImportanceEnum string
const (
	CloudprojectdatabasekafkaConnectcapabilitiesconnectorproperty_ImportanceEnum_high CloudprojectdatabasekafkaConnectcapabilitiesconnectorproperty_ImportanceEnum = "high"
	CloudprojectdatabasekafkaConnectcapabilitiesconnectorproperty_ImportanceEnum_low  = "low"
	CloudprojectdatabasekafkaConnectcapabilitiesconnectorproperty_ImportanceEnum_medium  = "medium"
)

type CloudprojectdatabasekafkaConnectconnector_StatusEnum string
const (
	CloudprojectdatabasekafkaConnectconnector_StatusEnum_CREATING CloudprojectdatabasekafkaConnectconnector_StatusEnum = "CREATING"
	CloudprojectdatabasekafkaConnectconnector_StatusEnum_FAILED  = "FAILED"
	CloudprojectdatabasekafkaConnectconnector_StatusEnum_PAUSED  = "PAUSED"
	CloudprojectdatabasekafkaConnectconnector_StatusEnum_RUNNING  = "RUNNING"
	CloudprojectdatabasekafkaConnectconnector_StatusEnum_UNASSIGNED  = "UNASSIGNED"
)

type CloudprojectdatabasekafkaConnectconnector_Task struct {
	Id     int64                                                    `json:"id"`
	Status CloudprojectdatabasekafkaConnectconnectortask_StatusEnum `json:"status"`
	Trace  string                                                   `json:"trace"`
}

type CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum string
const (
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_boolean CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum = "boolean"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_class  = "class"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_double  = "double"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_int16  = "int16"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_int32  = "int32"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_int64  = "int64"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_list  = "list"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_password  = "password"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_string  = "string"
	CloudprojectdatabasekafkaConnectconnectorproperty_TypeEnum_transform  = "transform"
)

type CloudprojectdatabasekafkaConnectconnectortask_StatusEnum string
const (
	CloudprojectdatabasekafkaConnectconnectortask_StatusEnum_FAILED CloudprojectdatabasekafkaConnectconnectortask_StatusEnum = "FAILED"
	CloudprojectdatabasekafkaConnectconnectortask_StatusEnum_PAUSED  = "PAUSED"
	CloudprojectdatabasekafkaConnectconnectortask_StatusEnum_RUNNING  = "RUNNING"
)

type Cloudprojectdatabasem3db_Namespace struct {
	Id                       string                                      `json:"id"`
	Name                     string                                      `json:"name"`
	Resolution               time.Duration                               `json:"resolution"`
	Retention                Cloudprojectdatabasem3dbnamespace_Retention `json:"retention"`
	SnapshotEnabled          bool                                        `json:"snapshotEnabled"`
	Type                     Cloudprojectdatabasem3dbnamespace_TypeEnum  `json:"type"`
	WritesToCommitLogEnabled bool                                        `json:"writesToCommitLogEnabled"`
}

type Cloudprojectdatabasem3db_NamespaceCreation struct {
	Id                       string                                      `json:"id"`
	Name                     string                                      `json:"name"`
	Resolution               time.Duration                               `json:"resolution"`
	Retention                Cloudprojectdatabasem3dbnamespace_Retention `json:"retention"`
	SnapshotEnabled          bool                                        `json:"snapshotEnabled"`
	Type                     Cloudprojectdatabasem3dbnamespace_TypeEnum  `json:"type"`
	WritesToCommitLogEnabled bool                                        `json:"writesToCommitLogEnabled"`
}

type Cloudprojectdatabasem3db_User struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Group     string                          `json:"group"`
	Id        string                          `json:"id"`
	Status    Cloudprojectdatabase_StatusEnum `json:"status"`
	Username  string                          `json:"username"`
}

type Cloudprojectdatabasem3db_UserCreation struct {
	Group string `json:"group"`
	Name  string `json:"name"`
}

type Cloudprojectdatabasem3db_UserWithPassword struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Group     string                          `json:"group"`
	Id        string                          `json:"id"`
	Password  string                          `json:"password"`
	Status    Cloudprojectdatabase_StatusEnum `json:"status"`
	Username  string                          `json:"username"`
}

type Cloudprojectdatabasem3dbnamespace_Retention struct {
	BlockDataExpirationDuration time.Duration `json:"blockDataExpirationDuration"`
	BlockSizeDuration           time.Duration `json:"blockSizeDuration"`
	BufferFutureDuration        time.Duration `json:"bufferFutureDuration"`
	BufferPastDuration          time.Duration `json:"bufferPastDuration"`
	PeriodDuration              time.Duration `json:"periodDuration"`
}

type Cloudprojectdatabasem3dbnamespace_TypeEnum string
const (
	Cloudprojectdatabasem3dbnamespace_TypeEnum_aggregated Cloudprojectdatabasem3dbnamespace_TypeEnum = "aggregated"
	Cloudprojectdatabasem3dbnamespace_TypeEnum_unaggregated  = "unaggregated"
)

type Cloudprojectdatabasemysql_QueryStatistics struct {
	Queries []Cloudprojectdatabasemysqlquerystatistics_Query `json:"queries"`
}

type Cloudprojectdatabasemysqlquerystatistics_Query struct {
	AvgTimerWait            int64     `json:"avgTimerWait"`
	CountStar               int64     `json:"countStar"`
	Digest                  string    `json:"digest"`
	DigestText              string    `json:"digestText"`
	FirstSeen               time.Time `json:"firstSeen"`
	LastSeen                time.Time `json:"lastSeen"`
	MaxTimerWait            int64     `json:"maxTimerWait"`
	MinTimerWait            int64     `json:"minTimerWait"`
	Quantile95              int64     `json:"quantile95"`
	Quantile99              int64     `json:"quantile99"`
	Quantile999             int64     `json:"quantile999"`
	QuerySampleSeen         time.Time `json:"querySampleSeen"`
	QuerySampleText         string    `json:"querySampleText"`
	QuerySampleTimerWait    int64     `json:"querySampleTimerWait"`
	SchemaName              string    `json:"schemaName"`
	SumCreatedTmpDiskTables int64     `json:"sumCreatedTmpDiskTables"`
	SumCreatedTmpTables     int64     `json:"sumCreatedTmpTables"`
	SumErrors               int64     `json:"sumErrors"`
	SumLockTime             int64     `json:"sumLockTime"`
	SumNoGoodIndexUsed      int64     `json:"sumNoGoodIndexUsed"`
	SumNoIndexUsed          int64     `json:"sumNoIndexUsed"`
	SumRowsAffected         int64     `json:"sumRowsAffected"`
	SumRowsExamined         int64     `json:"sumRowsExamined"`
	SumRowsSent             int64     `json:"sumRowsSent"`
	SumSelectFullJoin       int64     `json:"sumSelectFullJoin"`
	SumSelectFullRangeJoin  int64     `json:"sumSelectFullRangeJoin"`
	SumSelectRange          int64     `json:"sumSelectRange"`
	SumSelectRangeCheck     int64     `json:"sumSelectRangeCheck"`
	SumSelectScan           int64     `json:"sumSelectScan"`
	SumSortMergePasses      int64     `json:"sumSortMergePasses"`
	SumSortRange            int64     `json:"sumSortRange"`
	SumSortRows             int64     `json:"sumSortRows"`
	SumSortScan             int64     `json:"sumSortScan"`
	SumTimerWait            int64     `json:"sumTimerWait"`
	SumWarnings             int64     `json:"sumWarnings"`
}

type Cloudprojectdatabaseopensearch_Index struct {
	CreatedAt      time.Time `json:"createdAt"`
	Documents      int64     `json:"documents"`
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	ReplicasNumber int64     `json:"replicasNumber"`
	ShardsNumber   int64     `json:"shardsNumber"`
	Size           int64     `json:"size"`
}

type Cloudprojectdatabaseopensearch_Pattern struct {
	Id            string `json:"id"`
	MaxIndexCount int64  `json:"maxIndexCount"`
	Pattern       string `json:"pattern"`
}

type Cloudprojectdatabaseopensearch_Permissions struct {
	Names []string `json:"names"`
}

type Cloudprojectdatabaseopensearch_Service struct {
	AclsEnabled       bool                                                 `json:"aclsEnabled"`
	AdditionalUris    Cloudprojectdatabaseopensearchservice_AdditionalUris `json:"additionalUris"`
	BackupTime        time.Time                                            `json:"backupTime"`
	CreatedAt         time.Time                                            `json:"createdAt"`
	Description       string                                               `json:"description"`
	Disk              Cloudprojectdatabaseservice_Disk                     `json:"disk"`
	Domain            string                                               `json:"domain"`
	Endpoints         []Cloudprojectdatabaseservice_Endpoint               `json:"endpoints"`
	Engine            Cloudprojectdatabase_EngineEnum                      `json:"engine"`
	Flavor            string                                               `json:"flavor"`
	Id                string                                               `json:"id"`
	MaintenanceTime   time.Time                                            `json:"maintenanceTime"`
	MaintenanceWindow Cloudprojectdatabaseservice_MaintenanceWindow        `json:"maintenanceWindow"`
	NetworkId         string                                               `json:"networkId"`
	NetworkType       Cloudprojectdatabase_NetworkTypeEnum                 `json:"networkType"`
	NodeNumber        int64                                                `json:"nodeNumber"`
	Plan              string                                               `json:"plan"`
	Port              int64                                                `json:"port"`
	SslMode           string                                               `json:"sslMode"`
	Status            Cloudprojectdatabase_StatusEnum                      `json:"status"`
	SubnetId          string                                               `json:"subnetId"`
	Uri               string                                               `json:"uri"`
	Version           string                                               `json:"version"`
}

type Cloudprojectdatabaseopensearch_User struct {
	Acls      []Cloudprojectdatabaseopensearch_UserAcl `json:"acls"`
	CreatedAt time.Time                                `json:"createdAt"`
	Id        string                                   `json:"id"`
	Status    Cloudprojectdatabase_StatusEnum          `json:"status"`
	Username  string                                   `json:"username"`
}

type Cloudprojectdatabaseopensearch_UserAcl struct {
	Pattern    string `json:"pattern"`
	Permission string `json:"permission"`
}

type Cloudprojectdatabaseopensearch_UserCreation struct {
	Acls []Cloudprojectdatabaseopensearch_UserAcl `json:"acls"`
	Name string                                   `json:"name"`
}

type Cloudprojectdatabaseopensearch_UserWithPassword struct {
	Acls      []Cloudprojectdatabaseopensearch_UserAcl `json:"acls"`
	CreatedAt time.Time                                `json:"createdAt"`
	Id        string                                   `json:"id"`
	Password  string                                   `json:"password"`
	Status    Cloudprojectdatabase_StatusEnum          `json:"status"`
	Username  string                                   `json:"username"`
}

type Cloudprojectdatabaseopensearchservice_AdditionalUris struct {
	Kibana string `json:"kibana"`
}

type Cloudprojectdatabasepostgresql_ConnectionPool struct {
	DatabaseId string                                                   `json:"databaseId"`
	Id         string                                                   `json:"id"`
	Mode       Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum    `json:"mode"`
	Name       string                                                   `json:"name"`
	Port       int64                                                    `json:"port"`
	Size       int64                                                    `json:"size"`
	SslMode    Cloudprojectdatabasepostgresqlconnectionpool_SslModeEnum `json:"sslMode"`
	Uri        string                                                   `json:"uri"`
	UserId     string                                                   `json:"userId"`
}

type Cloudprojectdatabasepostgresql_ConnectionPoolCreation struct {
	DatabaseId string                                                `json:"databaseId"`
	Mode       Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum `json:"mode"`
	Name       string                                                `json:"name"`
	Size       int64                                                 `json:"size"`
	UserId     string                                                `json:"userId"`
}

type Cloudprojectdatabasepostgresql_QueryStatistics struct {
	Queries []Cloudprojectdatabasepostgresqlquerystatistics_Query `json:"queries"`
}

type Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum string
const (
	Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum_session Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum = "session"
	Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum_statement  = "statement"
	Cloudprojectdatabasepostgresqlconnectionpool_ModeEnum_transaction  = "transaction"
)

type Cloudprojectdatabasepostgresqlconnectionpool_SslModeEnum string
const (
	Cloudprojectdatabasepostgresqlconnectionpool_SslModeEnum_require Cloudprojectdatabasepostgresqlconnectionpool_SslModeEnum = "require"
)

type Cloudprojectdatabasepostgresqlquerystatistics_Query struct {
	BlkReadTime       float64                 `json:"blkReadTime"`
	BlkWriteTime      float64                 `json:"blkWriteTime"`
	Calls             int64                   `json:"calls"`
	DatabaseName      string                  `json:"databaseName"`
	LocalBlksDirtied  int64                   `json:"localBlksDirtied"`
	LocalBlksHit      int64                   `json:"localBlksHit"`
	LocalBlksRead     int64                   `json:"localBlksRead"`
	LocalBlksWritten  int64                   `json:"localBlksWritten"`
	MaxPlanTime       float64                 `json:"maxPlanTime"`
	MaxTime           float64                 `json:"maxTime"`
	MeanPlanTime      float64                 `json:"meanPlanTime"`
	MeanTime          float64                 `json:"meanTime"`
	MinPlanTime       float64                 `json:"minPlanTime"`
	MinTime           float64                 `json:"minTime"`
	Query             string                  `json:"query"`
	Rows              int64                   `json:"rows"`
	SharedBlksDirtied int64                   `json:"sharedBlksDirtied"`
	SharedBlksHit     int64                   `json:"sharedBlksHit"`
	SharedBlksRead    int64                   `json:"sharedBlksRead"`
	SharedBlksWritten int64                   `json:"sharedBlksWritten"`
	StddevPlanTime    float64                 `json:"stddevPlanTime"`
	StddevTime        float64                 `json:"stddevTime"`
	TempBlksRead      int64                   `json:"tempBlksRead"`
	TempBlksWritten   int64                   `json:"tempBlksWritten"`
	TotalPlanTime     float64                 `json:"totalPlanTime"`
	TotalTime         float64                 `json:"totalTime"`
	Username          string                  `json:"username"`
	WalBytes          types.UnitAndValueInt64 `json:"walBytes"`
	WalFpi            int64                   `json:"walFpi"`
	WalRecords        int64                   `json:"walRecords"`
}

type Cloudprojectdatabaseredis_User struct {
	Categories []string                        `json:"categories"`
	Channels   []string                        `json:"channels"`
	Commands   []string                        `json:"commands"`
	CreatedAt  time.Time                       `json:"createdAt"`
	Id         string                          `json:"id"`
	Keys       []string                        `json:"keys"`
	Status     Cloudprojectdatabase_StatusEnum `json:"status"`
	Username   string                          `json:"username"`
}

type Cloudprojectdatabaseredis_UserCreation struct {
	Categories []string `json:"categories"`
	Channels   []string `json:"channels"`
	Commands   []string `json:"commands"`
	Keys       []string `json:"keys"`
	Name       string   `json:"name"`
}

type Cloudprojectdatabaseredis_UserWithPassword struct {
	Categories []string                        `json:"categories"`
	Channels   []string                        `json:"channels"`
	Commands   []string                        `json:"commands"`
	CreatedAt  time.Time                       `json:"createdAt"`
	Id         string                          `json:"id"`
	Keys       []string                        `json:"keys"`
	Password   string                          `json:"password"`
	Status     Cloudprojectdatabase_StatusEnum `json:"status"`
	Username   string                          `json:"username"`
}

type Cloudprojectdatabaseservice_Backup struct {
	CreatedAt   time.Time                           `json:"createdAt"`
	Description string                              `json:"description"`
	Id          string                              `json:"id"`
	Region      string                              `json:"region"`
	Size        types.UnitAndValueInt64             `json:"size"`
	Status      Cloudprojectdatabase_StatusEnum     `json:"status"`
	Type        Cloudprojectdatabase_BackupTypeEnum `json:"type"`
}

type Cloudprojectdatabaseservice_Certificates struct {
	Ca string `json:"ca"`
}

type Cloudprojectdatabaseservice_CurrentQueries struct {
	Queries []Cloudprojectdatabaseservicecurrentqueries_Query `json:"queries"`
}

type Cloudprojectdatabaseservice_DataPoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type Cloudprojectdatabaseservice_Database struct {
	Default bool   `json:"default"`
	Id      string `json:"id"`
	Name    string `json:"name"`
}

type Cloudprojectdatabaseservice_Disk struct {
	Size int64  `json:"size"`
	Type string `json:"type"`
}

type Cloudprojectdatabaseservice_Endpoint struct {
	Component Cloudprojectdatabaseserviceendpoint_ComponentEnum `json:"component"`
	Domain    string                                            `json:"domain"`
	Path      string                                            `json:"path"`
	Port      int64                                             `json:"port"`
	Scheme    string                                            `json:"scheme"`
	Ssl       bool                                              `json:"ssl"`
	SslMode   string                                            `json:"sslMode"`
	Uri       string                                            `json:"uri"`
}

type Cloudprojectdatabaseservice_HostMetric struct {
	DataPoints []Cloudprojectdatabaseservice_DataPoint `json:"dataPoints"`
	Hostname   string                                  `json:"hostname"`
}

type Cloudprojectdatabaseservice_Integration struct {
	DestinationServiceId string                                            `json:"destinationServiceId"`
	Id                   string                                            `json:"id"`
	Parameters           map[string]string                                 `json:"parameters"`
	SourceServiceId      string                                            `json:"sourceServiceId"`
	Status               Cloudprojectdatabaseserviceintegration_StatusEnum `json:"status"`
	Type                 Cloudprojectdatabaseserviceintegration_TypeEnum   `json:"type"`
}

type Cloudprojectdatabaseservice_LogEntry struct {
	Hostname  string `json:"hostname"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type Cloudprojectdatabaseservice_Maintenance struct {
	AppliedAt   time.Time                                         `json:"appliedAt"`
	Description string                                            `json:"description"`
	Id          string                                            `json:"id"`
	ScheduledAt time.Time                                         `json:"scheduledAt"`
	Status      Cloudprojectdatabaseservicemaintenance_StatusEnum `json:"status"`
}

type Cloudprojectdatabaseservice_MaintenanceWindow struct {
	End   time.Time `json:"end"`
	Start time.Time `json:"start"`
}

type Cloudprojectdatabaseservice_Metric struct {
	Metrics []Cloudprojectdatabaseservice_HostMetric   `json:"metrics"`
	Name    string                                     `json:"name"`
	Units   Cloudprojectdatabaseservice_MetricUnitEnum `json:"units"`
}

type Cloudprojectdatabaseservice_MetricPeriodEnum string
const (
	Cloudprojectdatabaseservice_MetricPeriodEnum_lastDay Cloudprojectdatabaseservice_MetricPeriodEnum = "lastDay"
	Cloudprojectdatabaseservice_MetricPeriodEnum_lastHour  = "lastHour"
	Cloudprojectdatabaseservice_MetricPeriodEnum_lastMonth  = "lastMonth"
	Cloudprojectdatabaseservice_MetricPeriodEnum_lastWeek  = "lastWeek"
	Cloudprojectdatabaseservice_MetricPeriodEnum_lastYear  = "lastYear"
)

type Cloudprojectdatabaseservice_MetricUnitEnum string
const (
	Cloudprojectdatabaseservice_MetricUnitEnum_BYTES Cloudprojectdatabaseservice_MetricUnitEnum = "BYTES"
	Cloudprojectdatabaseservice_MetricUnitEnum_BYTES_PER_SECOND  = "BYTES_PER_SECOND"
	Cloudprojectdatabaseservice_MetricUnitEnum_GIGABYTES  = "GIGABYTES"
	Cloudprojectdatabaseservice_MetricUnitEnum_GIGABYTES_PER_HOUR  = "GIGABYTES_PER_HOUR"
	Cloudprojectdatabaseservice_MetricUnitEnum_MEGABYTES  = "MEGABYTES"
	Cloudprojectdatabaseservice_MetricUnitEnum_MEGABYTES_PER_SECOND  = "MEGABYTES_PER_SECOND"
	Cloudprojectdatabaseservice_MetricUnitEnum_MILLISECONDS  = "MILLISECONDS"
	Cloudprojectdatabaseservice_MetricUnitEnum_PERCENT  = "PERCENT"
	Cloudprojectdatabaseservice_MetricUnitEnum_SCALAR  = "SCALAR"
	Cloudprojectdatabaseservice_MetricUnitEnum_SCALAR_PER_SECOND  = "SCALAR_PER_SECOND"
	Cloudprojectdatabaseservice_MetricUnitEnum_SECONDS  = "SECONDS"
	Cloudprojectdatabaseservice_MetricUnitEnum_UNKNOWN  = "UNKNOWN"
)

type Cloudprojectdatabaseservice_Node struct {
	CreatedAt time.Time                                `json:"createdAt"`
	Flavor    string                                   `json:"flavor"`
	Id        string                                   `json:"id"`
	Name      string                                   `json:"name"`
	Port      int64                                    `json:"port"`
	Region    string                                   `json:"region"`
	Role      Cloudprojectdatabaseservicenode_RoleEnum `json:"role"`
	Status    Cloudprojectdatabase_StatusEnum          `json:"status"`
}

type Cloudprojectdatabaseservice_NodeCreation struct {
	CreatedAt time.Time                                `json:"createdAt"`
	Flavor    string                                   `json:"flavor"`
	Id        string                                   `json:"id"`
	Name      string                                   `json:"name"`
	Port      int64                                    `json:"port"`
	Region    string                                   `json:"region"`
	Role      Cloudprojectdatabaseservicenode_RoleEnum `json:"role"`
	Status    Cloudprojectdatabase_StatusEnum          `json:"status"`
}

type Cloudprojectdatabaseservice_NodePattern struct {
	Flavor string `json:"flavor"`
	Number int64  `json:"number"`
	Region string `json:"region"`
}

type Cloudprojectdatabaseservice_Replication struct {
	EmitHeartbeats         bool                                                   `json:"emitHeartbeats"`
	Enabled                bool                                                   `json:"enabled"`
	Id                     string                                                 `json:"id"`
	ReplicationPolicyClass Cloudprojectdatabaseservicereplication_PolicyClassEnum `json:"replicationPolicyClass"`
	SourceIntegration      string                                                 `json:"sourceIntegration"`
	SyncGroupOffsets       bool                                                   `json:"syncGroupOffsets"`
	SyncInterval           int64                                                  `json:"syncInterval"`
	TargetIntegration      string                                                 `json:"targetIntegration"`
	TopicExcludeList       []string                                               `json:"topicExcludeList"`
	Topics                 []string                                               `json:"topics"`
}

type Cloudprojectdatabaseservice_ReplicationCreation struct {
	EmitHeartbeats         bool                                                   `json:"emitHeartbeats"`
	Enabled                bool                                                   `json:"enabled"`
	ReplicationPolicyClass Cloudprojectdatabaseservicereplication_PolicyClassEnum `json:"replicationPolicyClass"`
	SourceIntegration      string                                                 `json:"sourceIntegration"`
	SyncGroupOffsets       bool                                                   `json:"syncGroupOffsets"`
	SyncInterval           int64                                                  `json:"syncInterval"`
	TargetIntegration      string                                                 `json:"targetIntegration"`
	TopicExcludeList       []string                                               `json:"topicExcludeList"`
	Topics                 []string                                               `json:"topics"`
}

type Cloudprojectdatabaseservice_Restore struct {
	PointInTime time.Time `json:"pointInTime"`
}

type Cloudprojectdatabaseservice_User struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Id        string                          `json:"id"`
	Status    Cloudprojectdatabase_StatusEnum `json:"status"`
	Username  string                          `json:"username"`
}

type Cloudprojectdatabaseservice_UserCreation struct {
	Name string `json:"name"`
}

type Cloudprojectdatabaseservice_UserWithPassword struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Id        string                          `json:"id"`
	Password  string                          `json:"password"`
	Status    Cloudprojectdatabase_StatusEnum `json:"status"`
	Username  string                          `json:"username"`
}

type Cloudprojectdatabaseservice_UserWithPasswordAndRoles struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Id        string                          `json:"id"`
	Password  string                          `json:"password"`
	Roles     []string                        `json:"roles"`
	Status    Cloudprojectdatabase_StatusEnum `json:"status"`
	Username  string                          `json:"username"`
}

type Cloudprojectdatabaseservice_UserWithRoles struct {
	CreatedAt time.Time                       `json:"createdAt"`
	Id        string                          `json:"id"`
	Roles     []string                        `json:"roles"`
	Status    Cloudprojectdatabase_StatusEnum `json:"status"`
	Username  string                          `json:"username"`
}

type Cloudprojectdatabaseservice_UserWithRolesCreation struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type Cloudprojectdatabaseservicecreation_BackupFork struct {
	Id          string    `json:"id"`
	PointInTime time.Time `json:"pointInTime"`
	ServiceId   string    `json:"serviceId"`
}

type Cloudprojectdatabaseservicecurrentqueries_Query struct {
	ApplicationName  string                                                      `json:"applicationName"`
	BackendStart     time.Time                                                   `json:"backendStart"`
	BackendType      string                                                      `json:"backendType"`
	BackendXid       int64                                                       `json:"backendXid"`
	BackendXmin      int64                                                       `json:"backendXmin"`
	ClientHostname   string                                                      `json:"clientHostname"`
	ClientIp         string                                                      `json:"clientIp"`
	ClientPort       int64                                                       `json:"clientPort"`
	DatabaseId       int64                                                       `json:"databaseId"`
	DatabaseName     string                                                      `json:"databaseName"`
	LeaderPid        int64                                                       `json:"leaderPid"`
	Pid              int64                                                       `json:"pid"`
	Query            string                                                      `json:"query"`
	QueryDuration    float64                                                     `json:"queryDuration"`
	QueryStart       time.Time                                                   `json:"queryStart"`
	State            Cloudprojectdatabaseservicecurrentqueries_StateEnum         `json:"state"`
	StateChange      time.Time                                                   `json:"stateChange"`
	TransactionStart time.Time                                                   `json:"transactionStart"`
	UserId           int64                                                       `json:"userId"`
	UserName         string                                                      `json:"userName"`
	WaitEvent        string                                                      `json:"waitEvent"`
	WaitEventType    Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum `json:"waitEventType"`
}

type Cloudprojectdatabaseservicecurrentqueries_StateEnum string
const (
	Cloudprojectdatabaseservicecurrentqueries_StateEnum_ACTIVE Cloudprojectdatabaseservicecurrentqueries_StateEnum = "ACTIVE"
	Cloudprojectdatabaseservicecurrentqueries_StateEnum_DISABLED  = "DISABLED"
	Cloudprojectdatabaseservicecurrentqueries_StateEnum_FASTPATH_FUNCTION_CALL  = "FASTPATH_FUNCTION_CALL"
	Cloudprojectdatabaseservicecurrentqueries_StateEnum_IDLE  = "IDLE"
	Cloudprojectdatabaseservicecurrentqueries_StateEnum_IDLE_IN_TRANSACTION  = "IDLE_IN_TRANSACTION"
	Cloudprojectdatabaseservicecurrentqueries_StateEnum_IDLE_IN_TRANSACTION_ABORTED  = "IDLE_IN_TRANSACTION_ABORTED"
)

type Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum string
const (
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_ACTIVITY Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum = "ACTIVITY"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_BUFFER_PIN  = "BUFFER_PIN"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_CLIENT  = "CLIENT"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_EXTENSION  = "EXTENSION"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_IO  = "IO"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_IPC  = "IPC"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_LOCK  = "LOCK"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_LWLOCK  = "LWLOCK"
	Cloudprojectdatabaseservicecurrentqueries_WaitEventTypeEnum_TIMEOUT  = "TIMEOUT"
)

type Cloudprojectdatabaseservicecurrentqueriesquery_CancelRequest struct {
	Pid       int64 `json:"pid"`
	Terminate bool  `json:"terminate"`
}

type Cloudprojectdatabaseservicecurrentqueriesquery_CancelResponse struct {
	Success bool `json:"success"`
}

type Cloudprojectdatabaseserviceendpoint_ComponentEnum string
const (
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_cassandra Cloudprojectdatabaseserviceendpoint_ComponentEnum = "cassandra"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_grafana  = "grafana"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_graphite  = "graphite"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_influxdb  = "influxdb"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_kafka  = "kafka"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_kafkaConnect  = "kafkaConnect"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_kafkaRestApi  = "kafkaRestApi"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_kafkaSASL  = "kafkaSASL"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_kibana  = "kibana"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_m3coordinator  = "m3coordinator"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mongodb  = "mongodb"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mongodbAnalytics  = "mongodbAnalytics"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mongodbSrv  = "mongodbSrv"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mongodbSrvAnalytics  = "mongodbSrvAnalytics"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mysql  = "mysql"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mysqlRead  = "mysqlRead"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_mysqlx  = "mysqlx"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_opensearch  = "opensearch"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_postgresql  = "postgresql"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_postgresqlRead  = "postgresqlRead"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_postgresqlReadReplica  = "postgresqlReadReplica"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_prometheusRead  = "prometheusRead"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_prometheusWrite  = "prometheusWrite"
	Cloudprojectdatabaseserviceendpoint_ComponentEnum_redis  = "redis"
)

type Cloudprojectdatabaseserviceintegration_StatusEnum string
const (
	Cloudprojectdatabaseserviceintegration_StatusEnum_READY Cloudprojectdatabaseserviceintegration_StatusEnum = "READY"
)

type Cloudprojectdatabaseserviceintegration_TypeEnum string
const (
	Cloudprojectdatabaseserviceintegration_TypeEnum_grafanaDashboard Cloudprojectdatabaseserviceintegration_TypeEnum = "grafanaDashboard"
	Cloudprojectdatabaseserviceintegration_TypeEnum_grafanaDatasource  = "grafanaDatasource"
	Cloudprojectdatabaseserviceintegration_TypeEnum_kafkaConnect  = "kafkaConnect"
	Cloudprojectdatabaseserviceintegration_TypeEnum_kafkaLogs  = "kafkaLogs"
	Cloudprojectdatabaseserviceintegration_TypeEnum_kafkaMirrorMaker  = "kafkaMirrorMaker"
	Cloudprojectdatabaseserviceintegration_TypeEnum_m3aggregator  = "m3aggregator"
	Cloudprojectdatabaseserviceintegration_TypeEnum_m3dbMetrics  = "m3dbMetrics"
	Cloudprojectdatabaseserviceintegration_TypeEnum_opensearchLogs  = "opensearchLogs"
	Cloudprojectdatabaseserviceintegration_TypeEnum_postgresqlMetrics  = "postgresqlMetrics"
)

type Cloudprojectdatabaseservicemaintenance_StatusEnum string
const (
	Cloudprojectdatabaseservicemaintenance_StatusEnum_APPLIED Cloudprojectdatabaseservicemaintenance_StatusEnum = "APPLIED"
	Cloudprojectdatabaseservicemaintenance_StatusEnum_APPLYING  = "APPLYING"
	Cloudprojectdatabaseservicemaintenance_StatusEnum_ERROR  = "ERROR"
	Cloudprojectdatabaseservicemaintenance_StatusEnum_SCHEDULED  = "SCHEDULED"
)

type Cloudprojectdatabaseservicenode_RoleEnum string
const (
	Cloudprojectdatabaseservicenode_RoleEnum_ANALYTICS Cloudprojectdatabaseservicenode_RoleEnum = "ANALYTICS"
	Cloudprojectdatabaseservicenode_RoleEnum_STANDARD  = "STANDARD"
)

type Cloudprojectdatabaseservicereplication_PolicyClassEnum string
const (
	Cloudprojectdatabaseservicereplication_PolicyClassEnum_orgapachekafkaconnectmirrorDefaultReplicationPolicy Cloudprojectdatabaseservicereplication_PolicyClassEnum = "org.apache.kafka.connect.mirror.DefaultReplicationPolicy"
	Cloudprojectdatabaseservicereplication_PolicyClassEnum_orgapachekafkaconnectmirrorIdentityReplicationPolicy  = "org.apache.kafka.connect.mirror.IdentityReplicationPolicy"
)

type CloudprojectfloatingIp_AssociatedEntity struct {
	GatewayId string                                          `json:"gatewayId"`
	Id        string                                          `json:"id"`
	Ip        string                                          `json:"ip"`
	Type      CloudprojectfloatingIpassociatedEntity_TypeEnum `json:"type"`
}

type CloudprojectfloatingIp_StatusEnum string
const (
	CloudprojectfloatingIp_StatusEnum_active CloudprojectfloatingIp_StatusEnum = "active"
	CloudprojectfloatingIp_StatusEnum_down  = "down"
	CloudprojectfloatingIp_StatusEnum_error  = "error"
)

type CloudprojectfloatingIpassociatedEntity_TypeEnum string
const (
	CloudprojectfloatingIpassociatedEntity_TypeEnum_dhcp CloudprojectfloatingIpassociatedEntity_TypeEnum = "dhcp"
	CloudprojectfloatingIpassociatedEntity_TypeEnum_instance  = "instance"
	CloudprojectfloatingIpassociatedEntity_TypeEnum_loadbalancer  = "loadbalancer"
	CloudprojectfloatingIpassociatedEntity_TypeEnum_routerInterface  = "routerInterface"
	CloudprojectfloatingIpassociatedEntity_TypeEnum_unknown  = "unknown"
)

type Cloudprojectio_Stream struct {
	Backlog     time.Duration                   `json:"backlog"`
	Description string                          `json:"description"`
	Id          string                          `json:"id"`
	Kind        Cloudprojectio_StreamKindEnum   `json:"kind"`
	Name        string                          `json:"name"`
	Regions     []string                        `json:"regions"`
	Retention   time.Duration                   `json:"retention"`
	Status      Cloudprojectio_StreamStatusEnum `json:"status"`
	Throttling  int64                           `json:"throttling"`
}

type Cloudprojectio_StreamCreation struct {
	Description string                        `json:"description"`
	Kind        Cloudprojectio_StreamKindEnum `json:"kind"`
	Name        string                        `json:"name"`
	Region      string                        `json:"region"`
}

type Cloudprojectio_StreamKindEnum string
const (
	Cloudprojectio_StreamKindEnum_NON_PERSISTENT Cloudprojectio_StreamKindEnum = "NON_PERSISTENT"
	Cloudprojectio_StreamKindEnum_PERSISTENT  = "PERSISTENT"
)

type Cloudprojectio_StreamStats struct {
	Usage float64 `json:"usage"`
}

type Cloudprojectio_StreamStatusEnum string
const (
	Cloudprojectio_StreamStatusEnum_ERROR Cloudprojectio_StreamStatusEnum = "ERROR"
	Cloudprojectio_StreamStatusEnum_INSTALLING  = "INSTALLING"
	Cloudprojectio_StreamStatusEnum_RUNNING  = "RUNNING"
)

type Cloudprojectiostream_Region struct {
	Endpoint Cloudprojectiostream_RegionEndpoint `json:"endpoint"`
	Region   string                              `json:"region"`
}

type Cloudprojectiostream_RegionEndpoint struct {
	Pulsar string `json:"pulsar"`
}

type Cloudprojectiostream_Subscription struct {
	Id   string                                    `json:"id"`
	Kind Cloudprojectiostream_SubscriptionKindEnum `json:"kind"`
	Name string                                    `json:"name"`
}

type Cloudprojectiostream_SubscriptionCreation struct {
	Name string `json:"name"`
}

type Cloudprojectiostream_SubscriptionKindEnum string
const (
	Cloudprojectiostream_SubscriptionKindEnum_EXCLUSIVE Cloudprojectiostream_SubscriptionKindEnum = "EXCLUSIVE"
	Cloudprojectiostream_SubscriptionKindEnum_FAILOVER  = "FAILOVER"
	Cloudprojectiostream_SubscriptionKindEnum_KEY_SHARED  = "KEY_SHARED"
	Cloudprojectiostream_SubscriptionKindEnum_SHARED  = "SHARED"
)

type Cloudprojectiostream_SubscriptionStats struct {
	Lag int64 `json:"lag"`
}

type Cloudprojectiostream_Token struct {
	Action Cloudprojectiostream_TokenActionEnum `json:"action"`
	Id     string                               `json:"id"`
	Token  string                               `json:"token"`
}

type Cloudprojectiostream_TokenActionEnum string
const (
	Cloudprojectiostream_TokenActionEnum_BOTH Cloudprojectiostream_TokenActionEnum = "BOTH"
	Cloudprojectiostream_TokenActionEnum_CONSUME  = "CONSUME"
	Cloudprojectiostream_TokenActionEnum_PRODUCE  = "PRODUCE"
)

type Cloudprojectiostream_TokenCreation struct {
	Action Cloudprojectiostream_TokenActionEnum `json:"action"`
}

type Cloudprojectloadbalancer_ActionDispatch struct {
	Name   string `json:"name"`
	Target string `json:"target"`
}

type Cloudprojectloadbalancer_ActionRedirect struct {
	Location   string                                                `json:"location"`
	Name       string                                                `json:"name"`
	StatusCode Cloudprojectloadbalanceraction_RedirectStatusCodeEnum `json:"statusCode"`
}

type Cloudprojectloadbalancer_ActionReject struct {
	Name       string                                              `json:"name"`
	StatusCode Cloudprojectloadbalanceraction_RejectStatusCodeEnum `json:"statusCode"`
}

type Cloudprojectloadbalancer_ActionRewrite struct {
	Location string `json:"location"`
	Name     string `json:"name"`
}

type Cloudprojectloadbalancer_Actions struct {
	Dispatch []Cloudprojectloadbalancer_ActionDispatch `json:"dispatch"`
	Redirect []Cloudprojectloadbalancer_ActionRedirect `json:"redirect"`
	Reject   []Cloudprojectloadbalancer_ActionReject   `json:"reject"`
	Rewrite  []Cloudprojectloadbalancer_ActionRewrite  `json:"rewrite"`
}

type Cloudprojectloadbalancer_Address struct {
	Ipv4 string `json:"ipv4"`
	Ipv6 string `json:"ipv6"`
}

type Cloudprojectloadbalancer_Addresses struct {
	Ipv4 []string `json:"ipv4"`
	Ipv6 []string `json:"ipv6"`
}

type Cloudprojectloadbalancer_ApplicationConfiguration struct {
	Actions      Cloudprojectloadbalancer_Actions                           `json:"actions"`
	Certificates []string                                                   `json:"certificates"`
	Conditions   []Cloudprojectloadbalancer_Condition                       `json:"conditions"`
	EntryPoints  []Cloudprojectloadbalancer_EntryPoint                      `json:"entryPoints"`
	Networking   Cloudprojectloadbalancerconfigurationnetworking_Networking `json:"networking"`
	Targets      []Cloudprojectloadbalancer_Target                          `json:"targets"`
	Version      int64                                                      `json:"version"`
}

type Cloudprojectloadbalancer_ApplicationConfigurationCreation struct {
	Actions      Cloudprojectloadbalancer_Actions                           `json:"actions"`
	Certificates []string                                                   `json:"certificates"`
	Conditions   []Cloudprojectloadbalancer_Condition                       `json:"conditions"`
	EntryPoints  []Cloudprojectloadbalancer_EntryPoint                      `json:"entryPoints"`
	Networking   Cloudprojectloadbalancerconfigurationnetworking_Networking `json:"networking"`
	Targets      []Cloudprojectloadbalancer_Target                          `json:"targets"`
	Version      int64                                                      `json:"version"`
}

type Cloudprojectloadbalancer_ApplicationLoadBalancerSizeCapability struct {
	Bandwidth         int64                             `json:"bandwidth"`
	MaximumConnection int64                             `json:"maximumConnection"`
	RequestsPerSecond int64                             `json:"requestsPerSecond"`
	Size              Cloudprojectloadbalancer_SizeEnum `json:"size"`
}

type Cloudprojectloadbalancer_Backend struct {
	Balancer      Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum `json:"balancer"`
	Name          string                                                `json:"name"`
	ProxyProtocol Cloudprojectloadbalancerbackend_ProxyProtocolEnum     `json:"proxyProtocol"`
	Servers       []Cloudprojectloadbalancer_Server                     `json:"servers"`
	Sticky        bool                                                  `json:"sticky"`
}

type Cloudprojectloadbalancer_BackendSelector struct {
	Name string `json:"name"`
}

type Cloudprojectloadbalancer_Condition struct {
	Key    string                                      `json:"key"`
	Match  Cloudprojectloadbalancercondition_MatchEnum `json:"match"`
	Name   string                                      `json:"name"`
	Negate bool                                        `json:"negate"`
	Type   Cloudprojectloadbalancercondition_TypeEnum  `json:"type"`
	Values []string                                    `json:"values"`
}

type Cloudprojectloadbalancer_Configuration struct {
	Backends     []Cloudprojectloadbalancer_Backend                         `json:"backends"`
	Certificates []string                                                   `json:"certificates"`
	Frontends    []Cloudprojectloadbalancer_Frontend                        `json:"frontends"`
	Networking   Cloudprojectloadbalancerconfigurationnetworking_Networking `json:"networking"`
	Version      int64                                                      `json:"version"`
}

type Cloudprojectloadbalancer_ConfigurationCreation struct {
	Backends     []Cloudprojectloadbalancer_Backend                         `json:"backends"`
	Certificates []string                                                   `json:"certificates"`
	Frontends    []Cloudprojectloadbalancer_Frontend                        `json:"frontends"`
	Networking   Cloudprojectloadbalancerconfigurationnetworking_Networking `json:"networking"`
	Version      int64                                                      `json:"version"`
}

type Cloudprojectloadbalancer_ConfigurationVersion struct {
	Applied int64 `json:"applied"`
	Latest  int64 `json:"latest"`
}

type Cloudprojectloadbalancer_EntryPoint struct {
	DefaultTarget string                               `json:"defaultTarget"`
	DisableH2     bool                                 `json:"disableH2"`
	Name          string                               `json:"name"`
	PortRanges    []Cloudprojectloadbalancer_PortRange `json:"portRanges"`
	Ports         []int64                              `json:"ports"`
	Rules         []Cloudprojectloadbalancer_Rule      `json:"rules"`
	Tls           bool                                 `json:"tls"`
}

type Cloudprojectloadbalancer_Frontend struct {
	Backends   []Cloudprojectloadbalancer_BackendSelector `json:"backends"`
	Mode       Cloudprojectloadbalancerfrontend_ModeEnum  `json:"mode"`
	Name       string                                     `json:"name"`
	Port       int64                                      `json:"port"`
	PortRanges []Cloudprojectloadbalancer_PortRange       `json:"portRanges"`
	Ports      []int64                                    `json:"ports"`
	Tls        bool                                       `json:"tls"`
	Whitelist  []string                                   `json:"whitelist"`
}

type Cloudprojectloadbalancer_LoadBalancerSizeCapability struct {
	Bandwidth              int64                             `json:"bandwidth"`
	MaximumConnection      int64                             `json:"maximumConnection"`
	NewConnectionPerSecond int64                             `json:"newConnectionPerSecond"`
	Size                   Cloudprojectloadbalancer_SizeEnum `json:"size"`
}

type Cloudprojectloadbalancer_PortRange struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type Cloudprojectloadbalancer_Region struct {
	Region string `json:"region"`
}

type Cloudprojectloadbalancer_Rule struct {
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
}

type Cloudprojectloadbalancer_Server struct {
	Ip      string `json:"ip"`
	Name    string `json:"name"`
	NoCheck bool   `json:"noCheck"`
	Port    int64  `json:"port"`
	Weight  int64  `json:"weight"`
}

type Cloudprojectloadbalancer_SizeEnum string
const (
	Cloudprojectloadbalancer_SizeEnum_L Cloudprojectloadbalancer_SizeEnum = "L"
	Cloudprojectloadbalancer_SizeEnum_M  = "M"
	Cloudprojectloadbalancer_SizeEnum_S  = "S"
)

type Cloudprojectloadbalancer_Stats struct {
	ConcurrentFlows         float64                                  `json:"concurrentFlows"`
	HttpRequestsPerSecond   float64                                  `json:"httpRequestsPerSecond"`
	Status                  Cloudprojectloadbalancerstats_StatusEnum `json:"status"`
	Targets                 []Cloudprojectloadbalancerstats_Target   `json:"targets"`
	TcpConnectionsPerSecond float64                                  `json:"tcpConnectionsPerSecond"`
	Throughput              Cloudprojectloadbalancerstats_Throughput `json:"throughput"`
}

type Cloudprojectloadbalancer_StatusEnum string
const (
	Cloudprojectloadbalancer_StatusEnum_APPLYING Cloudprojectloadbalancer_StatusEnum = "APPLYING"
	Cloudprojectloadbalancer_StatusEnum_CREATED  = "CREATED"
	Cloudprojectloadbalancer_StatusEnum_DELETING  = "DELETING"
	Cloudprojectloadbalancer_StatusEnum_ERROR  = "ERROR"
	Cloudprojectloadbalancer_StatusEnum_FROZEN  = "FROZEN"
	Cloudprojectloadbalancer_StatusEnum_RUNNING  = "RUNNING"
)

type Cloudprojectloadbalancer_Target struct {
	Balancer      Cloudprojectloadbalancertarget_BalancerAlgorithmEnum `json:"balancer"`
	Name          string                                               `json:"name"`
	ProxyProtocol Cloudprojectloadbalancertarget_ProxyProtocolEnum     `json:"proxyProtocol"`
	Servers       []Cloudprojectloadbalancer_Server                    `json:"servers"`
	Sticky        bool                                                 `json:"sticky"`
}

type Cloudprojectloadbalanceraction_RedirectStatusCodeEnum string
const (
	Cloudprojectloadbalanceraction_RedirectStatusCodeEnum__301 Cloudprojectloadbalanceraction_RedirectStatusCodeEnum = "301"
	Cloudprojectloadbalanceraction_RedirectStatusCodeEnum__302  = "302"
	Cloudprojectloadbalanceraction_RedirectStatusCodeEnum__303  = "303"
	Cloudprojectloadbalanceraction_RedirectStatusCodeEnum__307  = "307"
	Cloudprojectloadbalanceraction_RedirectStatusCodeEnum__308  = "308"
)

type Cloudprojectloadbalanceraction_RejectStatusCodeEnum string
const (
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__200 Cloudprojectloadbalanceraction_RejectStatusCodeEnum = "200"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__400  = "400"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__403  = "403"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__405  = "405"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__408  = "408"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__429  = "429"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__500  = "500"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__502  = "502"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__503  = "503"
	Cloudprojectloadbalanceraction_RejectStatusCodeEnum__504  = "504"
)

type Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum string
const (
	Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum_first Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum = "first"
	Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum_leastconn  = "leastconn"
	Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum_roundrobin  = "roundrobin"
	Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum_source  = "source"
	Cloudprojectloadbalancerbackend_BalancerAlgorithmEnum_staticrr  = "static-rr"
)

type Cloudprojectloadbalancerbackend_ProxyProtocolEnum string
const (
	Cloudprojectloadbalancerbackend_ProxyProtocolEnum_v1 Cloudprojectloadbalancerbackend_ProxyProtocolEnum = "v1"
	Cloudprojectloadbalancerbackend_ProxyProtocolEnum_v2  = "v2"
	Cloudprojectloadbalancerbackend_ProxyProtocolEnum_v2cn  = "v2-cn"
	Cloudprojectloadbalancerbackend_ProxyProtocolEnum_v2ssl  = "v2-ssl"
)

type Cloudprojectloadbalancercondition_MatchEnum string
const (
	Cloudprojectloadbalancercondition_MatchEnum_endwith Cloudprojectloadbalancercondition_MatchEnum = "end-with"
	Cloudprojectloadbalancercondition_MatchEnum_exists  = "exists"
	Cloudprojectloadbalancercondition_MatchEnum_is  = "is"
	Cloudprojectloadbalancercondition_MatchEnum_regex  = "regex"
	Cloudprojectloadbalancercondition_MatchEnum_startwith  = "start-with"
)

type Cloudprojectloadbalancercondition_TypeEnum string
const (
	Cloudprojectloadbalancercondition_TypeEnum_cookie Cloudprojectloadbalancercondition_TypeEnum = "cookie"
	Cloudprojectloadbalancercondition_TypeEnum_header  = "header"
	Cloudprojectloadbalancercondition_TypeEnum_host  = "host"
	Cloudprojectloadbalancercondition_TypeEnum_method  = "method"
	Cloudprojectloadbalancercondition_TypeEnum_path  = "path"
	Cloudprojectloadbalancercondition_TypeEnum_queryparam  = "query-param"
	Cloudprojectloadbalancercondition_TypeEnum_source  = "source"
)

type Cloudprojectloadbalancerconfigurationnetworking_Egress struct {
	Id   string                                            `json:"id"`
	Kind Cloudprojectloadbalancernetworkingegress_KindEnum `json:"kind"`
}

type Cloudprojectloadbalancerconfigurationnetworking_Ingress struct {
	Kind Cloudprojectloadbalancernetworkingingress_KindEnum `json:"kind"`
}

type Cloudprojectloadbalancerconfigurationnetworking_Networking struct {
	Egress  Cloudprojectloadbalancerconfigurationnetworking_Egress  `json:"egress"`
	Ingress Cloudprojectloadbalancerconfigurationnetworking_Ingress `json:"ingress"`
}

type Cloudprojectloadbalancerfrontend_ModeEnum string
const (
	Cloudprojectloadbalancerfrontend_ModeEnum_TCP Cloudprojectloadbalancerfrontend_ModeEnum = "TCP"
)

type Cloudprojectloadbalancernetworking_Egress struct {
	Id   string                                            `json:"id"`
	Kind Cloudprojectloadbalancernetworkingegress_KindEnum `json:"kind"`
}

type Cloudprojectloadbalancernetworking_EgressCreation struct {
	Kind Cloudprojectloadbalancernetworkingegress_KindEnum `json:"kind"`
}

type Cloudprojectloadbalancernetworking_Ingress struct {
	Kind Cloudprojectloadbalancernetworkingingress_KindEnum `json:"kind"`
}

type Cloudprojectloadbalancernetworking_IngressCreation struct {
	Kind Cloudprojectloadbalancernetworkingingress_KindEnum `json:"kind"`
}

type Cloudprojectloadbalancernetworking_Networking struct {
	Egress  Cloudprojectloadbalancernetworking_Egress  `json:"egress"`
	Ingress Cloudprojectloadbalancernetworking_Ingress `json:"ingress"`
}

type Cloudprojectloadbalancernetworking_NetworkingCreation struct {
	Egress  Cloudprojectloadbalancernetworking_EgressCreation  `json:"egress"`
	Ingress Cloudprojectloadbalancernetworking_IngressCreation `json:"ingress"`
}

type Cloudprojectloadbalancernetworkingegress_KindEnum string
const (
	Cloudprojectloadbalancernetworkingegress_KindEnum_public Cloudprojectloadbalancernetworkingegress_KindEnum = "public"
	Cloudprojectloadbalancernetworkingegress_KindEnum_vrack  = "vrack"
)

type Cloudprojectloadbalancernetworkingingress_KindEnum string
const (
	Cloudprojectloadbalancernetworkingingress_KindEnum_public Cloudprojectloadbalancernetworkingingress_KindEnum = "public"
)

type Cloudprojectloadbalancerstats_StatusEnum string
const (
	Cloudprojectloadbalancerstats_StatusEnum_HEALTHY Cloudprojectloadbalancerstats_StatusEnum = "HEALTHY"
	Cloudprojectloadbalancerstats_StatusEnum_NOT_AVAILABLE  = "NOT_AVAILABLE"
)

type Cloudprojectloadbalancerstats_Target struct {
	Name    string                                       `json:"name"`
	Servers []Cloudprojectloadbalancerstatstarget_Server `json:"servers"`
}

type Cloudprojectloadbalancerstats_Throughput struct {
	In  float64 `json:"in"`
	Out float64 `json:"out"`
}

type Cloudprojectloadbalancerstatstarget_Server struct {
	Name   string                                               `json:"name"`
	Status Cloudprojectloadbalancerstatstargetserver_StatusEnum `json:"status"`
}

type Cloudprojectloadbalancerstatstargetserver_StatusEnum string
const (
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_ERROR Cloudprojectloadbalancerstatstargetserver_StatusEnum = "ERROR"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_HEALTHY  = "HEALTHY"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_INIT  = "INIT"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_L4_CONNECTION_ERROR  = "L4_CONNECTION_ERROR"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_L4_TIMEOUT_ERROR  = "L4_TIMEOUT_ERROR"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_L7_PROTOCOL_ERROR  = "L7_PROTOCOL_ERROR"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_L7_RESPONSE_ERROR  = "L7_RESPONSE_ERROR"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_L7_TIMEOUT  = "L7_TIMEOUT"
	Cloudprojectloadbalancerstatstargetserver_StatusEnum_UNKNOWN  = "UNKNOWN"
)

type Cloudprojectloadbalancertarget_BalancerAlgorithmEnum string
const (
	Cloudprojectloadbalancertarget_BalancerAlgorithmEnum_first Cloudprojectloadbalancertarget_BalancerAlgorithmEnum = "first"
	Cloudprojectloadbalancertarget_BalancerAlgorithmEnum_leastconn  = "leastconn"
	Cloudprojectloadbalancertarget_BalancerAlgorithmEnum_roundrobin  = "roundrobin"
	Cloudprojectloadbalancertarget_BalancerAlgorithmEnum_source  = "source"
	Cloudprojectloadbalancertarget_BalancerAlgorithmEnum_staticrr  = "static-rr"
)

type Cloudprojectloadbalancertarget_ProxyProtocolEnum string
const (
	Cloudprojectloadbalancertarget_ProxyProtocolEnum_v1 Cloudprojectloadbalancertarget_ProxyProtocolEnum = "v1"
	Cloudprojectloadbalancertarget_ProxyProtocolEnum_v2  = "v2"
	Cloudprojectloadbalancertarget_ProxyProtocolEnum_v2cn  = "v2-cn"
	Cloudprojectloadbalancertarget_ProxyProtocolEnum_v2ssl  = "v2-ssl"
)

type Cloudprojectnetworkloadbalancer_ActionReject struct {
	Name string                                               `json:"name"`
	Type Cloudprojectnetworkloadbalanceraction_RejectTypeEnum `json:"type"`
}

type Cloudprojectnetworkloadbalancer_Actions struct {
	Dispatch []Cloudprojectloadbalancer_ActionDispatch      `json:"dispatch"`
	Reject   []Cloudprojectnetworkloadbalancer_ActionReject `json:"reject"`
}

type Cloudprojectnetworkloadbalancer_Condition struct {
	Key    string                                            `json:"key"`
	Match  Cloudprojectloadbalancercondition_MatchEnum       `json:"match"`
	Name   string                                            `json:"name"`
	Negate bool                                              `json:"negate"`
	Type   Cloudprojectnetworkloadbalancercondition_TypeEnum `json:"type"`
	Values []string                                          `json:"values"`
}

type Cloudprojectnetworkloadbalancer_Configuration struct {
	Actions     Cloudprojectnetworkloadbalancer_Actions                    `json:"actions"`
	Conditions  []Cloudprojectnetworkloadbalancer_Condition                `json:"conditions"`
	EntryPoints []Cloudprojectnetworkloadbalancer_EntryPoint               `json:"entryPoints"`
	Networking  Cloudprojectloadbalancerconfigurationnetworking_Networking `json:"networking"`
	Targets     []Cloudprojectloadbalancer_Target                          `json:"targets"`
	Version     int64                                                      `json:"version"`
}

type Cloudprojectnetworkloadbalancer_ConfigurationCreation struct {
	Actions     Cloudprojectnetworkloadbalancer_Actions                    `json:"actions"`
	Conditions  []Cloudprojectnetworkloadbalancer_Condition                `json:"conditions"`
	EntryPoints []Cloudprojectnetworkloadbalancer_EntryPoint               `json:"entryPoints"`
	Networking  Cloudprojectloadbalancerconfigurationnetworking_Networking `json:"networking"`
	Targets     []Cloudprojectloadbalancer_Target                          `json:"targets"`
	Version     int64                                                      `json:"version"`
}

type Cloudprojectnetworkloadbalancer_EntryPoint struct {
	DefaultTarget string                               `json:"defaultTarget"`
	Name          string                               `json:"name"`
	PortRanges    []Cloudprojectloadbalancer_PortRange `json:"portRanges"`
	Ports         []int64                              `json:"ports"`
	Rules         []Cloudprojectloadbalancer_Rule      `json:"rules"`
}

type Cloudprojectnetworkloadbalancer_LoadBalancerSizeCapability struct {
	Bandwidth              int64                             `json:"bandwidth"`
	MaximumConnection      int64                             `json:"maximumConnection"`
	NewConnectionPerSecond int64                             `json:"newConnectionPerSecond"`
	Size                   Cloudprojectloadbalancer_SizeEnum `json:"size"`
}

type Cloudprojectnetworkloadbalanceraction_RejectTypeEnum string
const (
	Cloudprojectnetworkloadbalanceraction_RejectTypeEnum_deny Cloudprojectnetworkloadbalanceraction_RejectTypeEnum = "deny"
	Cloudprojectnetworkloadbalanceraction_RejectTypeEnum_drop  = "drop"
)

type Cloudprojectnetworkloadbalancercondition_TypeEnum string
const (
	Cloudprojectnetworkloadbalancercondition_TypeEnum_source Cloudprojectnetworkloadbalancercondition_TypeEnum = "source"
)

type Cloudquota_AllowedQuota struct {
	Compute Cloudquota_ComputeQuota `json:"compute"`
	Name    string                  `json:"name"`
	Network Cloudquota_NetworkQuota `json:"network"`
	Volume  Cloudquota_VolumeQuota  `json:"volume"`
}

type Cloudquota_ComputeQuota struct {
	Cores     int64 `json:"cores"`
	Instances int64 `json:"instances"`
	Ram       int64 `json:"ram"`
}

type Cloudquota_InstanceUsageQuotas struct {
	MaxCores      int64 `json:"maxCores"`
	MaxInstances  int64 `json:"maxInstances"`
	MaxRam        int64 `json:"maxRam"`
	UsedCores     int64 `json:"usedCores"`
	UsedInstances int64 `json:"usedInstances"`
	UsedRAM       int64 `json:"usedRAM"`
}

type Cloudquota_KeypairQuotas struct {
	MaxCount int64 `json:"maxCount"`
}

type Cloudquota_NetworkQuota struct {
	Networks int64 `json:"networks"`
	Ports    int64 `json:"ports"`
	Subnets  int64 `json:"subnets"`
}

type Cloudquota_Quotas struct {
	Instance Cloudquota_InstanceUsageQuotas `json:"instance"`
	Keypair  Cloudquota_KeypairQuotas       `json:"keypair"`
	Region   string                         `json:"region"`
	Volume   Cloudquota_VolumeUsageQuotas   `json:"volume"`
}

type Cloudquota_VolumeQuota struct {
	Gigabytes int64 `json:"gigabytes"`
	Snapshots int64 `json:"snapshots"`
	Volumes   int64 `json:"volumes"`
}

type Cloudquota_VolumeUsageQuotas struct {
	MaxGigabytes  int64 `json:"maxGigabytes"`
	UsedGigabytes int64 `json:"usedGigabytes"`
	VolumeCount   int64 `json:"volumeCount"`
}

type Cloudquotastorage_Quota struct {
	BytesUsed      int64 `json:"bytesUsed"`
	ContainerCount int64 `json:"containerCount"`
	ObjectCount    int64 `json:"objectCount"`
	QuotaBytes     int64 `json:"quotaBytes"`
}

type Cloudquotastorage_QuotaUpdate struct {
	QuotaBytes int64 `json:"quotaBytes"`
}

type Cloudrole_Permission struct {
	Label string   `json:"label"`
	Roles []string `json:"roles"`
}

type Cloudrole_Role struct {
	Description string   `json:"description"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type Cloudrole_Roles struct {
	Roles    []Cloudrole_Role    `json:"roles"`
	Services []Cloudrole_Service `json:"services"`
}

type Cloudrole_Service struct {
	Name        string                 `json:"name"`
	Permissions []Cloudrole_Permission `json:"permissions"`
}

type Cloudsshkey_SshKey struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	PublicKey string   `json:"publicKey"`
	Regions   []string `json:"regions"`
}

type Cloudsshkey_SshKeyDetail struct {
	FingerPrint string   `json:"fingerPrint"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	PublicKey   string   `json:"publicKey"`
	Regions     []string `json:"regions"`
}

type Cloudstack_Content struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type Cloudstack_InstructionGuide struct {
	Content  []Cloudstack_Content `json:"content"`
	Language string               `json:"language"`
	Sections []Cloudstack_Section `json:"sections"`
	Title    string               `json:"title"`
}

type Cloudstack_Section struct {
	Content []Cloudstack_Content `json:"content"`
	Steps   []Cloudstack_Step    `json:"steps"`
	Title   string               `json:"title"`
}

type Cloudstack_Stack struct {
	Commit        string                        `json:"commit"`
	Description   string                        `json:"description"`
	GitRepository string                        `json:"gitRepository"`
	Instructions  []Cloudstack_InstructionGuide `json:"instructions"`
	Name          string                        `json:"name"`
	Release       string                        `json:"release"`
	Uuid          string                        `json:"uuid"`
}

type Cloudstack_Step struct {
	Content []Cloudstack_Content `json:"content"`
	Title   string               `json:"title"`
}

type Cloudstorage_AddContainerPolicy struct {
	ObjectKey string                      `json:"objectKey"`
	RoleName  Cloudstorage_PolicyRoleEnum `json:"roleName"`
}

type Cloudstorage_Container struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Region        string `json:"region"`
	StoredBytes   int64  `json:"storedBytes"`
	StoredObjects int64  `json:"storedObjects"`
}

type Cloudstorage_ContainerAccess struct {
	Endpoints []Cloudstorage_Endpoint `json:"endpoints"`
	Token     string                  `json:"token"`
}

type Cloudstorage_ContainerDetail struct {
	Archive       bool                           `json:"archive"`
	ContainerType Cloudstorage_TypeEnum          `json:"containerType"`
	Cors          []string                       `json:"cors"`
	Name          string                         `json:"name"`
	Objects       []Cloudstorage_ContainerObject `json:"objects"`
	Public        bool                           `json:"public"`
	Region        string                         `json:"region"`
	StaticUrl     string                         `json:"staticUrl"`
	StoredBytes   int64                          `json:"storedBytes"`
	StoredObjects int64                          `json:"storedObjects"`
}

type Cloudstorage_ContainerObject struct {
	ContentType    string                          `json:"contentType"`
	LastModified   time.Time                       `json:"lastModified"`
	Name           string                          `json:"name"`
	RetrievalDelay int64                           `json:"retrievalDelay"`
	RetrievalState Cloudstorage_RetrievalStateEnum `json:"retrievalState"`
	Size           int64                           `json:"size"`
}

type Cloudstorage_ContainerObjectTempURL struct {
	ExpirationDate time.Time `json:"expirationDate"`
	GetURL         string    `json:"getURL"`
}

type Cloudstorage_Endpoint struct {
	Region string `json:"region"`
	Url    string `json:"url"`
}

type Cloudstorage_PolicyRaw struct {
	Policy string `json:"policy"`
}

type Cloudstorage_PolicyRoleEnum string
const (
	Cloudstorage_PolicyRoleEnum_admin Cloudstorage_PolicyRoleEnum = "admin"
	Cloudstorage_PolicyRoleEnum_deny  = "deny"
	Cloudstorage_PolicyRoleEnum_readOnly  = "readOnly"
	Cloudstorage_PolicyRoleEnum_readWrite  = "readWrite"
)

type Cloudstorage_PresignedURL struct {
	Method Cloudstorage_PresignedURLMethodEnum `json:"method"`
	Url    string                              `json:"url"`
}

type Cloudstorage_PresignedURLInput struct {
	Expire int64                               `json:"expire"`
	Method Cloudstorage_PresignedURLMethodEnum `json:"method"`
	Object string                              `json:"object"`
}

type Cloudstorage_PresignedURLMethodEnum string
const (
	Cloudstorage_PresignedURLMethodEnum_GET Cloudstorage_PresignedURLMethodEnum = "GET"
	Cloudstorage_PresignedURLMethodEnum_PUT  = "PUT"
)

type Cloudstorage_RetrievalStateEnum string
const (
	Cloudstorage_RetrievalStateEnum_sealed Cloudstorage_RetrievalStateEnum = "sealed"
	Cloudstorage_RetrievalStateEnum_unsealed  = "unsealed"
	Cloudstorage_RetrievalStateEnum_unsealing  = "unsealing"
)

type Cloudstorage_RightEnum string
const (
	Cloudstorage_RightEnum_all Cloudstorage_RightEnum = "all"
	Cloudstorage_RightEnum_read  = "read"
	Cloudstorage_RightEnum_write  = "write"
)

type Cloudstorage_TypeEnum string
const (
	Cloudstorage_TypeEnum_private Cloudstorage_TypeEnum = "private"
	Cloudstorage_TypeEnum_public  = "public"
	Cloudstorage_TypeEnum_static  = "static"
)

type Cloudusage_PaymentTypeEnum string
const (
	Cloudusage_PaymentTypeEnum_post Cloudusage_PaymentTypeEnum = "post"
	Cloudusage_PaymentTypeEnum_pre  = "pre"
)

type Cloudusage_Period struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type Cloudusage_UsageBill struct {
	Bill_id      string                     `json:"bill_id"`
	Credit       float64                    `json:"credit"`
	Part         float64                    `json:"part"`
	Payment_type Cloudusage_PaymentTypeEnum `json:"payment_type"`
	Total        float64                    `json:"total"`
}

type Cloudusage_UsageCurrent struct {
	HourlyUsage    CloudbillingView_HourlyResources  `json:"hourlyUsage"`
	LastUpdate     time.Time                         `json:"lastUpdate"`
	MonthlyUsage   CloudbillingView_MonthlyResources `json:"monthlyUsage"`
	Period         Cloudusage_Period                 `json:"period"`
	ResourcesUsage []CloudbillingView_TypedResources `json:"resourcesUsage"`
}

type Cloudusage_UsageCurrentBills struct {
	Bills []Cloudusage_UsageBill `json:"bills"`
}

type Cloudusage_UsageForecast struct {
	HourlyUsage    CloudbillingView_HourlyResources  `json:"hourlyUsage"`
	LastUpdate     time.Time                         `json:"lastUpdate"`
	MonthlyUsage   CloudbillingView_MonthlyResources `json:"monthlyUsage"`
	Period         Cloudusage_Period                 `json:"period"`
	ResourcesUsage []CloudbillingView_TypedResources `json:"resourcesUsage"`
	UsableCredits  CloudbillingView_UsedCredits      `json:"usableCredits"`
}

type Cloudusage_UsageHistory struct {
	Id         string            `json:"id"`
	LastUpdate time.Time         `json:"lastUpdate"`
	Period     Cloudusage_Period `json:"period"`
}

type Cloudusage_UsageHistoryDetail struct {
	HourlyUsage    CloudbillingView_HourlyResources  `json:"hourlyUsage"`
	Id             string                            `json:"id"`
	LastUpdate     time.Time                         `json:"lastUpdate"`
	MonthlyUsage   CloudbillingView_MonthlyResources `json:"monthlyUsage"`
	Period         Cloudusage_Period                 `json:"period"`
	ResourcesUsage []CloudbillingView_TypedResources `json:"resourcesUsage"`
}

type Cloudusage_UsageHistoryDetailBills struct {
	Bills []Cloudusage_UsageBill `json:"bills"`
}

type Clouduser_Openrc struct {
	Content string `json:"content"`
}

type Clouduser_OpenrcVersionEnum string
const (
	Clouduser_OpenrcVersionEnum_v20 Clouduser_OpenrcVersionEnum = "v2.0"
	Clouduser_OpenrcVersionEnum_v3  = "v3"
)

type Clouduser_RCloneServiceEnum string
const (
	Clouduser_RCloneServiceEnum_storage Clouduser_RCloneServiceEnum = "storage"
	Clouduser_RCloneServiceEnum_storages3  = "storage-s3"
)

type Clouduser_Rclone struct {
	Content string `json:"content"`
}

type Clouduser_RoleEnum string
const (
	Clouduser_RoleEnum_admin Clouduser_RoleEnum = "admin"
	Clouduser_RoleEnum_administrator  = "administrator"
	Clouduser_RoleEnum_ai_training_operator  = "ai_training_operator"
	Clouduser_RoleEnum_ai_training_read  = "ai_training_read"
	Clouduser_RoleEnum_authentication  = "authentication"
	Clouduser_RoleEnum_backup_operator  = "backup_operator"
	Clouduser_RoleEnum_compute_operator  = "compute_operator"
	Clouduser_RoleEnum_image_operator  = "image_operator"
	Clouduser_RoleEnum_infrastructure_supervisor  = "infrastructure_supervisor"
	Clouduser_RoleEnum_network_operator  = "network_operator"
	Clouduser_RoleEnum_network_security_operator  = "network_security_operator"
	Clouduser_RoleEnum_objectstore_operator  = "objectstore_operator"
	Clouduser_RoleEnum_volume_operator  = "volume_operator"
)

type Clouduser_S3CredentialsWithSecret struct {
	Access   string `json:"access"`
	Secret   string `json:"secret"`
	TenantId string `json:"tenantId"`
	UserId   string `json:"userId"`
}

type Clouduser_User struct {
	CreationDate time.Time                `json:"creationDate"`
	Description  string                   `json:"description"`
	Id           int64                    `json:"id"`
	OpenstackId  string                   `json:"openstackId"`
	Roles        []Cloudrole_Role         `json:"roles"`
	Status       Clouduser_UserStatusEnum `json:"status"`
	Username     string                   `json:"username"`
}

type Clouduser_UserDetail struct {
	CreationDate time.Time                `json:"creationDate"`
	Description  string                   `json:"description"`
	Id           int64                    `json:"id"`
	OpenstackId  string                   `json:"openstackId"`
	Password     string                   `json:"password"`
	Roles        []Cloudrole_Role         `json:"roles"`
	Status       Clouduser_UserStatusEnum `json:"status"`
	Username     string                   `json:"username"`
}

type Clouduser_UserStatusEnum string
const (
	Clouduser_UserStatusEnum_creating Clouduser_UserStatusEnum = "creating"
	Clouduser_UserStatusEnum_deleted  = "deleted"
	Clouduser_UserStatusEnum_deleting  = "deleting"
	Clouduser_UserStatusEnum_ok  = "ok"
)

type Cloudvolume_Snapshot struct {
	CreationDate time.Time                      `json:"creationDate"`
	Description  string                         `json:"description"`
	Id           string                         `json:"id"`
	Name         string                         `json:"name"`
	PlanCode     string                         `json:"planCode"`
	Region       string                         `json:"region"`
	Size         int64                          `json:"size"`
	Status       Cloudvolume_SnapshotStatusEnum `json:"status"`
	VolumeId     string                         `json:"volumeId"`
}

type Cloudvolume_SnapshotStatusEnum string
const (
	Cloudvolume_SnapshotStatusEnum_available Cloudvolume_SnapshotStatusEnum = "available"
	Cloudvolume_SnapshotStatusEnum_creating  = "creating"
	Cloudvolume_SnapshotStatusEnum_deleting  = "deleting"
	Cloudvolume_SnapshotStatusEnum_error  = "error"
	Cloudvolume_SnapshotStatusEnum_error_deleting  = "error_deleting"
)

type Cloudvolume_Volume struct {
	AttachedTo   []string                   `json:"attachedTo"`
	Bootable     bool                       `json:"bootable"`
	CreationDate time.Time                  `json:"creationDate"`
	Description  string                     `json:"description"`
	Id           string                     `json:"id"`
	Name         string                     `json:"name"`
	PlanCode     string                     `json:"planCode"`
	Region       string                     `json:"region"`
	Size         int64                      `json:"size"`
	Status       string                     `json:"status"`
	Type         Cloudvolume_VolumeTypeEnum `json:"type"`
}

type Cloudvolume_VolumeTypeEnum string
const (
	Cloudvolume_VolumeTypeEnum_classic Cloudvolume_VolumeTypeEnum = "classic"
	Cloudvolume_VolumeTypeEnum_highspeed  = "high-speed"
	Cloudvolume_VolumeTypeEnum_highspeedgen2  = "high-speed-gen2"
)

type CloudvolumeBackup_VolumeBackup struct {
	CreationDate time.Time                                `json:"creationDate"`
	Id           string                                   `json:"id"`
	Name         string                                   `json:"name"`
	Size         int64                                    `json:"size"`
	Status       CloudvolumeBackup_VolumeBackupStatusEnum `json:"status"`
	VolumeId     string                                   `json:"volumeId"`
}

type CloudvolumeBackup_VolumeBackupCreation struct {
	Name     string `json:"name"`
	VolumeId string `json:"volumeId"`
}

type CloudvolumeBackup_VolumeBackupRestore struct {
	VolumeId string `json:"volumeId"`
}

type CloudvolumeBackup_VolumeBackupStatusEnum string
const (
	CloudvolumeBackup_VolumeBackupStatusEnum_creating CloudvolumeBackup_VolumeBackupStatusEnum = "creating"
	CloudvolumeBackup_VolumeBackupStatusEnum_deleting  = "deleting"
	CloudvolumeBackup_VolumeBackupStatusEnum_error  = "error"
	CloudvolumeBackup_VolumeBackupStatusEnum_ok  = "ok"
	CloudvolumeBackup_VolumeBackupStatusEnum_restoring  = "restoring"
)

type CloudvolumeBackup_VolumeCreationFromBackup struct {
	Name string `json:"name"`
}

type Nichandle_OvhSubsidiaryEnum string
const (
	Nichandle_OvhSubsidiaryEnum_CZ Nichandle_OvhSubsidiaryEnum = "CZ"
	Nichandle_OvhSubsidiaryEnum_DE  = "DE"
	Nichandle_OvhSubsidiaryEnum_ES  = "ES"
	Nichandle_OvhSubsidiaryEnum_EU  = "EU"
	Nichandle_OvhSubsidiaryEnum_FI  = "FI"
	Nichandle_OvhSubsidiaryEnum_FR  = "FR"
	Nichandle_OvhSubsidiaryEnum_GB  = "GB"
	Nichandle_OvhSubsidiaryEnum_IE  = "IE"
	Nichandle_OvhSubsidiaryEnum_IT  = "IT"
	Nichandle_OvhSubsidiaryEnum_LT  = "LT"
	Nichandle_OvhSubsidiaryEnum_MA  = "MA"
	Nichandle_OvhSubsidiaryEnum_NL  = "NL"
	Nichandle_OvhSubsidiaryEnum_PL  = "PL"
	Nichandle_OvhSubsidiaryEnum_PT  = "PT"
	Nichandle_OvhSubsidiaryEnum_SN  = "SN"
	Nichandle_OvhSubsidiaryEnum_TN  = "TN"
)

type Order_CurrencyCodeEnum string
const (
	Order_CurrencyCodeEnum_AUD Order_CurrencyCodeEnum = "AUD"
	Order_CurrencyCodeEnum_CAD  = "CAD"
	Order_CurrencyCodeEnum_CZK  = "CZK"
	Order_CurrencyCodeEnum_EUR  = "EUR"
	Order_CurrencyCodeEnum_GBP  = "GBP"
	Order_CurrencyCodeEnum_INR  = "INR"
	Order_CurrencyCodeEnum_LTL  = "LTL"
	Order_CurrencyCodeEnum_MAD  = "MAD"
	Order_CurrencyCodeEnum_NA  = "N/A"
	Order_CurrencyCodeEnum_PLN  = "PLN"
	Order_CurrencyCodeEnum_SGD  = "SGD"
	Order_CurrencyCodeEnum_TND  = "TND"
	Order_CurrencyCodeEnum_USD  = "USD"
	Order_CurrencyCodeEnum_XOF  = "XOF"
	Order_CurrencyCodeEnum_points  = "points"
)

type Order_Price struct {
	CurrencyCode Order_CurrencyCodeEnum `json:"currencyCode"`
	Text         string                 `json:"text"`
	Value        float64                `json:"value"`
}

type Service_RenewType struct {
	Automatic          bool  `json:"automatic"`
	DeleteAtExpiration bool  `json:"deleteAtExpiration"`
	Forced             bool  `json:"forced"`
	ManualPayment      bool  `json:"manualPayment"`
	Period             int64 `json:"period"`
}

type Service_RenewalTypeEnum string
const (
	Service_RenewalTypeEnum_automaticForcedProduct Service_RenewalTypeEnum = "automaticForcedProduct"
	Service_RenewalTypeEnum_automaticV2012  = "automaticV2012"
	Service_RenewalTypeEnum_automaticV2014  = "automaticV2014"
	Service_RenewalTypeEnum_automaticV2016  = "automaticV2016"
	Service_RenewalTypeEnum_manual  = "manual"
	Service_RenewalTypeEnum_oneShot  = "oneShot"
	Service_RenewalTypeEnum_option  = "option"
)

type Service_StateEnum string
const (
	Service_StateEnum_expired Service_StateEnum = "expired"
	Service_StateEnum_inCreation  = "inCreation"
	Service_StateEnum_ok  = "ok"
	Service_StateEnum_pendingDebt  = "pendingDebt"
	Service_StateEnum_unPaid  = "unPaid"
)

type Service_TerminationFutureUseEnum string
const (
	Service_TerminationFutureUseEnum_NOT_REPLACING_SERVICE Service_TerminationFutureUseEnum = "NOT_REPLACING_SERVICE"
	Service_TerminationFutureUseEnum_OTHER  = "OTHER"
	Service_TerminationFutureUseEnum_SUBSCRIBE_AN_OTHER_SERVICE  = "SUBSCRIBE_AN_OTHER_SERVICE"
	Service_TerminationFutureUseEnum_SUBSCRIBE_OTHER_KIND_OF_SERVICE_WITH_COMPETITOR  = "SUBSCRIBE_OTHER_KIND_OF_SERVICE_WITH_COMPETITOR"
	Service_TerminationFutureUseEnum_SUBSCRIBE_SIMILAR_SERVICE_WITH_COMPETITOR  = "SUBSCRIBE_SIMILAR_SERVICE_WITH_COMPETITOR"
)

type Service_TerminationReasonEnum string
const (
	Service_TerminationReasonEnum_FEATURES_DONT_SUIT_ME Service_TerminationReasonEnum = "FEATURES_DONT_SUIT_ME"
	Service_TerminationReasonEnum_LACK_OF_PERFORMANCES  = "LACK_OF_PERFORMANCES"
	Service_TerminationReasonEnum_MIGRATED_TO_ANOTHER_OVH_PRODUCT  = "MIGRATED_TO_ANOTHER_OVH_PRODUCT"
	Service_TerminationReasonEnum_MIGRATED_TO_COMPETITOR  = "MIGRATED_TO_COMPETITOR"
	Service_TerminationReasonEnum_NOT_ENOUGH_RECOGNITION  = "NOT_ENOUGH_RECOGNITION"
	Service_TerminationReasonEnum_NOT_NEEDED_ANYMORE  = "NOT_NEEDED_ANYMORE"
	Service_TerminationReasonEnum_NOT_RELIABLE  = "NOT_RELIABLE"
	Service_TerminationReasonEnum_NO_ANSWER  = "NO_ANSWER"
	Service_TerminationReasonEnum_OTHER  = "OTHER"
	Service_TerminationReasonEnum_PRODUCT_DIMENSION_DONT_SUIT_ME  = "PRODUCT_DIMENSION_DONT_SUIT_ME"
	Service_TerminationReasonEnum_PRODUCT_TOOLS_DONT_SUIT_ME  = "PRODUCT_TOOLS_DONT_SUIT_ME"
	Service_TerminationReasonEnum_TOO_EXPENSIVE  = "TOO_EXPENSIVE"
	Service_TerminationReasonEnum_TOO_HARD_TO_USE  = "TOO_HARD_TO_USE"
	Service_TerminationReasonEnum_UNSATIFIED_BY_CUSTOMER_SUPPORT  = "UNSATIFIED_BY_CUSTOMER_SUPPORT"
)

type Services_Service struct {
	CanDeleteAtExpiration bool                    `json:"canDeleteAtExpiration"`
	ContactAdmin          string                  `json:"contactAdmin"`
	ContactBilling        string                  `json:"contactBilling"`
	ContactTech           string                  `json:"contactTech"`
	Creation              time.Time               `json:"creation"`
	Domain                string                  `json:"domain"`
	EngagedUpTo           time.Time               `json:"engagedUpTo"`
	Expiration            time.Time               `json:"expiration"`
	PossibleRenewPeriod   []int64                 `json:"possibleRenewPeriod"`
	Renew                 Service_RenewType       `json:"renew"`
	RenewalType           Service_RenewalTypeEnum `json:"renewalType"`
	ServiceId             int64                   `json:"serviceId"`
	Status                Service_StateEnum       `json:"status"`
}
