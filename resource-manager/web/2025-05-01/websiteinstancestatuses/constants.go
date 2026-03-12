package websiteinstancestatuses

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRuntimeState string

const (
	SiteRuntimeStateREADY   SiteRuntimeState = "READY"
	SiteRuntimeStateSTOPPED SiteRuntimeState = "STOPPED"
	SiteRuntimeStateUNKNOWN SiteRuntimeState = "UNKNOWN"
)

func PossibleValuesForSiteRuntimeState() []string {
	return []string{
		string(SiteRuntimeStateREADY),
		string(SiteRuntimeStateSTOPPED),
		string(SiteRuntimeStateUNKNOWN),
	}
}

func (s *SiteRuntimeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteRuntimeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteRuntimeState(input string) (*SiteRuntimeState, error) {
	vals := map[string]SiteRuntimeState{
		"ready":   SiteRuntimeStateREADY,
		"stopped": SiteRuntimeStateSTOPPED,
		"unknown": SiteRuntimeStateUNKNOWN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteRuntimeState(input)
	return &out, nil
}
