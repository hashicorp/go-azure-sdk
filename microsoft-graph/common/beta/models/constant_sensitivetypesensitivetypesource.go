package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeSensitiveTypeSource string

const (
	SensitiveTypeSensitiveTypeSourceoutOfBox SensitiveTypeSensitiveTypeSource = "OutOfBox"
	SensitiveTypeSensitiveTypeSourcetenant   SensitiveTypeSensitiveTypeSource = "Tenant"
)

func PossibleValuesForSensitiveTypeSensitiveTypeSource() []string {
	return []string{
		string(SensitiveTypeSensitiveTypeSourceoutOfBox),
		string(SensitiveTypeSensitiveTypeSourcetenant),
	}
}

func parseSensitiveTypeSensitiveTypeSource(input string) (*SensitiveTypeSensitiveTypeSource, error) {
	vals := map[string]SensitiveTypeSensitiveTypeSource{
		"outofbox": SensitiveTypeSensitiveTypeSourceoutOfBox,
		"tenant":   SensitiveTypeSensitiveTypeSourcetenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeSensitiveTypeSource(input)
	return &out, nil
}
