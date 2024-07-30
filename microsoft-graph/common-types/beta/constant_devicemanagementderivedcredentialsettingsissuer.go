package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialSettingsIssuer string

const (
	DeviceManagementDerivedCredentialSettingsIssuer_EntrustDatacard DeviceManagementDerivedCredentialSettingsIssuer = "entrustDatacard"
	DeviceManagementDerivedCredentialSettingsIssuer_Intercede       DeviceManagementDerivedCredentialSettingsIssuer = "intercede"
	DeviceManagementDerivedCredentialSettingsIssuer_Purebred        DeviceManagementDerivedCredentialSettingsIssuer = "purebred"
	DeviceManagementDerivedCredentialSettingsIssuer_XTec            DeviceManagementDerivedCredentialSettingsIssuer = "xTec"
)

func PossibleValuesForDeviceManagementDerivedCredentialSettingsIssuer() []string {
	return []string{
		string(DeviceManagementDerivedCredentialSettingsIssuer_EntrustDatacard),
		string(DeviceManagementDerivedCredentialSettingsIssuer_Intercede),
		string(DeviceManagementDerivedCredentialSettingsIssuer_Purebred),
		string(DeviceManagementDerivedCredentialSettingsIssuer_XTec),
	}
}

func (s *DeviceManagementDerivedCredentialSettingsIssuer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementDerivedCredentialSettingsIssuer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementDerivedCredentialSettingsIssuer(input string) (*DeviceManagementDerivedCredentialSettingsIssuer, error) {
	vals := map[string]DeviceManagementDerivedCredentialSettingsIssuer{
		"entrustdatacard": DeviceManagementDerivedCredentialSettingsIssuer_EntrustDatacard,
		"intercede":       DeviceManagementDerivedCredentialSettingsIssuer_Intercede,
		"purebred":        DeviceManagementDerivedCredentialSettingsIssuer_Purebred,
		"xtec":            DeviceManagementDerivedCredentialSettingsIssuer_XTec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDerivedCredentialSettingsIssuer(input)
	return &out, nil
}
