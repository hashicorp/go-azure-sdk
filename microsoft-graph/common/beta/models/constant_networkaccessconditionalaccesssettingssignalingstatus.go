package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConditionalAccessSettingsSignalingStatus string

const (
	NetworkaccessConditionalAccessSettingsSignalingStatusdisabled NetworkaccessConditionalAccessSettingsSignalingStatus = "Disabled"
	NetworkaccessConditionalAccessSettingsSignalingStatusenabled  NetworkaccessConditionalAccessSettingsSignalingStatus = "Enabled"
)

func PossibleValuesForNetworkaccessConditionalAccessSettingsSignalingStatus() []string {
	return []string{
		string(NetworkaccessConditionalAccessSettingsSignalingStatusdisabled),
		string(NetworkaccessConditionalAccessSettingsSignalingStatusenabled),
	}
}

func parseNetworkaccessConditionalAccessSettingsSignalingStatus(input string) (*NetworkaccessConditionalAccessSettingsSignalingStatus, error) {
	vals := map[string]NetworkaccessConditionalAccessSettingsSignalingStatus{
		"disabled": NetworkaccessConditionalAccessSettingsSignalingStatusdisabled,
		"enabled":  NetworkaccessConditionalAccessSettingsSignalingStatusenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessConditionalAccessSettingsSignalingStatus(input)
	return &out, nil
}
