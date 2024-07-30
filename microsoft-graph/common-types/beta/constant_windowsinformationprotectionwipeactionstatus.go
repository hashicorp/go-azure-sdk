package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionWipeActionStatus string

const (
	WindowsInformationProtectionWipeActionStatus_Active       WindowsInformationProtectionWipeActionStatus = "active"
	WindowsInformationProtectionWipeActionStatus_Canceled     WindowsInformationProtectionWipeActionStatus = "canceled"
	WindowsInformationProtectionWipeActionStatus_Done         WindowsInformationProtectionWipeActionStatus = "done"
	WindowsInformationProtectionWipeActionStatus_Failed       WindowsInformationProtectionWipeActionStatus = "failed"
	WindowsInformationProtectionWipeActionStatus_None         WindowsInformationProtectionWipeActionStatus = "none"
	WindowsInformationProtectionWipeActionStatus_NotSupported WindowsInformationProtectionWipeActionStatus = "notSupported"
	WindowsInformationProtectionWipeActionStatus_Pending      WindowsInformationProtectionWipeActionStatus = "pending"
)

func PossibleValuesForWindowsInformationProtectionWipeActionStatus() []string {
	return []string{
		string(WindowsInformationProtectionWipeActionStatus_Active),
		string(WindowsInformationProtectionWipeActionStatus_Canceled),
		string(WindowsInformationProtectionWipeActionStatus_Done),
		string(WindowsInformationProtectionWipeActionStatus_Failed),
		string(WindowsInformationProtectionWipeActionStatus_None),
		string(WindowsInformationProtectionWipeActionStatus_NotSupported),
		string(WindowsInformationProtectionWipeActionStatus_Pending),
	}
}

func (s *WindowsInformationProtectionWipeActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionWipeActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionWipeActionStatus(input string) (*WindowsInformationProtectionWipeActionStatus, error) {
	vals := map[string]WindowsInformationProtectionWipeActionStatus{
		"active":       WindowsInformationProtectionWipeActionStatus_Active,
		"canceled":     WindowsInformationProtectionWipeActionStatus_Canceled,
		"done":         WindowsInformationProtectionWipeActionStatus_Done,
		"failed":       WindowsInformationProtectionWipeActionStatus_Failed,
		"none":         WindowsInformationProtectionWipeActionStatus_None,
		"notsupported": WindowsInformationProtectionWipeActionStatus_NotSupported,
		"pending":      WindowsInformationProtectionWipeActionStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionWipeActionStatus(input)
	return &out, nil
}
