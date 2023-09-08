package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityCredentialOrigin string

const (
	AppCredentialSignInActivityCredentialOriginapplication      AppCredentialSignInActivityCredentialOrigin = "Application"
	AppCredentialSignInActivityCredentialOriginservicePrincipal AppCredentialSignInActivityCredentialOrigin = "ServicePrincipal"
)

func PossibleValuesForAppCredentialSignInActivityCredentialOrigin() []string {
	return []string{
		string(AppCredentialSignInActivityCredentialOriginapplication),
		string(AppCredentialSignInActivityCredentialOriginservicePrincipal),
	}
}

func parseAppCredentialSignInActivityCredentialOrigin(input string) (*AppCredentialSignInActivityCredentialOrigin, error) {
	vals := map[string]AppCredentialSignInActivityCredentialOrigin{
		"application":      AppCredentialSignInActivityCredentialOriginapplication,
		"serviceprincipal": AppCredentialSignInActivityCredentialOriginservicePrincipal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialSignInActivityCredentialOrigin(input)
	return &out, nil
}
