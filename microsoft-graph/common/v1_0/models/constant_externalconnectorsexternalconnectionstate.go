package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalConnectionState string

const (
	ExternalConnectorsExternalConnectionStatedraft         ExternalConnectorsExternalConnectionState = "Draft"
	ExternalConnectorsExternalConnectionStatelimitExceeded ExternalConnectorsExternalConnectionState = "LimitExceeded"
	ExternalConnectorsExternalConnectionStateobsolete      ExternalConnectorsExternalConnectionState = "Obsolete"
	ExternalConnectorsExternalConnectionStateready         ExternalConnectorsExternalConnectionState = "Ready"
)

func PossibleValuesForExternalConnectorsExternalConnectionState() []string {
	return []string{
		string(ExternalConnectorsExternalConnectionStatedraft),
		string(ExternalConnectorsExternalConnectionStatelimitExceeded),
		string(ExternalConnectorsExternalConnectionStateobsolete),
		string(ExternalConnectorsExternalConnectionStateready),
	}
}

func parseExternalConnectorsExternalConnectionState(input string) (*ExternalConnectorsExternalConnectionState, error) {
	vals := map[string]ExternalConnectorsExternalConnectionState{
		"draft":         ExternalConnectorsExternalConnectionStatedraft,
		"limitexceeded": ExternalConnectorsExternalConnectionStatelimitExceeded,
		"obsolete":      ExternalConnectorsExternalConnectionStateobsolete,
		"ready":         ExternalConnectorsExternalConnectionStateready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalConnectionState(input)
	return &out, nil
}
