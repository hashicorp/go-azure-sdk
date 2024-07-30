package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LandingPageStatus string

const (
	LandingPageStatus_Archive LandingPageStatus = "archive"
	LandingPageStatus_Delete  LandingPageStatus = "delete"
	LandingPageStatus_Draft   LandingPageStatus = "draft"
	LandingPageStatus_Ready   LandingPageStatus = "ready"
	LandingPageStatus_Unknown LandingPageStatus = "unknown"
)

func PossibleValuesForLandingPageStatus() []string {
	return []string{
		string(LandingPageStatus_Archive),
		string(LandingPageStatus_Delete),
		string(LandingPageStatus_Draft),
		string(LandingPageStatus_Ready),
		string(LandingPageStatus_Unknown),
	}
}

func (s *LandingPageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLandingPageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLandingPageStatus(input string) (*LandingPageStatus, error) {
	vals := map[string]LandingPageStatus{
		"archive": LandingPageStatus_Archive,
		"delete":  LandingPageStatus_Delete,
		"draft":   LandingPageStatus_Draft,
		"ready":   LandingPageStatus_Ready,
		"unknown": LandingPageStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LandingPageStatus(input)
	return &out, nil
}
