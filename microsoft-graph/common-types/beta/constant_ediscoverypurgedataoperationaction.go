package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryPurgeDataOperationAction string

const (
	EdiscoveryPurgeDataOperationAction_AddToReviewSet     EdiscoveryPurgeDataOperationAction = "addToReviewSet"
	EdiscoveryPurgeDataOperationAction_ApplyTags          EdiscoveryPurgeDataOperationAction = "applyTags"
	EdiscoveryPurgeDataOperationAction_ContentExport      EdiscoveryPurgeDataOperationAction = "contentExport"
	EdiscoveryPurgeDataOperationAction_ConvertToPdf       EdiscoveryPurgeDataOperationAction = "convertToPdf"
	EdiscoveryPurgeDataOperationAction_EstimateStatistics EdiscoveryPurgeDataOperationAction = "estimateStatistics"
	EdiscoveryPurgeDataOperationAction_HoldUpdate         EdiscoveryPurgeDataOperationAction = "holdUpdate"
	EdiscoveryPurgeDataOperationAction_Index              EdiscoveryPurgeDataOperationAction = "index"
	EdiscoveryPurgeDataOperationAction_PurgeData          EdiscoveryPurgeDataOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryPurgeDataOperationAction() []string {
	return []string{
		string(EdiscoveryPurgeDataOperationAction_AddToReviewSet),
		string(EdiscoveryPurgeDataOperationAction_ApplyTags),
		string(EdiscoveryPurgeDataOperationAction_ContentExport),
		string(EdiscoveryPurgeDataOperationAction_ConvertToPdf),
		string(EdiscoveryPurgeDataOperationAction_EstimateStatistics),
		string(EdiscoveryPurgeDataOperationAction_HoldUpdate),
		string(EdiscoveryPurgeDataOperationAction_Index),
		string(EdiscoveryPurgeDataOperationAction_PurgeData),
	}
}

func (s *EdiscoveryPurgeDataOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryPurgeDataOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryPurgeDataOperationAction(input string) (*EdiscoveryPurgeDataOperationAction, error) {
	vals := map[string]EdiscoveryPurgeDataOperationAction{
		"addtoreviewset":     EdiscoveryPurgeDataOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryPurgeDataOperationAction_ApplyTags,
		"contentexport":      EdiscoveryPurgeDataOperationAction_ContentExport,
		"converttopdf":       EdiscoveryPurgeDataOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryPurgeDataOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryPurgeDataOperationAction_HoldUpdate,
		"index":              EdiscoveryPurgeDataOperationAction_Index,
		"purgedata":          EdiscoveryPurgeDataOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryPurgeDataOperationAction(input)
	return &out, nil
}
