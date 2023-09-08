package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileApprovalType string

const (
	WindowsDriverUpdateProfileApprovalTypeautomatic WindowsDriverUpdateProfileApprovalType = "Automatic"
	WindowsDriverUpdateProfileApprovalTypemanual    WindowsDriverUpdateProfileApprovalType = "Manual"
)

func PossibleValuesForWindowsDriverUpdateProfileApprovalType() []string {
	return []string{
		string(WindowsDriverUpdateProfileApprovalTypeautomatic),
		string(WindowsDriverUpdateProfileApprovalTypemanual),
	}
}

func parseWindowsDriverUpdateProfileApprovalType(input string) (*WindowsDriverUpdateProfileApprovalType, error) {
	vals := map[string]WindowsDriverUpdateProfileApprovalType{
		"automatic": WindowsDriverUpdateProfileApprovalTypeautomatic,
		"manual":    WindowsDriverUpdateProfileApprovalTypemanual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateProfileApprovalType(input)
	return &out, nil
}
