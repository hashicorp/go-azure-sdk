package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSettingsExternalUserLifecycleAction string

const (
	EntitlementManagementSettingsExternalUserLifecycleAction_BlockSignIn          EntitlementManagementSettingsExternalUserLifecycleAction = "blockSignIn"
	EntitlementManagementSettingsExternalUserLifecycleAction_BlockSignInAndDelete EntitlementManagementSettingsExternalUserLifecycleAction = "blockSignInAndDelete"
	EntitlementManagementSettingsExternalUserLifecycleAction_None                 EntitlementManagementSettingsExternalUserLifecycleAction = "none"
)

func PossibleValuesForEntitlementManagementSettingsExternalUserLifecycleAction() []string {
	return []string{
		string(EntitlementManagementSettingsExternalUserLifecycleAction_BlockSignIn),
		string(EntitlementManagementSettingsExternalUserLifecycleAction_BlockSignInAndDelete),
		string(EntitlementManagementSettingsExternalUserLifecycleAction_None),
	}
}

func (s *EntitlementManagementSettingsExternalUserLifecycleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntitlementManagementSettingsExternalUserLifecycleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntitlementManagementSettingsExternalUserLifecycleAction(input string) (*EntitlementManagementSettingsExternalUserLifecycleAction, error) {
	vals := map[string]EntitlementManagementSettingsExternalUserLifecycleAction{
		"blocksignin":          EntitlementManagementSettingsExternalUserLifecycleAction_BlockSignIn,
		"blocksigninanddelete": EntitlementManagementSettingsExternalUserLifecycleAction_BlockSignInAndDelete,
		"none":                 EntitlementManagementSettingsExternalUserLifecycleAction_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntitlementManagementSettingsExternalUserLifecycleAction(input)
	return &out, nil
}
