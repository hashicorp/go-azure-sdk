package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettingsPreviewVisibility string

const (
	IosNotificationSettingsPreviewVisibility_AlwaysShow     IosNotificationSettingsPreviewVisibility = "alwaysShow"
	IosNotificationSettingsPreviewVisibility_HideWhenLocked IosNotificationSettingsPreviewVisibility = "hideWhenLocked"
	IosNotificationSettingsPreviewVisibility_NeverShow      IosNotificationSettingsPreviewVisibility = "neverShow"
	IosNotificationSettingsPreviewVisibility_NotConfigured  IosNotificationSettingsPreviewVisibility = "notConfigured"
)

func PossibleValuesForIosNotificationSettingsPreviewVisibility() []string {
	return []string{
		string(IosNotificationSettingsPreviewVisibility_AlwaysShow),
		string(IosNotificationSettingsPreviewVisibility_HideWhenLocked),
		string(IosNotificationSettingsPreviewVisibility_NeverShow),
		string(IosNotificationSettingsPreviewVisibility_NotConfigured),
	}
}

func (s *IosNotificationSettingsPreviewVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosNotificationSettingsPreviewVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosNotificationSettingsPreviewVisibility(input string) (*IosNotificationSettingsPreviewVisibility, error) {
	vals := map[string]IosNotificationSettingsPreviewVisibility{
		"alwaysshow":     IosNotificationSettingsPreviewVisibility_AlwaysShow,
		"hidewhenlocked": IosNotificationSettingsPreviewVisibility_HideWhenLocked,
		"nevershow":      IosNotificationSettingsPreviewVisibility_NeverShow,
		"notconfigured":  IosNotificationSettingsPreviewVisibility_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosNotificationSettingsPreviewVisibility(input)
	return &out, nil
}
