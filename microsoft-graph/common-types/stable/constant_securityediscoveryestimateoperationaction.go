package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryEstimateOperationAction string

const (
	SecurityEdiscoveryEstimateOperationAction_AddToReviewSet     SecurityEdiscoveryEstimateOperationAction = "addToReviewSet"
	SecurityEdiscoveryEstimateOperationAction_ApplyTags          SecurityEdiscoveryEstimateOperationAction = "applyTags"
	SecurityEdiscoveryEstimateOperationAction_ContentExport      SecurityEdiscoveryEstimateOperationAction = "contentExport"
	SecurityEdiscoveryEstimateOperationAction_ConvertToPdf       SecurityEdiscoveryEstimateOperationAction = "convertToPdf"
	SecurityEdiscoveryEstimateOperationAction_EstimateStatistics SecurityEdiscoveryEstimateOperationAction = "estimateStatistics"
	SecurityEdiscoveryEstimateOperationAction_HoldUpdate         SecurityEdiscoveryEstimateOperationAction = "holdUpdate"
	SecurityEdiscoveryEstimateOperationAction_Index              SecurityEdiscoveryEstimateOperationAction = "index"
	SecurityEdiscoveryEstimateOperationAction_PurgeData          SecurityEdiscoveryEstimateOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryEstimateOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryEstimateOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryEstimateOperationAction_ApplyTags),
		string(SecurityEdiscoveryEstimateOperationAction_ContentExport),
		string(SecurityEdiscoveryEstimateOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryEstimateOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryEstimateOperationAction_HoldUpdate),
		string(SecurityEdiscoveryEstimateOperationAction_Index),
		string(SecurityEdiscoveryEstimateOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryEstimateOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryEstimateOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryEstimateOperationAction(input string) (*SecurityEdiscoveryEstimateOperationAction, error) {
	vals := map[string]SecurityEdiscoveryEstimateOperationAction{
		"addtoreviewset":     SecurityEdiscoveryEstimateOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryEstimateOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryEstimateOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryEstimateOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryEstimateOperationAction_EstimateStatistics,
		"holdupdate":         SecurityEdiscoveryEstimateOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryEstimateOperationAction_Index,
		"purgedata":          SecurityEdiscoveryEstimateOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryEstimateOperationAction(input)
	return &out, nil
}
