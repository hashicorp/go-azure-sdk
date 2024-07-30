package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSettingStatus string

const (
	AutomaticRepliesSettingStatus_AlwaysEnabled AutomaticRepliesSettingStatus = "alwaysEnabled"
	AutomaticRepliesSettingStatus_Disabled      AutomaticRepliesSettingStatus = "disabled"
	AutomaticRepliesSettingStatus_Scheduled     AutomaticRepliesSettingStatus = "scheduled"
)

func PossibleValuesForAutomaticRepliesSettingStatus() []string {
	return []string{
		string(AutomaticRepliesSettingStatus_AlwaysEnabled),
		string(AutomaticRepliesSettingStatus_Disabled),
		string(AutomaticRepliesSettingStatus_Scheduled),
	}
}

func (s *AutomaticRepliesSettingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticRepliesSettingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticRepliesSettingStatus(input string) (*AutomaticRepliesSettingStatus, error) {
	vals := map[string]AutomaticRepliesSettingStatus{
		"alwaysenabled": AutomaticRepliesSettingStatus_AlwaysEnabled,
		"disabled":      AutomaticRepliesSettingStatus_Disabled,
		"scheduled":     AutomaticRepliesSettingStatus_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesSettingStatus(input)
	return &out, nil
}
