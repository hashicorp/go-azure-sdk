package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeActionResultStatus string

const (
	DeviceScopeActionResultStatus_Failed    DeviceScopeActionResultStatus = "failed"
	DeviceScopeActionResultStatus_Succeeded DeviceScopeActionResultStatus = "succeeded"
)

func PossibleValuesForDeviceScopeActionResultStatus() []string {
	return []string{
		string(DeviceScopeActionResultStatus_Failed),
		string(DeviceScopeActionResultStatus_Succeeded),
	}
}

func (s *DeviceScopeActionResultStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceScopeActionResultStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceScopeActionResultStatus(input string) (*DeviceScopeActionResultStatus, error) {
	vals := map[string]DeviceScopeActionResultStatus{
		"failed":    DeviceScopeActionResultStatus_Failed,
		"succeeded": DeviceScopeActionResultStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceScopeActionResultStatus(input)
	return &out, nil
}
