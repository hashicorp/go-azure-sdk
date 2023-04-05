package privateendpointconnection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Status string

const (
	StatusApproved     Status = "Approved"
	StatusDisconnected Status = "Disconnected"
	StatusPending      Status = "Pending"
	StatusRejected     Status = "Rejected"
	StatusUnknown      Status = "Unknown"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusApproved),
		string(StatusDisconnected),
		string(StatusPending),
		string(StatusRejected),
		string(StatusUnknown),
	}
}
