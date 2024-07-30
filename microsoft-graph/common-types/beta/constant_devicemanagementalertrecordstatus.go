package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecordStatus string

const (
	DeviceManagementAlertRecordStatus_Active   DeviceManagementAlertRecordStatus = "active"
	DeviceManagementAlertRecordStatus_Resolved DeviceManagementAlertRecordStatus = "resolved"
)

func PossibleValuesForDeviceManagementAlertRecordStatus() []string {
	return []string{
		string(DeviceManagementAlertRecordStatus_Active),
		string(DeviceManagementAlertRecordStatus_Resolved),
	}
}

func (s *DeviceManagementAlertRecordStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertRecordStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertRecordStatus(input string) (*DeviceManagementAlertRecordStatus, error) {
	vals := map[string]DeviceManagementAlertRecordStatus{
		"active":   DeviceManagementAlertRecordStatus_Active,
		"resolved": DeviceManagementAlertRecordStatus_Resolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRecordStatus(input)
	return &out, nil
}
