package charges

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChargeSummaryKind string

const (
	ChargeSummaryKindLegacy ChargeSummaryKind = "legacy"
	ChargeSummaryKindModern ChargeSummaryKind = "modern"
)

func PossibleValuesForChargeSummaryKind() []string {
	return []string{
		string(ChargeSummaryKindLegacy),
		string(ChargeSummaryKindModern),
	}
}

func parseChargeSummaryKind(input string) (*ChargeSummaryKind, error) {
	vals := map[string]ChargeSummaryKind{
		"legacy": ChargeSummaryKindLegacy,
		"modern": ChargeSummaryKindModern,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChargeSummaryKind(input)
	return &out, nil
}
