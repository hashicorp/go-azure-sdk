package usersession

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationType string

const (
	ApplicationTypeDesktop   ApplicationType = "Desktop"
	ApplicationTypeRemoteApp ApplicationType = "RemoteApp"
)

func PossibleValuesForApplicationType() []string {
	return []string{
		string(ApplicationTypeDesktop),
		string(ApplicationTypeRemoteApp),
	}
}

func parseApplicationType(input string) (*ApplicationType, error) {
	vals := map[string]ApplicationType{
		"desktop":   ApplicationTypeDesktop,
		"remoteapp": ApplicationTypeRemoteApp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationType(input)
	return &out, nil
}

type SessionState string

const (
	SessionStateActive                 SessionState = "Active"
	SessionStateDisconnected           SessionState = "Disconnected"
	SessionStateLogOff                 SessionState = "LogOff"
	SessionStatePending                SessionState = "Pending"
	SessionStateUnknown                SessionState = "Unknown"
	SessionStateUserProfileDiskMounted SessionState = "UserProfileDiskMounted"
)

func PossibleValuesForSessionState() []string {
	return []string{
		string(SessionStateActive),
		string(SessionStateDisconnected),
		string(SessionStateLogOff),
		string(SessionStatePending),
		string(SessionStateUnknown),
		string(SessionStateUserProfileDiskMounted),
	}
}

func parseSessionState(input string) (*SessionState, error) {
	vals := map[string]SessionState{
		"active":                 SessionStateActive,
		"disconnected":           SessionStateDisconnected,
		"logoff":                 SessionStateLogOff,
		"pending":                SessionStatePending,
		"unknown":                SessionStateUnknown,
		"userprofilediskmounted": SessionStateUserProfileDiskMounted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionState(input)
	return &out, nil
}
