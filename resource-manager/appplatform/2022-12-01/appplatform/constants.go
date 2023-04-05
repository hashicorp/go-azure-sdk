package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiPortalProvisioningState string

const (
	ApiPortalProvisioningStateCreating  ApiPortalProvisioningState = "Creating"
	ApiPortalProvisioningStateDeleting  ApiPortalProvisioningState = "Deleting"
	ApiPortalProvisioningStateFailed    ApiPortalProvisioningState = "Failed"
	ApiPortalProvisioningStateSucceeded ApiPortalProvisioningState = "Succeeded"
	ApiPortalProvisioningStateUpdating  ApiPortalProvisioningState = "Updating"
)

func PossibleValuesForApiPortalProvisioningState() []string {
	return []string{
		string(ApiPortalProvisioningStateCreating),
		string(ApiPortalProvisioningStateDeleting),
		string(ApiPortalProvisioningStateFailed),
		string(ApiPortalProvisioningStateSucceeded),
		string(ApiPortalProvisioningStateUpdating),
	}
}

type AppResourceProvisioningState string

const (
	AppResourceProvisioningStateCreating  AppResourceProvisioningState = "Creating"
	AppResourceProvisioningStateDeleting  AppResourceProvisioningState = "Deleting"
	AppResourceProvisioningStateFailed    AppResourceProvisioningState = "Failed"
	AppResourceProvisioningStateSucceeded AppResourceProvisioningState = "Succeeded"
	AppResourceProvisioningStateUpdating  AppResourceProvisioningState = "Updating"
)

func PossibleValuesForAppResourceProvisioningState() []string {
	return []string{
		string(AppResourceProvisioningStateCreating),
		string(AppResourceProvisioningStateDeleting),
		string(AppResourceProvisioningStateFailed),
		string(AppResourceProvisioningStateSucceeded),
		string(AppResourceProvisioningStateUpdating),
	}
}

type BackendProtocol string

const (
	BackendProtocolDefault BackendProtocol = "Default"
	BackendProtocolGRPC    BackendProtocol = "GRPC"
)

func PossibleValuesForBackendProtocol() []string {
	return []string{
		string(BackendProtocolDefault),
		string(BackendProtocolGRPC),
	}
}

type BindingType string

const (
	BindingTypeApacheSkyWalking    BindingType = "ApacheSkyWalking"
	BindingTypeAppDynamics         BindingType = "AppDynamics"
	BindingTypeApplicationInsights BindingType = "ApplicationInsights"
	BindingTypeDynatrace           BindingType = "Dynatrace"
	BindingTypeElasticAPM          BindingType = "ElasticAPM"
	BindingTypeNewRelic            BindingType = "NewRelic"
)

func PossibleValuesForBindingType() []string {
	return []string{
		string(BindingTypeApacheSkyWalking),
		string(BindingTypeAppDynamics),
		string(BindingTypeApplicationInsights),
		string(BindingTypeDynatrace),
		string(BindingTypeElasticAPM),
		string(BindingTypeNewRelic),
	}
}

type BuildProvisioningState string

const (
	BuildProvisioningStateCreating  BuildProvisioningState = "Creating"
	BuildProvisioningStateDeleting  BuildProvisioningState = "Deleting"
	BuildProvisioningStateFailed    BuildProvisioningState = "Failed"
	BuildProvisioningStateSucceeded BuildProvisioningState = "Succeeded"
	BuildProvisioningStateUpdating  BuildProvisioningState = "Updating"
)

func PossibleValuesForBuildProvisioningState() []string {
	return []string{
		string(BuildProvisioningStateCreating),
		string(BuildProvisioningStateDeleting),
		string(BuildProvisioningStateFailed),
		string(BuildProvisioningStateSucceeded),
		string(BuildProvisioningStateUpdating),
	}
}

type BuildResultProvisioningState string

const (
	BuildResultProvisioningStateBuilding  BuildResultProvisioningState = "Building"
	BuildResultProvisioningStateDeleting  BuildResultProvisioningState = "Deleting"
	BuildResultProvisioningStateFailed    BuildResultProvisioningState = "Failed"
	BuildResultProvisioningStateQueuing   BuildResultProvisioningState = "Queuing"
	BuildResultProvisioningStateSucceeded BuildResultProvisioningState = "Succeeded"
)

