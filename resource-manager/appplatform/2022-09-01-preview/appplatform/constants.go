package appplatform

import "strings"

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

func parseApiPortalProvisioningState(input string) (*ApiPortalProvisioningState, error) {
	vals := map[string]ApiPortalProvisioningState{
		"creating":  ApiPortalProvisioningStateCreating,
		"deleting":  ApiPortalProvisioningStateDeleting,
		"failed":    ApiPortalProvisioningStateFailed,
		"succeeded": ApiPortalProvisioningStateSucceeded,
		"updating":  ApiPortalProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiPortalProvisioningState(input)
	return &out, nil
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

func parseAppResourceProvisioningState(input string) (*AppResourceProvisioningState, error) {
	vals := map[string]AppResourceProvisioningState{
		"creating":  AppResourceProvisioningStateCreating,
		"deleting":  AppResourceProvisioningStateDeleting,
		"failed":    AppResourceProvisioningStateFailed,
		"succeeded": AppResourceProvisioningStateSucceeded,
		"updating":  AppResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppResourceProvisioningState(input)
	return &out, nil
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

func parseBackendProtocol(input string) (*BackendProtocol, error) {
	vals := map[string]BackendProtocol{
		"default": BackendProtocolDefault,
		"grpc":    BackendProtocolGRPC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackendProtocol(input)
	return &out, nil
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

func parseBindingType(input string) (*BindingType, error) {
	vals := map[string]BindingType{
		"apacheskywalking":    BindingTypeApacheSkyWalking,
		"appdynamics":         BindingTypeAppDynamics,
		"applicationinsights": BindingTypeApplicationInsights,
		"dynatrace":           BindingTypeDynatrace,
		"elasticapm":          BindingTypeElasticAPM,
		"newrelic":            BindingTypeNewRelic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BindingType(input)
	return &out, nil
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

func parseBuildProvisioningState(input string) (*BuildProvisioningState, error) {
	vals := map[string]BuildProvisioningState{
		"creating":  BuildProvisioningStateCreating,
		"deleting":  BuildProvisioningStateDeleting,
		"failed":    BuildProvisioningStateFailed,
		"succeeded": BuildProvisioningStateSucceeded,
		"updating":  BuildProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildProvisioningState(input)
	return &out, nil
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

func parseBuildResultProvisioningState(input string) (*BuildResultProvisioningState, error) {
	vals := map[string]BuildResultProvisioningState{
		"building":  BuildResultProvisioningStateBuilding,
		"deleting":  BuildResultProvisioningStateDeleting,
		"failed":    BuildResultProvisioningStateFailed,
		"queuing":   BuildResultProvisioningStateQueuing,
		"succeeded": BuildResultProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildResultProvisioningState(input)
	return &out, nil
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

func parseBuildServiceProvisioningState(input string) (*BuildServiceProvisioningState, error) {
	vals := map[string]BuildServiceProvisioningState{
		"creating":  BuildServiceProvisioningStateCreating,
		"deleting":  BuildServiceProvisioningStateDeleting,
		"failed":    BuildServiceProvisioningStateFailed,
		"succeeded": BuildServiceProvisioningStateSucceeded,
		"updating":  BuildServiceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildServiceProvisioningState(input)
	return &out, nil
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

func parseBuilderProvisioningState(input string) (*BuilderProvisioningState, error) {
	vals := map[string]BuilderProvisioningState{
		"creating":  BuilderProvisioningStateCreating,
		"deleting":  BuilderProvisioningStateDeleting,
		"failed":    BuilderProvisioningStateFailed,
		"succeeded": BuilderProvisioningStateSucceeded,
		"updating":  BuilderProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuilderProvisioningState(input)
	return &out, nil
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

func parseBuildpackBindingProvisioningState(input string) (*BuildpackBindingProvisioningState, error) {
	vals := map[string]BuildpackBindingProvisioningState{
		"creating":  BuildpackBindingProvisioningStateCreating,
		"deleting":  BuildpackBindingProvisioningStateDeleting,
		"failed":    BuildpackBindingProvisioningStateFailed,
		"succeeded": BuildpackBindingProvisioningStateSucceeded,
		"updating":  BuildpackBindingProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuildpackBindingProvisioningState(input)
	return &out, nil
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

func parseCertificateResourceProvisioningState(input string) (*CertificateResourceProvisioningState, error) {
	vals := map[string]CertificateResourceProvisioningState{
		"creating":  CertificateResourceProvisioningStateCreating,
		"deleting":  CertificateResourceProvisioningStateDeleting,
		"failed":    CertificateResourceProvisioningStateFailed,
		"succeeded": CertificateResourceProvisioningStateSucceeded,
		"updating":  CertificateResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateResourceProvisioningState(input)
	return &out, nil
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

func parseConfigServerState(input string) (*ConfigServerState, error) {
	vals := map[string]ConfigServerState{
		"deleted":      ConfigServerStateDeleted,
		"failed":       ConfigServerStateFailed,
		"notavailable": ConfigServerStateNotAvailable,
		"succeeded":    ConfigServerStateSucceeded,
		"updating":     ConfigServerStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigServerState(input)
	return &out, nil
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

func parseConfigurationServiceProvisioningState(input string) (*ConfigurationServiceProvisioningState, error) {
	vals := map[string]ConfigurationServiceProvisioningState{
		"creating":  ConfigurationServiceProvisioningStateCreating,
		"deleting":  ConfigurationServiceProvisioningStateDeleting,
		"failed":    ConfigurationServiceProvisioningStateFailed,
		"succeeded": ConfigurationServiceProvisioningStateSucceeded,
		"updating":  ConfigurationServiceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationServiceProvisioningState(input)
	return &out, nil
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

func parseCustomDomainResourceProvisioningState(input string) (*CustomDomainResourceProvisioningState, error) {
	vals := map[string]CustomDomainResourceProvisioningState{
		"creating":  CustomDomainResourceProvisioningStateCreating,
		"deleting":  CustomDomainResourceProvisioningStateDeleting,
		"failed":    CustomDomainResourceProvisioningStateFailed,
		"succeeded": CustomDomainResourceProvisioningStateSucceeded,
		"updating":  CustomDomainResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomDomainResourceProvisioningState(input)
	return &out, nil
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

func parseDeploymentResourceProvisioningState(input string) (*DeploymentResourceProvisioningState, error) {
	vals := map[string]DeploymentResourceProvisioningState{
		"creating":  DeploymentResourceProvisioningStateCreating,
		"failed":    DeploymentResourceProvisioningStateFailed,
		"succeeded": DeploymentResourceProvisioningStateSucceeded,
		"updating":  DeploymentResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentResourceProvisioningState(input)
	return &out, nil
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

func parseDeploymentResourceStatus(input string) (*DeploymentResourceStatus, error) {
	vals := map[string]DeploymentResourceStatus{
		"running": DeploymentResourceStatusRunning,
		"stopped": DeploymentResourceStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentResourceStatus(input)
	return &out, nil
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

func parseGatewayProvisioningState(input string) (*GatewayProvisioningState, error) {
	vals := map[string]GatewayProvisioningState{
		"creating":  GatewayProvisioningStateCreating,
		"deleting":  GatewayProvisioningStateDeleting,
		"failed":    GatewayProvisioningStateFailed,
		"succeeded": GatewayProvisioningStateSucceeded,
		"updating":  GatewayProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GatewayProvisioningState(input)
	return &out, nil
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

func parseGatewayRouteConfigProtocol(input string) (*GatewayRouteConfigProtocol, error) {
	vals := map[string]GatewayRouteConfigProtocol{
		"http":  GatewayRouteConfigProtocolHTTP,
		"https": GatewayRouteConfigProtocolHTTPS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GatewayRouteConfigProtocol(input)
	return &out, nil
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

func parseHTTPSchemeType(input string) (*HTTPSchemeType, error) {
	vals := map[string]HTTPSchemeType{
		"http":  HTTPSchemeTypeHTTP,
		"https": HTTPSchemeTypeHTTPS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HTTPSchemeType(input)
	return &out, nil
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

func parseKPackBuildStageProvisioningState(input string) (*KPackBuildStageProvisioningState, error) {
	vals := map[string]KPackBuildStageProvisioningState{
		"failed":     KPackBuildStageProvisioningStateFailed,
		"notstarted": KPackBuildStageProvisioningStateNotStarted,
		"running":    KPackBuildStageProvisioningStateRunning,
		"succeeded":  KPackBuildStageProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KPackBuildStageProvisioningState(input)
	return &out, nil
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

func parseMonitoringSettingState(input string) (*MonitoringSettingState, error) {
	vals := map[string]MonitoringSettingState{
		"failed":       MonitoringSettingStateFailed,
		"notavailable": MonitoringSettingStateNotAvailable,
		"succeeded":    MonitoringSettingStateSucceeded,
		"updating":     MonitoringSettingStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringSettingState(input)
	return &out, nil
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

func parsePowerState(input string) (*PowerState, error) {
	vals := map[string]PowerState{
		"running": PowerStateRunning,
		"stopped": PowerStateStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PowerState(input)
	return &out, nil
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

func parseProbeActionType(input string) (*ProbeActionType, error) {
	vals := map[string]ProbeActionType{
		"execaction":      ProbeActionTypeExecAction,
		"httpgetaction":   ProbeActionTypeHTTPGetAction,
		"tcpsocketaction": ProbeActionTypeTCPSocketAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProbeActionType(input)
	return &out, nil
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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":   ProvisioningStateCreating,
		"deleted":    ProvisioningStateDeleted,
		"deleting":   ProvisioningStateDeleting,
		"failed":     ProvisioningStateFailed,
		"movefailed": ProvisioningStateMoveFailed,
		"moved":      ProvisioningStateMoved,
		"moving":     ProvisioningStateMoving,
		"starting":   ProvisioningStateStarting,
		"stopping":   ProvisioningStateStopping,
		"succeeded":  ProvisioningStateSucceeded,
		"updating":   ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
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

func parseResourceSkuRestrictionsReasonCode(input string) (*ResourceSkuRestrictionsReasonCode, error) {
	vals := map[string]ResourceSkuRestrictionsReasonCode{
		"notavailableforsubscription": ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription,
		"quotaid":                     ResourceSkuRestrictionsReasonCodeQuotaId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceSkuRestrictionsReasonCode(input)
	return &out, nil
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

func parseResourceSkuRestrictionsType(input string) (*ResourceSkuRestrictionsType, error) {
	vals := map[string]ResourceSkuRestrictionsType{
		"location": ResourceSkuRestrictionsTypeLocation,
		"zone":     ResourceSkuRestrictionsTypeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceSkuRestrictionsType(input)
	return &out, nil
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

func parseServiceRegistryProvisioningState(input string) (*ServiceRegistryProvisioningState, error) {
	vals := map[string]ServiceRegistryProvisioningState{
		"creating":  ServiceRegistryProvisioningStateCreating,
		"deleting":  ServiceRegistryProvisioningStateDeleting,
		"failed":    ServiceRegistryProvisioningStateFailed,
		"succeeded": ServiceRegistryProvisioningStateSucceeded,
		"updating":  ServiceRegistryProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceRegistryProvisioningState(input)
	return &out, nil
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

func parseSessionAffinity(input string) (*SessionAffinity, error) {
	vals := map[string]SessionAffinity{
		"cookie": SessionAffinityCookie,
		"none":   SessionAffinityNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionAffinity(input)
	return &out, nil
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

func parseSkuScaleType(input string) (*SkuScaleType, error) {
	vals := map[string]SkuScaleType{
		"automatic": SkuScaleTypeAutomatic,
		"manual":    SkuScaleTypeManual,
		"none":      SkuScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuScaleType(input)
	return &out, nil
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

func parseStorageType(input string) (*StorageType, error) {
	vals := map[string]StorageType{
		"storageaccount": StorageTypeStorageAccount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageType(input)
	return &out, nil
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

func parseSupportedRuntimePlatform(input string) (*SupportedRuntimePlatform, error) {
	vals := map[string]SupportedRuntimePlatform{
		"java":      SupportedRuntimePlatformJava,
		".net core": SupportedRuntimePlatformPointNETCore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportedRuntimePlatform(input)
	return &out, nil
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

func parseSupportedRuntimeValue(input string) (*SupportedRuntimeValue, error) {
	vals := map[string]SupportedRuntimeValue{
		"java_8":     SupportedRuntimeValueJavaEight,
		"java_11":    SupportedRuntimeValueJavaOneOne,
		"java_17":    SupportedRuntimeValueJavaOneSeven,
		"netcore_31": SupportedRuntimeValueNetCoreThreeOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportedRuntimeValue(input)
	return &out, nil
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

func parseTestKeyType(input string) (*TestKeyType, error) {
	vals := map[string]TestKeyType{
		"primary":   TestKeyTypePrimary,
		"secondary": TestKeyTypeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TestKeyType(input)
	return &out, nil
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

func parseTrafficDirection(input string) (*TrafficDirection, error) {
	vals := map[string]TrafficDirection{
		"inbound":  TrafficDirectionInbound,
		"outbound": TrafficDirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrafficDirection(input)
	return &out, nil
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

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"azurefilevolume": TypeAzureFileVolume,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
