package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRestrictedContentSharingDisabled string

const (
	OnlineMeetingRestrictedContentSharingDisabledwatermarkProtection OnlineMeetingRestrictedContentSharingDisabled = "WatermarkProtection"
)

func PossibleValuesForOnlineMeetingRestrictedContentSharingDisabled() []string {
	return []string{
		string(OnlineMeetingRestrictedContentSharingDisabledwatermarkProtection),
	}
}

func parseOnlineMeetingRestrictedContentSharingDisabled(input string) (*OnlineMeetingRestrictedContentSharingDisabled, error) {
	vals := map[string]OnlineMeetingRestrictedContentSharingDisabled{
		"watermarkprotection": OnlineMeetingRestrictedContentSharingDisabledwatermarkProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingRestrictedContentSharingDisabled(input)
	return &out, nil
}
