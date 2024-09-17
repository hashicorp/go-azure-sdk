package featuresupport

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportStatus string

const (
	SupportStatusDefaultOFF   SupportStatus = "DefaultOFF"
	SupportStatusDefaultON    SupportStatus = "DefaultON"
	SupportStatusInvalid      SupportStatus = "Invalid"
	SupportStatusNotSupported SupportStatus = "NotSupported"
	SupportStatusSupported    SupportStatus = "Supported"
)

func PossibleValuesForSupportStatus() []string {
	return []string{
		string(SupportStatusDefaultOFF),
		string(SupportStatusDefaultON),
		string(SupportStatusInvalid),
		string(SupportStatusNotSupported),
		string(SupportStatusSupported),
	}
}

func parseSupportStatus(input string) (*SupportStatus, error) {
	vals := map[string]SupportStatus{
		"defaultoff":   SupportStatusDefaultOFF,
		"defaulton":    SupportStatusDefaultON,
		"invalid":      SupportStatusInvalid,
		"notsupported": SupportStatusNotSupported,
		"supported":    SupportStatusSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportStatus(input)
	return &out, nil
}
