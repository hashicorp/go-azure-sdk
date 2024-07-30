package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSettingsUserType string

const (
	OutOfBoxExperienceSettingsUserType_Administrator OutOfBoxExperienceSettingsUserType = "administrator"
	OutOfBoxExperienceSettingsUserType_Standard      OutOfBoxExperienceSettingsUserType = "standard"
)

func PossibleValuesForOutOfBoxExperienceSettingsUserType() []string {
	return []string{
		string(OutOfBoxExperienceSettingsUserType_Administrator),
		string(OutOfBoxExperienceSettingsUserType_Standard),
	}
}

func (s *OutOfBoxExperienceSettingsUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutOfBoxExperienceSettingsUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutOfBoxExperienceSettingsUserType(input string) (*OutOfBoxExperienceSettingsUserType, error) {
	vals := map[string]OutOfBoxExperienceSettingsUserType{
		"administrator": OutOfBoxExperienceSettingsUserType_Administrator,
		"standard":      OutOfBoxExperienceSettingsUserType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutOfBoxExperienceSettingsUserType(input)
	return &out, nil
}
