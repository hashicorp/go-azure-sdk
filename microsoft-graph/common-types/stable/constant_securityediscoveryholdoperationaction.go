package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryHoldOperationAction string

const (
	SecurityEdiscoveryHoldOperationAction_AddToReviewSet     SecurityEdiscoveryHoldOperationAction = "addToReviewSet"
	SecurityEdiscoveryHoldOperationAction_ApplyTags          SecurityEdiscoveryHoldOperationAction = "applyTags"
	SecurityEdiscoveryHoldOperationAction_ContentExport      SecurityEdiscoveryHoldOperationAction = "contentExport"
	SecurityEdiscoveryHoldOperationAction_ConvertToPdf       SecurityEdiscoveryHoldOperationAction = "convertToPdf"
	SecurityEdiscoveryHoldOperationAction_EstimateStatistics SecurityEdiscoveryHoldOperationAction = "estimateStatistics"
	SecurityEdiscoveryHoldOperationAction_HoldUpdate         SecurityEdiscoveryHoldOperationAction = "holdUpdate"
	SecurityEdiscoveryHoldOperationAction_Index              SecurityEdiscoveryHoldOperationAction = "index"
	SecurityEdiscoveryHoldOperationAction_PurgeData          SecurityEdiscoveryHoldOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryHoldOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryHoldOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryHoldOperationAction_ApplyTags),
		string(SecurityEdiscoveryHoldOperationAction_ContentExport),
		string(SecurityEdiscoveryHoldOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryHoldOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryHoldOperationAction_HoldUpdate),
		string(SecurityEdiscoveryHoldOperationAction_Index),
		string(SecurityEdiscoveryHoldOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryHoldOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryHoldOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryHoldOperationAction(input string) (*SecurityEdiscoveryHoldOperationAction, error) {
	vals := map[string]SecurityEdiscoveryHoldOperationAction{
		"addtoreviewset":     SecurityEdiscoveryHoldOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryHoldOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryHoldOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryHoldOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryHoldOperationAction_EstimateStatistics,
		"holdupdate":         SecurityEdiscoveryHoldOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryHoldOperationAction_Index,
		"purgedata":          SecurityEdiscoveryHoldOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryHoldOperationAction(input)
	return &out, nil
}
