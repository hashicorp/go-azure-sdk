package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddWatermarkOrientation string

const (
	AddWatermarkOrientationdiagonal   AddWatermarkOrientation = "Diagonal"
	AddWatermarkOrientationhorizontal AddWatermarkOrientation = "Horizontal"
)

func PossibleValuesForAddWatermarkOrientation() []string {
	return []string{
		string(AddWatermarkOrientationdiagonal),
		string(AddWatermarkOrientationhorizontal),
	}
}

func parseAddWatermarkOrientation(input string) (*AddWatermarkOrientation, error) {
	vals := map[string]AddWatermarkOrientation{
		"diagonal":   AddWatermarkOrientationdiagonal,
		"horizontal": AddWatermarkOrientationhorizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddWatermarkOrientation(input)
	return &out, nil
}
