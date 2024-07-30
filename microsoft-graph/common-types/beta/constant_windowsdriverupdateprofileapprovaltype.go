package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileApprovalType string

const (
	WindowsDriverUpdateProfileApprovalType_Automatic WindowsDriverUpdateProfileApprovalType = "automatic"
	WindowsDriverUpdateProfileApprovalType_Manual    WindowsDriverUpdateProfileApprovalType = "manual"
)

func PossibleValuesForWindowsDriverUpdateProfileApprovalType() []string {
	return []string{
		string(WindowsDriverUpdateProfileApprovalType_Automatic),
		string(WindowsDriverUpdateProfileApprovalType_Manual),
	}
}

func (s *WindowsDriverUpdateProfileApprovalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDriverUpdateProfileApprovalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDriverUpdateProfileApprovalType(input string) (*WindowsDriverUpdateProfileApprovalType, error) {
	vals := map[string]WindowsDriverUpdateProfileApprovalType{
		"automatic": WindowsDriverUpdateProfileApprovalType_Automatic,
		"manual":    WindowsDriverUpdateProfileApprovalType_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateProfileApprovalType(input)
	return &out, nil
}
