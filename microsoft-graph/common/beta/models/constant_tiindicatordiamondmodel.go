package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorDiamondModel string

const (
	TiIndicatorDiamondModeladversary      TiIndicatorDiamondModel = "Adversary"
	TiIndicatorDiamondModelcapability     TiIndicatorDiamondModel = "Capability"
	TiIndicatorDiamondModelinfrastructure TiIndicatorDiamondModel = "Infrastructure"
	TiIndicatorDiamondModelunknown        TiIndicatorDiamondModel = "Unknown"
	TiIndicatorDiamondModelvictim         TiIndicatorDiamondModel = "Victim"
)

func PossibleValuesForTiIndicatorDiamondModel() []string {
	return []string{
		string(TiIndicatorDiamondModeladversary),
		string(TiIndicatorDiamondModelcapability),
		string(TiIndicatorDiamondModelinfrastructure),
		string(TiIndicatorDiamondModelunknown),
		string(TiIndicatorDiamondModelvictim),
	}
}

func parseTiIndicatorDiamondModel(input string) (*TiIndicatorDiamondModel, error) {
	vals := map[string]TiIndicatorDiamondModel{
		"adversary":      TiIndicatorDiamondModeladversary,
		"capability":     TiIndicatorDiamondModelcapability,
		"infrastructure": TiIndicatorDiamondModelinfrastructure,
		"unknown":        TiIndicatorDiamondModelunknown,
		"victim":         TiIndicatorDiamondModelvictim,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorDiamondModel(input)
	return &out, nil
}
