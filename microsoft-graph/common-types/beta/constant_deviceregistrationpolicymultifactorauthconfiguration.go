package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationPolicyMultiFactorAuthConfiguration string

const (
	DeviceRegistrationPolicyMultiFactorAuthConfiguration_NotRequired DeviceRegistrationPolicyMultiFactorAuthConfiguration = "notRequired"
	DeviceRegistrationPolicyMultiFactorAuthConfiguration_Required    DeviceRegistrationPolicyMultiFactorAuthConfiguration = "required"
)

func PossibleValuesForDeviceRegistrationPolicyMultiFactorAuthConfiguration() []string {
	return []string{
		string(DeviceRegistrationPolicyMultiFactorAuthConfiguration_NotRequired),
		string(DeviceRegistrationPolicyMultiFactorAuthConfiguration_Required),
	}
}

func (s *DeviceRegistrationPolicyMultiFactorAuthConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceRegistrationPolicyMultiFactorAuthConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceRegistrationPolicyMultiFactorAuthConfiguration(input string) (*DeviceRegistrationPolicyMultiFactorAuthConfiguration, error) {
	vals := map[string]DeviceRegistrationPolicyMultiFactorAuthConfiguration{
		"notrequired": DeviceRegistrationPolicyMultiFactorAuthConfiguration_NotRequired,
		"required":    DeviceRegistrationPolicyMultiFactorAuthConfiguration_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceRegistrationPolicyMultiFactorAuthConfiguration(input)
	return &out, nil
}
