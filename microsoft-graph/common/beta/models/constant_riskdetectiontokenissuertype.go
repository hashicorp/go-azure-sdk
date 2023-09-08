package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionTokenIssuerType string

const (
	RiskDetectionTokenIssuerTypeADFederationServices           RiskDetectionTokenIssuerType = "ADFederationServices"
	RiskDetectionTokenIssuerTypeADFederationServicesMFAAdapter RiskDetectionTokenIssuerType = "ADFederationServicesMFAAdapter"
	RiskDetectionTokenIssuerTypeAzureAD                        RiskDetectionTokenIssuerType = "AzureAD"
	RiskDetectionTokenIssuerTypeAzureADBackupAuth              RiskDetectionTokenIssuerType = "AzureADBackupAuth"
	RiskDetectionTokenIssuerTypeNPSExtension                   RiskDetectionTokenIssuerType = "NPSExtension"
)

func PossibleValuesForRiskDetectionTokenIssuerType() []string {
	return []string{
		string(RiskDetectionTokenIssuerTypeADFederationServices),
		string(RiskDetectionTokenIssuerTypeADFederationServicesMFAAdapter),
		string(RiskDetectionTokenIssuerTypeAzureAD),
		string(RiskDetectionTokenIssuerTypeAzureADBackupAuth),
		string(RiskDetectionTokenIssuerTypeNPSExtension),
	}
}

func parseRiskDetectionTokenIssuerType(input string) (*RiskDetectionTokenIssuerType, error) {
	vals := map[string]RiskDetectionTokenIssuerType{
		"adfederationservices":           RiskDetectionTokenIssuerTypeADFederationServices,
		"adfederationservicesmfaadapter": RiskDetectionTokenIssuerTypeADFederationServicesMFAAdapter,
		"azuread":                        RiskDetectionTokenIssuerTypeAzureAD,
		"azureadbackupauth":              RiskDetectionTokenIssuerTypeAzureADBackupAuth,
		"npsextension":                   RiskDetectionTokenIssuerTypeNPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionTokenIssuerType(input)
	return &out, nil
}
