package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode string

const (
	AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Default   AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode = "default"
	AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Postponed AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode = "postponed"
	AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Priority  AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode = "priority"
)

func PossibleValuesForAndroidManagedStoreAppAssignmentSettingsAutoUpdateMode() []string {
	return []string{
		string(AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Default),
		string(AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Postponed),
		string(AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Priority),
	}
}

func (s *AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAppAssignmentSettingsAutoUpdateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAppAssignmentSettingsAutoUpdateMode(input string) (*AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode, error) {
	vals := map[string]AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode{
		"default":   AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Default,
		"postponed": AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Postponed,
		"priority":  AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode_Priority,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode(input)
	return &out, nil
}
