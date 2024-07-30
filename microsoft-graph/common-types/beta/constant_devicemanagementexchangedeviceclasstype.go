package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeDeviceClassType string

const (
	DeviceManagementExchangeDeviceClassType_Family DeviceManagementExchangeDeviceClassType = "family"
	DeviceManagementExchangeDeviceClassType_Model  DeviceManagementExchangeDeviceClassType = "model"
)

func PossibleValuesForDeviceManagementExchangeDeviceClassType() []string {
	return []string{
		string(DeviceManagementExchangeDeviceClassType_Family),
		string(DeviceManagementExchangeDeviceClassType_Model),
	}
}

func (s *DeviceManagementExchangeDeviceClassType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeDeviceClassType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeDeviceClassType(input string) (*DeviceManagementExchangeDeviceClassType, error) {
	vals := map[string]DeviceManagementExchangeDeviceClassType{
		"family": DeviceManagementExchangeDeviceClassType_Family,
		"model":  DeviceManagementExchangeDeviceClassType_Model,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeDeviceClassType(input)
	return &out, nil
}
