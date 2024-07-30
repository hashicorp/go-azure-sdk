package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointSettingsImageTaggingOption string

const (
	SharepointSettingsImageTaggingOption_Basic    SharepointSettingsImageTaggingOption = "basic"
	SharepointSettingsImageTaggingOption_Disabled SharepointSettingsImageTaggingOption = "disabled"
	SharepointSettingsImageTaggingOption_Enhanced SharepointSettingsImageTaggingOption = "enhanced"
)

func PossibleValuesForSharepointSettingsImageTaggingOption() []string {
	return []string{
		string(SharepointSettingsImageTaggingOption_Basic),
		string(SharepointSettingsImageTaggingOption_Disabled),
		string(SharepointSettingsImageTaggingOption_Enhanced),
	}
}

func (s *SharepointSettingsImageTaggingOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharepointSettingsImageTaggingOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharepointSettingsImageTaggingOption(input string) (*SharepointSettingsImageTaggingOption, error) {
	vals := map[string]SharepointSettingsImageTaggingOption{
		"basic":    SharepointSettingsImageTaggingOption_Basic,
		"disabled": SharepointSettingsImageTaggingOption_Disabled,
		"enhanced": SharepointSettingsImageTaggingOption_Enhanced,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharepointSettingsImageTaggingOption(input)
	return &out, nil
}
