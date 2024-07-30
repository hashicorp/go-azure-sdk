package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationPasscodeRequiredType string

const (
	IosGeneralDeviceConfigurationPasscodeRequiredType_Alphanumeric  IosGeneralDeviceConfigurationPasscodeRequiredType = "alphanumeric"
	IosGeneralDeviceConfigurationPasscodeRequiredType_DeviceDefault IosGeneralDeviceConfigurationPasscodeRequiredType = "deviceDefault"
	IosGeneralDeviceConfigurationPasscodeRequiredType_Numeric       IosGeneralDeviceConfigurationPasscodeRequiredType = "numeric"
)

func PossibleValuesForIosGeneralDeviceConfigurationPasscodeRequiredType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationPasscodeRequiredType_Alphanumeric),
		string(IosGeneralDeviceConfigurationPasscodeRequiredType_DeviceDefault),
		string(IosGeneralDeviceConfigurationPasscodeRequiredType_Numeric),
	}
}

func (s *IosGeneralDeviceConfigurationPasscodeRequiredType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosGeneralDeviceConfigurationPasscodeRequiredType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosGeneralDeviceConfigurationPasscodeRequiredType(input string) (*IosGeneralDeviceConfigurationPasscodeRequiredType, error) {
	vals := map[string]IosGeneralDeviceConfigurationPasscodeRequiredType{
		"alphanumeric":  IosGeneralDeviceConfigurationPasscodeRequiredType_Alphanumeric,
		"devicedefault": IosGeneralDeviceConfigurationPasscodeRequiredType_DeviceDefault,
		"numeric":       IosGeneralDeviceConfigurationPasscodeRequiredType_Numeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationPasscodeRequiredType(input)
	return &out, nil
}
