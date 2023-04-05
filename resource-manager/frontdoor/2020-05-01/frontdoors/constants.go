package frontdoors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackendEnabledState string

const (
	BackendEnabledStateDisabled BackendEnabledState = "Disabled"
	BackendEnabledStateEnabled  BackendEnabledState = "Enabled"
)

func PossibleValuesForBackendEnabledState() []string {
	return []string{
		string(BackendEnabledStateDisabled),
		string(BackendEnabledStateEnabled),
	}
}

type CustomHTTPSProvisioningState string

const (
	CustomHTTPSProvisioningStateDisabled  CustomHTTPSProvisioningState = "Disabled"
	CustomHTTPSProvisioningStateDisabling CustomHTTPSProvisioningState = "Disabling"
	CustomHTTPSProvisioningStateEnabled   CustomHTTPSProvisioningState = "Enabled"
	CustomHTTPSProvisioningStateEnabling  CustomHTTPSProvisioningState = "Enabling"
	CustomHTTPSProvisioningStateFailed    CustomHTTPSProvisioningState = "Failed"
)

func PossibleValuesForCustomHTTPSProvisioningState() []string {
	return []string{
		string(CustomHTTPSProvisioningStateDisabled),
		string(CustomHTTPSProvisioningStateDisabling),
		string(CustomHTTPSProvisioningStateEnabled),
		string(CustomHTTPSProvisioningStateEnabling),
		string(CustomHTTPSProvisioningStateFailed),
	}
}

type CustomHTTPSProvisioningSubstate string

const (
	CustomHTTPSProvisioningSubstateCertificateDeleted                            CustomHTTPSProvisioningSubstate = "CertificateDeleted"
	CustomHTTPSProvisioningSubstateCertificateDeployed                           CustomHTTPSProvisioningSubstate = "CertificateDeployed"
	CustomHTTPSProvisioningSubstateDeletingCertificate                           CustomHTTPSProvisioningSubstate = "DeletingCertificate"
	CustomHTTPSProvisioningSubstateDeployingCertificate                          CustomHTTPSProvisioningSubstate = "DeployingCertificate"
	CustomHTTPSProvisioningSubstateDomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = "DomainControlValidationRequestApproved"
	CustomHTTPSProvisioningSubstateDomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = "DomainControlValidationRequestRejected"
	CustomHTTPSProvisioningSubstateDomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = "DomainControlValidationRequestTimedOut"
	CustomHTTPSProvisioningSubstateIssuingCertificate                            CustomHTTPSProvisioningSubstate = "IssuingCertificate"
	CustomHTTPSProvisioningSubstatePendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = "PendingDomainControlValidationREquestApproval"
	CustomHTTPSProvisioningSubstateSubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = "SubmittingDomainControlValidationRequest"
)

func PossibleValuesForCustomHTTPSProvisioningSubstate() []string {
	return []string{
		string(CustomHTTPSProvisioningSubstateCertificateDeleted),
		string(CustomHTTPSProvisioningSubstateCertificateDeployed),
		string(CustomHTTPSProvisioningSubstateDeletingCertificate),
		string(CustomHTTPSProvisioningSubstateDeployingCertificate),
		string(CustomHTTPSProvisioningSubstateDomainControlValidationRequestApproved),
		string(CustomHTTPSProvisioningSubstateDomainControlValidationRequestRejected),
		string(CustomHTTPSProvisioningSubstateDomainControlValidationRequestTimedOut),
		string(CustomHTTPSProvisioningSubstateIssuingCertificate),
		string(CustomHTTPSProvisioningSubstatePendingDomainControlValidationREquestApproval),
		string(CustomHTTPSProvisioningSubstateSubmittingDomainControlValidationRequest),
	}
}

type DynamicCompressionEnabled string

const (
	DynamicCompressionEnabledDisabled DynamicCompressionEnabled = "Disabled"
	DynamicCompressionEnabledEnabled  DynamicCompressionEnabled = "Enabled"
)

func PossibleValuesForDynamicCompressionEnabled() []string {
	return []string{
		string(DynamicCompressionEnabledDisabled),
		string(DynamicCompressionEnabledEnabled),
	}
}

type EnforceCertificateNameCheckEnabledState string

const (
	EnforceCertificateNameCheckEnabledStateDisabled EnforceCertificateNameCheckEnabledState = "Disabled"
	EnforceCertificateNameCheckEnabledStateEnabled  EnforceCertificateNameCheckEnabledState = "Enabled"
)

func PossibleValuesForEnforceCertificateNameCheckEnabledState() []string {
	return []string{
		string(EnforceCertificateNameCheckEnabledStateDisabled),
		string(EnforceCertificateNameCheckEnabledStateEnabled),
	}
}

