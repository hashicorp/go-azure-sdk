package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTeamsPublicationInfoCreationSourceKind string

const (
	PlannerTeamsPublicationInfoCreationSourceKindexternal    PlannerTeamsPublicationInfoCreationSourceKind = "External"
	PlannerTeamsPublicationInfoCreationSourceKindnone        PlannerTeamsPublicationInfoCreationSourceKind = "None"
	PlannerTeamsPublicationInfoCreationSourceKindpublication PlannerTeamsPublicationInfoCreationSourceKind = "Publication"
)

func PossibleValuesForPlannerTeamsPublicationInfoCreationSourceKind() []string {
	return []string{
		string(PlannerTeamsPublicationInfoCreationSourceKindexternal),
		string(PlannerTeamsPublicationInfoCreationSourceKindnone),
		string(PlannerTeamsPublicationInfoCreationSourceKindpublication),
	}
}

func parsePlannerTeamsPublicationInfoCreationSourceKind(input string) (*PlannerTeamsPublicationInfoCreationSourceKind, error) {
	vals := map[string]PlannerTeamsPublicationInfoCreationSourceKind{
		"external":    PlannerTeamsPublicationInfoCreationSourceKindexternal,
		"none":        PlannerTeamsPublicationInfoCreationSourceKindnone,
		"publication": PlannerTeamsPublicationInfoCreationSourceKindpublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTeamsPublicationInfoCreationSourceKind(input)
	return &out, nil
}
