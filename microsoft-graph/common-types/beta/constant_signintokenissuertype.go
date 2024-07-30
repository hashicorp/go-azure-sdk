package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInTokenIssuerType string

const (
	SignInTokenIssuerType_ADFederationServices           SignInTokenIssuerType = "ADFederationServices"
	SignInTokenIssuerType_ADFederationServicesMFAAdapter SignInTokenIssuerType = "ADFederationServicesMFAAdapter"
	SignInTokenIssuerType_AzureAD                        SignInTokenIssuerType = "AzureAD"
	SignInTokenIssuerType_AzureADBackupAuth              SignInTokenIssuerType = "AzureADBackupAuth"
	SignInTokenIssuerType_NPSExtension                   SignInTokenIssuerType = "NPSExtension"
)

func PossibleValuesForSignInTokenIssuerType() []string {
	return []string{
		string(SignInTokenIssuerType_ADFederationServices),
		string(SignInTokenIssuerType_ADFederationServicesMFAAdapter),
		string(SignInTokenIssuerType_AzureAD),
		string(SignInTokenIssuerType_AzureADBackupAuth),
		string(SignInTokenIssuerType_NPSExtension),
	}
}

func (s *SignInTokenIssuerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInTokenIssuerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInTokenIssuerType(input string) (*SignInTokenIssuerType, error) {
	vals := map[string]SignInTokenIssuerType{
		"adfederationservices":           SignInTokenIssuerType_ADFederationServices,
		"adfederationservicesmfaadapter": SignInTokenIssuerType_ADFederationServicesMFAAdapter,
		"azuread":                        SignInTokenIssuerType_AzureAD,
		"azureadbackupauth":              SignInTokenIssuerType_AzureADBackupAuth,
		"npsextension":                   SignInTokenIssuerType_NPSExtension,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInTokenIssuerType(input)
	return &out, nil
}
