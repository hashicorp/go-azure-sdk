package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryAddToReviewSetOperationAction string

const (
	EdiscoveryAddToReviewSetOperationAction_AddToReviewSet     EdiscoveryAddToReviewSetOperationAction = "addToReviewSet"
	EdiscoveryAddToReviewSetOperationAction_ApplyTags          EdiscoveryAddToReviewSetOperationAction = "applyTags"
	EdiscoveryAddToReviewSetOperationAction_ContentExport      EdiscoveryAddToReviewSetOperationAction = "contentExport"
	EdiscoveryAddToReviewSetOperationAction_ConvertToPdf       EdiscoveryAddToReviewSetOperationAction = "convertToPdf"
	EdiscoveryAddToReviewSetOperationAction_EstimateStatistics EdiscoveryAddToReviewSetOperationAction = "estimateStatistics"
	EdiscoveryAddToReviewSetOperationAction_HoldUpdate         EdiscoveryAddToReviewSetOperationAction = "holdUpdate"
	EdiscoveryAddToReviewSetOperationAction_Index              EdiscoveryAddToReviewSetOperationAction = "index"
	EdiscoveryAddToReviewSetOperationAction_PurgeData          EdiscoveryAddToReviewSetOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryAddToReviewSetOperationAction() []string {
	return []string{
		string(EdiscoveryAddToReviewSetOperationAction_AddToReviewSet),
		string(EdiscoveryAddToReviewSetOperationAction_ApplyTags),
		string(EdiscoveryAddToReviewSetOperationAction_ContentExport),
		string(EdiscoveryAddToReviewSetOperationAction_ConvertToPdf),
		string(EdiscoveryAddToReviewSetOperationAction_EstimateStatistics),
		string(EdiscoveryAddToReviewSetOperationAction_HoldUpdate),
		string(EdiscoveryAddToReviewSetOperationAction_Index),
		string(EdiscoveryAddToReviewSetOperationAction_PurgeData),
	}
}

func (s *EdiscoveryAddToReviewSetOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryAddToReviewSetOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryAddToReviewSetOperationAction(input string) (*EdiscoveryAddToReviewSetOperationAction, error) {
	vals := map[string]EdiscoveryAddToReviewSetOperationAction{
		"addtoreviewset":     EdiscoveryAddToReviewSetOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryAddToReviewSetOperationAction_ApplyTags,
		"contentexport":      EdiscoveryAddToReviewSetOperationAction_ContentExport,
		"converttopdf":       EdiscoveryAddToReviewSetOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryAddToReviewSetOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryAddToReviewSetOperationAction_HoldUpdate,
		"index":              EdiscoveryAddToReviewSetOperationAction_Index,
		"purgedata":          EdiscoveryAddToReviewSetOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryAddToReviewSetOperationAction(input)
	return &out, nil
}
