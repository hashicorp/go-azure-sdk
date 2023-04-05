package dataconnectorsconnect

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
