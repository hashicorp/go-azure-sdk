package registeredserverresource

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
