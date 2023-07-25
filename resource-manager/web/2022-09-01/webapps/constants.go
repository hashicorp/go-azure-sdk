package webapps

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

type AzureResourceType string

const (
	AzureResourceTypeTrafficManager AzureResourceType = "TrafficManager"
	AzureResourceTypeWebsite        AzureResourceType = "Website"
)

func PossibleValuesForAzureResourceType() []string {
	return []string{
		string(AzureResourceTypeTrafficManager),
		string(AzureResourceTypeWebsite),
	}
}

func parseAzureResourceType(input string) (*AzureResourceType, error) {
	vals := map[string]AzureResourceType{
		"trafficmanager": AzureResourceTypeTrafficManager,
		"website":        AzureResourceTypeWebsite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureResourceType(input)
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

type BackupItemStatus string

const (
	BackupItemStatusCreated            BackupItemStatus = "Created"
	BackupItemStatusDeleteFailed       BackupItemStatus = "DeleteFailed"
	BackupItemStatusDeleteInProgress   BackupItemStatus = "DeleteInProgress"
	BackupItemStatusDeleted            BackupItemStatus = "Deleted"
	BackupItemStatusFailed             BackupItemStatus = "Failed"
	BackupItemStatusInProgress         BackupItemStatus = "InProgress"
	BackupItemStatusPartiallySucceeded BackupItemStatus = "PartiallySucceeded"
	BackupItemStatusSkipped            BackupItemStatus = "Skipped"
	BackupItemStatusSucceeded          BackupItemStatus = "Succeeded"
	BackupItemStatusTimedOut           BackupItemStatus = "TimedOut"
)

func PossibleValuesForBackupItemStatus() []string {
	return []string{
		string(BackupItemStatusCreated),
		string(BackupItemStatusDeleteFailed),
		string(BackupItemStatusDeleteInProgress),
		string(BackupItemStatusDeleted),
		string(BackupItemStatusFailed),
		string(BackupItemStatusInProgress),
		string(BackupItemStatusPartiallySucceeded),
		string(BackupItemStatusSkipped),
		string(BackupItemStatusSucceeded),
		string(BackupItemStatusTimedOut),
	}
}

func parseBackupItemStatus(input string) (*BackupItemStatus, error) {
	vals := map[string]BackupItemStatus{
		"created":            BackupItemStatusCreated,
		"deletefailed":       BackupItemStatusDeleteFailed,
		"deleteinprogress":   BackupItemStatusDeleteInProgress,
		"deleted":            BackupItemStatusDeleted,
		"failed":             BackupItemStatusFailed,
		"inprogress":         BackupItemStatusInProgress,
		"partiallysucceeded": BackupItemStatusPartiallySucceeded,
		"skipped":            BackupItemStatusSkipped,
		"succeeded":          BackupItemStatusSucceeded,
		"timedout":           BackupItemStatusTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupItemStatus(input)
	return &out, nil
}

type BackupRestoreOperationType string

const (
	BackupRestoreOperationTypeClone      BackupRestoreOperationType = "Clone"
	BackupRestoreOperationTypeCloudFS    BackupRestoreOperationType = "CloudFS"
	BackupRestoreOperationTypeDefault    BackupRestoreOperationType = "Default"
	BackupRestoreOperationTypeRelocation BackupRestoreOperationType = "Relocation"
	BackupRestoreOperationTypeSnapshot   BackupRestoreOperationType = "Snapshot"
)

func PossibleValuesForBackupRestoreOperationType() []string {
	return []string{
		string(BackupRestoreOperationTypeClone),
		string(BackupRestoreOperationTypeCloudFS),
		string(BackupRestoreOperationTypeDefault),
		string(BackupRestoreOperationTypeRelocation),
		string(BackupRestoreOperationTypeSnapshot),
	}
}

func parseBackupRestoreOperationType(input string) (*BackupRestoreOperationType, error) {
	vals := map[string]BackupRestoreOperationType{
		"clone":      BackupRestoreOperationTypeClone,
		"cloudfs":    BackupRestoreOperationTypeCloudFS,
		"default":    BackupRestoreOperationTypeDefault,
		"relocation": BackupRestoreOperationTypeRelocation,
		"snapshot":   BackupRestoreOperationTypeSnapshot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupRestoreOperationType(input)
	return &out, nil
}

type BuiltInAuthenticationProvider string

const (
	BuiltInAuthenticationProviderAzureActiveDirectory BuiltInAuthenticationProvider = "AzureActiveDirectory"
	BuiltInAuthenticationProviderFacebook             BuiltInAuthenticationProvider = "Facebook"
	BuiltInAuthenticationProviderGithub               BuiltInAuthenticationProvider = "Github"
	BuiltInAuthenticationProviderGoogle               BuiltInAuthenticationProvider = "Google"
	BuiltInAuthenticationProviderMicrosoftAccount     BuiltInAuthenticationProvider = "MicrosoftAccount"
	BuiltInAuthenticationProviderTwitter              BuiltInAuthenticationProvider = "Twitter"
)

func PossibleValuesForBuiltInAuthenticationProvider() []string {
	return []string{
		string(BuiltInAuthenticationProviderAzureActiveDirectory),
		string(BuiltInAuthenticationProviderFacebook),
		string(BuiltInAuthenticationProviderGithub),
		string(BuiltInAuthenticationProviderGoogle),
		string(BuiltInAuthenticationProviderMicrosoftAccount),
		string(BuiltInAuthenticationProviderTwitter),
	}
}

func parseBuiltInAuthenticationProvider(input string) (*BuiltInAuthenticationProvider, error) {
	vals := map[string]BuiltInAuthenticationProvider{
		"azureactivedirectory": BuiltInAuthenticationProviderAzureActiveDirectory,
		"facebook":             BuiltInAuthenticationProviderFacebook,
		"github":               BuiltInAuthenticationProviderGithub,
		"google":               BuiltInAuthenticationProviderGoogle,
		"microsoftaccount":     BuiltInAuthenticationProviderMicrosoftAccount,
		"twitter":              BuiltInAuthenticationProviderTwitter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuiltInAuthenticationProvider(input)
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

type ClientCredentialMethod string

const (
	ClientCredentialMethodClientSecretPost ClientCredentialMethod = "ClientSecretPost"
)

func PossibleValuesForClientCredentialMethod() []string {
	return []string{
		string(ClientCredentialMethodClientSecretPost),
	}
}

func parseClientCredentialMethod(input string) (*ClientCredentialMethod, error) {
	vals := map[string]ClientCredentialMethod{
		"clientsecretpost": ClientCredentialMethodClientSecretPost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientCredentialMethod(input)
	return &out, nil
}

type CloneAbilityResult string

const (
	CloneAbilityResultCloneable          CloneAbilityResult = "Cloneable"
	CloneAbilityResultNotCloneable       CloneAbilityResult = "NotCloneable"
	CloneAbilityResultPartiallyCloneable CloneAbilityResult = "PartiallyCloneable"
)

func PossibleValuesForCloneAbilityResult() []string {
	return []string{
		string(CloneAbilityResultCloneable),
		string(CloneAbilityResultNotCloneable),
		string(CloneAbilityResultPartiallyCloneable),
	}
}

func parseCloneAbilityResult(input string) (*CloneAbilityResult, error) {
	vals := map[string]CloneAbilityResult{
		"cloneable":          CloneAbilityResultCloneable,
		"notcloneable":       CloneAbilityResultNotCloneable,
		"partiallycloneable": CloneAbilityResultPartiallyCloneable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloneAbilityResult(input)
	return &out, nil
}

type ConfigReferenceSource string

const (
	ConfigReferenceSourceKeyVault ConfigReferenceSource = "KeyVault"
)

func PossibleValuesForConfigReferenceSource() []string {
	return []string{
		string(ConfigReferenceSourceKeyVault),
	}
}

func parseConfigReferenceSource(input string) (*ConfigReferenceSource, error) {
	vals := map[string]ConfigReferenceSource{
		"keyvault": ConfigReferenceSourceKeyVault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigReferenceSource(input)
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

type ContinuousWebJobStatus string

const (
	ContinuousWebJobStatusInitializing   ContinuousWebJobStatus = "Initializing"
	ContinuousWebJobStatusPendingRestart ContinuousWebJobStatus = "PendingRestart"
	ContinuousWebJobStatusRunning        ContinuousWebJobStatus = "Running"
	ContinuousWebJobStatusStarting       ContinuousWebJobStatus = "Starting"
	ContinuousWebJobStatusStopped        ContinuousWebJobStatus = "Stopped"
)

func PossibleValuesForContinuousWebJobStatus() []string {
	return []string{
		string(ContinuousWebJobStatusInitializing),
		string(ContinuousWebJobStatusPendingRestart),
		string(ContinuousWebJobStatusRunning),
		string(ContinuousWebJobStatusStarting),
		string(ContinuousWebJobStatusStopped),
	}
}

func parseContinuousWebJobStatus(input string) (*ContinuousWebJobStatus, error) {
	vals := map[string]ContinuousWebJobStatus{
		"initializing":   ContinuousWebJobStatusInitializing,
		"pendingrestart": ContinuousWebJobStatusPendingRestart,
		"running":        ContinuousWebJobStatusRunning,
		"starting":       ContinuousWebJobStatusStarting,
		"stopped":        ContinuousWebJobStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContinuousWebJobStatus(input)
	return &out, nil
}

type CookieExpirationConvention string

const (
	CookieExpirationConventionFixedTime               CookieExpirationConvention = "FixedTime"
	CookieExpirationConventionIdentityProviderDerived CookieExpirationConvention = "IdentityProviderDerived"
)

func PossibleValuesForCookieExpirationConvention() []string {
	return []string{
		string(CookieExpirationConventionFixedTime),
		string(CookieExpirationConventionIdentityProviderDerived),
	}
}

func parseCookieExpirationConvention(input string) (*CookieExpirationConvention, error) {
	vals := map[string]CookieExpirationConvention{
		"fixedtime":               CookieExpirationConventionFixedTime,
		"identityproviderderived": CookieExpirationConventionIdentityProviderDerived,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CookieExpirationConvention(input)
	return &out, nil
}

type CustomHostNameDnsRecordType string

const (
	CustomHostNameDnsRecordTypeA     CustomHostNameDnsRecordType = "A"
	CustomHostNameDnsRecordTypeCName CustomHostNameDnsRecordType = "CName"
)

func PossibleValuesForCustomHostNameDnsRecordType() []string {
	return []string{
		string(CustomHostNameDnsRecordTypeA),
		string(CustomHostNameDnsRecordTypeCName),
	}
}

func parseCustomHostNameDnsRecordType(input string) (*CustomHostNameDnsRecordType, error) {
	vals := map[string]CustomHostNameDnsRecordType{
		"a":     CustomHostNameDnsRecordTypeA,
		"cname": CustomHostNameDnsRecordTypeCName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomHostNameDnsRecordType(input)
	return &out, nil
}

type DatabaseType string

const (
	DatabaseTypeLocalMySql DatabaseType = "LocalMySql"
	DatabaseTypeMySql      DatabaseType = "MySql"
	DatabaseTypePostgreSql DatabaseType = "PostgreSql"
	DatabaseTypeSqlAzure   DatabaseType = "SqlAzure"
)

func PossibleValuesForDatabaseType() []string {
	return []string{
		string(DatabaseTypeLocalMySql),
		string(DatabaseTypeMySql),
		string(DatabaseTypePostgreSql),
		string(DatabaseTypeSqlAzure),
	}
}

func parseDatabaseType(input string) (*DatabaseType, error) {
	vals := map[string]DatabaseType{
		"localmysql": DatabaseTypeLocalMySql,
		"mysql":      DatabaseTypeMySql,
		"postgresql": DatabaseTypePostgreSql,
		"sqlazure":   DatabaseTypeSqlAzure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseType(input)
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

type DeploymentBuildStatus string

const (
	DeploymentBuildStatusBuildAborted             DeploymentBuildStatus = "BuildAborted"
	DeploymentBuildStatusBuildFailed              DeploymentBuildStatus = "BuildFailed"
	DeploymentBuildStatusBuildInProgress          DeploymentBuildStatus = "BuildInProgress"
	DeploymentBuildStatusBuildPending             DeploymentBuildStatus = "BuildPending"
	DeploymentBuildStatusBuildRequestReceived     DeploymentBuildStatus = "BuildRequestReceived"
	DeploymentBuildStatusBuildSuccessful          DeploymentBuildStatus = "BuildSuccessful"
	DeploymentBuildStatusPostBuildRestartRequired DeploymentBuildStatus = "PostBuildRestartRequired"
	DeploymentBuildStatusRuntimeFailed            DeploymentBuildStatus = "RuntimeFailed"
	DeploymentBuildStatusRuntimeStarting          DeploymentBuildStatus = "RuntimeStarting"
	DeploymentBuildStatusRuntimeSuccessful        DeploymentBuildStatus = "RuntimeSuccessful"
	DeploymentBuildStatusStartPolling             DeploymentBuildStatus = "StartPolling"
	DeploymentBuildStatusStartPollingWithRestart  DeploymentBuildStatus = "StartPollingWithRestart"
	DeploymentBuildStatusTimedOut                 DeploymentBuildStatus = "TimedOut"
)

func PossibleValuesForDeploymentBuildStatus() []string {
	return []string{
		string(DeploymentBuildStatusBuildAborted),
		string(DeploymentBuildStatusBuildFailed),
		string(DeploymentBuildStatusBuildInProgress),
		string(DeploymentBuildStatusBuildPending),
		string(DeploymentBuildStatusBuildRequestReceived),
		string(DeploymentBuildStatusBuildSuccessful),
		string(DeploymentBuildStatusPostBuildRestartRequired),
		string(DeploymentBuildStatusRuntimeFailed),
		string(DeploymentBuildStatusRuntimeStarting),
		string(DeploymentBuildStatusRuntimeSuccessful),
		string(DeploymentBuildStatusStartPolling),
		string(DeploymentBuildStatusStartPollingWithRestart),
		string(DeploymentBuildStatusTimedOut),
	}
}

func parseDeploymentBuildStatus(input string) (*DeploymentBuildStatus, error) {
	vals := map[string]DeploymentBuildStatus{
		"buildaborted":             DeploymentBuildStatusBuildAborted,
		"buildfailed":              DeploymentBuildStatusBuildFailed,
		"buildinprogress":          DeploymentBuildStatusBuildInProgress,
		"buildpending":             DeploymentBuildStatusBuildPending,
		"buildrequestreceived":     DeploymentBuildStatusBuildRequestReceived,
		"buildsuccessful":          DeploymentBuildStatusBuildSuccessful,
		"postbuildrestartrequired": DeploymentBuildStatusPostBuildRestartRequired,
		"runtimefailed":            DeploymentBuildStatusRuntimeFailed,
		"runtimestarting":          DeploymentBuildStatusRuntimeStarting,
		"runtimesuccessful":        DeploymentBuildStatusRuntimeSuccessful,
		"startpolling":             DeploymentBuildStatusStartPolling,
		"startpollingwithrestart":  DeploymentBuildStatusStartPollingWithRestart,
		"timedout":                 DeploymentBuildStatusTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentBuildStatus(input)
	return &out, nil
}

type DnsVerificationTestResult string

const (
	DnsVerificationTestResultFailed  DnsVerificationTestResult = "Failed"
	DnsVerificationTestResultPassed  DnsVerificationTestResult = "Passed"
	DnsVerificationTestResultSkipped DnsVerificationTestResult = "Skipped"
)

func PossibleValuesForDnsVerificationTestResult() []string {
	return []string{
		string(DnsVerificationTestResultFailed),
		string(DnsVerificationTestResultPassed),
		string(DnsVerificationTestResultSkipped),
	}
}

func parseDnsVerificationTestResult(input string) (*DnsVerificationTestResult, error) {
	vals := map[string]DnsVerificationTestResult{
		"failed":  DnsVerificationTestResultFailed,
		"passed":  DnsVerificationTestResultPassed,
		"skipped": DnsVerificationTestResultSkipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DnsVerificationTestResult(input)
	return &out, nil
}

type ForwardProxyConvention string

const (
	ForwardProxyConventionCustom   ForwardProxyConvention = "Custom"
	ForwardProxyConventionNoProxy  ForwardProxyConvention = "NoProxy"
	ForwardProxyConventionStandard ForwardProxyConvention = "Standard"
)

func PossibleValuesForForwardProxyConvention() []string {
	return []string{
		string(ForwardProxyConventionCustom),
		string(ForwardProxyConventionNoProxy),
		string(ForwardProxyConventionStandard),
	}
}

func parseForwardProxyConvention(input string) (*ForwardProxyConvention, error) {
	vals := map[string]ForwardProxyConvention{
		"custom":   ForwardProxyConventionCustom,
		"noproxy":  ForwardProxyConventionNoProxy,
		"standard": ForwardProxyConventionStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForwardProxyConvention(input)
	return &out, nil
}

type FrequencyUnit string

const (
	FrequencyUnitDay  FrequencyUnit = "Day"
	FrequencyUnitHour FrequencyUnit = "Hour"
)

func PossibleValuesForFrequencyUnit() []string {
	return []string{
		string(FrequencyUnitDay),
		string(FrequencyUnitHour),
	}
}

func parseFrequencyUnit(input string) (*FrequencyUnit, error) {
	vals := map[string]FrequencyUnit{
		"day":  FrequencyUnitDay,
		"hour": FrequencyUnitHour,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FrequencyUnit(input)
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

type HostNameType string

const (
	HostNameTypeManaged  HostNameType = "Managed"
	HostNameTypeVerified HostNameType = "Verified"
)

func PossibleValuesForHostNameType() []string {
	return []string{
		string(HostNameTypeManaged),
		string(HostNameTypeVerified),
	}
}

func parseHostNameType(input string) (*HostNameType, error) {
	vals := map[string]HostNameType{
		"managed":  HostNameTypeManaged,
		"verified": HostNameTypeVerified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostNameType(input)
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

type LogLevel string

const (
	LogLevelError       LogLevel = "Error"
	LogLevelInformation LogLevel = "Information"
	LogLevelOff         LogLevel = "Off"
	LogLevelVerbose     LogLevel = "Verbose"
	LogLevelWarning     LogLevel = "Warning"
)

func PossibleValuesForLogLevel() []string {
	return []string{
		string(LogLevelError),
		string(LogLevelInformation),
		string(LogLevelOff),
		string(LogLevelVerbose),
		string(LogLevelWarning),
	}
}

func parseLogLevel(input string) (*LogLevel, error) {
	vals := map[string]LogLevel{
		"error":       LogLevelError,
		"information": LogLevelInformation,
		"off":         LogLevelOff,
		"verbose":     LogLevelVerbose,
		"warning":     LogLevelWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogLevel(input)
	return &out, nil
}

type MSDeployLogEntryType string

const (
	MSDeployLogEntryTypeError   MSDeployLogEntryType = "Error"
	MSDeployLogEntryTypeMessage MSDeployLogEntryType = "Message"
	MSDeployLogEntryTypeWarning MSDeployLogEntryType = "Warning"
)

func PossibleValuesForMSDeployLogEntryType() []string {
	return []string{
		string(MSDeployLogEntryTypeError),
		string(MSDeployLogEntryTypeMessage),
		string(MSDeployLogEntryTypeWarning),
	}
}

func parseMSDeployLogEntryType(input string) (*MSDeployLogEntryType, error) {
	vals := map[string]MSDeployLogEntryType{
		"error":   MSDeployLogEntryTypeError,
		"message": MSDeployLogEntryTypeMessage,
		"warning": MSDeployLogEntryTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MSDeployLogEntryType(input)
	return &out, nil
}

type MSDeployProvisioningState string

const (
	MSDeployProvisioningStateAccepted  MSDeployProvisioningState = "accepted"
	MSDeployProvisioningStateCanceled  MSDeployProvisioningState = "canceled"
	MSDeployProvisioningStateFailed    MSDeployProvisioningState = "failed"
	MSDeployProvisioningStateRunning   MSDeployProvisioningState = "running"
	MSDeployProvisioningStateSucceeded MSDeployProvisioningState = "succeeded"
)

func PossibleValuesForMSDeployProvisioningState() []string {
	return []string{
		string(MSDeployProvisioningStateAccepted),
		string(MSDeployProvisioningStateCanceled),
		string(MSDeployProvisioningStateFailed),
		string(MSDeployProvisioningStateRunning),
		string(MSDeployProvisioningStateSucceeded),
	}
}

func parseMSDeployProvisioningState(input string) (*MSDeployProvisioningState, error) {
	vals := map[string]MSDeployProvisioningState{
		"accepted":  MSDeployProvisioningStateAccepted,
		"canceled":  MSDeployProvisioningStateCanceled,
		"failed":    MSDeployProvisioningStateFailed,
		"running":   MSDeployProvisioningStateRunning,
		"succeeded": MSDeployProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MSDeployProvisioningState(input)
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

type MySqlMigrationType string

const (
	MySqlMigrationTypeLocalToRemote MySqlMigrationType = "LocalToRemote"
	MySqlMigrationTypeRemoteToLocal MySqlMigrationType = "RemoteToLocal"
)

func PossibleValuesForMySqlMigrationType() []string {
	return []string{
		string(MySqlMigrationTypeLocalToRemote),
		string(MySqlMigrationTypeRemoteToLocal),
	}
}

func parseMySqlMigrationType(input string) (*MySqlMigrationType, error) {
	vals := map[string]MySqlMigrationType{
		"localtoremote": MySqlMigrationTypeLocalToRemote,
		"remotetolocal": MySqlMigrationTypeRemoteToLocal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MySqlMigrationType(input)
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

type PublicCertificateLocation string

const (
	PublicCertificateLocationCurrentUserMy  PublicCertificateLocation = "CurrentUserMy"
	PublicCertificateLocationLocalMachineMy PublicCertificateLocation = "LocalMachineMy"
	PublicCertificateLocationUnknown        PublicCertificateLocation = "Unknown"
)

func PossibleValuesForPublicCertificateLocation() []string {
	return []string{
		string(PublicCertificateLocationCurrentUserMy),
		string(PublicCertificateLocationLocalMachineMy),
		string(PublicCertificateLocationUnknown),
	}
}

func parsePublicCertificateLocation(input string) (*PublicCertificateLocation, error) {
	vals := map[string]PublicCertificateLocation{
		"currentusermy":  PublicCertificateLocationCurrentUserMy,
		"localmachinemy": PublicCertificateLocationLocalMachineMy,
		"unknown":        PublicCertificateLocationUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicCertificateLocation(input)
	return &out, nil
}

type PublishingProfileFormat string

const (
	PublishingProfileFormatFileZillaThree PublishingProfileFormat = "FileZilla3"
	PublishingProfileFormatFtp            PublishingProfileFormat = "Ftp"
	PublishingProfileFormatWebDeploy      PublishingProfileFormat = "WebDeploy"
)

func PossibleValuesForPublishingProfileFormat() []string {
	return []string{
		string(PublishingProfileFormatFileZillaThree),
		string(PublishingProfileFormatFtp),
		string(PublishingProfileFormatWebDeploy),
	}
}

func parsePublishingProfileFormat(input string) (*PublishingProfileFormat, error) {
	vals := map[string]PublishingProfileFormat{
		"filezilla3": PublishingProfileFormatFileZillaThree,
		"ftp":        PublishingProfileFormatFtp,
		"webdeploy":  PublishingProfileFormatWebDeploy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublishingProfileFormat(input)
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

type ResolveStatus string

const (
	ResolveStatusAccessToKeyVaultDenied ResolveStatus = "AccessToKeyVaultDenied"
	ResolveStatusFetchTimedOut          ResolveStatus = "FetchTimedOut"
	ResolveStatusInitialized            ResolveStatus = "Initialized"
	ResolveStatusInvalidSyntax          ResolveStatus = "InvalidSyntax"
	ResolveStatusMSINotEnabled          ResolveStatus = "MSINotEnabled"
	ResolveStatusOtherReasons           ResolveStatus = "OtherReasons"
	ResolveStatusResolved               ResolveStatus = "Resolved"
	ResolveStatusSecretNotFound         ResolveStatus = "SecretNotFound"
	ResolveStatusSecretVersionNotFound  ResolveStatus = "SecretVersionNotFound"
	ResolveStatusUnauthorizedClient     ResolveStatus = "UnauthorizedClient"
	ResolveStatusVaultNotFound          ResolveStatus = "VaultNotFound"
)

func PossibleValuesForResolveStatus() []string {
	return []string{
		string(ResolveStatusAccessToKeyVaultDenied),
		string(ResolveStatusFetchTimedOut),
		string(ResolveStatusInitialized),
		string(ResolveStatusInvalidSyntax),
		string(ResolveStatusMSINotEnabled),
		string(ResolveStatusOtherReasons),
		string(ResolveStatusResolved),
		string(ResolveStatusSecretNotFound),
		string(ResolveStatusSecretVersionNotFound),
		string(ResolveStatusUnauthorizedClient),
		string(ResolveStatusVaultNotFound),
	}
}

func parseResolveStatus(input string) (*ResolveStatus, error) {
	vals := map[string]ResolveStatus{
		"accesstokeyvaultdenied": ResolveStatusAccessToKeyVaultDenied,
		"fetchtimedout":          ResolveStatusFetchTimedOut,
		"initialized":            ResolveStatusInitialized,
		"invalidsyntax":          ResolveStatusInvalidSyntax,
		"msinotenabled":          ResolveStatusMSINotEnabled,
		"otherreasons":           ResolveStatusOtherReasons,
		"resolved":               ResolveStatusResolved,
		"secretnotfound":         ResolveStatusSecretNotFound,
		"secretversionnotfound":  ResolveStatusSecretVersionNotFound,
		"unauthorizedclient":     ResolveStatusUnauthorizedClient,
		"vaultnotfound":          ResolveStatusVaultNotFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResolveStatus(input)
	return &out, nil
}

type RouteType string

const (
	RouteTypeDefault   RouteType = "DEFAULT"
	RouteTypeINHERITED RouteType = "INHERITED"
	RouteTypeSTATIC    RouteType = "STATIC"
)

func PossibleValuesForRouteType() []string {
	return []string{
		string(RouteTypeDefault),
		string(RouteTypeINHERITED),
		string(RouteTypeSTATIC),
	}
}

func parseRouteType(input string) (*RouteType, error) {
	vals := map[string]RouteType{
		"default":   RouteTypeDefault,
		"inherited": RouteTypeINHERITED,
		"static":    RouteTypeSTATIC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RouteType(input)
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

type SiteExtensionType string

const (
	SiteExtensionTypeGallery SiteExtensionType = "Gallery"
	SiteExtensionTypeWebRoot SiteExtensionType = "WebRoot"
)

func PossibleValuesForSiteExtensionType() []string {
	return []string{
		string(SiteExtensionTypeGallery),
		string(SiteExtensionTypeWebRoot),
	}
}

func parseSiteExtensionType(input string) (*SiteExtensionType, error) {
	vals := map[string]SiteExtensionType{
		"gallery": SiteExtensionTypeGallery,
		"webroot": SiteExtensionTypeWebRoot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteExtensionType(input)
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

type SiteRuntimeState string

const (
	SiteRuntimeStateREADY   SiteRuntimeState = "READY"
	SiteRuntimeStateSTOPPED SiteRuntimeState = "STOPPED"
	SiteRuntimeStateUNKNOWN SiteRuntimeState = "UNKNOWN"
)

func PossibleValuesForSiteRuntimeState() []string {
	return []string{
		string(SiteRuntimeStateREADY),
		string(SiteRuntimeStateSTOPPED),
		string(SiteRuntimeStateUNKNOWN),
	}
}

func parseSiteRuntimeState(input string) (*SiteRuntimeState, error) {
	vals := map[string]SiteRuntimeState{
		"ready":   SiteRuntimeStateREADY,
		"stopped": SiteRuntimeStateSTOPPED,
		"unknown": SiteRuntimeStateUNKNOWN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteRuntimeState(input)
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

type TriggeredWebJobStatus string

const (
	TriggeredWebJobStatusError   TriggeredWebJobStatus = "Error"
	TriggeredWebJobStatusFailed  TriggeredWebJobStatus = "Failed"
	TriggeredWebJobStatusSuccess TriggeredWebJobStatus = "Success"
)

func PossibleValuesForTriggeredWebJobStatus() []string {
	return []string{
		string(TriggeredWebJobStatusError),
		string(TriggeredWebJobStatusFailed),
		string(TriggeredWebJobStatusSuccess),
	}
}

func parseTriggeredWebJobStatus(input string) (*TriggeredWebJobStatus, error) {
	vals := map[string]TriggeredWebJobStatus{
		"error":   TriggeredWebJobStatusError,
		"failed":  TriggeredWebJobStatusFailed,
		"success": TriggeredWebJobStatusSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggeredWebJobStatus(input)
	return &out, nil
}

type UnauthenticatedClientAction string

const (
	UnauthenticatedClientActionAllowAnonymous      UnauthenticatedClientAction = "AllowAnonymous"
	UnauthenticatedClientActionRedirectToLoginPage UnauthenticatedClientAction = "RedirectToLoginPage"
)

func PossibleValuesForUnauthenticatedClientAction() []string {
	return []string{
		string(UnauthenticatedClientActionAllowAnonymous),
		string(UnauthenticatedClientActionRedirectToLoginPage),
	}
}

func parseUnauthenticatedClientAction(input string) (*UnauthenticatedClientAction, error) {
	vals := map[string]UnauthenticatedClientAction{
		"allowanonymous":      UnauthenticatedClientActionAllowAnonymous,
		"redirecttologinpage": UnauthenticatedClientActionRedirectToLoginPage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnauthenticatedClientAction(input)
	return &out, nil
}

type UnauthenticatedClientActionV2 string

const (
	UnauthenticatedClientActionV2AllowAnonymous      UnauthenticatedClientActionV2 = "AllowAnonymous"
	UnauthenticatedClientActionV2RedirectToLoginPage UnauthenticatedClientActionV2 = "RedirectToLoginPage"
	UnauthenticatedClientActionV2ReturnFourZeroOne   UnauthenticatedClientActionV2 = "Return401"
	UnauthenticatedClientActionV2ReturnFourZeroThree UnauthenticatedClientActionV2 = "Return403"
)

func PossibleValuesForUnauthenticatedClientActionV2() []string {
	return []string{
		string(UnauthenticatedClientActionV2AllowAnonymous),
		string(UnauthenticatedClientActionV2RedirectToLoginPage),
		string(UnauthenticatedClientActionV2ReturnFourZeroOne),
		string(UnauthenticatedClientActionV2ReturnFourZeroThree),
	}
}

func parseUnauthenticatedClientActionV2(input string) (*UnauthenticatedClientActionV2, error) {
	vals := map[string]UnauthenticatedClientActionV2{
		"allowanonymous":      UnauthenticatedClientActionV2AllowAnonymous,
		"redirecttologinpage": UnauthenticatedClientActionV2RedirectToLoginPage,
		"return401":           UnauthenticatedClientActionV2ReturnFourZeroOne,
		"return403":           UnauthenticatedClientActionV2ReturnFourZeroThree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnauthenticatedClientActionV2(input)
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

type WebJobType string

const (
	WebJobTypeContinuous WebJobType = "Continuous"
	WebJobTypeTriggered  WebJobType = "Triggered"
)

func PossibleValuesForWebJobType() []string {
	return []string{
		string(WebJobTypeContinuous),
		string(WebJobTypeTriggered),
	}
}

func parseWebJobType(input string) (*WebJobType, error) {
	vals := map[string]WebJobType{
		"continuous": WebJobTypeContinuous,
		"triggered":  WebJobTypeTriggered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebJobType(input)
	return &out, nil
}

type WorkflowHealthState string

const (
	WorkflowHealthStateHealthy      WorkflowHealthState = "Healthy"
	WorkflowHealthStateNotSpecified WorkflowHealthState = "NotSpecified"
	WorkflowHealthStateUnhealthy    WorkflowHealthState = "Unhealthy"
	WorkflowHealthStateUnknown      WorkflowHealthState = "Unknown"
)

func PossibleValuesForWorkflowHealthState() []string {
	return []string{
		string(WorkflowHealthStateHealthy),
		string(WorkflowHealthStateNotSpecified),
		string(WorkflowHealthStateUnhealthy),
		string(WorkflowHealthStateUnknown),
	}
}

func parseWorkflowHealthState(input string) (*WorkflowHealthState, error) {
	vals := map[string]WorkflowHealthState{
		"healthy":      WorkflowHealthStateHealthy,
		"notspecified": WorkflowHealthStateNotSpecified,
		"unhealthy":    WorkflowHealthStateUnhealthy,
		"unknown":      WorkflowHealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowHealthState(input)
	return &out, nil
}

type WorkflowState string

const (
	WorkflowStateCompleted    WorkflowState = "Completed"
	WorkflowStateDeleted      WorkflowState = "Deleted"
	WorkflowStateDisabled     WorkflowState = "Disabled"
	WorkflowStateEnabled      WorkflowState = "Enabled"
	WorkflowStateNotSpecified WorkflowState = "NotSpecified"
	WorkflowStateSuspended    WorkflowState = "Suspended"
)

func PossibleValuesForWorkflowState() []string {
	return []string{
		string(WorkflowStateCompleted),
		string(WorkflowStateDeleted),
		string(WorkflowStateDisabled),
		string(WorkflowStateEnabled),
		string(WorkflowStateNotSpecified),
		string(WorkflowStateSuspended),
	}
}

func parseWorkflowState(input string) (*WorkflowState, error) {
	vals := map[string]WorkflowState{
		"completed":    WorkflowStateCompleted,
		"deleted":      WorkflowStateDeleted,
		"disabled":     WorkflowStateDisabled,
		"enabled":      WorkflowStateEnabled,
		"notspecified": WorkflowStateNotSpecified,
		"suspended":    WorkflowStateSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowState(input)
	return &out, nil
}
