package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDeviceDiscoveryTaskStatus string

const (
	UnmanagedDeviceDiscoveryTaskStatus_Active    UnmanagedDeviceDiscoveryTaskStatus = "active"
	UnmanagedDeviceDiscoveryTaskStatus_Completed UnmanagedDeviceDiscoveryTaskStatus = "completed"
	UnmanagedDeviceDiscoveryTaskStatus_Pending   UnmanagedDeviceDiscoveryTaskStatus = "pending"
	UnmanagedDeviceDiscoveryTaskStatus_Rejected  UnmanagedDeviceDiscoveryTaskStatus = "rejected"
	UnmanagedDeviceDiscoveryTaskStatus_Unknown   UnmanagedDeviceDiscoveryTaskStatus = "unknown"
)

func PossibleValuesForUnmanagedDeviceDiscoveryTaskStatus() []string {
	return []string{
		string(UnmanagedDeviceDiscoveryTaskStatus_Active),
		string(UnmanagedDeviceDiscoveryTaskStatus_Completed),
		string(UnmanagedDeviceDiscoveryTaskStatus_Pending),
		string(UnmanagedDeviceDiscoveryTaskStatus_Rejected),
		string(UnmanagedDeviceDiscoveryTaskStatus_Unknown),
	}
}

func (s *UnmanagedDeviceDiscoveryTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnmanagedDeviceDiscoveryTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnmanagedDeviceDiscoveryTaskStatus(input string) (*UnmanagedDeviceDiscoveryTaskStatus, error) {
	vals := map[string]UnmanagedDeviceDiscoveryTaskStatus{
		"active":    UnmanagedDeviceDiscoveryTaskStatus_Active,
		"completed": UnmanagedDeviceDiscoveryTaskStatus_Completed,
		"pending":   UnmanagedDeviceDiscoveryTaskStatus_Pending,
		"rejected":  UnmanagedDeviceDiscoveryTaskStatus_Rejected,
		"unknown":   UnmanagedDeviceDiscoveryTaskStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmanagedDeviceDiscoveryTaskStatus(input)
	return &out, nil
}
