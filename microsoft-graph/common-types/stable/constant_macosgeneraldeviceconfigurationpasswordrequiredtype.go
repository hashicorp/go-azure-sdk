package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfigurationPasswordRequiredType string

const (
	MacOSGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric  MacOSGeneralDeviceConfigurationPasswordRequiredType = "alphanumeric"
	MacOSGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault MacOSGeneralDeviceConfigurationPasswordRequiredType = "deviceDefault"
	MacOSGeneralDeviceConfigurationPasswordRequiredType_Numeric       MacOSGeneralDeviceConfigurationPasswordRequiredType = "numeric"
)

func PossibleValuesForMacOSGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(MacOSGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric),
		string(MacOSGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault),
		string(MacOSGeneralDeviceConfigurationPasswordRequiredType_Numeric),
	}
}

func (s *MacOSGeneralDeviceConfigurationPasswordRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSGeneralDeviceConfigurationPasswordRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSGeneralDeviceConfigurationPasswordRequiredType(input string) (*MacOSGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]MacOSGeneralDeviceConfigurationPasswordRequiredType{
		"alphanumeric":  MacOSGeneralDeviceConfigurationPasswordRequiredType_Alphanumeric,
		"devicedefault": MacOSGeneralDeviceConfigurationPasswordRequiredType_DeviceDefault,
		"numeric":       MacOSGeneralDeviceConfigurationPasswordRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
