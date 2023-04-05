package listkeys

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationKeyName string

const (
	AutomationKeyNamePrimary   AutomationKeyName = "Primary"
	AutomationKeyNameSecondary AutomationKeyName = "Secondary"
)

func PossibleValuesForAutomationKeyName() []string {
	return []string{
		string(AutomationKeyNamePrimary),
		string(AutomationKeyNameSecondary),
	}
}

func (s *AutomationKeyName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAutomationKeyName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AutomationKeyName(decoded)
	return nil
}

type AutomationKeyPermissions string

const (
	AutomationKeyPermissionsFull AutomationKeyPermissions = "Full"
	AutomationKeyPermissionsRead AutomationKeyPermissions = "Read"
)

func PossibleValuesForAutomationKeyPermissions() []string {
	return []string{
		string(AutomationKeyPermissionsFull),
		string(AutomationKeyPermissionsRead),
	}
}

func (s *AutomationKeyPermissions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAutomationKeyPermissions() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AutomationKeyPermissions(decoded)
	return nil
}
