package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationFeatureAreas string

const (
	RecommendationFeatureAreasaccessReviews     RecommendationFeatureAreas = "AccessReviews"
	RecommendationFeatureAreasapplications      RecommendationFeatureAreas = "Applications"
	RecommendationFeatureAreasconditionalAccess RecommendationFeatureAreas = "ConditionalAccess"
	RecommendationFeatureAreasdevices           RecommendationFeatureAreas = "Devices"
	RecommendationFeatureAreasgovernance        RecommendationFeatureAreas = "Governance"
	RecommendationFeatureAreasgroups            RecommendationFeatureAreas = "Groups"
	RecommendationFeatureAreasusers             RecommendationFeatureAreas = "Users"
)

func PossibleValuesForRecommendationFeatureAreas() []string {
	return []string{
		string(RecommendationFeatureAreasaccessReviews),
		string(RecommendationFeatureAreasapplications),
		string(RecommendationFeatureAreasconditionalAccess),
		string(RecommendationFeatureAreasdevices),
		string(RecommendationFeatureAreasgovernance),
		string(RecommendationFeatureAreasgroups),
		string(RecommendationFeatureAreasusers),
	}
}

func parseRecommendationFeatureAreas(input string) (*RecommendationFeatureAreas, error) {
	vals := map[string]RecommendationFeatureAreas{
		"accessreviews":     RecommendationFeatureAreasaccessReviews,
		"applications":      RecommendationFeatureAreasapplications,
		"conditionalaccess": RecommendationFeatureAreasconditionalAccess,
		"devices":           RecommendationFeatureAreasdevices,
		"governance":        RecommendationFeatureAreasgovernance,
		"groups":            RecommendationFeatureAreasgroups,
		"users":             RecommendationFeatureAreasusers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationFeatureAreas(input)
	return &out, nil
}
