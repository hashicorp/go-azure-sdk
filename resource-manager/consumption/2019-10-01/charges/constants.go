package charges

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
