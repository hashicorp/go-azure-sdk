package usersession

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
