package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseHoldOperationAction string

const (
	EdiscoveryCaseHoldOperationAction_AddToReviewSet     EdiscoveryCaseHoldOperationAction = "addToReviewSet"
	EdiscoveryCaseHoldOperationAction_ApplyTags          EdiscoveryCaseHoldOperationAction = "applyTags"
	EdiscoveryCaseHoldOperationAction_ContentExport      EdiscoveryCaseHoldOperationAction = "contentExport"
	EdiscoveryCaseHoldOperationAction_ConvertToPdf       EdiscoveryCaseHoldOperationAction = "convertToPdf"
	EdiscoveryCaseHoldOperationAction_EstimateStatistics EdiscoveryCaseHoldOperationAction = "estimateStatistics"
	EdiscoveryCaseHoldOperationAction_HoldUpdate         EdiscoveryCaseHoldOperationAction = "holdUpdate"
	EdiscoveryCaseHoldOperationAction_Index              EdiscoveryCaseHoldOperationAction = "index"
	EdiscoveryCaseHoldOperationAction_PurgeData          EdiscoveryCaseHoldOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryCaseHoldOperationAction() []string {
	return []string{
		string(EdiscoveryCaseHoldOperationAction_AddToReviewSet),
		string(EdiscoveryCaseHoldOperationAction_ApplyTags),
		string(EdiscoveryCaseHoldOperationAction_ContentExport),
		string(EdiscoveryCaseHoldOperationAction_ConvertToPdf),
		string(EdiscoveryCaseHoldOperationAction_EstimateStatistics),
		string(EdiscoveryCaseHoldOperationAction_HoldUpdate),
		string(EdiscoveryCaseHoldOperationAction_Index),
		string(EdiscoveryCaseHoldOperationAction_PurgeData),
	}
}

func (s *EdiscoveryCaseHoldOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseHoldOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseHoldOperationAction(input string) (*EdiscoveryCaseHoldOperationAction, error) {
	vals := map[string]EdiscoveryCaseHoldOperationAction{
		"addtoreviewset":     EdiscoveryCaseHoldOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryCaseHoldOperationAction_ApplyTags,
		"contentexport":      EdiscoveryCaseHoldOperationAction_ContentExport,
		"converttopdf":       EdiscoveryCaseHoldOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryCaseHoldOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryCaseHoldOperationAction_HoldUpdate,
		"index":              EdiscoveryCaseHoldOperationAction_Index,
		"purgedata":          EdiscoveryCaseHoldOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseHoldOperationAction(input)
	return &out, nil
}
