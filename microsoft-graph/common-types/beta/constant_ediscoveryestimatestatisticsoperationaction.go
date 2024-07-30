package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryEstimateStatisticsOperationAction string

const (
	EdiscoveryEstimateStatisticsOperationAction_AddToReviewSet     EdiscoveryEstimateStatisticsOperationAction = "addToReviewSet"
	EdiscoveryEstimateStatisticsOperationAction_ApplyTags          EdiscoveryEstimateStatisticsOperationAction = "applyTags"
	EdiscoveryEstimateStatisticsOperationAction_ContentExport      EdiscoveryEstimateStatisticsOperationAction = "contentExport"
	EdiscoveryEstimateStatisticsOperationAction_ConvertToPdf       EdiscoveryEstimateStatisticsOperationAction = "convertToPdf"
	EdiscoveryEstimateStatisticsOperationAction_EstimateStatistics EdiscoveryEstimateStatisticsOperationAction = "estimateStatistics"
	EdiscoveryEstimateStatisticsOperationAction_HoldUpdate         EdiscoveryEstimateStatisticsOperationAction = "holdUpdate"
	EdiscoveryEstimateStatisticsOperationAction_Index              EdiscoveryEstimateStatisticsOperationAction = "index"
	EdiscoveryEstimateStatisticsOperationAction_PurgeData          EdiscoveryEstimateStatisticsOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryEstimateStatisticsOperationAction() []string {
	return []string{
		string(EdiscoveryEstimateStatisticsOperationAction_AddToReviewSet),
		string(EdiscoveryEstimateStatisticsOperationAction_ApplyTags),
		string(EdiscoveryEstimateStatisticsOperationAction_ContentExport),
		string(EdiscoveryEstimateStatisticsOperationAction_ConvertToPdf),
		string(EdiscoveryEstimateStatisticsOperationAction_EstimateStatistics),
		string(EdiscoveryEstimateStatisticsOperationAction_HoldUpdate),
		string(EdiscoveryEstimateStatisticsOperationAction_Index),
		string(EdiscoveryEstimateStatisticsOperationAction_PurgeData),
	}
}

func (s *EdiscoveryEstimateStatisticsOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryEstimateStatisticsOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryEstimateStatisticsOperationAction(input string) (*EdiscoveryEstimateStatisticsOperationAction, error) {
	vals := map[string]EdiscoveryEstimateStatisticsOperationAction{
		"addtoreviewset":     EdiscoveryEstimateStatisticsOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryEstimateStatisticsOperationAction_ApplyTags,
		"contentexport":      EdiscoveryEstimateStatisticsOperationAction_ContentExport,
		"converttopdf":       EdiscoveryEstimateStatisticsOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryEstimateStatisticsOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryEstimateStatisticsOperationAction_HoldUpdate,
		"index":              EdiscoveryEstimateStatisticsOperationAction_Index,
		"purgedata":          EdiscoveryEstimateStatisticsOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryEstimateStatisticsOperationAction(input)
	return &out, nil
}
