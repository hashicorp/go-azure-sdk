package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseIndexOperationAction string

const (
	EdiscoveryCaseIndexOperationAction_AddToReviewSet     EdiscoveryCaseIndexOperationAction = "addToReviewSet"
	EdiscoveryCaseIndexOperationAction_ApplyTags          EdiscoveryCaseIndexOperationAction = "applyTags"
	EdiscoveryCaseIndexOperationAction_ContentExport      EdiscoveryCaseIndexOperationAction = "contentExport"
	EdiscoveryCaseIndexOperationAction_ConvertToPdf       EdiscoveryCaseIndexOperationAction = "convertToPdf"
	EdiscoveryCaseIndexOperationAction_EstimateStatistics EdiscoveryCaseIndexOperationAction = "estimateStatistics"
	EdiscoveryCaseIndexOperationAction_HoldUpdate         EdiscoveryCaseIndexOperationAction = "holdUpdate"
	EdiscoveryCaseIndexOperationAction_Index              EdiscoveryCaseIndexOperationAction = "index"
	EdiscoveryCaseIndexOperationAction_PurgeData          EdiscoveryCaseIndexOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryCaseIndexOperationAction() []string {
	return []string{
		string(EdiscoveryCaseIndexOperationAction_AddToReviewSet),
		string(EdiscoveryCaseIndexOperationAction_ApplyTags),
		string(EdiscoveryCaseIndexOperationAction_ContentExport),
		string(EdiscoveryCaseIndexOperationAction_ConvertToPdf),
		string(EdiscoveryCaseIndexOperationAction_EstimateStatistics),
		string(EdiscoveryCaseIndexOperationAction_HoldUpdate),
		string(EdiscoveryCaseIndexOperationAction_Index),
		string(EdiscoveryCaseIndexOperationAction_PurgeData),
	}
}

func (s *EdiscoveryCaseIndexOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseIndexOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseIndexOperationAction(input string) (*EdiscoveryCaseIndexOperationAction, error) {
	vals := map[string]EdiscoveryCaseIndexOperationAction{
		"addtoreviewset":     EdiscoveryCaseIndexOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryCaseIndexOperationAction_ApplyTags,
		"contentexport":      EdiscoveryCaseIndexOperationAction_ContentExport,
		"converttopdf":       EdiscoveryCaseIndexOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryCaseIndexOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryCaseIndexOperationAction_HoldUpdate,
		"index":              EdiscoveryCaseIndexOperationAction_Index,
		"purgedata":          EdiscoveryCaseIndexOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseIndexOperationAction(input)
	return &out, nil
}
