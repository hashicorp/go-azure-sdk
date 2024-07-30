package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoDeny      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "autoDeny"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoGrant     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "autoGrant"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_DeviceDefault AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "deviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_Prompt        AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "prompt"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoDeny),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoGrant),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_DeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_Prompt),
	}
}

func (s *AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy{
		"autodeny":      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoDeny,
		"autogrant":     AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoGrant,
		"devicedefault": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_DeviceDefault,
		"prompt":        AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input)
	return &out, nil
}
