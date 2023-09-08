package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsRankingHintImportanceScore string

const (
	ExternalConnectorsRankingHintImportanceScorehigh     ExternalConnectorsRankingHintImportanceScore = "High"
	ExternalConnectorsRankingHintImportanceScorelow      ExternalConnectorsRankingHintImportanceScore = "Low"
	ExternalConnectorsRankingHintImportanceScoremedium   ExternalConnectorsRankingHintImportanceScore = "Medium"
	ExternalConnectorsRankingHintImportanceScoreveryHigh ExternalConnectorsRankingHintImportanceScore = "VeryHigh"
)

func PossibleValuesForExternalConnectorsRankingHintImportanceScore() []string {
	return []string{
		string(ExternalConnectorsRankingHintImportanceScorehigh),
		string(ExternalConnectorsRankingHintImportanceScorelow),
		string(ExternalConnectorsRankingHintImportanceScoremedium),
		string(ExternalConnectorsRankingHintImportanceScoreveryHigh),
	}
}

func parseExternalConnectorsRankingHintImportanceScore(input string) (*ExternalConnectorsRankingHintImportanceScore, error) {
	vals := map[string]ExternalConnectorsRankingHintImportanceScore{
		"high":     ExternalConnectorsRankingHintImportanceScorehigh,
		"low":      ExternalConnectorsRankingHintImportanceScorelow,
		"medium":   ExternalConnectorsRankingHintImportanceScoremedium,
		"veryhigh": ExternalConnectorsRankingHintImportanceScoreveryHigh,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsRankingHintImportanceScore(input)
	return &out, nil
}
