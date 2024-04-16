package quotainformation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LimitType string

const (
	LimitTypeLimitValue LimitType = "LimitValue"
)

func PossibleValuesForLimitType() []string {
	return []string{
		string(LimitTypeLimitValue),
	}
}

func (s *LimitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLimitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLimitType(input string) (*LimitType, error) {
	vals := map[string]LimitType{
		"limitvalue": LimitTypeLimitValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LimitType(input)
	return &out, nil
}

type QuotaLimitTypes string

const (
	QuotaLimitTypesIndependent QuotaLimitTypes = "Independent"
	QuotaLimitTypesShared      QuotaLimitTypes = "Shared"
)

func PossibleValuesForQuotaLimitTypes() []string {
	return []string{
		string(QuotaLimitTypesIndependent),
		string(QuotaLimitTypesShared),
	}
}

func (s *QuotaLimitTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQuotaLimitTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQuotaLimitTypes(input string) (*QuotaLimitTypes, error) {
	vals := map[string]QuotaLimitTypes{
		"independent": QuotaLimitTypesIndependent,
		"shared":      QuotaLimitTypesShared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QuotaLimitTypes(input)
	return &out, nil
}
