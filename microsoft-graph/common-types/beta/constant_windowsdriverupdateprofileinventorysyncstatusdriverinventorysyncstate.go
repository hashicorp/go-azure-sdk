package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState string

const (
	WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Failure WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState = "failure"
	WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Pending WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState = "pending"
	WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Success WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState = "success"
)

func PossibleValuesForWindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState() []string {
	return []string{
		string(WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Failure),
		string(WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Pending),
		string(WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Success),
	}
}

func (s *WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState(input string) (*WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState, error) {
	vals := map[string]WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState{
		"failure": WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Failure,
		"pending": WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Pending,
		"success": WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateProfileInventorySyncStatusDriverInventorySyncState(input)
	return &out, nil
}
