package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDeviceDiscoveryTaskPriority string

const (
	UnmanagedDeviceDiscoveryTaskPriority_High UnmanagedDeviceDiscoveryTaskPriority = "high"
	UnmanagedDeviceDiscoveryTaskPriority_Low  UnmanagedDeviceDiscoveryTaskPriority = "low"
	UnmanagedDeviceDiscoveryTaskPriority_None UnmanagedDeviceDiscoveryTaskPriority = "none"
)

func PossibleValuesForUnmanagedDeviceDiscoveryTaskPriority() []string {
	return []string{
		string(UnmanagedDeviceDiscoveryTaskPriority_High),
		string(UnmanagedDeviceDiscoveryTaskPriority_Low),
		string(UnmanagedDeviceDiscoveryTaskPriority_None),
	}
}

func (s *UnmanagedDeviceDiscoveryTaskPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnmanagedDeviceDiscoveryTaskPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnmanagedDeviceDiscoveryTaskPriority(input string) (*UnmanagedDeviceDiscoveryTaskPriority, error) {
	vals := map[string]UnmanagedDeviceDiscoveryTaskPriority{
		"high": UnmanagedDeviceDiscoveryTaskPriority_High,
		"low":  UnmanagedDeviceDiscoveryTaskPriority_Low,
		"none": UnmanagedDeviceDiscoveryTaskPriority_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmanagedDeviceDiscoveryTaskPriority(input)
	return &out, nil
}
