package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpirationPatternType string

const (
	ExpirationPatternTypeafterDateTime ExpirationPatternType = "AfterDateTime"
	ExpirationPatternTypeafterDuration ExpirationPatternType = "AfterDuration"
	ExpirationPatternTypenoExpiration  ExpirationPatternType = "NoExpiration"
	ExpirationPatternTypenotSpecified  ExpirationPatternType = "NotSpecified"
)

func PossibleValuesForExpirationPatternType() []string {
	return []string{
		string(ExpirationPatternTypeafterDateTime),
		string(ExpirationPatternTypeafterDuration),
		string(ExpirationPatternTypenoExpiration),
		string(ExpirationPatternTypenotSpecified),
	}
}

func parseExpirationPatternType(input string) (*ExpirationPatternType, error) {
	vals := map[string]ExpirationPatternType{
		"afterdatetime": ExpirationPatternTypeafterDateTime,
		"afterduration": ExpirationPatternTypeafterDuration,
		"noexpiration":  ExpirationPatternTypenoExpiration,
		"notspecified":  ExpirationPatternTypenotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpirationPatternType(input)
	return &out, nil
}
