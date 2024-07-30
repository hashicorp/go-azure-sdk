package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRequestState string

const (
	AccessPackageResourceRequestState_Canceled           AccessPackageResourceRequestState = "canceled"
	AccessPackageResourceRequestState_Delivered          AccessPackageResourceRequestState = "delivered"
	AccessPackageResourceRequestState_Delivering         AccessPackageResourceRequestState = "delivering"
	AccessPackageResourceRequestState_DeliveryFailed     AccessPackageResourceRequestState = "deliveryFailed"
	AccessPackageResourceRequestState_Denied             AccessPackageResourceRequestState = "denied"
	AccessPackageResourceRequestState_PartiallyDelivered AccessPackageResourceRequestState = "partiallyDelivered"
	AccessPackageResourceRequestState_PendingApproval    AccessPackageResourceRequestState = "pendingApproval"
	AccessPackageResourceRequestState_Scheduled          AccessPackageResourceRequestState = "scheduled"
	AccessPackageResourceRequestState_Submitted          AccessPackageResourceRequestState = "submitted"
)

func PossibleValuesForAccessPackageResourceRequestState() []string {
	return []string{
		string(AccessPackageResourceRequestState_Canceled),
		string(AccessPackageResourceRequestState_Delivered),
		string(AccessPackageResourceRequestState_Delivering),
		string(AccessPackageResourceRequestState_DeliveryFailed),
		string(AccessPackageResourceRequestState_Denied),
		string(AccessPackageResourceRequestState_PartiallyDelivered),
		string(AccessPackageResourceRequestState_PendingApproval),
		string(AccessPackageResourceRequestState_Scheduled),
		string(AccessPackageResourceRequestState_Submitted),
	}
}

func (s *AccessPackageResourceRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageResourceRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageResourceRequestState(input string) (*AccessPackageResourceRequestState, error) {
	vals := map[string]AccessPackageResourceRequestState{
		"canceled":           AccessPackageResourceRequestState_Canceled,
		"delivered":          AccessPackageResourceRequestState_Delivered,
		"delivering":         AccessPackageResourceRequestState_Delivering,
		"deliveryfailed":     AccessPackageResourceRequestState_DeliveryFailed,
		"denied":             AccessPackageResourceRequestState_Denied,
		"partiallydelivered": AccessPackageResourceRequestState_PartiallyDelivered,
		"pendingapproval":    AccessPackageResourceRequestState_PendingApproval,
		"scheduled":          AccessPackageResourceRequestState_Scheduled,
		"submitted":          AccessPackageResourceRequestState_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageResourceRequestState(input)
	return &out, nil
}
