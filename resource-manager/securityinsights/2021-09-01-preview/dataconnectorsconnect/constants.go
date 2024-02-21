package dataconnectorsconnect

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectAuthKind string

const (
	ConnectAuthKindAPIKey   ConnectAuthKind = "APIKey"
	ConnectAuthKindBasic    ConnectAuthKind = "Basic"
	ConnectAuthKindOAuthTwo ConnectAuthKind = "OAuth2"
)

func PossibleValuesForConnectAuthKind() []string {
	return []string{
		string(ConnectAuthKindAPIKey),
		string(ConnectAuthKindBasic),
		string(ConnectAuthKindOAuthTwo),
	}
}

func (s *ConnectAuthKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectAuthKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectAuthKind(input string) (*ConnectAuthKind, error) {
	vals := map[string]ConnectAuthKind{
		"apikey": ConnectAuthKindAPIKey,
		"basic":  ConnectAuthKindBasic,
		"oauth2": ConnectAuthKindOAuthTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectAuthKind(input)
	return &out, nil
}