type FrontDoorCertificateSource string

const (
	FrontDoorCertificateSourceAzureKeyVault FrontDoorCertificateSource = "AzureKeyVault"
	FrontDoorCertificateSourceFrontDoor     FrontDoorCertificateSource = "FrontDoor"
)

func PossibleValuesForFrontDoorCertificateSource() []string {
	return []string{
		string(FrontDoorCertificateSourceAzureKeyVault),
		string(FrontDoorCertificateSourceFrontDoor),
	}
}

type FrontDoorCertificateType string

const (
	FrontDoorCertificateTypeDedicated FrontDoorCertificateType = "Dedicated"
)

func PossibleValuesForFrontDoorCertificateType() []string {
	return []string{
		string(FrontDoorCertificateTypeDedicated),
	}
}

type FrontDoorEnabledState string

const (
	FrontDoorEnabledStateDisabled FrontDoorEnabledState = "Disabled"
	FrontDoorEnabledStateEnabled  FrontDoorEnabledState = "Enabled"
)

func PossibleValuesForFrontDoorEnabledState() []string {
	return []string{
		string(FrontDoorEnabledStateDisabled),
		string(FrontDoorEnabledStateEnabled),
	}
}

type FrontDoorForwardingProtocol string

const (
	FrontDoorForwardingProtocolHTTPOnly     FrontDoorForwardingProtocol = "HttpOnly"
	FrontDoorForwardingProtocolHTTPSOnly    FrontDoorForwardingProtocol = "HttpsOnly"
	FrontDoorForwardingProtocolMatchRequest FrontDoorForwardingProtocol = "MatchRequest"
)

func PossibleValuesForFrontDoorForwardingProtocol() []string {
	return []string{
		string(FrontDoorForwardingProtocolHTTPOnly),
		string(FrontDoorForwardingProtocolHTTPSOnly),
		string(FrontDoorForwardingProtocolMatchRequest),
	}
}

type FrontDoorHealthProbeMethod string

const (
	FrontDoorHealthProbeMethodGET  FrontDoorHealthProbeMethod = "GET"
	FrontDoorHealthProbeMethodHEAD FrontDoorHealthProbeMethod = "HEAD"
)

func PossibleValuesForFrontDoorHealthProbeMethod() []string {
	return []string{
		string(FrontDoorHealthProbeMethodGET),
		string(FrontDoorHealthProbeMethodHEAD),
	}
}

type FrontDoorProtocol string

const (
	FrontDoorProtocolHTTP  FrontDoorProtocol = "Http"
	FrontDoorProtocolHTTPS FrontDoorProtocol = "Https"
)

func PossibleValuesForFrontDoorProtocol() []string {
	return []string{
		string(FrontDoorProtocolHTTP),
		string(FrontDoorProtocolHTTPS),
	}
}

type FrontDoorQuery string

const (
	FrontDoorQueryStripAll       FrontDoorQuery = "StripAll"
	FrontDoorQueryStripAllExcept FrontDoorQuery = "StripAllExcept"
	FrontDoorQueryStripNone      FrontDoorQuery = "StripNone"
	FrontDoorQueryStripOnly      FrontDoorQuery = "StripOnly"
)

func PossibleValuesForFrontDoorQuery() []string {
	return []string{
		string(FrontDoorQueryStripAll),
		string(FrontDoorQueryStripAllExcept),
		string(FrontDoorQueryStripNone),
		string(FrontDoorQueryStripOnly),
	}
}

type FrontDoorRedirectProtocol string

const (
	FrontDoorRedirectProtocolHTTPOnly     FrontDoorRedirectProtocol = "HttpOnly"
	FrontDoorRedirectProtocolHTTPSOnly    FrontDoorRedirectProtocol = "HttpsOnly"
	FrontDoorRedirectProtocolMatchRequest FrontDoorRedirectProtocol = "MatchRequest"
)

func PossibleValuesForFrontDoorRedirectProtocol() []string {
	return []string{
		string(FrontDoorRedirectProtocolHTTPOnly),
		string(FrontDoorRedirectProtocolHTTPSOnly),
		string(FrontDoorRedirectProtocolMatchRequest),
	}
}

type FrontDoorRedirectType string

const (
	FrontDoorRedirectTypeFound             FrontDoorRedirectType = "Found"
	FrontDoorRedirectTypeMoved             FrontDoorRedirectType = "Moved"
	FrontDoorRedirectTypePermanentRedirect FrontDoorRedirectType = "PermanentRedirect"
	FrontDoorRedirectTypeTemporaryRedirect FrontDoorRedirectType = "TemporaryRedirect"
)

