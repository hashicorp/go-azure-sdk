package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryAddToReviewSetOperationAction string

const (
	SecurityEdiscoveryAddToReviewSetOperationAction_AddToReviewSet     SecurityEdiscoveryAddToReviewSetOperationAction = "addToReviewSet"
	SecurityEdiscoveryAddToReviewSetOperationAction_ApplyTags          SecurityEdiscoveryAddToReviewSetOperationAction = "applyTags"
	SecurityEdiscoveryAddToReviewSetOperationAction_ContentExport      SecurityEdiscoveryAddToReviewSetOperationAction = "contentExport"
	SecurityEdiscoveryAddToReviewSetOperationAction_ConvertToPdf       SecurityEdiscoveryAddToReviewSetOperationAction = "convertToPdf"
	SecurityEdiscoveryAddToReviewSetOperationAction_EstimateStatistics SecurityEdiscoveryAddToReviewSetOperationAction = "estimateStatistics"
	SecurityEdiscoveryAddToReviewSetOperationAction_ExportReport       SecurityEdiscoveryAddToReviewSetOperationAction = "exportReport"
	SecurityEdiscoveryAddToReviewSetOperationAction_ExportResult       SecurityEdiscoveryAddToReviewSetOperationAction = "exportResult"
	SecurityEdiscoveryAddToReviewSetOperationAction_HoldUpdate         SecurityEdiscoveryAddToReviewSetOperationAction = "holdUpdate"
	SecurityEdiscoveryAddToReviewSetOperationAction_Index              SecurityEdiscoveryAddToReviewSetOperationAction = "index"
	SecurityEdiscoveryAddToReviewSetOperationAction_PurgeData          SecurityEdiscoveryAddToReviewSetOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryAddToReviewSetOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryAddToReviewSetOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_ApplyTags),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_ContentExport),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_ExportReport),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_ExportResult),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_HoldUpdate),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_Index),
		string(SecurityEdiscoveryAddToReviewSetOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryAddToReviewSetOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryAddToReviewSetOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryAddToReviewSetOperationAction(input string) (*SecurityEdiscoveryAddToReviewSetOperationAction, error) {
	vals := map[string]SecurityEdiscoveryAddToReviewSetOperationAction{
		"addtoreviewset":     SecurityEdiscoveryAddToReviewSetOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryAddToReviewSetOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryAddToReviewSetOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryAddToReviewSetOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryAddToReviewSetOperationAction_EstimateStatistics,
		"exportreport":       SecurityEdiscoveryAddToReviewSetOperationAction_ExportReport,
		"exportresult":       SecurityEdiscoveryAddToReviewSetOperationAction_ExportResult,
		"holdupdate":         SecurityEdiscoveryAddToReviewSetOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryAddToReviewSetOperationAction_Index,
		"purgedata":          SecurityEdiscoveryAddToReviewSetOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryAddToReviewSetOperationAction(input)
	return &out, nil
}
