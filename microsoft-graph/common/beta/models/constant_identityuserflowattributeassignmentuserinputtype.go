package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeAssignmentUserInputType string

const (
	IdentityUserFlowAttributeAssignmentUserInputTypecheckboxMultiSelect  IdentityUserFlowAttributeAssignmentUserInputType = "CheckboxMultiSelect"
	IdentityUserFlowAttributeAssignmentUserInputTypedateTimeDropdown     IdentityUserFlowAttributeAssignmentUserInputType = "DateTimeDropdown"
	IdentityUserFlowAttributeAssignmentUserInputTypedropdownSingleSelect IdentityUserFlowAttributeAssignmentUserInputType = "DropdownSingleSelect"
	IdentityUserFlowAttributeAssignmentUserInputTypeemailBox             IdentityUserFlowAttributeAssignmentUserInputType = "EmailBox"
	IdentityUserFlowAttributeAssignmentUserInputTyperadioSingleSelect    IdentityUserFlowAttributeAssignmentUserInputType = "RadioSingleSelect"
	IdentityUserFlowAttributeAssignmentUserInputTypetextBox              IdentityUserFlowAttributeAssignmentUserInputType = "TextBox"
)

func PossibleValuesForIdentityUserFlowAttributeAssignmentUserInputType() []string {
	return []string{
		string(IdentityUserFlowAttributeAssignmentUserInputTypecheckboxMultiSelect),
		string(IdentityUserFlowAttributeAssignmentUserInputTypedateTimeDropdown),
		string(IdentityUserFlowAttributeAssignmentUserInputTypedropdownSingleSelect),
		string(IdentityUserFlowAttributeAssignmentUserInputTypeemailBox),
		string(IdentityUserFlowAttributeAssignmentUserInputTyperadioSingleSelect),
		string(IdentityUserFlowAttributeAssignmentUserInputTypetextBox),
	}
}

func parseIdentityUserFlowAttributeAssignmentUserInputType(input string) (*IdentityUserFlowAttributeAssignmentUserInputType, error) {
	vals := map[string]IdentityUserFlowAttributeAssignmentUserInputType{
		"checkboxmultiselect":  IdentityUserFlowAttributeAssignmentUserInputTypecheckboxMultiSelect,
		"datetimedropdown":     IdentityUserFlowAttributeAssignmentUserInputTypedateTimeDropdown,
		"dropdownsingleselect": IdentityUserFlowAttributeAssignmentUserInputTypedropdownSingleSelect,
		"emailbox":             IdentityUserFlowAttributeAssignmentUserInputTypeemailBox,
		"radiosingleselect":    IdentityUserFlowAttributeAssignmentUserInputTyperadioSingleSelect,
		"textbox":              IdentityUserFlowAttributeAssignmentUserInputTypetextBox,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeAssignmentUserInputType(input)
	return &out, nil
}
