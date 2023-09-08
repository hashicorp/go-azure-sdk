package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestState string

const (
	AccessPackageAssignmentRequestStatecanceled           AccessPackageAssignmentRequestState = "Canceled"
	AccessPackageAssignmentRequestStatedelivered          AccessPackageAssignmentRequestState = "Delivered"
	AccessPackageAssignmentRequestStatedelivering         AccessPackageAssignmentRequestState = "Delivering"
	AccessPackageAssignmentRequestStatedeliveryFailed     AccessPackageAssignmentRequestState = "DeliveryFailed"
	AccessPackageAssignmentRequestStatedenied             AccessPackageAssignmentRequestState = "Denied"
	AccessPackageAssignmentRequestStatepartiallyDelivered AccessPackageAssignmentRequestState = "PartiallyDelivered"
	AccessPackageAssignmentRequestStatependingApproval    AccessPackageAssignmentRequestState = "PendingApproval"
	AccessPackageAssignmentRequestStatescheduled          AccessPackageAssignmentRequestState = "Scheduled"
	AccessPackageAssignmentRequestStatesubmitted          AccessPackageAssignmentRequestState = "Submitted"
)

func PossibleValuesForAccessPackageAssignmentRequestState() []string {
	return []string{
		string(AccessPackageAssignmentRequestStatecanceled),
		string(AccessPackageAssignmentRequestStatedelivered),
		string(AccessPackageAssignmentRequestStatedelivering),
		string(AccessPackageAssignmentRequestStatedeliveryFailed),
		string(AccessPackageAssignmentRequestStatedenied),
		string(AccessPackageAssignmentRequestStatepartiallyDelivered),
		string(AccessPackageAssignmentRequestStatependingApproval),
		string(AccessPackageAssignmentRequestStatescheduled),
		string(AccessPackageAssignmentRequestStatesubmitted),
	}
}

func parseAccessPackageAssignmentRequestState(input string) (*AccessPackageAssignmentRequestState, error) {
	vals := map[string]AccessPackageAssignmentRequestState{
		"canceled":           AccessPackageAssignmentRequestStatecanceled,
		"delivered":          AccessPackageAssignmentRequestStatedelivered,
		"delivering":         AccessPackageAssignmentRequestStatedelivering,
		"deliveryfailed":     AccessPackageAssignmentRequestStatedeliveryFailed,
		"denied":             AccessPackageAssignmentRequestStatedenied,
		"partiallydelivered": AccessPackageAssignmentRequestStatepartiallyDelivered,
		"pendingapproval":    AccessPackageAssignmentRequestStatependingApproval,
		"scheduled":          AccessPackageAssignmentRequestStatescheduled,
		"submitted":          AccessPackageAssignmentRequestStatesubmitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestState(input)
	return &out, nil
}
