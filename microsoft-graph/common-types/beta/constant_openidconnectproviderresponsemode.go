package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderResponseMode string

const (
	OpenIdConnectProviderResponseMode_Formpost OpenIdConnectProviderResponseMode = "form_post"
	OpenIdConnectProviderResponseMode_Query    OpenIdConnectProviderResponseMode = "query"
)

func PossibleValuesForOpenIdConnectProviderResponseMode() []string {
	return []string{
		string(OpenIdConnectProviderResponseMode_Formpost),
		string(OpenIdConnectProviderResponseMode_Query),
	}
}

func (s *OpenIdConnectProviderResponseMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenIdConnectProviderResponseMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenIdConnectProviderResponseMode(input string) (*OpenIdConnectProviderResponseMode, error) {
	vals := map[string]OpenIdConnectProviderResponseMode{
		"form_post": OpenIdConnectProviderResponseMode_Formpost,
		"query":     OpenIdConnectProviderResponseMode_Query,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectProviderResponseMode(input)
	return &out, nil
}
