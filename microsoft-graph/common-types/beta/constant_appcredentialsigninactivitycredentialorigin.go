package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityCredentialOrigin string

const (
	AppCredentialSignInActivityCredentialOrigin_Application      AppCredentialSignInActivityCredentialOrigin = "application"
	AppCredentialSignInActivityCredentialOrigin_ServicePrincipal AppCredentialSignInActivityCredentialOrigin = "servicePrincipal"
)

func PossibleValuesForAppCredentialSignInActivityCredentialOrigin() []string {
	return []string{
		string(AppCredentialSignInActivityCredentialOrigin_Application),
		string(AppCredentialSignInActivityCredentialOrigin_ServicePrincipal),
	}
}

func (s *AppCredentialSignInActivityCredentialOrigin) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppCredentialSignInActivityCredentialOrigin(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppCredentialSignInActivityCredentialOrigin(input string) (*AppCredentialSignInActivityCredentialOrigin, error) {
	vals := map[string]AppCredentialSignInActivityCredentialOrigin{
		"application":      AppCredentialSignInActivityCredentialOrigin_Application,
		"serviceprincipal": AppCredentialSignInActivityCredentialOrigin_ServicePrincipal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialSignInActivityCredentialOrigin(input)
	return &out, nil
}
