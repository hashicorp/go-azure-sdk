package checkdataconnectorrequirements

import "strings"

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

func parseDataConnectorAuthorizationState(input string) (*DataConnectorAuthorizationState, error) {
	vals := map[string]DataConnectorAuthorizationState{
		"invalid": DataConnectorAuthorizationStateInvalid,
		"valid":   DataConnectorAuthorizationStateValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataConnectorAuthorizationState(input)
	return &out, nil
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

func parseDataConnectorKind(input string) (*DataConnectorKind, error) {
	vals := map[string]DataConnectorKind{
		"apipolling":                                DataConnectorKindAPIPolling,
		"amazonwebservicescloudtrail":               DataConnectorKindAmazonWebServicesCloudTrail,
		"amazonwebservicess3":                       DataConnectorKindAmazonWebServicesSThree,
		"azureactivedirectory":                      DataConnectorKindAzureActiveDirectory,
		"azureadvancedthreatprotection":             DataConnectorKindAzureAdvancedThreatProtection,
		"azuresecuritycenter":                       DataConnectorKindAzureSecurityCenter,
		"dynamics365":                               DataConnectorKindDynamicsThreeSixFive,
		"genericui":                                 DataConnectorKindGenericUI,
		"microsoftcloudappsecurity":                 DataConnectorKindMicrosoftCloudAppSecurity,
		"microsoftdefenderadvancedthreatprotection": DataConnectorKindMicrosoftDefenderAdvancedThreatProtection,
		"microsoftthreatintelligence":               DataConnectorKindMicrosoftThreatIntelligence,
		"microsoftthreatprotection":                 DataConnectorKindMicrosoftThreatProtection,
		"officeatp":                                 DataConnectorKindOfficeATP,
		"officeirm":                                 DataConnectorKindOfficeIRM,
		"office365":                                 DataConnectorKindOfficeThreeSixFive,
		"threatintelligence":                        DataConnectorKindThreatIntelligence,
		"threatintelligencetaxii":                   DataConnectorKindThreatIntelligenceTaxii,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataConnectorKind(input)
	return &out, nil
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

func parseDataConnectorLicenseState(input string) (*DataConnectorLicenseState, error) {
	vals := map[string]DataConnectorLicenseState{
		"invalid": DataConnectorLicenseStateInvalid,
		"unknown": DataConnectorLicenseStateUnknown,
		"valid":   DataConnectorLicenseStateValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataConnectorLicenseState(input)
	return &out, nil
}
