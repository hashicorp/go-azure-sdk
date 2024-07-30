package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEBookAssignmentInstallIntent string

const (
	ManagedEBookAssignmentInstallIntent_Available                  ManagedEBookAssignmentInstallIntent = "available"
	ManagedEBookAssignmentInstallIntent_AvailableWithoutEnrollment ManagedEBookAssignmentInstallIntent = "availableWithoutEnrollment"
	ManagedEBookAssignmentInstallIntent_Required                   ManagedEBookAssignmentInstallIntent = "required"
	ManagedEBookAssignmentInstallIntent_Uninstall                  ManagedEBookAssignmentInstallIntent = "uninstall"
)

func PossibleValuesForManagedEBookAssignmentInstallIntent() []string {
	return []string{
		string(ManagedEBookAssignmentInstallIntent_Available),
		string(ManagedEBookAssignmentInstallIntent_AvailableWithoutEnrollment),
		string(ManagedEBookAssignmentInstallIntent_Required),
		string(ManagedEBookAssignmentInstallIntent_Uninstall),
	}
}

func (s *ManagedEBookAssignmentInstallIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedEBookAssignmentInstallIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedEBookAssignmentInstallIntent(input string) (*ManagedEBookAssignmentInstallIntent, error) {
	vals := map[string]ManagedEBookAssignmentInstallIntent{
		"available":                  ManagedEBookAssignmentInstallIntent_Available,
		"availablewithoutenrollment": ManagedEBookAssignmentInstallIntent_AvailableWithoutEnrollment,
		"required":                   ManagedEBookAssignmentInstallIntent_Required,
		"uninstall":                  ManagedEBookAssignmentInstallIntent_Uninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedEBookAssignmentInstallIntent(input)
	return &out, nil
}
