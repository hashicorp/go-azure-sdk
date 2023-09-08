package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextDetailsDisplayLinkType string

const (
	PlannerPlanContextDetailsDisplayLinkTypeloopPage       PlannerPlanContextDetailsDisplayLinkType = "LoopPage"
	PlannerPlanContextDetailsDisplayLinkTypemeetingNotes   PlannerPlanContextDetailsDisplayLinkType = "MeetingNotes"
	PlannerPlanContextDetailsDisplayLinkTypeother          PlannerPlanContextDetailsDisplayLinkType = "Other"
	PlannerPlanContextDetailsDisplayLinkTypeproject        PlannerPlanContextDetailsDisplayLinkType = "Project"
	PlannerPlanContextDetailsDisplayLinkTypesharePointPage PlannerPlanContextDetailsDisplayLinkType = "SharePointPage"
	PlannerPlanContextDetailsDisplayLinkTypeteamsTab       PlannerPlanContextDetailsDisplayLinkType = "TeamsTab"
)

func PossibleValuesForPlannerPlanContextDetailsDisplayLinkType() []string {
	return []string{
		string(PlannerPlanContextDetailsDisplayLinkTypeloopPage),
		string(PlannerPlanContextDetailsDisplayLinkTypemeetingNotes),
		string(PlannerPlanContextDetailsDisplayLinkTypeother),
		string(PlannerPlanContextDetailsDisplayLinkTypeproject),
		string(PlannerPlanContextDetailsDisplayLinkTypesharePointPage),
		string(PlannerPlanContextDetailsDisplayLinkTypeteamsTab),
	}
}

func parsePlannerPlanContextDetailsDisplayLinkType(input string) (*PlannerPlanContextDetailsDisplayLinkType, error) {
	vals := map[string]PlannerPlanContextDetailsDisplayLinkType{
		"looppage":       PlannerPlanContextDetailsDisplayLinkTypeloopPage,
		"meetingnotes":   PlannerPlanContextDetailsDisplayLinkTypemeetingNotes,
		"other":          PlannerPlanContextDetailsDisplayLinkTypeother,
		"project":        PlannerPlanContextDetailsDisplayLinkTypeproject,
		"sharepointpage": PlannerPlanContextDetailsDisplayLinkTypesharePointPage,
		"teamstab":       PlannerPlanContextDetailsDisplayLinkTypeteamsTab,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContextDetailsDisplayLinkType(input)
	return &out, nil
}
