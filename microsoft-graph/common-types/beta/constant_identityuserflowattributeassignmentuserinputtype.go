package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeAssignmentUserInputType string

const (
	IdentityUserFlowAttributeAssignmentUserInputType_CheckboxMultiSelect  IdentityUserFlowAttributeAssignmentUserInputType = "checkboxMultiSelect"
	IdentityUserFlowAttributeAssignmentUserInputType_DateTimeDropdown     IdentityUserFlowAttributeAssignmentUserInputType = "dateTimeDropdown"
	IdentityUserFlowAttributeAssignmentUserInputType_DropdownSingleSelect IdentityUserFlowAttributeAssignmentUserInputType = "dropdownSingleSelect"
	IdentityUserFlowAttributeAssignmentUserInputType_EmailBox             IdentityUserFlowAttributeAssignmentUserInputType = "emailBox"
	IdentityUserFlowAttributeAssignmentUserInputType_RadioSingleSelect    IdentityUserFlowAttributeAssignmentUserInputType = "radioSingleSelect"
	IdentityUserFlowAttributeAssignmentUserInputType_TextBox              IdentityUserFlowAttributeAssignmentUserInputType = "textBox"
)

func PossibleValuesForIdentityUserFlowAttributeAssignmentUserInputType() []string {
	return []string{
		string(IdentityUserFlowAttributeAssignmentUserInputType_CheckboxMultiSelect),
		string(IdentityUserFlowAttributeAssignmentUserInputType_DateTimeDropdown),
		string(IdentityUserFlowAttributeAssignmentUserInputType_DropdownSingleSelect),
		string(IdentityUserFlowAttributeAssignmentUserInputType_EmailBox),
		string(IdentityUserFlowAttributeAssignmentUserInputType_RadioSingleSelect),
		string(IdentityUserFlowAttributeAssignmentUserInputType_TextBox),
	}
}

func (s *IdentityUserFlowAttributeAssignmentUserInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityUserFlowAttributeAssignmentUserInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityUserFlowAttributeAssignmentUserInputType(input string) (*IdentityUserFlowAttributeAssignmentUserInputType, error) {
	vals := map[string]IdentityUserFlowAttributeAssignmentUserInputType{
		"checkboxmultiselect":  IdentityUserFlowAttributeAssignmentUserInputType_CheckboxMultiSelect,
		"datetimedropdown":     IdentityUserFlowAttributeAssignmentUserInputType_DateTimeDropdown,
		"dropdownsingleselect": IdentityUserFlowAttributeAssignmentUserInputType_DropdownSingleSelect,
		"emailbox":             IdentityUserFlowAttributeAssignmentUserInputType_EmailBox,
		"radiosingleselect":    IdentityUserFlowAttributeAssignmentUserInputType_RadioSingleSelect,
		"textbox":              IdentityUserFlowAttributeAssignmentUserInputType_TextBox,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeAssignmentUserInputType(input)
	return &out, nil
}
