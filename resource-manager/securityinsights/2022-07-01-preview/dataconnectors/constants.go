package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityStatus int64

const (
	AvailabilityStatusOne AvailabilityStatus = 1
)

func PossibleValuesForAvailabilityStatus() []int64 {
	return []int64{
		int64(AvailabilityStatusOne),
	}
}

type ConnectivityType string

const (
	ConnectivityTypeIsConnectedQuery ConnectivityType = "IsConnectedQuery"
)

func PossibleValuesForConnectivityType() []string {
	return []string{
		string(ConnectivityTypeIsConnectedQuery),
	}
}

type DataConnectorKind string

const (
	DataConnectorKindAPIPolling                                DataConnectorKind = "APIPolling"
	DataConnectorKindAmazonWebServicesCloudTrail               DataConnectorKind = "AmazonWebServicesCloudTrail"
	DataConnectorKindAmazonWebServicesSThree                   DataConnectorKind = "AmazonWebServicesS3"
	DataConnectorKindAzureActiveDirectory                      DataConnectorKind = "AzureActiveDirectory"
	DataConnectorKindAzureAdvancedThreatProtection             DataConnectorKind = "AzureAdvancedThreatProtection"
	DataConnectorKindAzureSecurityCenter                       DataConnectorKind = "AzureSecurityCenter"
	DataConnectorKindDynamicsThreeSixFive                      DataConnectorKind = "Dynamics365"
	DataConnectorKindGenericUI                                 DataConnectorKind = "GenericUI"
	DataConnectorKindIOT                                       DataConnectorKind = "IOT"
	DataConnectorKindMicrosoftCloudAppSecurity                 DataConnectorKind = "MicrosoftCloudAppSecurity"
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection DataConnectorKind = "MicrosoftDefenderAdvancedThreatProtection"
	DataConnectorKindMicrosoftThreatIntelligence               DataConnectorKind = "MicrosoftThreatIntelligence"
	DataConnectorKindMicrosoftThreatProtection                 DataConnectorKind = "MicrosoftThreatProtection"
	DataConnectorKindOfficeATP                                 DataConnectorKind = "OfficeATP"
	DataConnectorKindOfficeIRM                                 DataConnectorKind = "OfficeIRM"
	DataConnectorKindOfficePowerBI                             DataConnectorKind = "OfficePowerBI"
	DataConnectorKindOfficeThreeSixFive                        DataConnectorKind = "Office365"
	DataConnectorKindOfficeThreeSixFiveProject                 DataConnectorKind = "Office365Project"
	DataConnectorKindThreatIntelligence                        DataConnectorKind = "ThreatIntelligence"
	DataConnectorKindThreatIntelligenceTaxii                   DataConnectorKind = "ThreatIntelligenceTaxii"
)

func PossibleValuesForDataConnectorKind() []string {
	return []string{
		string(DataConnectorKindAPIPolling),
		string(DataConnectorKindAmazonWebServicesCloudTrail),
		string(DataConnectorKindAmazonWebServicesSThree),
		string(DataConnectorKindAzureActiveDirectory),
		string(DataConnectorKindAzureAdvancedThreatProtection),
		string(DataConnectorKindAzureSecurityCenter),
		string(DataConnectorKindDynamicsThreeSixFive),
		string(DataConnectorKindGenericUI),
		string(DataConnectorKindIOT),
		string(DataConnectorKindMicrosoftCloudAppSecurity),
		string(DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
		string(DataConnectorKindMicrosoftThreatIntelligence),
		string(DataConnectorKindMicrosoftThreatProtection),
		string(DataConnectorKindOfficeATP),
		string(DataConnectorKindOfficeIRM),
		string(DataConnectorKindOfficePowerBI),
		string(DataConnectorKindOfficeThreeSixFive),
		string(DataConnectorKindOfficeThreeSixFiveProject),
		string(DataConnectorKindThreatIntelligence),
		string(DataConnectorKindThreatIntelligenceTaxii),
	}
}

type DataTypeState string

const (
	DataTypeStateDisabled DataTypeState = "Disabled"
	DataTypeStateEnabled  DataTypeState = "Enabled"
)

func PossibleValuesForDataTypeState() []string {
	return []string{
		string(DataTypeStateDisabled),
		string(DataTypeStateEnabled),
	}
}

type PermissionProviderScope string

const (
	PermissionProviderScopeResourceGroup PermissionProviderScope = "ResourceGroup"
	PermissionProviderScopeSubscription  PermissionProviderScope = "Subscription"
	PermissionProviderScopeWorkspace     PermissionProviderScope = "Workspace"
)

func PossibleValuesForPermissionProviderScope() []string {
	return []string{
		string(PermissionProviderScopeResourceGroup),
		string(PermissionProviderScopeSubscription),
		string(PermissionProviderScopeWorkspace),
	}
}

type PollingFrequency string

const (
	PollingFrequencyOnceADay    PollingFrequency = "OnceADay"
	PollingFrequencyOnceAMinute PollingFrequency = "OnceAMinute"
	PollingFrequencyOnceAnHour  PollingFrequency = "OnceAnHour"
)

func PossibleValuesForPollingFrequency() []string {
	return []string{
		string(PollingFrequencyOnceADay),
		string(PollingFrequencyOnceAMinute),
		string(PollingFrequencyOnceAnHour),
	}
}

type ProviderName string

const (
	ProviderNameMicrosoftPointAuthorizationPolicyAssignments           ProviderName = "Microsoft.Authorization/policyAssignments"
	ProviderNameMicrosoftPointOperationalInsightsSolutions             ProviderName = "Microsoft.OperationalInsights/solutions"
	ProviderNameMicrosoftPointOperationalInsightsWorkspaces            ProviderName = "Microsoft.OperationalInsights/workspaces"
	ProviderNameMicrosoftPointOperationalInsightsWorkspacesDatasources ProviderName = "Microsoft.OperationalInsights/workspaces/datasources"
	ProviderNameMicrosoftPointOperationalInsightsWorkspacesSharedKeys  ProviderName = "Microsoft.OperationalInsights/workspaces/sharedKeys"
	ProviderNameMicrosoftPointaadiamDiagnosticSettings                 ProviderName = "microsoft.aadiam/diagnosticSettings"
)

func PossibleValuesForProviderName() []string {
	return []string{
		string(ProviderNameMicrosoftPointAuthorizationPolicyAssignments),
		string(ProviderNameMicrosoftPointOperationalInsightsSolutions),
		string(ProviderNameMicrosoftPointOperationalInsightsWorkspaces),
		string(ProviderNameMicrosoftPointOperationalInsightsWorkspacesDatasources),
		string(ProviderNameMicrosoftPointOperationalInsightsWorkspacesSharedKeys),
		string(ProviderNameMicrosoftPointaadiamDiagnosticSettings),
	}
}

type SettingType string

const (
	SettingTypeCopyableLabel         SettingType = "CopyableLabel"
	SettingTypeInfoMessage           SettingType = "InfoMessage"
	SettingTypeInstructionStepsGroup SettingType = "InstructionStepsGroup"
)

func PossibleValuesForSettingType() []string {
	return []string{
		string(SettingTypeCopyableLabel),
		string(SettingTypeInfoMessage),
		string(SettingTypeInstructionStepsGroup),
	}
}
