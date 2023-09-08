package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesSettingExternalAudience string

const (
	AutomaticRepliesSettingExternalAudienceall          AutomaticRepliesSettingExternalAudience = "All"
	AutomaticRepliesSettingExternalAudiencecontactsOnly AutomaticRepliesSettingExternalAudience = "ContactsOnly"
	AutomaticRepliesSettingExternalAudiencenone         AutomaticRepliesSettingExternalAudience = "None"
)

func PossibleValuesForAutomaticRepliesSettingExternalAudience() []string {
	return []string{
		string(AutomaticRepliesSettingExternalAudienceall),
		string(AutomaticRepliesSettingExternalAudiencecontactsOnly),
		string(AutomaticRepliesSettingExternalAudiencenone),
	}
}

func parseAutomaticRepliesSettingExternalAudience(input string) (*AutomaticRepliesSettingExternalAudience, error) {
	vals := map[string]AutomaticRepliesSettingExternalAudience{
		"all":          AutomaticRepliesSettingExternalAudienceall,
		"contactsonly": AutomaticRepliesSettingExternalAudiencecontactsOnly,
		"none":         AutomaticRepliesSettingExternalAudiencenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesSettingExternalAudience(input)
	return &out, nil
}
