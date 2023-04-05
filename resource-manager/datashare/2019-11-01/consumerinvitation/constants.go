package consumerinvitation

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
