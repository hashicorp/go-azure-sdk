package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectIdentityProviderResponseType string

const (
	OpenIdConnectIdentityProviderResponseType_Code    OpenIdConnectIdentityProviderResponseType = "code"
	OpenIdConnectIdentityProviderResponseType_Idtoken OpenIdConnectIdentityProviderResponseType = "id_token"
	OpenIdConnectIdentityProviderResponseType_Token   OpenIdConnectIdentityProviderResponseType = "token"
)

func PossibleValuesForOpenIdConnectIdentityProviderResponseType() []string {
	return []string{
		string(OpenIdConnectIdentityProviderResponseType_Code),
		string(OpenIdConnectIdentityProviderResponseType_Idtoken),
		string(OpenIdConnectIdentityProviderResponseType_Token),
	}
}

func (s *OpenIdConnectIdentityProviderResponseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenIdConnectIdentityProviderResponseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenIdConnectIdentityProviderResponseType(input string) (*OpenIdConnectIdentityProviderResponseType, error) {
	vals := map[string]OpenIdConnectIdentityProviderResponseType{
		"code":     OpenIdConnectIdentityProviderResponseType_Code,
		"id_token": OpenIdConnectIdentityProviderResponseType_Idtoken,
		"token":    OpenIdConnectIdentityProviderResponseType_Token,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectIdentityProviderResponseType(input)
	return &out, nil
}
