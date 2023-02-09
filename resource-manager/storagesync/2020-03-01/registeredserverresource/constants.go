package registeredserverresource

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredServerAgentVersionStatus string

const (
	RegisteredServerAgentVersionStatusBlocked    RegisteredServerAgentVersionStatus = "Blocked"
	RegisteredServerAgentVersionStatusExpired    RegisteredServerAgentVersionStatus = "Expired"
	RegisteredServerAgentVersionStatusNearExpiry RegisteredServerAgentVersionStatus = "NearExpiry"
	RegisteredServerAgentVersionStatusOk         RegisteredServerAgentVersionStatus = "Ok"
)

func PossibleValuesForRegisteredServerAgentVersionStatus() []string {
	return []string{
		string(RegisteredServerAgentVersionStatusBlocked),
		string(RegisteredServerAgentVersionStatusExpired),
		string(RegisteredServerAgentVersionStatusNearExpiry),
		string(RegisteredServerAgentVersionStatusOk),
	}
}

func parseRegisteredServerAgentVersionStatus(input string) (*RegisteredServerAgentVersionStatus, error) {
	vals := map[string]RegisteredServerAgentVersionStatus{
		"blocked":    RegisteredServerAgentVersionStatusBlocked,
		"expired":    RegisteredServerAgentVersionStatusExpired,
		"nearexpiry": RegisteredServerAgentVersionStatusNearExpiry,
		"ok":         RegisteredServerAgentVersionStatusOk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegisteredServerAgentVersionStatus(input)
	return &out, nil
}
