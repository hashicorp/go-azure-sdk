package appserviceenvironments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoHealActionType string

const (
	AutoHealActionTypeCustomAction AutoHealActionType = "CustomAction"
	AutoHealActionTypeLogEvent     AutoHealActionType = "LogEvent"
	AutoHealActionTypeRecycle      AutoHealActionType = "Recycle"
)

func PossibleValuesForAutoHealActionType() []string {
	return []string{
		string(AutoHealActionTypeCustomAction),
		string(AutoHealActionTypeLogEvent),
		string(AutoHealActionTypeRecycle),
	}
}

func parseAutoHealActionType(input string) (*AutoHealActionType, error) {
	vals := map[string]AutoHealActionType{
		"customaction": AutoHealActionTypeCustomAction,
		"logevent":     AutoHealActionTypeLogEvent,
		"recycle":      AutoHealActionTypeRecycle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoHealActionType(input)
	return &out, nil
}

type AzureStorageState string

const (
	AzureStorageStateInvalidCredentials AzureStorageState = "InvalidCredentials"
	AzureStorageStateInvalidShare       AzureStorageState = "InvalidShare"
	AzureStorageStateNotValidated       AzureStorageState = "NotValidated"
	AzureStorageStateOk                 AzureStorageState = "Ok"
)

func PossibleValuesForAzureStorageState() []string {
	return []string{
		string(AzureStorageStateInvalidCredentials),
		string(AzureStorageStateInvalidShare),
		string(AzureStorageStateNotValidated),
		string(AzureStorageStateOk),
	}
}

func parseAzureStorageState(input string) (*AzureStorageState, error) {
	vals := map[string]AzureStorageState{
		"invalidcredentials": AzureStorageStateInvalidCredentials,
		"invalidshare":       AzureStorageStateInvalidShare,
		"notvalidated":       AzureStorageStateNotValidated,
		"ok":                 AzureStorageStateOk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureStorageState(input)
	return &out, nil
}

type AzureStorageType string

const (
	AzureStorageTypeAzureBlob  AzureStorageType = "AzureBlob"
	AzureStorageTypeAzureFiles AzureStorageType = "AzureFiles"
)

func PossibleValuesForAzureStorageType() []string {
	return []string{
		string(AzureStorageTypeAzureBlob),
		string(AzureStorageTypeAzureFiles),
	}
}

func parseAzureStorageType(input string) (*AzureStorageType, error) {
	vals := map[string]AzureStorageType{
		"azureblob":  AzureStorageTypeAzureBlob,
		"azurefiles": AzureStorageTypeAzureFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureStorageType(input)
	return &out, nil
}

type ClientCertMode string

const (
	ClientCertModeOptional                ClientCertMode = "Optional"
	ClientCertModeOptionalInteractiveUser ClientCertMode = "OptionalInteractiveUser"
	ClientCertModeRequired                ClientCertMode = "Required"
)

func PossibleValuesForClientCertMode() []string {
	return []string{
		string(ClientCertModeOptional),
		string(ClientCertModeOptionalInteractiveUser),
		string(ClientCertModeRequired),
	}
}

func parseClientCertMode(input string) (*ClientCertMode, error) {
	vals := map[string]ClientCertMode{
		"optional":                ClientCertModeOptional,
		"optionalinteractiveuser": ClientCertModeOptionalInteractiveUser,
		"required":                ClientCertModeRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientCertMode(input)
	return &out, nil
}

type ComputeModeOptions string

const (
	ComputeModeOptionsDedicated ComputeModeOptions = "Dedicated"
	ComputeModeOptionsDynamic   ComputeModeOptions = "Dynamic"
	ComputeModeOptionsShared    ComputeModeOptions = "Shared"
)

func PossibleValuesForComputeModeOptions() []string {
	return []string{
		string(ComputeModeOptionsDedicated),
		string(ComputeModeOptionsDynamic),
		string(ComputeModeOptionsShared),
	}
}

func parseComputeModeOptions(input string) (*ComputeModeOptions, error) {
	vals := map[string]ComputeModeOptions{
		"dedicated": ComputeModeOptionsDedicated,
		"dynamic":   ComputeModeOptionsDynamic,
		"shared":    ComputeModeOptionsShared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeModeOptions(input)
	return &out, nil
}

type ConnectionStringType string

const (
	ConnectionStringTypeApiHub          ConnectionStringType = "ApiHub"
	ConnectionStringTypeCustom          ConnectionStringType = "Custom"
	ConnectionStringTypeDocDb           ConnectionStringType = "DocDb"
	ConnectionStringTypeEventHub        ConnectionStringType = "EventHub"
	ConnectionStringTypeMySql           ConnectionStringType = "MySql"
	ConnectionStringTypeNotificationHub ConnectionStringType = "NotificationHub"
	ConnectionStringTypePostgreSQL      ConnectionStringType = "PostgreSQL"
	ConnectionStringTypeRedisCache      ConnectionStringType = "RedisCache"
	ConnectionStringTypeSQLAzure        ConnectionStringType = "SQLAzure"
	ConnectionStringTypeSQLServer       ConnectionStringType = "SQLServer"
	ConnectionStringTypeServiceBus      ConnectionStringType = "ServiceBus"
)

func PossibleValuesForConnectionStringType() []string {
	return []string{
		string(ConnectionStringTypeApiHub),
		string(ConnectionStringTypeCustom),
		string(ConnectionStringTypeDocDb),
		string(ConnectionStringTypeEventHub),
		string(ConnectionStringTypeMySql),
		string(ConnectionStringTypeNotificationHub),
		string(ConnectionStringTypePostgreSQL),
		string(ConnectionStringTypeRedisCache),
		string(ConnectionStringTypeSQLAzure),
		string(ConnectionStringTypeSQLServer),
		string(ConnectionStringTypeServiceBus),
	}
}

func parseConnectionStringType(input string) (*ConnectionStringType, error) {
	vals := map[string]ConnectionStringType{
		"apihub":          ConnectionStringTypeApiHub,
		"custom":          ConnectionStringTypeCustom,
		"docdb":           ConnectionStringTypeDocDb,
		"eventhub":        ConnectionStringTypeEventHub,
		"mysql":           ConnectionStringTypeMySql,
		"notificationhub": ConnectionStringTypeNotificationHub,
		"postgresql":      ConnectionStringTypePostgreSQL,
		"rediscache":      ConnectionStringTypeRedisCache,
		"sqlazure":        ConnectionStringTypeSQLAzure,
		"sqlserver":       ConnectionStringTypeSQLServer,
		"servicebus":      ConnectionStringTypeServiceBus,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionStringType(input)
	return &out, nil
}

type CustomDnsSuffixProvisioningState string

const (
	CustomDnsSuffixProvisioningStateDegraded   CustomDnsSuffixProvisioningState = "Degraded"
	CustomDnsSuffixProvisioningStateFailed     CustomDnsSuffixProvisioningState = "Failed"
	CustomDnsSuffixProvisioningStateInProgress CustomDnsSuffixProvisioningState = "InProgress"
	CustomDnsSuffixProvisioningStateSucceeded  CustomDnsSuffixProvisioningState = "Succeeded"
)

func PossibleValuesForCustomDnsSuffixProvisioningState() []string {
	return []string{
		string(CustomDnsSuffixProvisioningStateDegraded),
		string(CustomDnsSuffixProvisioningStateFailed),
		string(CustomDnsSuffixProvisioningStateInProgress),
		string(CustomDnsSuffixProvisioningStateSucceeded),
	}
}

func parseCustomDnsSuffixProvisioningState(input string) (*CustomDnsSuffixProvisioningState, error) {
	vals := map[string]CustomDnsSuffixProvisioningState{
		"degraded":   CustomDnsSuffixProvisioningStateDegraded,
		"failed":     CustomDnsSuffixProvisioningStateFailed,
		"inprogress": CustomDnsSuffixProvisioningStateInProgress,
		"succeeded":  CustomDnsSuffixProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomDnsSuffixProvisioningState(input)
	return &out, nil
}

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

func parseDefaultAction(input string) (*DefaultAction, error) {
	vals := map[string]DefaultAction{
		"allow": DefaultActionAllow,
		"deny":  DefaultActionDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultAction(input)
	return &out, nil
}

type FtpsState string

const (
	FtpsStateAllAllowed FtpsState = "AllAllowed"
	FtpsStateDisabled   FtpsState = "Disabled"
	FtpsStateFtpsOnly   FtpsState = "FtpsOnly"
)

func PossibleValuesForFtpsState() []string {
	return []string{
		string(FtpsStateAllAllowed),
		string(FtpsStateDisabled),
		string(FtpsStateFtpsOnly),
	}
}

func parseFtpsState(input string) (*FtpsState, error) {
	vals := map[string]FtpsState{
		"allallowed": FtpsStateAllAllowed,
		"disabled":   FtpsStateDisabled,
		"ftpsonly":   FtpsStateFtpsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FtpsState(input)
	return &out, nil
}

type HostType string

const (
	HostTypeRepository HostType = "Repository"
	HostTypeStandard   HostType = "Standard"
)

func PossibleValuesForHostType() []string {
	return []string{
		string(HostTypeRepository),
		string(HostTypeStandard),
	}
}

func parseHostType(input string) (*HostType, error) {
	vals := map[string]HostType{
		"repository": HostTypeRepository,
		"standard":   HostTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostType(input)
	return &out, nil
}

type HostingEnvironmentStatus string

const (
	HostingEnvironmentStatusDeleting  HostingEnvironmentStatus = "Deleting"
	HostingEnvironmentStatusPreparing HostingEnvironmentStatus = "Preparing"
	HostingEnvironmentStatusReady     HostingEnvironmentStatus = "Ready"
	HostingEnvironmentStatusScaling   HostingEnvironmentStatus = "Scaling"
)

func PossibleValuesForHostingEnvironmentStatus() []string {
	return []string{
		string(HostingEnvironmentStatusDeleting),
		string(HostingEnvironmentStatusPreparing),
		string(HostingEnvironmentStatusReady),
		string(HostingEnvironmentStatusScaling),
	}
}

func parseHostingEnvironmentStatus(input string) (*HostingEnvironmentStatus, error) {
	vals := map[string]HostingEnvironmentStatus{
		"deleting":  HostingEnvironmentStatusDeleting,
		"preparing": HostingEnvironmentStatusPreparing,
		"ready":     HostingEnvironmentStatusReady,
		"scaling":   HostingEnvironmentStatusScaling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostingEnvironmentStatus(input)
	return &out, nil
}

type IPFilterTag string

const (
	IPFilterTagDefault    IPFilterTag = "Default"
	IPFilterTagServiceTag IPFilterTag = "ServiceTag"
	IPFilterTagXffProxy   IPFilterTag = "XffProxy"
)

func PossibleValuesForIPFilterTag() []string {
	return []string{
		string(IPFilterTagDefault),
		string(IPFilterTagServiceTag),
		string(IPFilterTagXffProxy),
	}
}

func parseIPFilterTag(input string) (*IPFilterTag, error) {
	vals := map[string]IPFilterTag{
		"default":    IPFilterTagDefault,
		"servicetag": IPFilterTagServiceTag,
		"xffproxy":   IPFilterTagXffProxy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPFilterTag(input)
	return &out, nil
}

type LoadBalancingMode string

const (
	LoadBalancingModeNone          LoadBalancingMode = "None"
	LoadBalancingModePublishing    LoadBalancingMode = "Publishing"
	LoadBalancingModeWeb           LoadBalancingMode = "Web"
	LoadBalancingModeWebPublishing LoadBalancingMode = "Web, Publishing"
)

func PossibleValuesForLoadBalancingMode() []string {
	return []string{
		string(LoadBalancingModeNone),
		string(LoadBalancingModePublishing),
		string(LoadBalancingModeWeb),
		string(LoadBalancingModeWebPublishing),
	}
}

func parseLoadBalancingMode(input string) (*LoadBalancingMode, error) {
	vals := map[string]LoadBalancingMode{
		"none":            LoadBalancingModeNone,
		"publishing":      LoadBalancingModePublishing,
		"web":             LoadBalancingModeWeb,
		"web, publishing": LoadBalancingModeWebPublishing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoadBalancingMode(input)
	return &out, nil
}

type ManagedPipelineMode string

const (
	ManagedPipelineModeClassic    ManagedPipelineMode = "Classic"
	ManagedPipelineModeIntegrated ManagedPipelineMode = "Integrated"
)

func PossibleValuesForManagedPipelineMode() []string {
	return []string{
		string(ManagedPipelineModeClassic),
		string(ManagedPipelineModeIntegrated),
	}
}

func parseManagedPipelineMode(input string) (*ManagedPipelineMode, error) {
	vals := map[string]ManagedPipelineMode{
		"classic":    ManagedPipelineModeClassic,
		"integrated": ManagedPipelineModeIntegrated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedPipelineMode(input)
	return &out, nil
}

type OperationStatus string

const (
	OperationStatusCreated    OperationStatus = "Created"
	OperationStatusFailed     OperationStatus = "Failed"
	OperationStatusInProgress OperationStatus = "InProgress"
	OperationStatusSucceeded  OperationStatus = "Succeeded"
	OperationStatusTimedOut   OperationStatus = "TimedOut"
)

func PossibleValuesForOperationStatus() []string {
	return []string{
		string(OperationStatusCreated),
		string(OperationStatusFailed),
		string(OperationStatusInProgress),
		string(OperationStatusSucceeded),
		string(OperationStatusTimedOut),
	}
}

func parseOperationStatus(input string) (*OperationStatus, error) {
	vals := map[string]OperationStatus{
		"created":    OperationStatusCreated,
		"failed":     OperationStatusFailed,
		"inprogress": OperationStatusInProgress,
		"succeeded":  OperationStatusSucceeded,
		"timedout":   OperationStatusTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled   ProvisioningState = "Canceled"
	ProvisioningStateDeleting   ProvisioningState = "Deleting"
	ProvisioningStateFailed     ProvisioningState = "Failed"
	ProvisioningStateInProgress ProvisioningState = "InProgress"
	ProvisioningStateSucceeded  ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":   ProvisioningStateCanceled,
		"deleting":   ProvisioningStateDeleting,
		"failed":     ProvisioningStateFailed,
		"inprogress": ProvisioningStateInProgress,
		"succeeded":  ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type RedundancyMode string

const (
	RedundancyModeActiveActive RedundancyMode = "ActiveActive"
	RedundancyModeFailover     RedundancyMode = "Failover"
	RedundancyModeGeoRedundant RedundancyMode = "GeoRedundant"
	RedundancyModeManual       RedundancyMode = "Manual"
	RedundancyModeNone         RedundancyMode = "None"
)

func PossibleValuesForRedundancyMode() []string {
	return []string{
		string(RedundancyModeActiveActive),
		string(RedundancyModeFailover),
		string(RedundancyModeGeoRedundant),
		string(RedundancyModeManual),
		string(RedundancyModeNone),
	}
}

func parseRedundancyMode(input string) (*RedundancyMode, error) {
	vals := map[string]RedundancyMode{
		"activeactive": RedundancyModeActiveActive,
		"failover":     RedundancyModeFailover,
		"georedundant": RedundancyModeGeoRedundant,
		"manual":       RedundancyModeManual,
		"none":         RedundancyModeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RedundancyMode(input)
	return &out, nil
}

type ScmType string

const (
	ScmTypeBitbucketGit ScmType = "BitbucketGit"
	ScmTypeBitbucketHg  ScmType = "BitbucketHg"
	ScmTypeCodePlexGit  ScmType = "CodePlexGit"
	ScmTypeCodePlexHg   ScmType = "CodePlexHg"
	ScmTypeDropbox      ScmType = "Dropbox"
	ScmTypeExternalGit  ScmType = "ExternalGit"
	ScmTypeExternalHg   ScmType = "ExternalHg"
	ScmTypeGitHub       ScmType = "GitHub"
	ScmTypeLocalGit     ScmType = "LocalGit"
	ScmTypeNone         ScmType = "None"
	ScmTypeOneDrive     ScmType = "OneDrive"
	ScmTypeTfs          ScmType = "Tfs"
	ScmTypeVSO          ScmType = "VSO"
	ScmTypeVSTSRM       ScmType = "VSTSRM"
)

func PossibleValuesForScmType() []string {
	return []string{
		string(ScmTypeBitbucketGit),
		string(ScmTypeBitbucketHg),
		string(ScmTypeCodePlexGit),
		string(ScmTypeCodePlexHg),
		string(ScmTypeDropbox),
		string(ScmTypeExternalGit),
		string(ScmTypeExternalHg),
		string(ScmTypeGitHub),
		string(ScmTypeLocalGit),
		string(ScmTypeNone),
		string(ScmTypeOneDrive),
		string(ScmTypeTfs),
		string(ScmTypeVSO),
		string(ScmTypeVSTSRM),
	}
}

func parseScmType(input string) (*ScmType, error) {
	vals := map[string]ScmType{
		"bitbucketgit": ScmTypeBitbucketGit,
		"bitbuckethg":  ScmTypeBitbucketHg,
		"codeplexgit":  ScmTypeCodePlexGit,
		"codeplexhg":   ScmTypeCodePlexHg,
		"dropbox":      ScmTypeDropbox,
		"externalgit":  ScmTypeExternalGit,
		"externalhg":   ScmTypeExternalHg,
		"github":       ScmTypeGitHub,
		"localgit":     ScmTypeLocalGit,
		"none":         ScmTypeNone,
		"onedrive":     ScmTypeOneDrive,
		"tfs":          ScmTypeTfs,
		"vso":          ScmTypeVSO,
		"vstsrm":       ScmTypeVSTSRM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScmType(input)
	return &out, nil
}

type SiteAvailabilityState string

const (
	SiteAvailabilityStateDisasterRecoveryMode SiteAvailabilityState = "DisasterRecoveryMode"
	SiteAvailabilityStateLimited              SiteAvailabilityState = "Limited"
	SiteAvailabilityStateNormal               SiteAvailabilityState = "Normal"
)

func PossibleValuesForSiteAvailabilityState() []string {
	return []string{
		string(SiteAvailabilityStateDisasterRecoveryMode),
		string(SiteAvailabilityStateLimited),
		string(SiteAvailabilityStateNormal),
	}
}

func parseSiteAvailabilityState(input string) (*SiteAvailabilityState, error) {
	vals := map[string]SiteAvailabilityState{
		"disasterrecoverymode": SiteAvailabilityStateDisasterRecoveryMode,
		"limited":              SiteAvailabilityStateLimited,
		"normal":               SiteAvailabilityStateNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteAvailabilityState(input)
	return &out, nil
}

type SiteLoadBalancing string

const (
	SiteLoadBalancingLeastRequests        SiteLoadBalancing = "LeastRequests"
	SiteLoadBalancingLeastResponseTime    SiteLoadBalancing = "LeastResponseTime"
	SiteLoadBalancingPerSiteRoundRobin    SiteLoadBalancing = "PerSiteRoundRobin"
	SiteLoadBalancingRequestHash          SiteLoadBalancing = "RequestHash"
	SiteLoadBalancingWeightedRoundRobin   SiteLoadBalancing = "WeightedRoundRobin"
	SiteLoadBalancingWeightedTotalTraffic SiteLoadBalancing = "WeightedTotalTraffic"
)

func PossibleValuesForSiteLoadBalancing() []string {
	return []string{
		string(SiteLoadBalancingLeastRequests),
		string(SiteLoadBalancingLeastResponseTime),
		string(SiteLoadBalancingPerSiteRoundRobin),
		string(SiteLoadBalancingRequestHash),
		string(SiteLoadBalancingWeightedRoundRobin),
		string(SiteLoadBalancingWeightedTotalTraffic),
	}
}

func parseSiteLoadBalancing(input string) (*SiteLoadBalancing, error) {
	vals := map[string]SiteLoadBalancing{
		"leastrequests":        SiteLoadBalancingLeastRequests,
		"leastresponsetime":    SiteLoadBalancingLeastResponseTime,
		"persiteroundrobin":    SiteLoadBalancingPerSiteRoundRobin,
		"requesthash":          SiteLoadBalancingRequestHash,
		"weightedroundrobin":   SiteLoadBalancingWeightedRoundRobin,
		"weightedtotaltraffic": SiteLoadBalancingWeightedTotalTraffic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteLoadBalancing(input)
	return &out, nil
}

type SslState string

const (
	SslStateDisabled       SslState = "Disabled"
	SslStateIPBasedEnabled SslState = "IpBasedEnabled"
	SslStateSniEnabled     SslState = "SniEnabled"
)

func PossibleValuesForSslState() []string {
	return []string{
		string(SslStateDisabled),
		string(SslStateIPBasedEnabled),
		string(SslStateSniEnabled),
	}
}

func parseSslState(input string) (*SslState, error) {
	vals := map[string]SslState{
		"disabled":       SslStateDisabled,
		"ipbasedenabled": SslStateIPBasedEnabled,
		"snienabled":     SslStateSniEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SslState(input)
	return &out, nil
}

type StatusOptions string

const (
	StatusOptionsCreating StatusOptions = "Creating"
	StatusOptionsPending  StatusOptions = "Pending"
	StatusOptionsReady    StatusOptions = "Ready"
)

func PossibleValuesForStatusOptions() []string {
	return []string{
		string(StatusOptionsCreating),
		string(StatusOptionsPending),
		string(StatusOptionsReady),
	}
}

func parseStatusOptions(input string) (*StatusOptions, error) {
	vals := map[string]StatusOptions{
		"creating": StatusOptionsCreating,
		"pending":  StatusOptionsPending,
		"ready":    StatusOptionsReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusOptions(input)
	return &out, nil
}

type SupportedTlsVersions string

const (
	SupportedTlsVersionsOnePointOne  SupportedTlsVersions = "1.1"
	SupportedTlsVersionsOnePointTwo  SupportedTlsVersions = "1.2"
	SupportedTlsVersionsOnePointZero SupportedTlsVersions = "1.0"
)

func PossibleValuesForSupportedTlsVersions() []string {
	return []string{
		string(SupportedTlsVersionsOnePointOne),
		string(SupportedTlsVersionsOnePointTwo),
		string(SupportedTlsVersionsOnePointZero),
	}
}

func parseSupportedTlsVersions(input string) (*SupportedTlsVersions, error) {
	vals := map[string]SupportedTlsVersions{
		"1.1": SupportedTlsVersionsOnePointOne,
		"1.2": SupportedTlsVersionsOnePointTwo,
		"1.0": SupportedTlsVersionsOnePointZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportedTlsVersions(input)
	return &out, nil
}

type UpgradeAvailability string

const (
	UpgradeAvailabilityNone  UpgradeAvailability = "None"
	UpgradeAvailabilityReady UpgradeAvailability = "Ready"
)

func PossibleValuesForUpgradeAvailability() []string {
	return []string{
		string(UpgradeAvailabilityNone),
		string(UpgradeAvailabilityReady),
	}
}

func parseUpgradeAvailability(input string) (*UpgradeAvailability, error) {
	vals := map[string]UpgradeAvailability{
		"none":  UpgradeAvailabilityNone,
		"ready": UpgradeAvailabilityReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradeAvailability(input)
	return &out, nil
}

type UpgradePreference string

const (
	UpgradePreferenceEarly  UpgradePreference = "Early"
	UpgradePreferenceLate   UpgradePreference = "Late"
	UpgradePreferenceManual UpgradePreference = "Manual"
	UpgradePreferenceNone   UpgradePreference = "None"
)

func PossibleValuesForUpgradePreference() []string {
	return []string{
		string(UpgradePreferenceEarly),
		string(UpgradePreferenceLate),
		string(UpgradePreferenceManual),
		string(UpgradePreferenceNone),
	}
}

func parseUpgradePreference(input string) (*UpgradePreference, error) {
	vals := map[string]UpgradePreference{
		"early":  UpgradePreferenceEarly,
		"late":   UpgradePreferenceLate,
		"manual": UpgradePreferenceManual,
		"none":   UpgradePreferenceNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradePreference(input)
	return &out, nil
}

type UsageState string

const (
	UsageStateExceeded UsageState = "Exceeded"
	UsageStateNormal   UsageState = "Normal"
)

func PossibleValuesForUsageState() []string {
	return []string{
		string(UsageStateExceeded),
		string(UsageStateNormal),
	}
}

func parseUsageState(input string) (*UsageState, error) {
	vals := map[string]UsageState{
		"exceeded": UsageStateExceeded,
		"normal":   UsageStateNormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageState(input)
	return &out, nil
}

type WorkerSizeOptions string

const (
	WorkerSizeOptionsDOne             WorkerSizeOptions = "D1"
	WorkerSizeOptionsDThree           WorkerSizeOptions = "D3"
	WorkerSizeOptionsDTwo             WorkerSizeOptions = "D2"
	WorkerSizeOptionsDefault          WorkerSizeOptions = "Default"
	WorkerSizeOptionsLarge            WorkerSizeOptions = "Large"
	WorkerSizeOptionsLargeVThree      WorkerSizeOptions = "LargeV3"
	WorkerSizeOptionsMedium           WorkerSizeOptions = "Medium"
	WorkerSizeOptionsMediumVThree     WorkerSizeOptions = "MediumV3"
	WorkerSizeOptionsNestedSmall      WorkerSizeOptions = "NestedSmall"
	WorkerSizeOptionsNestedSmallLinux WorkerSizeOptions = "NestedSmallLinux"
	WorkerSizeOptionsSmall            WorkerSizeOptions = "Small"
	WorkerSizeOptionsSmallVThree      WorkerSizeOptions = "SmallV3"
)

func PossibleValuesForWorkerSizeOptions() []string {
	return []string{
		string(WorkerSizeOptionsDOne),
		string(WorkerSizeOptionsDThree),
		string(WorkerSizeOptionsDTwo),
		string(WorkerSizeOptionsDefault),
		string(WorkerSizeOptionsLarge),
		string(WorkerSizeOptionsLargeVThree),
		string(WorkerSizeOptionsMedium),
		string(WorkerSizeOptionsMediumVThree),
		string(WorkerSizeOptionsNestedSmall),
		string(WorkerSizeOptionsNestedSmallLinux),
		string(WorkerSizeOptionsSmall),
		string(WorkerSizeOptionsSmallVThree),
	}
}

func parseWorkerSizeOptions(input string) (*WorkerSizeOptions, error) {
	vals := map[string]WorkerSizeOptions{
		"d1":               WorkerSizeOptionsDOne,
		"d3":               WorkerSizeOptionsDThree,
		"d2":               WorkerSizeOptionsDTwo,
		"default":          WorkerSizeOptionsDefault,
		"large":            WorkerSizeOptionsLarge,
		"largev3":          WorkerSizeOptionsLargeVThree,
		"medium":           WorkerSizeOptionsMedium,
		"mediumv3":         WorkerSizeOptionsMediumVThree,
		"nestedsmall":      WorkerSizeOptionsNestedSmall,
		"nestedsmalllinux": WorkerSizeOptionsNestedSmallLinux,
		"small":            WorkerSizeOptionsSmall,
		"smallv3":          WorkerSizeOptionsSmallVThree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkerSizeOptions(input)
	return &out, nil
}
