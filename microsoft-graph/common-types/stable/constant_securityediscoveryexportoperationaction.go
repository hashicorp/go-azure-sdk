package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationAction string

const (
	SecurityEdiscoveryExportOperationAction_AddToReviewSet     SecurityEdiscoveryExportOperationAction = "addToReviewSet"
	SecurityEdiscoveryExportOperationAction_ApplyTags          SecurityEdiscoveryExportOperationAction = "applyTags"
	SecurityEdiscoveryExportOperationAction_ContentExport      SecurityEdiscoveryExportOperationAction = "contentExport"
	SecurityEdiscoveryExportOperationAction_ConvertToPdf       SecurityEdiscoveryExportOperationAction = "convertToPdf"
	SecurityEdiscoveryExportOperationAction_EstimateStatistics SecurityEdiscoveryExportOperationAction = "estimateStatistics"
	SecurityEdiscoveryExportOperationAction_HoldUpdate         SecurityEdiscoveryExportOperationAction = "holdUpdate"
	SecurityEdiscoveryExportOperationAction_Index              SecurityEdiscoveryExportOperationAction = "index"
	SecurityEdiscoveryExportOperationAction_PurgeData          SecurityEdiscoveryExportOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryExportOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryExportOperationAction_ApplyTags),
		string(SecurityEdiscoveryExportOperationAction_ContentExport),
		string(SecurityEdiscoveryExportOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryExportOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryExportOperationAction_HoldUpdate),
		string(SecurityEdiscoveryExportOperationAction_Index),
		string(SecurityEdiscoveryExportOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryExportOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryExportOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryExportOperationAction(input string) (*SecurityEdiscoveryExportOperationAction, error) {
	vals := map[string]SecurityEdiscoveryExportOperationAction{
		"addtoreviewset":     SecurityEdiscoveryExportOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryExportOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryExportOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryExportOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryExportOperationAction_EstimateStatistics,
		"holdupdate":         SecurityEdiscoveryExportOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryExportOperationAction_Index,
		"purgedata":          SecurityEdiscoveryExportOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationAction(input)
	return &out, nil
}