func PossibleValuesForFrontDoorRedirectType() []string {
	return []string{
		string(FrontDoorRedirectTypeFound),
		string(FrontDoorRedirectTypeMoved),
		string(FrontDoorRedirectTypePermanentRedirect),
		string(FrontDoorRedirectTypeTemporaryRedirect),
	}
}

type FrontDoorResourceState string

const (
	FrontDoorResourceStateCreating  FrontDoorResourceState = "Creating"
	FrontDoorResourceStateDeleting  FrontDoorResourceState = "Deleting"
	FrontDoorResourceStateDisabled  FrontDoorResourceState = "Disabled"
	FrontDoorResourceStateDisabling FrontDoorResourceState = "Disabling"
	FrontDoorResourceStateEnabled   FrontDoorResourceState = "Enabled"
	FrontDoorResourceStateEnabling  FrontDoorResourceState = "Enabling"
)

func PossibleValuesForFrontDoorResourceState() []string {
	return []string{
		string(FrontDoorResourceStateCreating),
		string(FrontDoorResourceStateDeleting),
		string(FrontDoorResourceStateDisabled),
		string(FrontDoorResourceStateDisabling),
		string(FrontDoorResourceStateEnabled),
		string(FrontDoorResourceStateEnabling),
	}
}

type FrontDoorTlsProtocolType string

const (
	FrontDoorTlsProtocolTypeServerNameIndication FrontDoorTlsProtocolType = "ServerNameIndication"
)

func PossibleValuesForFrontDoorTlsProtocolType() []string {
	return []string{
		string(FrontDoorTlsProtocolTypeServerNameIndication),
	}
}

type HeaderActionType string

const (
	HeaderActionTypeAppend    HeaderActionType = "Append"
	HeaderActionTypeDelete    HeaderActionType = "Delete"
	HeaderActionTypeOverwrite HeaderActionType = "Overwrite"
)

func PossibleValuesForHeaderActionType() []string {
	return []string{
		string(HeaderActionTypeAppend),
		string(HeaderActionTypeDelete),
		string(HeaderActionTypeOverwrite),
	}
}

type HealthProbeEnabled string

const (
	HealthProbeEnabledDisabled HealthProbeEnabled = "Disabled"
	HealthProbeEnabledEnabled  HealthProbeEnabled = "Enabled"
)

func PossibleValuesForHealthProbeEnabled() []string {
	return []string{
		string(HealthProbeEnabledDisabled),
		string(HealthProbeEnabledEnabled),
	}
}

type MatchProcessingBehavior string

const (
	MatchProcessingBehaviorContinue MatchProcessingBehavior = "Continue"
	MatchProcessingBehaviorStop     MatchProcessingBehavior = "Stop"
)

func PossibleValuesForMatchProcessingBehavior() []string {
	return []string{
		string(MatchProcessingBehaviorContinue),
		string(MatchProcessingBehaviorStop),
	}
}

type MinimumTLSVersion string

const (
	MinimumTLSVersionOnePointTwo  MinimumTLSVersion = "1.2"
	MinimumTLSVersionOnePointZero MinimumTLSVersion = "1.0"
)

func PossibleValuesForMinimumTLSVersion() []string {
	return []string{
		string(MinimumTLSVersionOnePointTwo),
		string(MinimumTLSVersionOnePointZero),
	}
}

type PrivateEndpointStatus string

const (
	PrivateEndpointStatusApproved     PrivateEndpointStatus = "Approved"
	PrivateEndpointStatusDisconnected PrivateEndpointStatus = "Disconnected"
	PrivateEndpointStatusPending      PrivateEndpointStatus = "Pending"
	PrivateEndpointStatusRejected     PrivateEndpointStatus = "Rejected"
	PrivateEndpointStatusTimeout      PrivateEndpointStatus = "Timeout"
)

func PossibleValuesForPrivateEndpointStatus() []string {
	return []string{
		string(PrivateEndpointStatusApproved),
		string(PrivateEndpointStatusDisconnected),
		string(PrivateEndpointStatusPending),
		string(PrivateEndpointStatusRejected),
		string(PrivateEndpointStatusTimeout),
	}
}

type RoutingRuleEnabledState string

const (
	RoutingRuleEnabledStateDisabled RoutingRuleEnabledState = "Disabled"
	RoutingRuleEnabledStateEnabled  RoutingRuleEnabledState = "Enabled"
)

func PossibleValuesForRoutingRuleEnabledState() []string {
	return []string{
		string(RoutingRuleEnabledStateDisabled),
		string(RoutingRuleEnabledStateEnabled),
	}
}

type RulesEngineMatchVariable string

