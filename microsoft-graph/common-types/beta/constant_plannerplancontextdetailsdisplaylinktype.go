package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextDetailsDisplayLinkType string

const (
	PlannerPlanContextDetailsDisplayLinkType_LoopPage       PlannerPlanContextDetailsDisplayLinkType = "loopPage"
	PlannerPlanContextDetailsDisplayLinkType_MeetingNotes   PlannerPlanContextDetailsDisplayLinkType = "meetingNotes"
	PlannerPlanContextDetailsDisplayLinkType_Other          PlannerPlanContextDetailsDisplayLinkType = "other"
	PlannerPlanContextDetailsDisplayLinkType_Project        PlannerPlanContextDetailsDisplayLinkType = "project"
	PlannerPlanContextDetailsDisplayLinkType_SharePointPage PlannerPlanContextDetailsDisplayLinkType = "sharePointPage"
	PlannerPlanContextDetailsDisplayLinkType_TeamsTab       PlannerPlanContextDetailsDisplayLinkType = "teamsTab"
)

func PossibleValuesForPlannerPlanContextDetailsDisplayLinkType() []string {
	return []string{
		string(PlannerPlanContextDetailsDisplayLinkType_LoopPage),
		string(PlannerPlanContextDetailsDisplayLinkType_MeetingNotes),
		string(PlannerPlanContextDetailsDisplayLinkType_Other),
		string(PlannerPlanContextDetailsDisplayLinkType_Project),
		string(PlannerPlanContextDetailsDisplayLinkType_SharePointPage),
		string(PlannerPlanContextDetailsDisplayLinkType_TeamsTab),
	}
}

func (s *PlannerPlanContextDetailsDisplayLinkType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanContextDetailsDisplayLinkType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanContextDetailsDisplayLinkType(input string) (*PlannerPlanContextDetailsDisplayLinkType, error) {
	vals := map[string]PlannerPlanContextDetailsDisplayLinkType{
		"looppage":       PlannerPlanContextDetailsDisplayLinkType_LoopPage,
		"meetingnotes":   PlannerPlanContextDetailsDisplayLinkType_MeetingNotes,
		"other":          PlannerPlanContextDetailsDisplayLinkType_Other,
		"project":        PlannerPlanContextDetailsDisplayLinkType_Project,
		"sharepointpage": PlannerPlanContextDetailsDisplayLinkType_SharePointPage,
		"teamstab":       PlannerPlanContextDetailsDisplayLinkType_TeamsTab,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContextDetailsDisplayLinkType(input)
	return &out, nil
}
