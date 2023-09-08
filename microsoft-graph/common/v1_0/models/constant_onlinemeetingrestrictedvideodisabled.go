package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRestrictedVideoDisabled string

const (
	OnlineMeetingRestrictedVideoDisabledwatermarkProtection OnlineMeetingRestrictedVideoDisabled = "WatermarkProtection"
)

func PossibleValuesForOnlineMeetingRestrictedVideoDisabled() []string {
	return []string{
		string(OnlineMeetingRestrictedVideoDisabledwatermarkProtection),
	}
}

func parseOnlineMeetingRestrictedVideoDisabled(input string) (*OnlineMeetingRestrictedVideoDisabled, error) {
	vals := map[string]OnlineMeetingRestrictedVideoDisabled{
		"watermarkprotection": OnlineMeetingRestrictedVideoDisabledwatermarkProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingRestrictedVideoDisabled(input)
	return &out, nil
}
