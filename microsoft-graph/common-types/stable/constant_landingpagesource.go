package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LandingPageSource string

const (
	LandingPageSource_Global  LandingPageSource = "global"
	LandingPageSource_Tenant  LandingPageSource = "tenant"
	LandingPageSource_Unknown LandingPageSource = "unknown"
)

func PossibleValuesForLandingPageSource() []string {
	return []string{
		string(LandingPageSource_Global),
		string(LandingPageSource_Tenant),
		string(LandingPageSource_Unknown),
	}
}

func (s *LandingPageSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLandingPageSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLandingPageSource(input string) (*LandingPageSource, error) {
	vals := map[string]LandingPageSource{
		"global":  LandingPageSource_Global,
		"tenant":  LandingPageSource_Tenant,
		"unknown": LandingPageSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LandingPageSource(input)
	return &out, nil
}
