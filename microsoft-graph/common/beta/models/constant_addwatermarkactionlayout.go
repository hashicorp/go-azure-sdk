package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddWatermarkActionLayout string

const (
	AddWatermarkActionLayoutdiagonal   AddWatermarkActionLayout = "Diagonal"
	AddWatermarkActionLayouthorizontal AddWatermarkActionLayout = "Horizontal"
)

func PossibleValuesForAddWatermarkActionLayout() []string {
	return []string{
		string(AddWatermarkActionLayoutdiagonal),
		string(AddWatermarkActionLayouthorizontal),
	}
}

func parseAddWatermarkActionLayout(input string) (*AddWatermarkActionLayout, error) {
	vals := map[string]AddWatermarkActionLayout{
		"diagonal":   AddWatermarkActionLayoutdiagonal,
		"horizontal": AddWatermarkActionLayouthorizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddWatermarkActionLayout(input)
	return &out, nil
}
