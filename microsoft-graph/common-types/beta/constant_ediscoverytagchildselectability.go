package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTagChildSelectability string

const (
	EdiscoveryTagChildSelectability_Many EdiscoveryTagChildSelectability = "Many"
	EdiscoveryTagChildSelectability_One  EdiscoveryTagChildSelectability = "One"
)

func PossibleValuesForEdiscoveryTagChildSelectability() []string {
	return []string{
		string(EdiscoveryTagChildSelectability_Many),
		string(EdiscoveryTagChildSelectability_One),
	}
}

func (s *EdiscoveryTagChildSelectability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryTagChildSelectability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryTagChildSelectability(input string) (*EdiscoveryTagChildSelectability, error) {
	vals := map[string]EdiscoveryTagChildSelectability{
		"many": EdiscoveryTagChildSelectability_Many,
		"one":  EdiscoveryTagChildSelectability_One,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryTagChildSelectability(input)
	return &out, nil
}
