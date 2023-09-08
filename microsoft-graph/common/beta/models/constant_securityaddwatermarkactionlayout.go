package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddWatermarkActionLayout string

const (
	SecurityAddWatermarkActionLayoutdiagonal   SecurityAddWatermarkActionLayout = "Diagonal"
	SecurityAddWatermarkActionLayouthorizontal SecurityAddWatermarkActionLayout = "Horizontal"
)

func PossibleValuesForSecurityAddWatermarkActionLayout() []string {
	return []string{
		string(SecurityAddWatermarkActionLayoutdiagonal),
		string(SecurityAddWatermarkActionLayouthorizontal),
	}
}

func parseSecurityAddWatermarkActionLayout(input string) (*SecurityAddWatermarkActionLayout, error) {
	vals := map[string]SecurityAddWatermarkActionLayout{
		"diagonal":   SecurityAddWatermarkActionLayoutdiagonal,
		"horizontal": SecurityAddWatermarkActionLayouthorizontal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAddWatermarkActionLayout(input)
	return &out, nil
}
