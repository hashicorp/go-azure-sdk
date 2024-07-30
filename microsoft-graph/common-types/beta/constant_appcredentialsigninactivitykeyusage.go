package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityKeyUsage string

const (
	AppCredentialSignInActivityKeyUsage_Sign   AppCredentialSignInActivityKeyUsage = "sign"
	AppCredentialSignInActivityKeyUsage_Verify AppCredentialSignInActivityKeyUsage = "verify"
)

func PossibleValuesForAppCredentialSignInActivityKeyUsage() []string {
	return []string{
		string(AppCredentialSignInActivityKeyUsage_Sign),
		string(AppCredentialSignInActivityKeyUsage_Verify),
	}
}

func (s *AppCredentialSignInActivityKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppCredentialSignInActivityKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppCredentialSignInActivityKeyUsage(input string) (*AppCredentialSignInActivityKeyUsage, error) {
	vals := map[string]AppCredentialSignInActivityKeyUsage{
		"sign":   AppCredentialSignInActivityKeyUsage_Sign,
		"verify": AppCredentialSignInActivityKeyUsage_Verify,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialSignInActivityKeyUsage(input)
	return &out, nil
}
