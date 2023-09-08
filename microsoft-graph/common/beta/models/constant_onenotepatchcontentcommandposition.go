package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchContentCommandPosition string

const (
	OnenotePatchContentCommandPositionAfter  OnenotePatchContentCommandPosition = "After"
	OnenotePatchContentCommandPositionBefore OnenotePatchContentCommandPosition = "Before"
)

func PossibleValuesForOnenotePatchContentCommandPosition() []string {
	return []string{
		string(OnenotePatchContentCommandPositionAfter),
		string(OnenotePatchContentCommandPositionBefore),
	}
}

func parseOnenotePatchContentCommandPosition(input string) (*OnenotePatchContentCommandPosition, error) {
	vals := map[string]OnenotePatchContentCommandPosition{
		"after":  OnenotePatchContentCommandPositionAfter,
		"before": OnenotePatchContentCommandPositionBefore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenotePatchContentCommandPosition(input)
	return &out, nil
}
