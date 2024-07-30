package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSoftwareUpdateStatusSoftwareFreshness string

const (
	TeamworkSoftwareUpdateStatusSoftwareFreshness_Latest          TeamworkSoftwareUpdateStatusSoftwareFreshness = "latest"
	TeamworkSoftwareUpdateStatusSoftwareFreshness_Unknown         TeamworkSoftwareUpdateStatusSoftwareFreshness = "unknown"
	TeamworkSoftwareUpdateStatusSoftwareFreshness_UpdateAvailable TeamworkSoftwareUpdateStatusSoftwareFreshness = "updateAvailable"
)

func PossibleValuesForTeamworkSoftwareUpdateStatusSoftwareFreshness() []string {
	return []string{
		string(TeamworkSoftwareUpdateStatusSoftwareFreshness_Latest),
		string(TeamworkSoftwareUpdateStatusSoftwareFreshness_Unknown),
		string(TeamworkSoftwareUpdateStatusSoftwareFreshness_UpdateAvailable),
	}
}

func (s *TeamworkSoftwareUpdateStatusSoftwareFreshness) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkSoftwareUpdateStatusSoftwareFreshness(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkSoftwareUpdateStatusSoftwareFreshness(input string) (*TeamworkSoftwareUpdateStatusSoftwareFreshness, error) {
	vals := map[string]TeamworkSoftwareUpdateStatusSoftwareFreshness{
		"latest":          TeamworkSoftwareUpdateStatusSoftwareFreshness_Latest,
		"unknown":         TeamworkSoftwareUpdateStatusSoftwareFreshness_Unknown,
		"updateavailable": TeamworkSoftwareUpdateStatusSoftwareFreshness_UpdateAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkSoftwareUpdateStatusSoftwareFreshness(input)
	return &out, nil
}