func PossibleValuesForBuildResultProvisioningState() []string {
	return []string{
		string(BuildResultProvisioningStateBuilding),
		string(BuildResultProvisioningStateDeleting),
		string(BuildResultProvisioningStateFailed),
		string(BuildResultProvisioningStateQueuing),
		string(BuildResultProvisioningStateSucceeded),
	}
}

type BuildServiceProvisioningState string

const (
	BuildServiceProvisioningStateCreating  BuildServiceProvisioningState = "Creating"
	BuildServiceProvisioningStateDeleting  BuildServiceProvisioningState = "Deleting"
	BuildServiceProvisioningStateFailed    BuildServiceProvisioningState = "Failed"
	BuildServiceProvisioningStateSucceeded BuildServiceProvisioningState = "Succeeded"
	BuildServiceProvisioningStateUpdating  BuildServiceProvisioningState = "Updating"
)

func PossibleValuesForBuildServiceProvisioningState() []string {
	return []string{
		string(BuildServiceProvisioningStateCreating),
		string(BuildServiceProvisioningStateDeleting),
		string(BuildServiceProvisioningStateFailed),
		string(BuildServiceProvisioningStateSucceeded),
		string(BuildServiceProvisioningStateUpdating),
	}
}

type BuilderProvisioningState string

const (
	BuilderProvisioningStateCreating  BuilderProvisioningState = "Creating"
	BuilderProvisioningStateDeleting  BuilderProvisioningState = "Deleting"
	BuilderProvisioningStateFailed    BuilderProvisioningState = "Failed"
	BuilderProvisioningStateSucceeded BuilderProvisioningState = "Succeeded"
	BuilderProvisioningStateUpdating  BuilderProvisioningState = "Updating"
)

func PossibleValuesForBuilderProvisioningState() []string {
	return []string{
		string(BuilderProvisioningStateCreating),
		string(BuilderProvisioningStateDeleting),
		string(BuilderProvisioningStateFailed),
		string(BuilderProvisioningStateSucceeded),
		string(BuilderProvisioningStateUpdating),
	}
}

type BuildpackBindingProvisioningState string

const (
	BuildpackBindingProvisioningStateCreating  BuildpackBindingProvisioningState = "Creating"
	BuildpackBindingProvisioningStateDeleting  BuildpackBindingProvisioningState = "Deleting"
	BuildpackBindingProvisioningStateFailed    BuildpackBindingProvisioningState = "Failed"
	BuildpackBindingProvisioningStateSucceeded BuildpackBindingProvisioningState = "Succeeded"
	BuildpackBindingProvisioningStateUpdating  BuildpackBindingProvisioningState = "Updating"
)

func PossibleValuesForBuildpackBindingProvisioningState() []string {
	return []string{
		string(BuildpackBindingProvisioningStateCreating),
		string(BuildpackBindingProvisioningStateDeleting),
		string(BuildpackBindingProvisioningStateFailed),
		string(BuildpackBindingProvisioningStateSucceeded),
		string(BuildpackBindingProvisioningStateUpdating),
	}
}

type CertificateResourceProvisioningState string

const (
	CertificateResourceProvisioningStateCreating  CertificateResourceProvisioningState = "Creating"
	CertificateResourceProvisioningStateDeleting  CertificateResourceProvisioningState = "Deleting"
	CertificateResourceProvisioningStateFailed    CertificateResourceProvisioningState = "Failed"
	CertificateResourceProvisioningStateSucceeded CertificateResourceProvisioningState = "Succeeded"
	CertificateResourceProvisioningStateUpdating  CertificateResourceProvisioningState = "Updating"
)

func PossibleValuesForCertificateResourceProvisioningState() []string {
	return []string{
		string(CertificateResourceProvisioningStateCreating),
		string(CertificateResourceProvisioningStateDeleting),
		string(CertificateResourceProvisioningStateFailed),
		string(CertificateResourceProvisioningStateSucceeded),
		string(CertificateResourceProvisioningStateUpdating),
	}
}

type ConfigServerState string

const (
	ConfigServerStateDeleted      ConfigServerState = "Deleted"
	ConfigServerStateFailed       ConfigServerState = "Failed"
	ConfigServerStateNotAvailable ConfigServerState = "NotAvailable"
	ConfigServerStateSucceeded    ConfigServerState = "Succeeded"
	ConfigServerStateUpdating     ConfigServerState = "Updating"
)

