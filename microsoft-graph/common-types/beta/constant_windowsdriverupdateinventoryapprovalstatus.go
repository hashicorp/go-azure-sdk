package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateInventoryApprovalStatus string

const (
	WindowsDriverUpdateInventoryApprovalStatus_Approved    WindowsDriverUpdateInventoryApprovalStatus = "approved"
	WindowsDriverUpdateInventoryApprovalStatus_Declined    WindowsDriverUpdateInventoryApprovalStatus = "declined"
	WindowsDriverUpdateInventoryApprovalStatus_NeedsReview WindowsDriverUpdateInventoryApprovalStatus = "needsReview"
	WindowsDriverUpdateInventoryApprovalStatus_Suspended   WindowsDriverUpdateInventoryApprovalStatus = "suspended"
)

func PossibleValuesForWindowsDriverUpdateInventoryApprovalStatus() []string {
	return []string{
		string(WindowsDriverUpdateInventoryApprovalStatus_Approved),
		string(WindowsDriverUpdateInventoryApprovalStatus_Declined),
		string(WindowsDriverUpdateInventoryApprovalStatus_NeedsReview),
		string(WindowsDriverUpdateInventoryApprovalStatus_Suspended),
	}
}

func (s *WindowsDriverUpdateInventoryApprovalStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDriverUpdateInventoryApprovalStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDriverUpdateInventoryApprovalStatus(input string) (*WindowsDriverUpdateInventoryApprovalStatus, error) {
	vals := map[string]WindowsDriverUpdateInventoryApprovalStatus{
		"approved":    WindowsDriverUpdateInventoryApprovalStatus_Approved,
		"declined":    WindowsDriverUpdateInventoryApprovalStatus_Declined,
		"needsreview": WindowsDriverUpdateInventoryApprovalStatus_NeedsReview,
		"suspended":   WindowsDriverUpdateInventoryApprovalStatus_Suspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateInventoryApprovalStatus(input)
	return &out, nil
}
