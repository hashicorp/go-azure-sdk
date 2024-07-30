package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateStateStatus string

const (
	WindowsUpdateStateStatus_Failed              WindowsUpdateStateStatus = "failed"
	WindowsUpdateStateStatus_PendingInstallation WindowsUpdateStateStatus = "pendingInstallation"
	WindowsUpdateStateStatus_PendingReboot       WindowsUpdateStateStatus = "pendingReboot"
	WindowsUpdateStateStatus_UpToDate            WindowsUpdateStateStatus = "upToDate"
)

func PossibleValuesForWindowsUpdateStateStatus() []string {
	return []string{
		string(WindowsUpdateStateStatus_Failed),
		string(WindowsUpdateStateStatus_PendingInstallation),
		string(WindowsUpdateStateStatus_PendingReboot),
		string(WindowsUpdateStateStatus_UpToDate),
	}
}

func (s *WindowsUpdateStateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateStateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateStateStatus(input string) (*WindowsUpdateStateStatus, error) {
	vals := map[string]WindowsUpdateStateStatus{
		"failed":              WindowsUpdateStateStatus_Failed,
		"pendinginstallation": WindowsUpdateStateStatus_PendingInstallation,
		"pendingreboot":       WindowsUpdateStateStatus_PendingReboot,
		"uptodate":            WindowsUpdateStateStatus_UpToDate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateStateStatus(input)
	return &out, nil
}
