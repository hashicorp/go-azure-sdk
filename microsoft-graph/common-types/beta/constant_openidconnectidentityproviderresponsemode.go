package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectIdentityProviderResponseMode string

const (
	OpenIdConnectIdentityProviderResponseMode_Formpost OpenIdConnectIdentityProviderResponseMode = "form_post"
	OpenIdConnectIdentityProviderResponseMode_Query    OpenIdConnectIdentityProviderResponseMode = "query"
)

func PossibleValuesForOpenIdConnectIdentityProviderResponseMode() []string {
	return []string{
		string(OpenIdConnectIdentityProviderResponseMode_Formpost),
		string(OpenIdConnectIdentityProviderResponseMode_Query),
	}
}

func (s *OpenIdConnectIdentityProviderResponseMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenIdConnectIdentityProviderResponseMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenIdConnectIdentityProviderResponseMode(input string) (*OpenIdConnectIdentityProviderResponseMode, error) {
	vals := map[string]OpenIdConnectIdentityProviderResponseMode{
		"form_post": OpenIdConnectIdentityProviderResponseMode_Formpost,
		"query":     OpenIdConnectIdentityProviderResponseMode_Query,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectIdentityProviderResponseMode(input)
	return &out, nil
}
