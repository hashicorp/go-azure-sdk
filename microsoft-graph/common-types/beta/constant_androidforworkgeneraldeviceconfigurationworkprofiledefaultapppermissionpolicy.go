package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoDeny      AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "autoDeny"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoGrant     AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "autoGrant"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_DeviceDefault AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "deviceDefault"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_Prompt        AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy = "prompt"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoDeny),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoGrant),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_DeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_Prompt),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy{
		"autodeny":      AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoDeny,
		"autogrant":     AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_AutoGrant,
		"devicedefault": AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_DeviceDefault,
		"prompt":        AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy_Prompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy(input)
	return &out, nil
}
