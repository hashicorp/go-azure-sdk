package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsRankingHintImportanceScore string

const (
	ExternalConnectorsRankingHintImportanceScore_High     ExternalConnectorsRankingHintImportanceScore = "high"
	ExternalConnectorsRankingHintImportanceScore_Low      ExternalConnectorsRankingHintImportanceScore = "low"
	ExternalConnectorsRankingHintImportanceScore_Medium   ExternalConnectorsRankingHintImportanceScore = "medium"
	ExternalConnectorsRankingHintImportanceScore_VeryHigh ExternalConnectorsRankingHintImportanceScore = "veryHigh"
)

func PossibleValuesForExternalConnectorsRankingHintImportanceScore() []string {
	return []string{
		string(ExternalConnectorsRankingHintImportanceScore_High),
		string(ExternalConnectorsRankingHintImportanceScore_Low),
		string(ExternalConnectorsRankingHintImportanceScore_Medium),
		string(ExternalConnectorsRankingHintImportanceScore_VeryHigh),
	}
}

func (s *ExternalConnectorsRankingHintImportanceScore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsRankingHintImportanceScore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsRankingHintImportanceScore(input string) (*ExternalConnectorsRankingHintImportanceScore, error) {
	vals := map[string]ExternalConnectorsRankingHintImportanceScore{
		"high":     ExternalConnectorsRankingHintImportanceScore_High,
		"low":      ExternalConnectorsRankingHintImportanceScore_Low,
		"medium":   ExternalConnectorsRankingHintImportanceScore_Medium,
		"veryhigh": ExternalConnectorsRankingHintImportanceScore_VeryHigh,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsRankingHintImportanceScore(input)
	return &out, nil
}
