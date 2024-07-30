package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionTokenIssuerType string

const (
	RiskDetectionTokenIssuerType_ADFederationServices           RiskDetectionTokenIssuerType = "ADFederationServices"
	RiskDetectionTokenIssuerType_ADFederationServicesMFAAdapter RiskDetectionTokenIssuerType = "ADFederationServicesMFAAdapter"
	RiskDetectionTokenIssuerType_AzureAD                        RiskDetectionTokenIssuerType = "AzureAD"
	RiskDetectionTokenIssuerType_AzureADBackupAuth              RiskDetectionTokenIssuerType = "AzureADBackupAuth"
	RiskDetectionTokenIssuerType_NPSExtension                   RiskDetectionTokenIssuerType = "NPSExtension"
)

func PossibleValuesForRiskDetectionTokenIssuerType() []string {
	return []string{
		string(RiskDetectionTokenIssuerType_ADFederationServices),
		string(RiskDetectionTokenIssuerType_ADFederationServicesMFAAdapter),
		string(RiskDetectionTokenIssuerType_AzureAD),
		string(RiskDetectionTokenIssuerType_AzureADBackupAuth),
		string(RiskDetectionTokenIssuerType_NPSExtension),
	}
}

func (s *RiskDetectionTokenIssuerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionTokenIssuerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionTokenIssuerType(input string) (*RiskDetectionTokenIssuerType, error) {
	vals := map[string]RiskDetectionTokenIssuerType{
		"adfederationservices":           RiskDetectionTokenIssuerType_ADFederationServices,
		"adfederationservicesmfaadapter": RiskDetectionTokenIssuerType_ADFederationServicesMFAAdapter,
		"azuread":                        RiskDetectionTokenIssuerType_AzureAD,
		"azureadbackupauth":              RiskDetectionTokenIssuerType_AzureADBackupAuth,
		"npsextension":                   RiskDetectionTokenIssuerType_NPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionTokenIssuerType(input)
	return &out, nil
}
