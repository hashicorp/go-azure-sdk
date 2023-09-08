package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSettingStatus string

const (
	AutomaticRepliesSettingStatusalwaysEnabled AutomaticRepliesSettingStatus = "AlwaysEnabled"
	AutomaticRepliesSettingStatusdisabled      AutomaticRepliesSettingStatus = "Disabled"
	AutomaticRepliesSettingStatusscheduled     AutomaticRepliesSettingStatus = "Scheduled"
)

func PossibleValuesForAutomaticRepliesSettingStatus() []string {
	return []string{
		string(AutomaticRepliesSettingStatusalwaysEnabled),
		string(AutomaticRepliesSettingStatusdisabled),
		string(AutomaticRepliesSettingStatusscheduled),
	}
}

func parseAutomaticRepliesSettingStatus(input string) (*AutomaticRepliesSettingStatus, error) {
	vals := map[string]AutomaticRepliesSettingStatus{
		"alwaysenabled": AutomaticRepliesSettingStatusalwaysEnabled,
		"disabled":      AutomaticRepliesSettingStatusdisabled,
		"scheduled":     AutomaticRepliesSettingStatusscheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesSettingStatus(input)
	return &out, nil
}
