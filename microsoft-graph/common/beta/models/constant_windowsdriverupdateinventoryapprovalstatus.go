package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateInventoryApprovalStatus string

const (
	WindowsDriverUpdateInventoryApprovalStatusapproved    WindowsDriverUpdateInventoryApprovalStatus = "Approved"
	WindowsDriverUpdateInventoryApprovalStatusdeclined    WindowsDriverUpdateInventoryApprovalStatus = "Declined"
	WindowsDriverUpdateInventoryApprovalStatusneedsReview WindowsDriverUpdateInventoryApprovalStatus = "NeedsReview"
	WindowsDriverUpdateInventoryApprovalStatussuspended   WindowsDriverUpdateInventoryApprovalStatus = "Suspended"
)

func PossibleValuesForWindowsDriverUpdateInventoryApprovalStatus() []string {
	return []string{
		string(WindowsDriverUpdateInventoryApprovalStatusapproved),
		string(WindowsDriverUpdateInventoryApprovalStatusdeclined),
		string(WindowsDriverUpdateInventoryApprovalStatusneedsReview),
		string(WindowsDriverUpdateInventoryApprovalStatussuspended),
	}
}

func parseWindowsDriverUpdateInventoryApprovalStatus(input string) (*WindowsDriverUpdateInventoryApprovalStatus, error) {
	vals := map[string]WindowsDriverUpdateInventoryApprovalStatus{
		"approved":    WindowsDriverUpdateInventoryApprovalStatusapproved,
		"declined":    WindowsDriverUpdateInventoryApprovalStatusdeclined,
		"needsreview": WindowsDriverUpdateInventoryApprovalStatusneedsReview,
		"suspended":   WindowsDriverUpdateInventoryApprovalStatussuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateInventoryApprovalStatus(input)
	return &out, nil
}
