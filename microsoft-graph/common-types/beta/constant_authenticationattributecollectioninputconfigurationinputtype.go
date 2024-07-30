package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputConfigurationInputType string

const (
	AuthenticationAttributeCollectionInputConfigurationInputType_Boolean             AuthenticationAttributeCollectionInputConfigurationInputType = "boolean"
	AuthenticationAttributeCollectionInputConfigurationInputType_CheckboxMultiSelect AuthenticationAttributeCollectionInputConfigurationInputType = "checkboxMultiSelect"
	AuthenticationAttributeCollectionInputConfigurationInputType_RadioSingleSelect   AuthenticationAttributeCollectionInputConfigurationInputType = "radioSingleSelect"
	AuthenticationAttributeCollectionInputConfigurationInputType_Text                AuthenticationAttributeCollectionInputConfigurationInputType = "text"
)

func PossibleValuesForAuthenticationAttributeCollectionInputConfigurationInputType() []string {
	return []string{
		string(AuthenticationAttributeCollectionInputConfigurationInputType_Boolean),
		string(AuthenticationAttributeCollectionInputConfigurationInputType_CheckboxMultiSelect),
		string(AuthenticationAttributeCollectionInputConfigurationInputType_RadioSingleSelect),
		string(AuthenticationAttributeCollectionInputConfigurationInputType_Text),
	}
}

func (s *AuthenticationAttributeCollectionInputConfigurationInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAttributeCollectionInputConfigurationInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAttributeCollectionInputConfigurationInputType(input string) (*AuthenticationAttributeCollectionInputConfigurationInputType, error) {
	vals := map[string]AuthenticationAttributeCollectionInputConfigurationInputType{
		"boolean":             AuthenticationAttributeCollectionInputConfigurationInputType_Boolean,
		"checkboxmultiselect": AuthenticationAttributeCollectionInputConfigurationInputType_CheckboxMultiSelect,
		"radiosingleselect":   AuthenticationAttributeCollectionInputConfigurationInputType_RadioSingleSelect,
		"text":                AuthenticationAttributeCollectionInputConfigurationInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAttributeCollectionInputConfigurationInputType(input)
	return &out, nil
}
