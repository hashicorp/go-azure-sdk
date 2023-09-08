package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityKeyUsage string

const (
	AppCredentialSignInActivityKeyUsagesign   AppCredentialSignInActivityKeyUsage = "Sign"
	AppCredentialSignInActivityKeyUsageverify AppCredentialSignInActivityKeyUsage = "Verify"
)

func PossibleValuesForAppCredentialSignInActivityKeyUsage() []string {
	return []string{
		string(AppCredentialSignInActivityKeyUsagesign),
		string(AppCredentialSignInActivityKeyUsageverify),
	}
}

func parseAppCredentialSignInActivityKeyUsage(input string) (*AppCredentialSignInActivityKeyUsage, error) {
	vals := map[string]AppCredentialSignInActivityKeyUsage{
		"sign":   AppCredentialSignInActivityKeyUsagesign,
		"verify": AppCredentialSignInActivityKeyUsageverify,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialSignInActivityKeyUsage(input)
	return &out, nil
}
