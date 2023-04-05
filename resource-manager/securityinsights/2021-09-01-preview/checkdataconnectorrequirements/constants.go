package checkdataconnectorrequirements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorAuthorizationState string

const (
	DataConnectorAuthorizationStateInvalid DataConnectorAuthorizationState = "Invalid"
	DataConnectorAuthorizationStateValid   DataConnectorAuthorizationState = "Valid"
)

func PossibleValuesForDataConnectorAuthorizationState() []string {
	return []string{
		string(DataConnectorAuthorizationStateInvalid),
		string(DataConnectorAuthorizationStateValid),
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
	DataConnectorKindMicrosoftCloudAppSecurity                 DataConnectorKind = "MicrosoftCloudAppSecurity"
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection DataConnectorKind = "MicrosoftDefenderAdvancedThreatProtection"
	DataConnectorKindMicrosoftThreatIntelligence               DataConnectorKind = "MicrosoftThreatIntelligence"
	DataConnectorKindMicrosoftThreatProtection                 DataConnectorKind = "MicrosoftThreatProtection"
	DataConnectorKindOfficeATP                                 DataConnectorKind = "OfficeATP"
	DataConnectorKindOfficeIRM                                 DataConnectorKind = "OfficeIRM"
	DataConnectorKindOfficeThreeSixFive                        DataConnectorKind = "Office365"
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
		string(DataConnectorKindMicrosoftCloudAppSecurity),
		string(DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
		string(DataConnectorKindMicrosoftThreatIntelligence),
		string(DataConnectorKindMicrosoftThreatProtection),
		string(DataConnectorKindOfficeATP),
		string(DataConnectorKindOfficeIRM),
		string(DataConnectorKindOfficeThreeSixFive),
		string(DataConnectorKindThreatIntelligence),
		string(DataConnectorKindThreatIntelligenceTaxii),
	}
}

type DataConnectorLicenseState string

const (
	DataConnectorLicenseStateInvalid DataConnectorLicenseState = "Invalid"
	DataConnectorLicenseStateUnknown DataConnectorLicenseState = "Unknown"
	DataConnectorLicenseStateValid   DataConnectorLicenseState = "Valid"
)

func PossibleValuesForDataConnectorLicenseState() []string {
	return []string{
		string(DataConnectorLicenseStateInvalid),
		string(DataConnectorLicenseStateUnknown),
		string(DataConnectorLicenseStateValid),
	}
}
