package invitation

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationStatus string

const (
	InvitationStatusAccepted  InvitationStatus = "Accepted"
	InvitationStatusPending   InvitationStatus = "Pending"
	InvitationStatusRejected  InvitationStatus = "Rejected"
	InvitationStatusWithdrawn InvitationStatus = "Withdrawn"
)

func PossibleValuesForInvitationStatus() []string {
	return []string{
		string(InvitationStatusAccepted),
		string(InvitationStatusPending),
		string(InvitationStatusRejected),
		string(InvitationStatusWithdrawn),
	}
}

func parseInvitationStatus(input string) (*InvitationStatus, error) {
	vals := map[string]InvitationStatus{
		"accepted":  InvitationStatusAccepted,
		"pending":   InvitationStatusPending,
		"rejected":  InvitationStatusRejected,
		"withdrawn": InvitationStatusWithdrawn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvitationStatus(input)
	return &out, nil
}
