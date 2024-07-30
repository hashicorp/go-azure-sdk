package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityKeyType string

const (
	AppCredentialSignInActivityKeyType_Certificate  AppCredentialSignInActivityKeyType = "certificate"
	AppCredentialSignInActivityKeyType_ClientSecret AppCredentialSignInActivityKeyType = "clientSecret"
)

func PossibleValuesForAppCredentialSignInActivityKeyType() []string {
	return []string{
		string(AppCredentialSignInActivityKeyType_Certificate),
		string(AppCredentialSignInActivityKeyType_ClientSecret),
	}
}

func (s *AppCredentialSignInActivityKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppCredentialSignInActivityKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppCredentialSignInActivityKeyType(input string) (*AppCredentialSignInActivityKeyType, error) {
	vals := map[string]AppCredentialSignInActivityKeyType{
		"certificate":  AppCredentialSignInActivityKeyType_Certificate,
		"clientsecret": AppCredentialSignInActivityKeyType_ClientSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialSignInActivityKeyType(input)
	return &out, nil
}
