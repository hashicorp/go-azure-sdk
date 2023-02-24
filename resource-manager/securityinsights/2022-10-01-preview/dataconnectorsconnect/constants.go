package dataconnectorsconnect

import "strings"

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
