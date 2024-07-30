package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorTlpLevel string

const (
	TiIndicatorTlpLevel_Amber   TiIndicatorTlpLevel = "amber"
	TiIndicatorTlpLevel_Green   TiIndicatorTlpLevel = "green"
	TiIndicatorTlpLevel_Red     TiIndicatorTlpLevel = "red"
	TiIndicatorTlpLevel_Unknown TiIndicatorTlpLevel = "unknown"
	TiIndicatorTlpLevel_White   TiIndicatorTlpLevel = "white"
)

func PossibleValuesForTiIndicatorTlpLevel() []string {
	return []string{
		string(TiIndicatorTlpLevel_Amber),
		string(TiIndicatorTlpLevel_Green),
		string(TiIndicatorTlpLevel_Red),
		string(TiIndicatorTlpLevel_Unknown),
		string(TiIndicatorTlpLevel_White),
	}
}

func (s *TiIndicatorTlpLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTiIndicatorTlpLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTiIndicatorTlpLevel(input string) (*TiIndicatorTlpLevel, error) {
	vals := map[string]TiIndicatorTlpLevel{
		"amber":   TiIndicatorTlpLevel_Amber,
		"green":   TiIndicatorTlpLevel_Green,
		"red":     TiIndicatorTlpLevel_Red,
		"unknown": TiIndicatorTlpLevel_Unknown,
		"white":   TiIndicatorTlpLevel_White,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorTlpLevel(input)
	return &out, nil
}
