package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *DataConnectorAuthorizationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataConnectorAuthorizationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
	DataConnectorKindGCP                                       DataConnectorKind = "GCP"
	DataConnectorKindGenericUI                                 DataConnectorKind = "GenericUI"
	DataConnectorKindIOT                                       DataConnectorKind = "IOT"
	DataConnectorKindMicrosoftCloudAppSecurity                 DataConnectorKind = "MicrosoftCloudAppSecurity"
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection DataConnectorKind = "MicrosoftDefenderAdvancedThreatProtection"
	DataConnectorKindMicrosoftPurviewInformationProtection     DataConnectorKind = "MicrosoftPurviewInformationProtection"
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
		string(DataConnectorKindGCP),
		string(DataConnectorKindGenericUI),
		string(DataConnectorKindIOT),
		string(DataConnectorKindMicrosoftCloudAppSecurity),
		string(DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
		string(DataConnectorKindMicrosoftPurviewInformationProtection),
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

func (s *DataConnectorKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataConnectorKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataConnectorKind(input string) (*DataConnectorKind, error) {
	vals := map[string]DataConnectorKind{
		"apipolling":                    DataConnectorKindAPIPolling,
		"amazonwebservicescloudtrail":   DataConnectorKindAmazonWebServicesCloudTrail,
		"amazonwebservicess3":           DataConnectorKindAmazonWebServicesSThree,
		"azureactivedirectory":          DataConnectorKindAzureActiveDirectory,
		"azureadvancedthreatprotection": DataConnectorKindAzureAdvancedThreatProtection,
		"azuresecuritycenter":           DataConnectorKindAzureSecurityCenter,
		"dynamics365":                   DataConnectorKindDynamicsThreeSixFive,
		"gcp":                           DataConnectorKindGCP,
		"genericui":                     DataConnectorKindGenericUI,
		"iot":                           DataConnectorKindIOT,
		"microsoftcloudappsecurity":     DataConnectorKindMicrosoftCloudAppSecurity,
		"microsoftdefenderadvancedthreatprotection": DataConnectorKindMicrosoftDefenderAdvancedThreatProtection,
		"microsoftpurviewinformationprotection":     DataConnectorKindMicrosoftPurviewInformationProtection,
		"microsoftthreatintelligence":               DataConnectorKindMicrosoftThreatIntelligence,
		"microsoftthreatprotection":                 DataConnectorKindMicrosoftThreatProtection,
		"officeatp":                                 DataConnectorKindOfficeATP,
		"officeirm":                                 DataConnectorKindOfficeIRM,
		"officepowerbi":                             DataConnectorKindOfficePowerBI,
		"office365":                                 DataConnectorKindOfficeThreeSixFive,
		"office365project":                          DataConnectorKindOfficeThreeSixFiveProject,
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

func (s *DataConnectorLicenseState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataConnectorLicenseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
