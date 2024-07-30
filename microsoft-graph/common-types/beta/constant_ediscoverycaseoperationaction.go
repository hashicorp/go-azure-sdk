package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseOperationAction string

const (
	EdiscoveryCaseOperationAction_AddToReviewSet     EdiscoveryCaseOperationAction = "addToReviewSet"
	EdiscoveryCaseOperationAction_ApplyTags          EdiscoveryCaseOperationAction = "applyTags"
	EdiscoveryCaseOperationAction_ContentExport      EdiscoveryCaseOperationAction = "contentExport"
	EdiscoveryCaseOperationAction_ConvertToPdf       EdiscoveryCaseOperationAction = "convertToPdf"
	EdiscoveryCaseOperationAction_EstimateStatistics EdiscoveryCaseOperationAction = "estimateStatistics"
	EdiscoveryCaseOperationAction_HoldUpdate         EdiscoveryCaseOperationAction = "holdUpdate"
	EdiscoveryCaseOperationAction_Index              EdiscoveryCaseOperationAction = "index"
	EdiscoveryCaseOperationAction_PurgeData          EdiscoveryCaseOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryCaseOperationAction() []string {
	return []string{
		string(EdiscoveryCaseOperationAction_AddToReviewSet),
		string(EdiscoveryCaseOperationAction_ApplyTags),
		string(EdiscoveryCaseOperationAction_ContentExport),
		string(EdiscoveryCaseOperationAction_ConvertToPdf),
		string(EdiscoveryCaseOperationAction_EstimateStatistics),
		string(EdiscoveryCaseOperationAction_HoldUpdate),
		string(EdiscoveryCaseOperationAction_Index),
		string(EdiscoveryCaseOperationAction_PurgeData),
	}
}

func (s *EdiscoveryCaseOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseOperationAction(input string) (*EdiscoveryCaseOperationAction, error) {
	vals := map[string]EdiscoveryCaseOperationAction{
		"addtoreviewset":     EdiscoveryCaseOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryCaseOperationAction_ApplyTags,
		"contentexport":      EdiscoveryCaseOperationAction_ContentExport,
		"converttopdf":       EdiscoveryCaseOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryCaseOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryCaseOperationAction_HoldUpdate,
		"index":              EdiscoveryCaseOperationAction_Index,
		"purgedata":          EdiscoveryCaseOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseOperationAction(input)
	return &out, nil
}
