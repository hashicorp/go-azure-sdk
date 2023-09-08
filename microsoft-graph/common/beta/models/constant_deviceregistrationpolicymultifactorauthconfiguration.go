package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationPolicyMultiFactorAuthConfiguration string

const (
	DeviceRegistrationPolicyMultiFactorAuthConfigurationnotRequired DeviceRegistrationPolicyMultiFactorAuthConfiguration = "NotRequired"
	DeviceRegistrationPolicyMultiFactorAuthConfigurationrequired    DeviceRegistrationPolicyMultiFactorAuthConfiguration = "Required"
)

func PossibleValuesForDeviceRegistrationPolicyMultiFactorAuthConfiguration() []string {
	return []string{
		string(DeviceRegistrationPolicyMultiFactorAuthConfigurationnotRequired),
		string(DeviceRegistrationPolicyMultiFactorAuthConfigurationrequired),
	}
}

func parseDeviceRegistrationPolicyMultiFactorAuthConfiguration(input string) (*DeviceRegistrationPolicyMultiFactorAuthConfiguration, error) {
	vals := map[string]DeviceRegistrationPolicyMultiFactorAuthConfiguration{
		"notrequired": DeviceRegistrationPolicyMultiFactorAuthConfigurationnotRequired,
		"required":    DeviceRegistrationPolicyMultiFactorAuthConfigurationrequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceRegistrationPolicyMultiFactorAuthConfiguration(input)
	return &out, nil
}