func PossibleValuesForConfigServerState() []string {
	return []string{
		string(ConfigServerStateDeleted),
		string(ConfigServerStateFailed),
		string(ConfigServerStateNotAvailable),
		string(ConfigServerStateSucceeded),
		string(ConfigServerStateUpdating),
	}
}

type ConfigurationServiceProvisioningState string

const (
	ConfigurationServiceProvisioningStateCreating  ConfigurationServiceProvisioningState = "Creating"
	ConfigurationServiceProvisioningStateDeleting  ConfigurationServiceProvisioningState = "Deleting"
	ConfigurationServiceProvisioningStateFailed    ConfigurationServiceProvisioningState = "Failed"
	ConfigurationServiceProvisioningStateSucceeded ConfigurationServiceProvisioningState = "Succeeded"
	ConfigurationServiceProvisioningStateUpdating  ConfigurationServiceProvisioningState = "Updating"
)

func PossibleValuesForConfigurationServiceProvisioningState() []string {
	return []string{
		string(ConfigurationServiceProvisioningStateCreating),
		string(ConfigurationServiceProvisioningStateDeleting),
		string(ConfigurationServiceProvisioningStateFailed),
		string(ConfigurationServiceProvisioningStateSucceeded),
		string(ConfigurationServiceProvisioningStateUpdating),
	}
}

type CustomDomainResourceProvisioningState string

const (
	CustomDomainResourceProvisioningStateCreating  CustomDomainResourceProvisioningState = "Creating"
	CustomDomainResourceProvisioningStateDeleting  CustomDomainResourceProvisioningState = "Deleting"
	CustomDomainResourceProvisioningStateFailed    CustomDomainResourceProvisioningState = "Failed"
	CustomDomainResourceProvisioningStateSucceeded CustomDomainResourceProvisioningState = "Succeeded"
	CustomDomainResourceProvisioningStateUpdating  CustomDomainResourceProvisioningState = "Updating"
)

func PossibleValuesForCustomDomainResourceProvisioningState() []string {
	return []string{
		string(CustomDomainResourceProvisioningStateCreating),
		string(CustomDomainResourceProvisioningStateDeleting),
		string(CustomDomainResourceProvisioningStateFailed),
		string(CustomDomainResourceProvisioningStateSucceeded),
		string(CustomDomainResourceProvisioningStateUpdating),
	}
}

type DeploymentResourceProvisioningState string

const (
	DeploymentResourceProvisioningStateCreating  DeploymentResourceProvisioningState = "Creating"
	DeploymentResourceProvisioningStateFailed    DeploymentResourceProvisioningState = "Failed"
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = "Succeeded"
	DeploymentResourceProvisioningStateUpdating  DeploymentResourceProvisioningState = "Updating"
)

func PossibleValuesForDeploymentResourceProvisioningState() []string {
	return []string{
		string(DeploymentResourceProvisioningStateCreating),
		string(DeploymentResourceProvisioningStateFailed),
		string(DeploymentResourceProvisioningStateSucceeded),
		string(DeploymentResourceProvisioningStateUpdating),
	}
}

type DeploymentResourceStatus string

const (
	DeploymentResourceStatusRunning DeploymentResourceStatus = "Running"
	DeploymentResourceStatusStopped DeploymentResourceStatus = "Stopped"
)

func PossibleValuesForDeploymentResourceStatus() []string {
	return []string{
		string(DeploymentResourceStatusRunning),
		string(DeploymentResourceStatusStopped),
	}
}

type GatewayProvisioningState string

const (
	GatewayProvisioningStateCreating  GatewayProvisioningState = "Creating"
	GatewayProvisioningStateDeleting  GatewayProvisioningState = "Deleting"
	GatewayProvisioningStateFailed    GatewayProvisioningState = "Failed"
	GatewayProvisioningStateSucceeded GatewayProvisioningState = "Succeeded"
	GatewayProvisioningStateUpdating  GatewayProvisioningState = "Updating"
)

func PossibleValuesForGatewayProvisioningState() []string {
	return []string{
		string(GatewayProvisioningStateCreating),
		string(GatewayProvisioningStateDeleting),
		string(GatewayProvisioningStateFailed),
		string(GatewayProvisioningStateSucceeded),
		string(GatewayProvisioningStateUpdating),
	}
}

type GatewayRouteConfigProtocol string

