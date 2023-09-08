package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionTokenIssuerType string

const (
	ServicePrincipalRiskDetectionTokenIssuerTypeADFederationServices           ServicePrincipalRiskDetectionTokenIssuerType = "ADFederationServices"
	ServicePrincipalRiskDetectionTokenIssuerTypeADFederationServicesMFAAdapter ServicePrincipalRiskDetectionTokenIssuerType = "ADFederationServicesMFAAdapter"
	ServicePrincipalRiskDetectionTokenIssuerTypeAzureAD                        ServicePrincipalRiskDetectionTokenIssuerType = "AzureAD"
	ServicePrincipalRiskDetectionTokenIssuerTypeAzureADBackupAuth              ServicePrincipalRiskDetectionTokenIssuerType = "AzureADBackupAuth"
	ServicePrincipalRiskDetectionTokenIssuerTypeNPSExtension                   ServicePrincipalRiskDetectionTokenIssuerType = "NPSExtension"
)

func PossibleValuesForServicePrincipalRiskDetectionTokenIssuerType() []string {
	return []string{
		string(ServicePrincipalRiskDetectionTokenIssuerTypeADFederationServices),
		string(ServicePrincipalRiskDetectionTokenIssuerTypeADFederationServicesMFAAdapter),
		string(ServicePrincipalRiskDetectionTokenIssuerTypeAzureAD),
		string(ServicePrincipalRiskDetectionTokenIssuerTypeAzureADBackupAuth),
		string(ServicePrincipalRiskDetectionTokenIssuerTypeNPSExtension),
	}
}

func parseServicePrincipalRiskDetectionTokenIssuerType(input string) (*ServicePrincipalRiskDetectionTokenIssuerType, error) {
	vals := map[string]ServicePrincipalRiskDetectionTokenIssuerType{
		"adfederationservices":           ServicePrincipalRiskDetectionTokenIssuerTypeADFederationServices,
		"adfederationservicesmfaadapter": ServicePrincipalRiskDetectionTokenIssuerTypeADFederationServicesMFAAdapter,
		"azuread":                        ServicePrincipalRiskDetectionTokenIssuerTypeAzureAD,
		"azureadbackupauth":              ServicePrincipalRiskDetectionTokenIssuerTypeAzureADBackupAuth,
		"npsextension":                   ServicePrincipalRiskDetectionTokenIssuerTypeNPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionTokenIssuerType(input)
	return &out, nil
}
