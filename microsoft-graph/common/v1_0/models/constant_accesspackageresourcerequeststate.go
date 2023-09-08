package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRequestState string

const (
	AccessPackageResourceRequestStatecanceled           AccessPackageResourceRequestState = "Canceled"
	AccessPackageResourceRequestStatedelivered          AccessPackageResourceRequestState = "Delivered"
	AccessPackageResourceRequestStatedelivering         AccessPackageResourceRequestState = "Delivering"
	AccessPackageResourceRequestStatedeliveryFailed     AccessPackageResourceRequestState = "DeliveryFailed"
	AccessPackageResourceRequestStatedenied             AccessPackageResourceRequestState = "Denied"
	AccessPackageResourceRequestStatepartiallyDelivered AccessPackageResourceRequestState = "PartiallyDelivered"
	AccessPackageResourceRequestStatependingApproval    AccessPackageResourceRequestState = "PendingApproval"
	AccessPackageResourceRequestStatescheduled          AccessPackageResourceRequestState = "Scheduled"
	AccessPackageResourceRequestStatesubmitted          AccessPackageResourceRequestState = "Submitted"
)

func PossibleValuesForAccessPackageResourceRequestState() []string {
	return []string{
		string(AccessPackageResourceRequestStatecanceled),
		string(AccessPackageResourceRequestStatedelivered),
		string(AccessPackageResourceRequestStatedelivering),
		string(AccessPackageResourceRequestStatedeliveryFailed),
		string(AccessPackageResourceRequestStatedenied),
		string(AccessPackageResourceRequestStatepartiallyDelivered),
		string(AccessPackageResourceRequestStatependingApproval),
		string(AccessPackageResourceRequestStatescheduled),
		string(AccessPackageResourceRequestStatesubmitted),
	}
}

func parseAccessPackageResourceRequestState(input string) (*AccessPackageResourceRequestState, error) {
	vals := map[string]AccessPackageResourceRequestState{
		"canceled":           AccessPackageResourceRequestStatecanceled,
		"delivered":          AccessPackageResourceRequestStatedelivered,
		"delivering":         AccessPackageResourceRequestStatedelivering,
		"deliveryfailed":     AccessPackageResourceRequestStatedeliveryFailed,
		"denied":             AccessPackageResourceRequestStatedenied,
		"partiallydelivered": AccessPackageResourceRequestStatepartiallyDelivered,
		"pendingapproval":    AccessPackageResourceRequestStatependingApproval,
		"scheduled":          AccessPackageResourceRequestStatescheduled,
		"submitted":          AccessPackageResourceRequestStatesubmitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageResourceRequestState(input)
	return &out, nil
}