const (
	GatewayRouteConfigProtocolHTTP  GatewayRouteConfigProtocol = "HTTP"
	GatewayRouteConfigProtocolHTTPS GatewayRouteConfigProtocol = "HTTPS"
)

func PossibleValuesForGatewayRouteConfigProtocol() []string {
	return []string{
		string(GatewayRouteConfigProtocolHTTP),
		string(GatewayRouteConfigProtocolHTTPS),
	}
}

type HTTPSchemeType string

const (
	HTTPSchemeTypeHTTP  HTTPSchemeType = "HTTP"
	HTTPSchemeTypeHTTPS HTTPSchemeType = "HTTPS"
)

func PossibleValuesForHTTPSchemeType() []string {
	return []string{
		string(HTTPSchemeTypeHTTP),
		string(HTTPSchemeTypeHTTPS),
	}
}

type KPackBuildStageProvisioningState string

const (
	KPackBuildStageProvisioningStateFailed     KPackBuildStageProvisioningState = "Failed"
	KPackBuildStageProvisioningStateNotStarted KPackBuildStageProvisioningState = "NotStarted"
	KPackBuildStageProvisioningStateRunning    KPackBuildStageProvisioningState = "Running"
	KPackBuildStageProvisioningStateSucceeded  KPackBuildStageProvisioningState = "Succeeded"
)

func PossibleValuesForKPackBuildStageProvisioningState() []string {
	return []string{
		string(KPackBuildStageProvisioningStateFailed),
		string(KPackBuildStageProvisioningStateNotStarted),
		string(KPackBuildStageProvisioningStateRunning),
		string(KPackBuildStageProvisioningStateSucceeded),
	}
}

type MonitoringSettingState string

const (
	MonitoringSettingStateFailed       MonitoringSettingState = "Failed"
	MonitoringSettingStateNotAvailable MonitoringSettingState = "NotAvailable"
	MonitoringSettingStateSucceeded    MonitoringSettingState = "Succeeded"
	MonitoringSettingStateUpdating     MonitoringSettingState = "Updating"
)

func PossibleValuesForMonitoringSettingState() []string {
	return []string{
		string(MonitoringSettingStateFailed),
		string(MonitoringSettingStateNotAvailable),
		string(MonitoringSettingStateSucceeded),
		string(MonitoringSettingStateUpdating),
	}
}

type PowerState string

const (
	PowerStateRunning PowerState = "Running"
	PowerStateStopped PowerState = "Stopped"
)

func PossibleValuesForPowerState() []string {
	return []string{
		string(PowerStateRunning),
		string(PowerStateStopped),
	}
}

type ProbeActionType string

const (
	ProbeActionTypeExecAction      ProbeActionType = "ExecAction"
	ProbeActionTypeHTTPGetAction   ProbeActionType = "HTTPGetAction"
	ProbeActionTypeTCPSocketAction ProbeActionType = "TCPSocketAction"
)

