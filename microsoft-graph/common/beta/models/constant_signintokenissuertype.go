package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInTokenIssuerType string

const (
	SignInTokenIssuerTypeADFederationServices           SignInTokenIssuerType = "ADFederationServices"
	SignInTokenIssuerTypeADFederationServicesMFAAdapter SignInTokenIssuerType = "ADFederationServicesMFAAdapter"
	SignInTokenIssuerTypeAzureAD                        SignInTokenIssuerType = "AzureAD"
	SignInTokenIssuerTypeAzureADBackupAuth              SignInTokenIssuerType = "AzureADBackupAuth"
	SignInTokenIssuerTypeNPSExtension                   SignInTokenIssuerType = "NPSExtension"
)

func PossibleValuesForSignInTokenIssuerType() []string {
	return []string{
		string(SignInTokenIssuerTypeADFederationServices),
		string(SignInTokenIssuerTypeADFederationServicesMFAAdapter),
		string(SignInTokenIssuerTypeAzureAD),
		string(SignInTokenIssuerTypeAzureADBackupAuth),
		string(SignInTokenIssuerTypeNPSExtension),
	}
}

func parseSignInTokenIssuerType(input string) (*SignInTokenIssuerType, error) {
	vals := map[string]SignInTokenIssuerType{
		"adfederationservices":           SignInTokenIssuerTypeADFederationServices,
		"adfederationservicesmfaadapter": SignInTokenIssuerTypeADFederationServicesMFAAdapter,
		"azuread":                        SignInTokenIssuerTypeAzureAD,
		"azureadbackupauth":              SignInTokenIssuerTypeAzureADBackupAuth,
		"npsextension":                   SignInTokenIssuerTypeNPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInTokenIssuerType(input)
	return &out, nil
}
