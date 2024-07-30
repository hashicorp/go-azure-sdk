package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorDiamondModel string

const (
	TiIndicatorDiamondModel_Adversary      TiIndicatorDiamondModel = "adversary"
	TiIndicatorDiamondModel_Capability     TiIndicatorDiamondModel = "capability"
	TiIndicatorDiamondModel_Infrastructure TiIndicatorDiamondModel = "infrastructure"
	TiIndicatorDiamondModel_Unknown        TiIndicatorDiamondModel = "unknown"
	TiIndicatorDiamondModel_Victim         TiIndicatorDiamondModel = "victim"
)

func PossibleValuesForTiIndicatorDiamondModel() []string {
	return []string{
		string(TiIndicatorDiamondModel_Adversary),
		string(TiIndicatorDiamondModel_Capability),
		string(TiIndicatorDiamondModel_Infrastructure),
		string(TiIndicatorDiamondModel_Unknown),
		string(TiIndicatorDiamondModel_Victim),
	}
}

func (s *TiIndicatorDiamondModel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTiIndicatorDiamondModel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTiIndicatorDiamondModel(input string) (*TiIndicatorDiamondModel, error) {
	vals := map[string]TiIndicatorDiamondModel{
		"adversary":      TiIndicatorDiamondModel_Adversary,
		"capability":     TiIndicatorDiamondModel_Capability,
		"infrastructure": TiIndicatorDiamondModel_Infrastructure,
		"unknown":        TiIndicatorDiamondModel_Unknown,
		"victim":         TiIndicatorDiamondModel_Victim,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorDiamondModel(input)
	return &out, nil
}
