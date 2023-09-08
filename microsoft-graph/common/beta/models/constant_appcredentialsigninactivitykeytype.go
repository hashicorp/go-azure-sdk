package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityKeyType string

const (
	AppCredentialSignInActivityKeyTypecertificate  AppCredentialSignInActivityKeyType = "Certificate"
	AppCredentialSignInActivityKeyTypeclientSecret AppCredentialSignInActivityKeyType = "ClientSecret"
)

func PossibleValuesForAppCredentialSignInActivityKeyType() []string {
	return []string{
		string(AppCredentialSignInActivityKeyTypecertificate),
		string(AppCredentialSignInActivityKeyTypeclientSecret),
	}
}

func parseAppCredentialSignInActivityKeyType(input string) (*AppCredentialSignInActivityKeyType, error) {
	vals := map[string]AppCredentialSignInActivityKeyType{
		"certificate":  AppCredentialSignInActivityKeyTypecertificate,
		"clientsecret": AppCredentialSignInActivityKeyTypeclientSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialSignInActivityKeyType(input)
	return &out, nil
}
