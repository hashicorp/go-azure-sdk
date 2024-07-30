package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionResponseStatus string

const (
	DeviceLogCollectionResponseStatus_Completed DeviceLogCollectionResponseStatus = "completed"
	DeviceLogCollectionResponseStatus_Failed    DeviceLogCollectionResponseStatus = "failed"
	DeviceLogCollectionResponseStatus_Pending   DeviceLogCollectionResponseStatus = "pending"
)

func PossibleValuesForDeviceLogCollectionResponseStatus() []string {
	return []string{
		string(DeviceLogCollectionResponseStatus_Completed),
		string(DeviceLogCollectionResponseStatus_Failed),
		string(DeviceLogCollectionResponseStatus_Pending),
	}
}

func (s *DeviceLogCollectionResponseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceLogCollectionResponseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceLogCollectionResponseStatus(input string) (*DeviceLogCollectionResponseStatus, error) {
	vals := map[string]DeviceLogCollectionResponseStatus{
		"completed": DeviceLogCollectionResponseStatus_Completed,
		"failed":    DeviceLogCollectionResponseStatus_Failed,
		"pending":   DeviceLogCollectionResponseStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionResponseStatus(input)
	return &out, nil
}