const (
	RulesEngineMatchVariableIsMobile                 RulesEngineMatchVariable = "IsMobile"
	RulesEngineMatchVariablePostArgs                 RulesEngineMatchVariable = "PostArgs"
	RulesEngineMatchVariableQueryString              RulesEngineMatchVariable = "QueryString"
	RulesEngineMatchVariableRemoteAddr               RulesEngineMatchVariable = "RemoteAddr"
	RulesEngineMatchVariableRequestBody              RulesEngineMatchVariable = "RequestBody"
	RulesEngineMatchVariableRequestFilename          RulesEngineMatchVariable = "RequestFilename"
	RulesEngineMatchVariableRequestFilenameExtension RulesEngineMatchVariable = "RequestFilenameExtension"
	RulesEngineMatchVariableRequestHeader            RulesEngineMatchVariable = "RequestHeader"
	RulesEngineMatchVariableRequestMethod            RulesEngineMatchVariable = "RequestMethod"
	RulesEngineMatchVariableRequestPath              RulesEngineMatchVariable = "RequestPath"
	RulesEngineMatchVariableRequestScheme            RulesEngineMatchVariable = "RequestScheme"
	RulesEngineMatchVariableRequestUri               RulesEngineMatchVariable = "RequestUri"
)

func PossibleValuesForRulesEngineMatchVariable() []string {
	return []string{
		string(RulesEngineMatchVariableIsMobile),
		string(RulesEngineMatchVariablePostArgs),
		string(RulesEngineMatchVariableQueryString),
		string(RulesEngineMatchVariableRemoteAddr),
		string(RulesEngineMatchVariableRequestBody),
		string(RulesEngineMatchVariableRequestFilename),
		string(RulesEngineMatchVariableRequestFilenameExtension),
		string(RulesEngineMatchVariableRequestHeader),
		string(RulesEngineMatchVariableRequestMethod),
		string(RulesEngineMatchVariableRequestPath),
		string(RulesEngineMatchVariableRequestScheme),
		string(RulesEngineMatchVariableRequestUri),
	}
}

type RulesEngineOperator string

const (
	RulesEngineOperatorAny                RulesEngineOperator = "Any"
	RulesEngineOperatorBeginsWith         RulesEngineOperator = "BeginsWith"
	RulesEngineOperatorContains           RulesEngineOperator = "Contains"
	RulesEngineOperatorEndsWith           RulesEngineOperator = "EndsWith"
	RulesEngineOperatorEqual              RulesEngineOperator = "Equal"
	RulesEngineOperatorGeoMatch           RulesEngineOperator = "GeoMatch"
	RulesEngineOperatorGreaterThan        RulesEngineOperator = "GreaterThan"
	RulesEngineOperatorGreaterThanOrEqual RulesEngineOperator = "GreaterThanOrEqual"
	RulesEngineOperatorIPMatch            RulesEngineOperator = "IPMatch"
	RulesEngineOperatorLessThan           RulesEngineOperator = "LessThan"
	RulesEngineOperatorLessThanOrEqual    RulesEngineOperator = "LessThanOrEqual"
)

func PossibleValuesForRulesEngineOperator() []string {
	return []string{
		string(RulesEngineOperatorAny),
		string(RulesEngineOperatorBeginsWith),
		string(RulesEngineOperatorContains),
		string(RulesEngineOperatorEndsWith),
		string(RulesEngineOperatorEqual),
		string(RulesEngineOperatorGeoMatch),
		string(RulesEngineOperatorGreaterThan),
		string(RulesEngineOperatorGreaterThanOrEqual),
		string(RulesEngineOperatorIPMatch),
		string(RulesEngineOperatorLessThan),
		string(RulesEngineOperatorLessThanOrEqual),
	}
}

type SessionAffinityEnabledState string

const (
	SessionAffinityEnabledStateDisabled SessionAffinityEnabledState = "Disabled"
	SessionAffinityEnabledStateEnabled  SessionAffinityEnabledState = "Enabled"
)

func PossibleValuesForSessionAffinityEnabledState() []string {
	return []string{
		string(SessionAffinityEnabledStateDisabled),
		string(SessionAffinityEnabledStateEnabled),
	}
}

type Transform string

const (
	TransformLowercase   Transform = "Lowercase"
	TransformRemoveNulls Transform = "RemoveNulls"
	TransformTrim        Transform = "Trim"
	TransformUppercase   Transform = "Uppercase"
	TransformUrlDecode   Transform = "UrlDecode"
	TransformUrlEncode   Transform = "UrlEncode"
)

func PossibleValuesForTransform() []string {
	return []string{
		string(TransformLowercase),
		string(TransformRemoveNulls),
		string(TransformTrim),
		string(TransformUppercase),
		string(TransformUrlDecode),
		string(TransformUrlEncode),
	}
}
