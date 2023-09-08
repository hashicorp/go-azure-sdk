package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalBucketSourceCreationSourceKind string

const (
	PlannerExternalBucketSourceCreationSourceKindexternal    PlannerExternalBucketSourceCreationSourceKind = "External"
	PlannerExternalBucketSourceCreationSourceKindnone        PlannerExternalBucketSourceCreationSourceKind = "None"
	PlannerExternalBucketSourceCreationSourceKindpublication PlannerExternalBucketSourceCreationSourceKind = "Publication"
)

func PossibleValuesForPlannerExternalBucketSourceCreationSourceKind() []string {
	return []string{
		string(PlannerExternalBucketSourceCreationSourceKindexternal),
		string(PlannerExternalBucketSourceCreationSourceKindnone),
		string(PlannerExternalBucketSourceCreationSourceKindpublication),
	}
}

func parsePlannerExternalBucketSourceCreationSourceKind(input string) (*PlannerExternalBucketSourceCreationSourceKind, error) {
	vals := map[string]PlannerExternalBucketSourceCreationSourceKind{
		"external":    PlannerExternalBucketSourceCreationSourceKindexternal,
		"none":        PlannerExternalBucketSourceCreationSourceKindnone,
		"publication": PlannerExternalBucketSourceCreationSourceKindpublication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalBucketSourceCreationSourceKind(input)
	return &out, nil
}
