package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestState string

const (
	AccessPackageAssignmentRequestState_Canceled           AccessPackageAssignmentRequestState = "canceled"
	AccessPackageAssignmentRequestState_Delivered          AccessPackageAssignmentRequestState = "delivered"
	AccessPackageAssignmentRequestState_Delivering         AccessPackageAssignmentRequestState = "delivering"
	AccessPackageAssignmentRequestState_DeliveryFailed     AccessPackageAssignmentRequestState = "deliveryFailed"
	AccessPackageAssignmentRequestState_Denied             AccessPackageAssignmentRequestState = "denied"
	AccessPackageAssignmentRequestState_PartiallyDelivered AccessPackageAssignmentRequestState = "partiallyDelivered"
	AccessPackageAssignmentRequestState_PendingApproval    AccessPackageAssignmentRequestState = "pendingApproval"
	AccessPackageAssignmentRequestState_Scheduled          AccessPackageAssignmentRequestState = "scheduled"
	AccessPackageAssignmentRequestState_Submitted          AccessPackageAssignmentRequestState = "submitted"
)

func PossibleValuesForAccessPackageAssignmentRequestState() []string {
	return []string{
		string(AccessPackageAssignmentRequestState_Canceled),
		string(AccessPackageAssignmentRequestState_Delivered),
		string(AccessPackageAssignmentRequestState_Delivering),
		string(AccessPackageAssignmentRequestState_DeliveryFailed),
		string(AccessPackageAssignmentRequestState_Denied),
		string(AccessPackageAssignmentRequestState_PartiallyDelivered),
		string(AccessPackageAssignmentRequestState_PendingApproval),
		string(AccessPackageAssignmentRequestState_Scheduled),
		string(AccessPackageAssignmentRequestState_Submitted),
	}
}

func (s *AccessPackageAssignmentRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentRequestState(input string) (*AccessPackageAssignmentRequestState, error) {
	vals := map[string]AccessPackageAssignmentRequestState{
		"canceled":           AccessPackageAssignmentRequestState_Canceled,
		"delivered":          AccessPackageAssignmentRequestState_Delivered,
		"delivering":         AccessPackageAssignmentRequestState_Delivering,
		"deliveryfailed":     AccessPackageAssignmentRequestState_DeliveryFailed,
		"denied":             AccessPackageAssignmentRequestState_Denied,
		"partiallydelivered": AccessPackageAssignmentRequestState_PartiallyDelivered,
		"pendingapproval":    AccessPackageAssignmentRequestState_PendingApproval,
		"scheduled":          AccessPackageAssignmentRequestState_Scheduled,
		"submitted":          AccessPackageAssignmentRequestState_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestState(input)
	return &out, nil
}
