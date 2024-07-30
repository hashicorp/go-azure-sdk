package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionTokenIssuerType string

const (
	ServicePrincipalRiskDetectionTokenIssuerType_ADFederationServices           ServicePrincipalRiskDetectionTokenIssuerType = "ADFederationServices"
	ServicePrincipalRiskDetectionTokenIssuerType_ADFederationServicesMFAAdapter ServicePrincipalRiskDetectionTokenIssuerType = "ADFederationServicesMFAAdapter"
	ServicePrincipalRiskDetectionTokenIssuerType_AzureAD                        ServicePrincipalRiskDetectionTokenIssuerType = "AzureAD"
	ServicePrincipalRiskDetectionTokenIssuerType_AzureADBackupAuth              ServicePrincipalRiskDetectionTokenIssuerType = "AzureADBackupAuth"
	ServicePrincipalRiskDetectionTokenIssuerType_NPSExtension                   ServicePrincipalRiskDetectionTokenIssuerType = "NPSExtension"
)

func PossibleValuesForServicePrincipalRiskDetectionTokenIssuerType() []string {
	return []string{
		string(ServicePrincipalRiskDetectionTokenIssuerType_ADFederationServices),
		string(ServicePrincipalRiskDetectionTokenIssuerType_ADFederationServicesMFAAdapter),
		string(ServicePrincipalRiskDetectionTokenIssuerType_AzureAD),
		string(ServicePrincipalRiskDetectionTokenIssuerType_AzureADBackupAuth),
		string(ServicePrincipalRiskDetectionTokenIssuerType_NPSExtension),
	}
}

func (s *ServicePrincipalRiskDetectionTokenIssuerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePrincipalRiskDetectionTokenIssuerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePrincipalRiskDetectionTokenIssuerType(input string) (*ServicePrincipalRiskDetectionTokenIssuerType, error) {
	vals := map[string]ServicePrincipalRiskDetectionTokenIssuerType{
		"adfederationservices":           ServicePrincipalRiskDetectionTokenIssuerType_ADFederationServices,
		"adfederationservicesmfaadapter": ServicePrincipalRiskDetectionTokenIssuerType_ADFederationServicesMFAAdapter,
		"azuread":                        ServicePrincipalRiskDetectionTokenIssuerType_AzureAD,
		"azureadbackupauth":              ServicePrincipalRiskDetectionTokenIssuerType_AzureADBackupAuth,
		"npsextension":                   ServicePrincipalRiskDetectionTokenIssuerType_NPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionTokenIssuerType(input)
	return &out, nil
}
