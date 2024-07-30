package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTagOperationAction string

const (
	EdiscoveryTagOperationAction_AddToReviewSet     EdiscoveryTagOperationAction = "addToReviewSet"
	EdiscoveryTagOperationAction_ApplyTags          EdiscoveryTagOperationAction = "applyTags"
	EdiscoveryTagOperationAction_ContentExport      EdiscoveryTagOperationAction = "contentExport"
	EdiscoveryTagOperationAction_ConvertToPdf       EdiscoveryTagOperationAction = "convertToPdf"
	EdiscoveryTagOperationAction_EstimateStatistics EdiscoveryTagOperationAction = "estimateStatistics"
	EdiscoveryTagOperationAction_HoldUpdate         EdiscoveryTagOperationAction = "holdUpdate"
	EdiscoveryTagOperationAction_Index              EdiscoveryTagOperationAction = "index"
	EdiscoveryTagOperationAction_PurgeData          EdiscoveryTagOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryTagOperationAction() []string {
	return []string{
		string(EdiscoveryTagOperationAction_AddToReviewSet),
		string(EdiscoveryTagOperationAction_ApplyTags),
		string(EdiscoveryTagOperationAction_ContentExport),
		string(EdiscoveryTagOperationAction_ConvertToPdf),
		string(EdiscoveryTagOperationAction_EstimateStatistics),
		string(EdiscoveryTagOperationAction_HoldUpdate),
		string(EdiscoveryTagOperationAction_Index),
		string(EdiscoveryTagOperationAction_PurgeData),
	}
}

func (s *EdiscoveryTagOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryTagOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryTagOperationAction(input string) (*EdiscoveryTagOperationAction, error) {
	vals := map[string]EdiscoveryTagOperationAction{
		"addtoreviewset":     EdiscoveryTagOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryTagOperationAction_ApplyTags,
		"contentexport":      EdiscoveryTagOperationAction_ContentExport,
		"converttopdf":       EdiscoveryTagOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryTagOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryTagOperationAction_HoldUpdate,
		"index":              EdiscoveryTagOperationAction_Index,
		"purgedata":          EdiscoveryTagOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryTagOperationAction(input)
	return &out, nil
}