func PossibleValuesForProbeActionType() []string {
	return []string{
		string(ProbeActionTypeExecAction),
		string(ProbeActionTypeHTTPGetAction),
		string(ProbeActionTypeTCPSocketAction),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating   ProvisioningState = "Creating"
	ProvisioningStateDeleted    ProvisioningState = "Deleted"
	ProvisioningStateDeleting   ProvisioningState = "Deleting"
	ProvisioningStateFailed     ProvisioningState = "Failed"
	ProvisioningStateMoveFailed ProvisioningState = "MoveFailed"
	ProvisioningStateMoved      ProvisioningState = "Moved"
	ProvisioningStateMoving     ProvisioningState = "Moving"
	ProvisioningStateStarting   ProvisioningState = "Starting"
	ProvisioningStateStopping   ProvisioningState = "Stopping"
	ProvisioningStateSucceeded  ProvisioningState = "Succeeded"
	ProvisioningStateUpdating   ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoveFailed),
		string(ProvisioningStateMoved),
		string(ProvisioningStateMoving),
		string(ProvisioningStateStarting),
		string(ProvisioningStateStopping),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type ResourceSkuRestrictionsReasonCode string

const (
	ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	ResourceSkuRestrictionsReasonCodeQuotaId                     ResourceSkuRestrictionsReasonCode = "QuotaId"
)

func PossibleValuesForResourceSkuRestrictionsReasonCode() []string {
	return []string{
		string(ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription),
		string(ResourceSkuRestrictionsReasonCodeQuotaId),
	}
}

type ResourceSkuRestrictionsType string

const (
	ResourceSkuRestrictionsTypeLocation ResourceSkuRestrictionsType = "Location"
	ResourceSkuRestrictionsTypeZone     ResourceSkuRestrictionsType = "Zone"
)

func PossibleValuesForResourceSkuRestrictionsType() []string {
	return []string{
		string(ResourceSkuRestrictionsTypeLocation),
		string(ResourceSkuRestrictionsTypeZone),
	}
}

type ServiceRegistryProvisioningState string

const (
	ServiceRegistryProvisioningStateCreating  ServiceRegistryProvisioningState = "Creating"
	ServiceRegistryProvisioningStateDeleting  ServiceRegistryProvisioningState = "Deleting"
	ServiceRegistryProvisioningStateFailed    ServiceRegistryProvisioningState = "Failed"
	ServiceRegistryProvisioningStateSucceeded ServiceRegistryProvisioningState = "Succeeded"
	ServiceRegistryProvisioningStateUpdating  ServiceRegistryProvisioningState = "Updating"
)

func PossibleValuesForServiceRegistryProvisioningState() []string {
	return []string{
		string(ServiceRegistryProvisioningStateCreating),
		string(ServiceRegistryProvisioningStateDeleting),
		string(ServiceRegistryProvisioningStateFailed),
		string(ServiceRegistryProvisioningStateSucceeded),
		string(ServiceRegistryProvisioningStateUpdating),
	}
}

type SessionAffinity string

const (
	SessionAffinityCookie SessionAffinity = "Cookie"
	SessionAffinityNone   SessionAffinity = "None"
)

func PossibleValuesForSessionAffinity() []string {
	return []string{
		string(SessionAffinityCookie),
		string(SessionAffinityNone),
	}
}

type SkuScaleType string

const (
	SkuScaleTypeAutomatic SkuScaleType = "Automatic"
	SkuScaleTypeManual    SkuScaleType = "Manual"
	SkuScaleTypeNone      SkuScaleType = "None"
)

func PossibleValuesForSkuScaleType() []string {
	return []string{
		string(SkuScaleTypeAutomatic),
		string(SkuScaleTypeManual),
		string(SkuScaleTypeNone),
	}
}

type StorageType string

const (
	StorageTypeStorageAccount StorageType = "StorageAccount"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypeStorageAccount),
	}
}

type SupportedRuntimePlatform string

const (
	SupportedRuntimePlatformJava         SupportedRuntimePlatform = "Java"
	SupportedRuntimePlatformPointNETCore SupportedRuntimePlatform = ".NET Core"
)

func PossibleValuesForSupportedRuntimePlatform() []string {
	return []string{
		string(SupportedRuntimePlatformJava),
		string(SupportedRuntimePlatformPointNETCore),
	}
}

type SupportedRuntimeValue string

const (
	SupportedRuntimeValueJavaEight       SupportedRuntimeValue = "Java_8"
	SupportedRuntimeValueJavaOneOne      SupportedRuntimeValue = "Java_11"
	SupportedRuntimeValueJavaOneSeven    SupportedRuntimeValue = "Java_17"
	SupportedRuntimeValueNetCoreThreeOne SupportedRuntimeValue = "NetCore_31"
)

func PossibleValuesForSupportedRuntimeValue() []string {
	return []string{
		string(SupportedRuntimeValueJavaEight),
		string(SupportedRuntimeValueJavaOneOne),
		string(SupportedRuntimeValueJavaOneSeven),
		string(SupportedRuntimeValueNetCoreThreeOne),
	}
}

type TestKeyType string

const (
	TestKeyTypePrimary   TestKeyType = "Primary"
	TestKeyTypeSecondary TestKeyType = "Secondary"
)

func PossibleValuesForTestKeyType() []string {
	return []string{
		string(TestKeyTypePrimary),
		string(TestKeyTypeSecondary),
	}
}

type TrafficDirection string

const (
	TrafficDirectionInbound  TrafficDirection = "Inbound"
	TrafficDirectionOutbound TrafficDirection = "Outbound"
)

func PossibleValuesForTrafficDirection() []string {
	return []string{
		string(TrafficDirectionInbound),
		string(TrafficDirectionOutbound),
	}
}

type Type string

const (
	TypeAzureFileVolume Type = "AzureFileVolume"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeAzureFileVolume),
	}
}
