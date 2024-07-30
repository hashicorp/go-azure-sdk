package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryIndexOperationAction string

const (
	SecurityEdiscoveryIndexOperationAction_AddToReviewSet     SecurityEdiscoveryIndexOperationAction = "addToReviewSet"
	SecurityEdiscoveryIndexOperationAction_ApplyTags          SecurityEdiscoveryIndexOperationAction = "applyTags"
	SecurityEdiscoveryIndexOperationAction_ContentExport      SecurityEdiscoveryIndexOperationAction = "contentExport"
	SecurityEdiscoveryIndexOperationAction_ConvertToPdf       SecurityEdiscoveryIndexOperationAction = "convertToPdf"
	SecurityEdiscoveryIndexOperationAction_EstimateStatistics SecurityEdiscoveryIndexOperationAction = "estimateStatistics"
	SecurityEdiscoveryIndexOperationAction_ExportReport       SecurityEdiscoveryIndexOperationAction = "exportReport"
	SecurityEdiscoveryIndexOperationAction_ExportResult       SecurityEdiscoveryIndexOperationAction = "exportResult"
	SecurityEdiscoveryIndexOperationAction_HoldUpdate         SecurityEdiscoveryIndexOperationAction = "holdUpdate"
	SecurityEdiscoveryIndexOperationAction_Index              SecurityEdiscoveryIndexOperationAction = "index"
	SecurityEdiscoveryIndexOperationAction_PurgeData          SecurityEdiscoveryIndexOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryIndexOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryIndexOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryIndexOperationAction_ApplyTags),
		string(SecurityEdiscoveryIndexOperationAction_ContentExport),
		string(SecurityEdiscoveryIndexOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryIndexOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryIndexOperationAction_ExportReport),
		string(SecurityEdiscoveryIndexOperationAction_ExportResult),
		string(SecurityEdiscoveryIndexOperationAction_HoldUpdate),
		string(SecurityEdiscoveryIndexOperationAction_Index),
		string(SecurityEdiscoveryIndexOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryIndexOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryIndexOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryIndexOperationAction(input string) (*SecurityEdiscoveryIndexOperationAction, error) {
	vals := map[string]SecurityEdiscoveryIndexOperationAction{
		"addtoreviewset":     SecurityEdiscoveryIndexOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryIndexOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryIndexOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryIndexOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryIndexOperationAction_EstimateStatistics,
		"exportreport":       SecurityEdiscoveryIndexOperationAction_ExportReport,
		"exportresult":       SecurityEdiscoveryIndexOperationAction_ExportResult,
		"holdupdate":         SecurityEdiscoveryIndexOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryIndexOperationAction_Index,
		"purgedata":          SecurityEdiscoveryIndexOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryIndexOperationAction(input)
	return &out, nil
}
