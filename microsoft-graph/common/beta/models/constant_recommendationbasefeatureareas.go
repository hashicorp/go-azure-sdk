package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseFeatureAreas string

const (
	RecommendationBaseFeatureAreasaccessReviews     RecommendationBaseFeatureAreas = "AccessReviews"
	RecommendationBaseFeatureAreasapplications      RecommendationBaseFeatureAreas = "Applications"
	RecommendationBaseFeatureAreasconditionalAccess RecommendationBaseFeatureAreas = "ConditionalAccess"
	RecommendationBaseFeatureAreasdevices           RecommendationBaseFeatureAreas = "Devices"
	RecommendationBaseFeatureAreasgovernance        RecommendationBaseFeatureAreas = "Governance"
	RecommendationBaseFeatureAreasgroups            RecommendationBaseFeatureAreas = "Groups"
	RecommendationBaseFeatureAreasusers             RecommendationBaseFeatureAreas = "Users"
)

func PossibleValuesForRecommendationBaseFeatureAreas() []string {
	return []string{
		string(RecommendationBaseFeatureAreasaccessReviews),
		string(RecommendationBaseFeatureAreasapplications),
		string(RecommendationBaseFeatureAreasconditionalAccess),
		string(RecommendationBaseFeatureAreasdevices),
		string(RecommendationBaseFeatureAreasgovernance),
		string(RecommendationBaseFeatureAreasgroups),
		string(RecommendationBaseFeatureAreasusers),
	}
}

func parseRecommendationBaseFeatureAreas(input string) (*RecommendationBaseFeatureAreas, error) {
	vals := map[string]RecommendationBaseFeatureAreas{
		"accessreviews":     RecommendationBaseFeatureAreasaccessReviews,
		"applications":      RecommendationBaseFeatureAreasapplications,
		"conditionalaccess": RecommendationBaseFeatureAreasconditionalAccess,
		"devices":           RecommendationBaseFeatureAreasdevices,
		"governance":        RecommendationBaseFeatureAreasgovernance,
		"groups":            RecommendationBaseFeatureAreasgroups,
		"users":             RecommendationBaseFeatureAreasusers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseFeatureAreas(input)
	return &out, nil
}
