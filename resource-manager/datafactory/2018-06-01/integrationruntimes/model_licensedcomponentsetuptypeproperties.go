package integrationruntimes

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicensedComponentSetupTypeProperties struct {
	ComponentName string     `json:"componentName"`
	LicenseKey    SecretBase `json:"licenseKey"`
}

var _ json.Unmarshaler = &LicensedComponentSetupTypeProperties{}

func (s *LicensedComponentSetupTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias LicensedComponentSetupTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into LicensedComponentSetupTypeProperties: %+v", err)
	}

	s.ComponentName = decoded.ComponentName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LicensedComponentSetupTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["licenseKey"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LicenseKey' for 'LicensedComponentSetupTypeProperties': %+v", err)
		}
		s.LicenseKey = impl
	}
	return nil
}
