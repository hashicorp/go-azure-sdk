package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseOperationAction string

const (
	SecurityCaseOperationAction_AddToReviewSet     SecurityCaseOperationAction = "addToReviewSet"
	SecurityCaseOperationAction_ApplyTags          SecurityCaseOperationAction = "applyTags"
	SecurityCaseOperationAction_ContentExport      SecurityCaseOperationAction = "contentExport"
	SecurityCaseOperationAction_ConvertToPdf       SecurityCaseOperationAction = "convertToPdf"
	SecurityCaseOperationAction_EstimateStatistics SecurityCaseOperationAction = "estimateStatistics"
	SecurityCaseOperationAction_HoldUpdate         SecurityCaseOperationAction = "holdUpdate"
	SecurityCaseOperationAction_Index              SecurityCaseOperationAction = "index"
	SecurityCaseOperationAction_PurgeData          SecurityCaseOperationAction = "purgeData"
)

func PossibleValuesForSecurityCaseOperationAction() []string {
	return []string{
		string(SecurityCaseOperationAction_AddToReviewSet),
		string(SecurityCaseOperationAction_ApplyTags),
		string(SecurityCaseOperationAction_ContentExport),
		string(SecurityCaseOperationAction_ConvertToPdf),
		string(SecurityCaseOperationAction_EstimateStatistics),
		string(SecurityCaseOperationAction_HoldUpdate),
		string(SecurityCaseOperationAction_Index),
		string(SecurityCaseOperationAction_PurgeData),
	}
}

func (s *SecurityCaseOperationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCaseOperationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCaseOperationAction(input string) (*SecurityCaseOperationAction, error) {
	vals := map[string]SecurityCaseOperationAction{
		"addtoreviewset":     SecurityCaseOperationAction_AddToReviewSet,
		"applytags":          SecurityCaseOperationAction_ApplyTags,
		"contentexport":      SecurityCaseOperationAction_ContentExport,
		"converttopdf":       SecurityCaseOperationAction_ConvertToPdf,
		"estimatestatistics": SecurityCaseOperationAction_EstimateStatistics,
		"holdupdate":         SecurityCaseOperationAction_HoldUpdate,
		"index":              SecurityCaseOperationAction_Index,
		"purgedata":          SecurityCaseOperationAction_PurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseOperationAction(input)
	return &out, nil
}
