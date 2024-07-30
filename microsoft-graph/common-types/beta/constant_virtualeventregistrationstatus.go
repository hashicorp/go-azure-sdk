package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationStatus string

const (
	VirtualEventRegistrationStatus_Canceled            VirtualEventRegistrationStatus = "canceled"
	VirtualEventRegistrationStatus_PendingApproval     VirtualEventRegistrationStatus = "pendingApproval"
	VirtualEventRegistrationStatus_Registered          VirtualEventRegistrationStatus = "registered"
	VirtualEventRegistrationStatus_RejectedByOrganizer VirtualEventRegistrationStatus = "rejectedByOrganizer"
	VirtualEventRegistrationStatus_Waitlisted          VirtualEventRegistrationStatus = "waitlisted"
)

func PossibleValuesForVirtualEventRegistrationStatus() []string {
	return []string{
		string(VirtualEventRegistrationStatus_Canceled),
		string(VirtualEventRegistrationStatus_PendingApproval),
		string(VirtualEventRegistrationStatus_Registered),
		string(VirtualEventRegistrationStatus_RejectedByOrganizer),
		string(VirtualEventRegistrationStatus_Waitlisted),
	}
}

func (s *VirtualEventRegistrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventRegistrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventRegistrationStatus(input string) (*VirtualEventRegistrationStatus, error) {
	vals := map[string]VirtualEventRegistrationStatus{
		"canceled":            VirtualEventRegistrationStatus_Canceled,
		"pendingapproval":     VirtualEventRegistrationStatus_PendingApproval,
		"registered":          VirtualEventRegistrationStatus_Registered,
		"rejectedbyorganizer": VirtualEventRegistrationStatus_RejectedByOrganizer,
		"waitlisted":          VirtualEventRegistrationStatus_Waitlisted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventRegistrationStatus(input)
	return &out, nil
}
