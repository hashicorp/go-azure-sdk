package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryPurgeDataOperationAction string

const (
	SecurityEdiscoveryPurgeDataOperationAction_AddToReviewSet     SecurityEdiscoveryPurgeDataOperationAction = "addToReviewSet"
	SecurityEdiscoveryPurgeDataOperationAction_ApplyTags          SecurityEdiscoveryPurgeDataOperationAction = "applyTags"
	SecurityEdiscoveryPurgeDataOperationAction_ContentExport      SecurityEdiscoveryPurgeDataOperationAction = "contentExport"
	SecurityEdiscoveryPurgeDataOperationAction_ConvertToPdf       SecurityEdiscoveryPurgeDataOperationAction = "convertToPdf"
	SecurityEdiscoveryPurgeDataOperationAction_EstimateStatistics SecurityEdiscoveryPurgeDataOperationAction = "estimateStatistics"
	SecurityEdiscoveryPurgeDataOperationAction_HoldUpdate         SecurityEdiscoveryPurgeDataOperationAction = "holdUpdate"
	SecurityEdiscoveryPurgeDataOperationAction_Index              SecurityEdiscoveryPurgeDataOperationAction = "index"
	SecurityEdiscoveryPurgeDataOperationAction_PurgeData          SecurityEdiscoveryPurgeDataOperationAction = "purgeData"
)

func PossibleValuesForSecurityEdiscoveryPurgeDataOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryPurgeDataOperationAction_AddToReviewSet),
		string(SecurityEdiscoveryPurgeDataOperationAction_ApplyTags),
		string(SecurityEdiscoveryPurgeDataOperationAction_ContentExport),
		string(SecurityEdiscoveryPurgeDataOperationAction_ConvertToPdf),
		string(SecurityEdiscoveryPurgeDataOperationAction_EstimateStatistics),
		string(SecurityEdiscoveryPurgeDataOperationAction_HoldUpdate),
		string(SecurityEdiscoveryPurgeDataOperationAction_Index),
		string(SecurityEdiscoveryPurgeDataOperationAction_PurgeData),
	}
}

func (s *SecurityEdiscoveryPurgeDataOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoveryPurgeDataOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoveryPurgeDataOperationAction(input string) (*SecurityEdiscoveryPurgeDataOperationAction, error) {
	vals := map[string]SecurityEdiscoveryPurgeDataOperationAction{
		"addtoreviewset":     SecurityEdiscoveryPurgeDataOperationAction_AddToReviewSet,
		"applytags":          SecurityEdiscoveryPurgeDataOperationAction_ApplyTags,
		"contentexport":      SecurityEdiscoveryPurgeDataOperationAction_ContentExport,
		"converttopdf":       SecurityEdiscoveryPurgeDataOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityEdiscoveryPurgeDataOperationAction_EstimateStatistics,
		"holdupdate":         SecurityEdiscoveryPurgeDataOperationAction_HoldUpdate,
		"index":              SecurityEdiscoveryPurgeDataOperationAction_Index,
		"purgedata":          SecurityEdiscoveryPurgeDataOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryPurgeDataOperationAction(input)
	return &out, nil
}
