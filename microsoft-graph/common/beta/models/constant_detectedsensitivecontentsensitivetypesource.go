package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentSensitiveTypeSource string

const (
	DetectedSensitiveContentSensitiveTypeSourceoutOfBox DetectedSensitiveContentSensitiveTypeSource = "OutOfBox"
	DetectedSensitiveContentSensitiveTypeSourcetenant   DetectedSensitiveContentSensitiveTypeSource = "Tenant"
)

func PossibleValuesForDetectedSensitiveContentSensitiveTypeSource() []string {
	return []string{
		string(DetectedSensitiveContentSensitiveTypeSourceoutOfBox),
		string(DetectedSensitiveContentSensitiveTypeSourcetenant),
	}
}

func parseDetectedSensitiveContentSensitiveTypeSource(input string) (*DetectedSensitiveContentSensitiveTypeSource, error) {
	vals := map[string]DetectedSensitiveContentSensitiveTypeSource{
		"outofbox": DetectedSensitiveContentSensitiveTypeSourceoutOfBox,
		"tenant":   DetectedSensitiveContentSensitiveTypeSourcetenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedSensitiveContentSensitiveTypeSource(input)
	return &out, nil
}
