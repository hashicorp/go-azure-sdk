package myworkbooksapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategoryType string

const (
	CategoryTypePerformance CategoryType = "performance"
	CategoryTypeRetention   CategoryType = "retention"
	CategoryTypeTSG         CategoryType = "TSG"
	CategoryTypeWorkbook    CategoryType = "workbook"
)

func PossibleValuesForCategoryType() []string {
	return []string{
		string(CategoryTypePerformance),
		string(CategoryTypeRetention),
		string(CategoryTypeTSG),
		string(CategoryTypeWorkbook),
	}
}

func (s *CategoryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategoryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategoryType(input string) (*CategoryType, error) {
	vals := map[string]CategoryType{
		"performance": CategoryTypePerformance,
		"retention":   CategoryTypeRetention,
		"tsg":         CategoryTypeTSG,
		"workbook":    CategoryTypeWorkbook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CategoryType(input)
	return &out, nil
}

type SharedTypeKind string

const (
	SharedTypeKindShared SharedTypeKind = "shared"
	SharedTypeKindUser   SharedTypeKind = "user"
)

func PossibleValuesForSharedTypeKind() []string {
	return []string{
		string(SharedTypeKindShared),
		string(SharedTypeKindUser),
	}
}

func (s *SharedTypeKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedTypeKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharedTypeKind(input string) (*SharedTypeKind, error) {
	vals := map[string]SharedTypeKind{
		"shared": SharedTypeKindShared,
		"user":   SharedTypeKindUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedTypeKind(input)
	return &out, nil
}
