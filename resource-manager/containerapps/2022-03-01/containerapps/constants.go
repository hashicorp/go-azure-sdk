package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple ActiveRevisionsMode = "Multiple"
	ActiveRevisionsModeSingle   ActiveRevisionsMode = "Single"
)

func PossibleValuesForActiveRevisionsMode() []string {
	return []string{
		string(ActiveRevisionsModeMultiple),
		string(ActiveRevisionsModeSingle),
	}
}

type AppProtocol string

const (
	AppProtocolGrpc AppProtocol = "grpc"
	AppProtocolHTTP AppProtocol = "http"
)

func PossibleValuesForAppProtocol() []string {
	return []string{
		string(AppProtocolGrpc),
		string(AppProtocolHTTP),
	}
}

type BindingType string

const (
	BindingTypeDisabled   BindingType = "Disabled"
	BindingTypeSniEnabled BindingType = "SniEnabled"
)

func PossibleValuesForBindingType() []string {
	return []string{
		string(BindingTypeDisabled),
		string(BindingTypeSniEnabled),
	}
}

type ContainerAppProvisioningState string

const (
	ContainerAppProvisioningStateCanceled   ContainerAppProvisioningState = "Canceled"
	ContainerAppProvisioningStateDeleting   ContainerAppProvisioningState = "Deleting"
	ContainerAppProvisioningStateFailed     ContainerAppProvisioningState = "Failed"
	ContainerAppProvisioningStateInProgress ContainerAppProvisioningState = "InProgress"
	ContainerAppProvisioningStateSucceeded  ContainerAppProvisioningState = "Succeeded"
)

func PossibleValuesForContainerAppProvisioningState() []string {
	return []string{
		string(ContainerAppProvisioningStateCanceled),
		string(ContainerAppProvisioningStateDeleting),
		string(ContainerAppProvisioningStateFailed),
		string(ContainerAppProvisioningStateInProgress),
		string(ContainerAppProvisioningStateSucceeded),
	}
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

type IngressTransportMethod string

const (
	IngressTransportMethodAuto    IngressTransportMethod = "auto"
	IngressTransportMethodHTTP    IngressTransportMethod = "http"
	IngressTransportMethodHTTPTwo IngressTransportMethod = "http2"
)

func PossibleValuesForIngressTransportMethod() []string {
	return []string{
		string(IngressTransportMethodAuto),
		string(IngressTransportMethodHTTP),
		string(IngressTransportMethodHTTPTwo),
	}
}

type Scheme string

const (
	SchemeHTTP  Scheme = "HTTP"
	SchemeHTTPS Scheme = "HTTPS"
)

func PossibleValuesForScheme() []string {
	return []string{
		string(SchemeHTTP),
		string(SchemeHTTPS),
	}
}

type StorageType string

const (
	StorageTypeAzureFile StorageType = "AzureFile"
	StorageTypeEmptyDir  StorageType = "EmptyDir"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypeAzureFile),
		string(StorageTypeEmptyDir),
	}
}

type Type string

const (
	TypeLiveness  Type = "Liveness"
	TypeReadiness Type = "Readiness"
	TypeStartup   Type = "Startup"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeLiveness),
		string(TypeReadiness),
		string(TypeStartup),
	}
}
