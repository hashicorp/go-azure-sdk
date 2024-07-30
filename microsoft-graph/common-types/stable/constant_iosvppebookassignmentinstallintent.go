package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppEBookAssignmentInstallIntent string

const (
	IosVppEBookAssignmentInstallIntent_Available                  IosVppEBookAssignmentInstallIntent = "available"
	IosVppEBookAssignmentInstallIntent_AvailableWithoutEnrollment IosVppEBookAssignmentInstallIntent = "availableWithoutEnrollment"
	IosVppEBookAssignmentInstallIntent_Required                   IosVppEBookAssignmentInstallIntent = "required"
	IosVppEBookAssignmentInstallIntent_Uninstall                  IosVppEBookAssignmentInstallIntent = "uninstall"
)

func PossibleValuesForIosVppEBookAssignmentInstallIntent() []string {
	return []string{
		string(IosVppEBookAssignmentInstallIntent_Available),
		string(IosVppEBookAssignmentInstallIntent_AvailableWithoutEnrollment),
		string(IosVppEBookAssignmentInstallIntent_Required),
		string(IosVppEBookAssignmentInstallIntent_Uninstall),
	}
}

func (s *IosVppEBookAssignmentInstallIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVppEBookAssignmentInstallIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVppEBookAssignmentInstallIntent(input string) (*IosVppEBookAssignmentInstallIntent, error) {
	vals := map[string]IosVppEBookAssignmentInstallIntent{
		"available":                  IosVppEBookAssignmentInstallIntent_Available,
		"availablewithoutenrollment": IosVppEBookAssignmentInstallIntent_AvailableWithoutEnrollment,
		"required":                   IosVppEBookAssignmentInstallIntent_Required,
		"uninstall":                  IosVppEBookAssignmentInstallIntent_Uninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppEBookAssignmentInstallIntent(input)
	return &out, nil
}
