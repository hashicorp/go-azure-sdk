package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputConfigurationInputType string

const (
	AuthenticationAttributeCollectionInputConfigurationInputTypeboolean             AuthenticationAttributeCollectionInputConfigurationInputType = "Boolean"
	AuthenticationAttributeCollectionInputConfigurationInputTypecheckboxMultiSelect AuthenticationAttributeCollectionInputConfigurationInputType = "CheckboxMultiSelect"
	AuthenticationAttributeCollectionInputConfigurationInputTyperadioSingleSelect   AuthenticationAttributeCollectionInputConfigurationInputType = "RadioSingleSelect"
	AuthenticationAttributeCollectionInputConfigurationInputTypetext                AuthenticationAttributeCollectionInputConfigurationInputType = "Text"
)

func PossibleValuesForAuthenticationAttributeCollectionInputConfigurationInputType() []string {
	return []string{
		string(AuthenticationAttributeCollectionInputConfigurationInputTypeboolean),
		string(AuthenticationAttributeCollectionInputConfigurationInputTypecheckboxMultiSelect),
		string(AuthenticationAttributeCollectionInputConfigurationInputTyperadioSingleSelect),
		string(AuthenticationAttributeCollectionInputConfigurationInputTypetext),
	}
}

func parseAuthenticationAttributeCollectionInputConfigurationInputType(input string) (*AuthenticationAttributeCollectionInputConfigurationInputType, error) {
	vals := map[string]AuthenticationAttributeCollectionInputConfigurationInputType{
		"boolean":             AuthenticationAttributeCollectionInputConfigurationInputTypeboolean,
		"checkboxmultiselect": AuthenticationAttributeCollectionInputConfigurationInputTypecheckboxMultiSelect,
		"radiosingleselect":   AuthenticationAttributeCollectionInputConfigurationInputTyperadioSingleSelect,
		"text":                AuthenticationAttributeCollectionInputConfigurationInputTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAttributeCollectionInputConfigurationInputType(input)
	return &out, nil
}
