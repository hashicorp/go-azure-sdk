package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectionState string

const (
	ExternalConnectionStatedraft         ExternalConnectionState = "Draft"
	ExternalConnectionStatelimitExceeded ExternalConnectionState = "LimitExceeded"
	ExternalConnectionStateobsolete      ExternalConnectionState = "Obsolete"
	ExternalConnectionStateready         ExternalConnectionState = "Ready"
)

func PossibleValuesForExternalConnectionState() []string {
	return []string{
		string(ExternalConnectionStatedraft),
		string(ExternalConnectionStatelimitExceeded),
		string(ExternalConnectionStateobsolete),
		string(ExternalConnectionStateready),
	}
}

func parseExternalConnectionState(input string) (*ExternalConnectionState, error) {
	vals := map[string]ExternalConnectionState{
		"draft":         ExternalConnectionStatedraft,
		"limitexceeded": ExternalConnectionStatelimitExceeded,
		"obsolete":      ExternalConnectionStateobsolete,
		"ready":         ExternalConnectionStateready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectionState(input)
	return &out, nil
}
