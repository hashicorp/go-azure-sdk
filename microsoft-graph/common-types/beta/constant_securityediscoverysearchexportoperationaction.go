package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperationAction string

const (
	SecurityEdiscoverySearchExportOperationAction_AddToReviewSet     SecurityEdiscoverySearchExportOperationAction = "addToReviewSet"
	SecurityEdiscoverySearchExportOperationAction_ApplyTags          SecurityEdiscoverySearchExportOperationAction = "applyTags"
	SecurityEdiscoverySearchExportOperationAction_ContentExport      SecurityEdiscoverySearchExportOperationAction = "contentExport"
	SecurityEdiscoverySearchExportOperationAction_ConvertToPdf       SecurityEdiscoverySearchExportOperationAction = "convertToPdf"
	SecurityEdiscoverySearchExportOperationAction_EstimateStatistics SecurityEdiscoverySearchExportOperationAction = "estimateStatistics"
	SecurityEdiscoverySearchExportOperationAction_ExportReport       SecurityEdiscoverySearchExportOperationAction = "exportReport"
	SecurityEdiscoverySearchExportOperationAction_ExportResult       SecurityEdiscoverySearchExportOperationAction = "exportResult"
	SecurityEdiscoverySearchExportOperationAction_HoldUpdate         SecurityEdiscoverySearchExportOperationAction = "holdUpdate"
	SecurityEdiscoverySearchExportOperationAction_Index              SecurityEdiscoverySearchExportOperationAction = "index"
	SecurityEdiscoverySearchExportOperationAction_PurgeData          SecurityEdiscoverySearchExportOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoverySearchExportOperationAction() []string {
	return []string{
		string(SecurityEdiscoverySearchExportOperationAction_AddToReviewSet),
		string(SecurityEdiscoverySearchExportOperationAction_ApplyTags),
		string(SecurityEdiscoverySearchExportOperationAction_ContentExport),
		string(SecurityEdiscoverySearchExportOperationAction_ConvertToPdf),
		string(SecurityEdiscoverySearchExportOperationAction_EstimateStatistics),
		string(SecurityEdiscoverySearchExportOperationAction_ExportReport),
		string(SecurityEdiscoverySearchExportOperationAction_ExportResult),
		string(SecurityEdiscoverySearchExportOperationAction_HoldUpdate),
		string(SecurityEdiscoverySearchExportOperationAction_Index),
		string(SecurityEdiscoverySearchExportOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoverySearchExportOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchExportOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchExportOperationAction(input string) (*SecurityEdiscoverySearchExportOperationAction, error) {
	vals := map[string]SecurityEdiscoverySearchExportOperationAction{
		"addtoreviewset":     SecurityEdiscoverySearchExportOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoverySearchExportOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoverySearchExportOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoverySearchExportOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoverySearchExportOperationAction_EstimateStatistics,
		"exportreport":       SecurityEdiscoverySearchExportOperationAction_ExportReport,
		"exportresult":       SecurityEdiscoverySearchExportOperationAction_ExportResult,
		"holdupdate":         SecurityEdiscoverySearchExportOperationAction_HoldUpdate,
		"index":              SecurityEdiscoverySearchExportOperationAction_Index,
		"purgedata":          SecurityEdiscoverySearchExportOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchExportOperationAction(input)
	return &out, nil
}
