package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryTagOperationAction string

const (
	SecurityEdiscoveryTagOperationAction_AddToReviewSet     SecurityEdiscoveryTagOperationAction = "addToReviewSet"
	SecurityEdiscoveryTagOperationAction_ApplyTags          SecurityEdiscoveryTagOperationAction = "applyTags"
	SecurityEdiscoveryTagOperationAction_ContentExport      SecurityEdiscoveryTagOperationAction = "contentExport"
	SecurityEdiscoveryTagOperationAction_ConvertToPdf       SecurityEdiscoveryTagOperationAction = "convertToPdf"
	SecurityEdiscoveryTagOperationAction_EstimateStatistics SecurityEdiscoveryTagOperationAction = "estimateStatistics"
	SecurityEdiscoveryTagOperationAction_ExportReport       SecurityEdiscoveryTagOperationAction = "exportReport"
	SecurityEdiscoveryTagOperationAction_ExportResult       SecurityEdiscoveryTagOperationAction = "exportResult"
	SecurityEdiscoveryTagOperationAction_HoldUpdate         SecurityEdiscoveryTagOperationAction = "holdUpdate"
	SecurityEdiscoveryTagOperationAction_Index              SecurityEdiscoveryTagOperationAction = "index"
	SecurityEdiscoveryTagOperationAction_PurgeData          SecurityEdiscoveryTagOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryTagOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryTagOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryTagOperationAction_ApplyTags),
		string(SecurityEdiscoveryTagOperationAction_ContentExport),
		string(SecurityEdiscoveryTagOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryTagOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryTagOperationAction_ExportReport),
		string(SecurityEdiscoveryTagOperationAction_ExportResult),
		string(SecurityEdiscoveryTagOperationAction_HoldUpdate),
		string(SecurityEdiscoveryTagOperationAction_Index),
		string(SecurityEdiscoveryTagOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryTagOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryTagOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryTagOperationAction(input string) (*SecurityEdiscoveryTagOperationAction, error) {
	vals := map[string]SecurityEdiscoveryTagOperationAction{
		"addtoreviewset":     SecurityEdiscoveryTagOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryTagOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryTagOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryTagOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryTagOperationAction_EstimateStatistics,
		"exportreport":       SecurityEdiscoveryTagOperationAction_ExportReport,
		"exportresult":       SecurityEdiscoveryTagOperationAction_ExportResult,
		"holdupdate":         SecurityEdiscoveryTagOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryTagOperationAction_Index,
		"purgedata":          SecurityEdiscoveryTagOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryTagOperationAction(input)
	return &out, nil
}
