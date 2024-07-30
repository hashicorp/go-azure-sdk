package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationAction string

const (
	EdiscoveryCaseExportOperationAction_AddToReviewSet     EdiscoveryCaseExportOperationAction = "addToReviewSet"
	EdiscoveryCaseExportOperationAction_ApplyTags          EdiscoveryCaseExportOperationAction = "applyTags"
	EdiscoveryCaseExportOperationAction_ContentExport      EdiscoveryCaseExportOperationAction = "contentExport"
	EdiscoveryCaseExportOperationAction_ConvertToPdf       EdiscoveryCaseExportOperationAction = "convertToPdf"
	EdiscoveryCaseExportOperationAction_EstimateStatistics EdiscoveryCaseExportOperationAction = "estimateStatistics"
	EdiscoveryCaseExportOperationAction_HoldUpdate         EdiscoveryCaseExportOperationAction = "holdUpdate"
	EdiscoveryCaseExportOperationAction_Index              EdiscoveryCaseExportOperationAction = "index"
	EdiscoveryCaseExportOperationAction_PurgeData          EdiscoveryCaseExportOperationAction = "purgeData"
)

func PossibleValuesForEdiscoveryCaseExportOperationAction() []string {
	return []string{
		string(EdiscoveryCaseExportOperationAction_AddToReviewSet),
		string(EdiscoveryCaseExportOperationAction_ApplyTags),
		string(EdiscoveryCaseExportOperationAction_ContentExport),
		string(EdiscoveryCaseExportOperationAction_ConvertToPdf),
		string(EdiscoveryCaseExportOperationAction_EstimateStatistics),
		string(EdiscoveryCaseExportOperationAction_HoldUpdate),
		string(EdiscoveryCaseExportOperationAction_Index),
		string(EdiscoveryCaseExportOperationAction_PurgeData),
	}
}

func (s *EdiscoveryCaseExportOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseExportOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseExportOperationAction(input string) (*EdiscoveryCaseExportOperationAction, error) {
	vals := map[string]EdiscoveryCaseExportOperationAction{
		"addtoreviewset":     EdiscoveryCaseExportOperationAction_AddToReviewSet,
		"applytags":          EdiscoveryCaseExportOperationAction_ApplyTags,
		"contentexport":      EdiscoveryCaseExportOperationAction_ContentExport,
		"converttopdf":       EdiscoveryCaseExportOperationAction_ConvertToPdf,
		"estimatestatistics": EdiscoveryCaseExportOperationAction_EstimateStatistics,
		"holdupdate":         EdiscoveryCaseExportOperationAction_HoldUpdate,
		"index":              EdiscoveryCaseExportOperationAction_Index,
		"purgedata":          EdiscoveryCaseExportOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationAction(input)
	return &out, nil
}
