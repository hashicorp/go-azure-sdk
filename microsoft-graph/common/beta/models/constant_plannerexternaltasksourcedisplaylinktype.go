package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalTaskSourceDisplayLinkType string

const (
	PlannerExternalTaskSourceDisplayLinkTypedefault PlannerExternalTaskSourceDisplayLinkType = "Default"
	PlannerExternalTaskSourceDisplayLinkTypenone    PlannerExternalTaskSourceDisplayLinkType = "None"
)

func PossibleValuesForPlannerExternalTaskSourceDisplayLinkType() []string {
	return []string{
		string(PlannerExternalTaskSourceDisplayLinkTypedefault),
		string(PlannerExternalTaskSourceDisplayLinkTypenone),
	}
}

func parsePlannerExternalTaskSourceDisplayLinkType(input string) (*PlannerExternalTaskSourceDisplayLinkType, error) {
	vals := map[string]PlannerExternalTaskSourceDisplayLinkType{
		"default": PlannerExternalTaskSourceDisplayLinkTypedefault,
		"none":    PlannerExternalTaskSourceDisplayLinkTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalTaskSourceDisplayLinkType(input)
	return &out, nil
}
