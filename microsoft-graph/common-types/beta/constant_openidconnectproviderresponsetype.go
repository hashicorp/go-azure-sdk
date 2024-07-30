package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderResponseType string

const (
	OpenIdConnectProviderResponseType_Code    OpenIdConnectProviderResponseType = "code"
	OpenIdConnectProviderResponseType_Idtoken OpenIdConnectProviderResponseType = "id_token"
	OpenIdConnectProviderResponseType_Token   OpenIdConnectProviderResponseType = "token"
)

func PossibleValuesForOpenIdConnectProviderResponseType() []string {
	return []string{
		string(OpenIdConnectProviderResponseType_Code),
		string(OpenIdConnectProviderResponseType_Idtoken),
		string(OpenIdConnectProviderResponseType_Token),
	}
}

func (s *OpenIdConnectProviderResponseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenIdConnectProviderResponseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenIdConnectProviderResponseType(input string) (*OpenIdConnectProviderResponseType, error) {
	vals := map[string]OpenIdConnectProviderResponseType{
		"code":     OpenIdConnectProviderResponseType_Code,
		"id_token": OpenIdConnectProviderResponseType_Idtoken,
		"token":    OpenIdConnectProviderResponseType_Token,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectProviderResponseType(input)
	return &out, nil
}
