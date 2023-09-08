package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettingsPreviewVisibility string

const (
	IosNotificationSettingsPreviewVisibilityalwaysShow     IosNotificationSettingsPreviewVisibility = "AlwaysShow"
	IosNotificationSettingsPreviewVisibilityhideWhenLocked IosNotificationSettingsPreviewVisibility = "HideWhenLocked"
	IosNotificationSettingsPreviewVisibilityneverShow      IosNotificationSettingsPreviewVisibility = "NeverShow"
	IosNotificationSettingsPreviewVisibilitynotConfigured  IosNotificationSettingsPreviewVisibility = "NotConfigured"
)

func PossibleValuesForIosNotificationSettingsPreviewVisibility() []string {
	return []string{
		string(IosNotificationSettingsPreviewVisibilityalwaysShow),
		string(IosNotificationSettingsPreviewVisibilityhideWhenLocked),
		string(IosNotificationSettingsPreviewVisibilityneverShow),
		string(IosNotificationSettingsPreviewVisibilitynotConfigured),
	}
}

func parseIosNotificationSettingsPreviewVisibility(input string) (*IosNotificationSettingsPreviewVisibility, error) {
	vals := map[string]IosNotificationSettingsPreviewVisibility{
		"alwaysshow":     IosNotificationSettingsPreviewVisibilityalwaysShow,
		"hidewhenlocked": IosNotificationSettingsPreviewVisibilityhideWhenLocked,
		"nevershow":      IosNotificationSettingsPreviewVisibilityneverShow,
		"notconfigured":  IosNotificationSettingsPreviewVisibilitynotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosNotificationSettingsPreviewVisibility(input)
	return &out, nil
}
