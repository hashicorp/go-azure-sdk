package dataconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CcpAuthType string

const (
	CcpAuthTypeAPIKey     CcpAuthType = "APIKey"
	CcpAuthTypeAWS        CcpAuthType = "AWS"
	CcpAuthTypeBasic      CcpAuthType = "Basic"
	CcpAuthTypeGCP        CcpAuthType = "GCP"
	CcpAuthTypeGitHub     CcpAuthType = "GitHub"
	CcpAuthTypeJwtToken   CcpAuthType = "JwtToken"
	CcpAuthTypeNone       CcpAuthType = "None"
	CcpAuthTypeOAuthTwo   CcpAuthType = "OAuth2"
	CcpAuthTypeOracle     CcpAuthType = "Oracle"
	CcpAuthTypeServiceBus CcpAuthType = "ServiceBus"
	CcpAuthTypeSession    CcpAuthType = "Session"
)

func PossibleValuesForCcpAuthType() []string {
	return []string{
		string(CcpAuthTypeAPIKey),
		string(CcpAuthTypeAWS),
		string(CcpAuthTypeBasic),
		string(CcpAuthTypeGCP),
		string(CcpAuthTypeGitHub),
		string(CcpAuthTypeJwtToken),
		string(CcpAuthTypeNone),
		string(CcpAuthTypeOAuthTwo),
		string(CcpAuthTypeOracle),
		string(CcpAuthTypeServiceBus),
		string(CcpAuthTypeSession),
	}
}

