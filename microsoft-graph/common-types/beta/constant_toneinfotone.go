package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ToneInfoTone string

const (
	ToneInfoTone_A     ToneInfoTone = "a"
	ToneInfoTone_B     ToneInfoTone = "b"
	ToneInfoTone_C     ToneInfoTone = "c"
	ToneInfoTone_D     ToneInfoTone = "d"
	ToneInfoTone_Flash ToneInfoTone = "flash"
	ToneInfoTone_Pound ToneInfoTone = "pound"
	ToneInfoTone_Star  ToneInfoTone = "star"
	ToneInfoTone_Tone0 ToneInfoTone = "tone0"
	ToneInfoTone_Tone1 ToneInfoTone = "tone1"
	ToneInfoTone_Tone2 ToneInfoTone = "tone2"
	ToneInfoTone_Tone3 ToneInfoTone = "tone3"
	ToneInfoTone_Tone4 ToneInfoTone = "tone4"
	ToneInfoTone_Tone5 ToneInfoTone = "tone5"
	ToneInfoTone_Tone6 ToneInfoTone = "tone6"
	ToneInfoTone_Tone7 ToneInfoTone = "tone7"
	ToneInfoTone_Tone8 ToneInfoTone = "tone8"
	ToneInfoTone_Tone9 ToneInfoTone = "tone9"
)

func PossibleValuesForToneInfoTone() []string {
	return []string{
		string(ToneInfoTone_A),
		string(ToneInfoTone_B),
		string(ToneInfoTone_C),
		string(ToneInfoTone_D),
		string(ToneInfoTone_Flash),
		string(ToneInfoTone_Pound),
		string(ToneInfoTone_Star),
		string(ToneInfoTone_Tone0),
		string(ToneInfoTone_Tone1),
		string(ToneInfoTone_Tone2),
		string(ToneInfoTone_Tone3),
		string(ToneInfoTone_Tone4),
		string(ToneInfoTone_Tone5),
		string(ToneInfoTone_Tone6),
		string(ToneInfoTone_Tone7),
		string(ToneInfoTone_Tone8),
		string(ToneInfoTone_Tone9),
	}
}

func (s *ToneInfoTone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseToneInfoTone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseToneInfoTone(input string) (*ToneInfoTone, error) {
	vals := map[string]ToneInfoTone{
		"a":     ToneInfoTone_A,
		"b":     ToneInfoTone_B,
		"c":     ToneInfoTone_C,
		"d":     ToneInfoTone_D,
		"flash": ToneInfoTone_Flash,
		"pound": ToneInfoTone_Pound,
		"star":  ToneInfoTone_Star,
		"tone0": ToneInfoTone_Tone0,
		"tone1": ToneInfoTone_Tone1,
		"tone2": ToneInfoTone_Tone2,
		"tone3": ToneInfoTone_Tone3,
		"tone4": ToneInfoTone_Tone4,
		"tone5": ToneInfoTone_Tone5,
		"tone6": ToneInfoTone_Tone6,
		"tone7": ToneInfoTone_Tone7,
		"tone8": ToneInfoTone_Tone8,
		"tone9": ToneInfoTone_Tone9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ToneInfoTone(input)
	return &out, nil
}
