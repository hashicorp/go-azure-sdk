package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentState string

const (
	AccessPackageAssignmentStatedelivered          AccessPackageAssignmentState = "Delivered"
	AccessPackageAssignmentStatedelivering         AccessPackageAssignmentState = "Delivering"
	AccessPackageAssignmentStatedeliveryFailed     AccessPackageAssignmentState = "DeliveryFailed"
	AccessPackageAssignmentStateexpired            AccessPackageAssignmentState = "Expired"
	AccessPackageAssignmentStatepartiallyDelivered AccessPackageAssignmentState = "PartiallyDelivered"
)

func PossibleValuesForAccessPackageAssignmentState() []string {
	return []string{
		string(AccessPackageAssignmentStatedelivered),
		string(AccessPackageAssignmentStatedelivering),
		string(AccessPackageAssignmentStatedeliveryFailed),
		string(AccessPackageAssignmentStateexpired),
		string(AccessPackageAssignmentStatepartiallyDelivered),
	}
}

func parseAccessPackageAssignmentState(input string) (*AccessPackageAssignmentState, error) {
	vals := map[string]AccessPackageAssignmentState{
		"delivered":          AccessPackageAssignmentStatedelivered,
		"delivering":         AccessPackageAssignmentStatedelivering,
		"deliveryfailed":     AccessPackageAssignmentStatedeliveryFailed,
		"expired":            AccessPackageAssignmentStateexpired,
		"partiallydelivered": AccessPackageAssignmentStatepartiallyDelivered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentState(input)
	return &out, nil
}