func (s *CcpAuthType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCcpAuthType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCcpAuthType(input string) (*CcpAuthType, error) {
	vals := map[string]CcpAuthType{
		"apikey":     CcpAuthTypeAPIKey,
		"aws":        CcpAuthTypeAWS,
		"basic":      CcpAuthTypeBasic,
		"gcp":        CcpAuthTypeGCP,
		"github":     CcpAuthTypeGitHub,
		"jwttoken":   CcpAuthTypeJwtToken,
		"none":       CcpAuthTypeNone,
		"oauth2":     CcpAuthTypeOAuthTwo,
		"oracle":     CcpAuthTypeOracle,
		"servicebus": CcpAuthTypeServiceBus,
		"session":    CcpAuthTypeSession,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CcpAuthType(input)
	return &out, nil
}

type DataConnectorKind string

const (
	DataConnectorKindAmazonWebServicesCloudTrail                   DataConnectorKind = "AmazonWebServicesCloudTrail"
	DataConnectorKindAzureActiveDirectory                          DataConnectorKind = "AzureActiveDirectory"
	DataConnectorKindAzureAdvancedThreatProtection                 DataConnectorKind = "AzureAdvancedThreatProtection"
	DataConnectorKindAzureSecurityCenter                           DataConnectorKind = "AzureSecurityCenter"
	DataConnectorKindMicrosoftCloudAppSecurity                     DataConnectorKind = "MicrosoftCloudAppSecurity"
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection     DataConnectorKind = "MicrosoftDefenderAdvancedThreatProtection"
	DataConnectorKindMicrosoftThreatIntelligence                   DataConnectorKind = "MicrosoftThreatIntelligence"
	DataConnectorKindOfficeThreeSixFive                            DataConnectorKind = "Office365"
	DataConnectorKindPremiumMicrosoftDefenderForThreatIntelligence DataConnectorKind = "PremiumMicrosoftDefenderForThreatIntelligence"
	DataConnectorKindRestApiPoller                                 DataConnectorKind = "RestApiPoller"
	DataConnectorKindThreatIntelligence                            DataConnectorKind = "ThreatIntelligence"
)

func PossibleValuesForDataConnectorKind() []string {
	return []string{
		string(DataConnectorKindAmazonWebServicesCloudTrail),
		string(DataConnectorKindAzureActiveDirectory),
		string(DataConnectorKindAzureAdvancedThreatProtection),
		string(DataConnectorKindAzureSecurityCenter),
		string(DataConnectorKindMicrosoftCloudAppSecurity),
		string(DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
		string(DataConnectorKindMicrosoftThreatIntelligence),
		string(DataConnectorKindOfficeThreeSixFive),
		string(DataConnectorKindPremiumMicrosoftDefenderForThreatIntelligence),
		string(DataConnectorKindRestApiPoller),
		string(DataConnectorKindThreatIntelligence),
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
		"amazonwebservicescloudtrail":                   DataConnectorKindAmazonWebServicesCloudTrail,
		"azureactivedirectory":                          DataConnectorKindAzureActiveDirectory,
		"azureadvancedthreatprotection":                 DataConnectorKindAzureAdvancedThreatProtection,
		"azuresecuritycenter":                           DataConnectorKindAzureSecurityCenter,
		"microsoftcloudappsecurity":                     DataConnectorKindMicrosoftCloudAppSecurity,
		"microsoftdefenderadvancedthreatprotection":     DataConnectorKindMicrosoftDefenderAdvancedThreatProtection,
		"microsoftthreatintelligence":                   DataConnectorKindMicrosoftThreatIntelligence,
		"office365":                                     DataConnectorKindOfficeThreeSixFive,
		"premiummicrosoftdefenderforthreatintelligence": DataConnectorKindPremiumMicrosoftDefenderForThreatIntelligence,
		"restapipoller":                                 DataConnectorKindRestApiPoller,
		"threatintelligence":                            DataConnectorKindThreatIntelligence,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataConnectorKind(input)
	return &out, nil
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

func (s *DataTypeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataTypeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataTypeState(input string) (*DataTypeState, error) {
	vals := map[string]DataTypeState{
		"disabled": DataTypeStateDisabled,
		"enabled":  DataTypeStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataTypeState(input)
	return &out, nil
}

type HTTPMethodVerb string

const (
	HTTPMethodVerbDELETE HTTPMethodVerb = "DELETE"
	HTTPMethodVerbGET    HTTPMethodVerb = "GET"
	HTTPMethodVerbPOST   HTTPMethodVerb = "POST"
	HTTPMethodVerbPUT    HTTPMethodVerb = "PUT"
)

func PossibleValuesForHTTPMethodVerb() []string {
	return []string{
		string(HTTPMethodVerbDELETE),
		string(HTTPMethodVerbGET),
		string(HTTPMethodVerbPOST),
		string(HTTPMethodVerbPUT),
	}
}

func (s *HTTPMethodVerb) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHTTPMethodVerb(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHTTPMethodVerb(input string) (*HTTPMethodVerb, error) {
	vals := map[string]HTTPMethodVerb{
		"delete": HTTPMethodVerbDELETE,
		"get":    HTTPMethodVerbGET,
		"post":   HTTPMethodVerbPOST,
		"put":    HTTPMethodVerbPUT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HTTPMethodVerb(input)
	return &out, nil
}

type RestApiPollerRequestPagingKind string

const (
	RestApiPollerRequestPagingKindCountBasedPaging     RestApiPollerRequestPagingKind = "CountBasedPaging"
	RestApiPollerRequestPagingKindLinkHeader           RestApiPollerRequestPagingKind = "LinkHeader"
	RestApiPollerRequestPagingKindNextPageToken        RestApiPollerRequestPagingKind = "NextPageToken"
	RestApiPollerRequestPagingKindNextPageURL          RestApiPollerRequestPagingKind = "NextPageUrl"
	RestApiPollerRequestPagingKindOffset               RestApiPollerRequestPagingKind = "Offset"
	RestApiPollerRequestPagingKindPersistentLinkHeader RestApiPollerRequestPagingKind = "PersistentLinkHeader"
	RestApiPollerRequestPagingKindPersistentToken      RestApiPollerRequestPagingKind = "PersistentToken"
)

func PossibleValuesForRestApiPollerRequestPagingKind() []string {
	return []string{
		string(RestApiPollerRequestPagingKindCountBasedPaging),
		string(RestApiPollerRequestPagingKindLinkHeader),
		string(RestApiPollerRequestPagingKindNextPageToken),
		string(RestApiPollerRequestPagingKindNextPageURL),
		string(RestApiPollerRequestPagingKindOffset),
		string(RestApiPollerRequestPagingKindPersistentLinkHeader),
		string(RestApiPollerRequestPagingKindPersistentToken),
	}
}

func (s *RestApiPollerRequestPagingKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestApiPollerRequestPagingKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestApiPollerRequestPagingKind(input string) (*RestApiPollerRequestPagingKind, error) {
	vals := map[string]RestApiPollerRequestPagingKind{
		"countbasedpaging":     RestApiPollerRequestPagingKindCountBasedPaging,
		"linkheader":           RestApiPollerRequestPagingKindLinkHeader,
		"nextpagetoken":        RestApiPollerRequestPagingKindNextPageToken,
		"nextpageurl":          RestApiPollerRequestPagingKindNextPageURL,
		"offset":               RestApiPollerRequestPagingKindOffset,
		"persistentlinkheader": RestApiPollerRequestPagingKindPersistentLinkHeader,
		"persistenttoken":      RestApiPollerRequestPagingKindPersistentToken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestApiPollerRequestPagingKind(input)
	return &out, nil
}
