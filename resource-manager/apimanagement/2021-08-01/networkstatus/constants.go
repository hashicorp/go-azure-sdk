package networkstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectivityStatusType string

const (
	ConnectivityStatusTypeFailure      ConnectivityStatusType = "failure"
	ConnectivityStatusTypeInitializing ConnectivityStatusType = "initializing"
	ConnectivityStatusTypeSuccess      ConnectivityStatusType = "success"
)

func PossibleValuesForConnectivityStatusType() []string {
	return []string{
		string(ConnectivityStatusTypeFailure),
		string(ConnectivityStatusTypeInitializing),
		string(ConnectivityStatusTypeSuccess),
	}
}
