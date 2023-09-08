package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseOperationAction string

const (
	SecurityCaseOperationActionaddToReviewSet     SecurityCaseOperationAction = "AddToReviewSet"
	SecurityCaseOperationActionapplyTags          SecurityCaseOperationAction = "ApplyTags"
	SecurityCaseOperationActioncontentExport      SecurityCaseOperationAction = "ContentExport"
	SecurityCaseOperationActionconvertToPdf       SecurityCaseOperationAction = "ConvertToPdf"
	SecurityCaseOperationActionestimateStatistics SecurityCaseOperationAction = "EstimateStatistics"
	SecurityCaseOperationActionholdUpdate         SecurityCaseOperationAction = "HoldUpdate"
	SecurityCaseOperationActionindex              SecurityCaseOperationAction = "Index"
	SecurityCaseOperationActionpurgeData          SecurityCaseOperationAction = "PurgeData"
)

func PossibleValuesForSecurityCaseOperationAction() []string {
	return []string{
		string(SecurityCaseOperationActionaddToReviewSet),
		string(SecurityCaseOperationActionapplyTags),
		string(SecurityCaseOperationActioncontentExport),
		string(SecurityCaseOperationActionconvertToPdf),
		string(SecurityCaseOperationActionestimateStatistics),
		string(SecurityCaseOperationActionholdUpdate),
		string(SecurityCaseOperationActionindex),
		string(SecurityCaseOperationActionpurgeData),
	}
}

func parseSecurityCaseOperationAction(input string) (*SecurityCaseOperationAction, error) {
	vals := map[string]SecurityCaseOperationAction{
		"addtoreviewset":     SecurityCaseOperationActionaddToReviewSet,
		"applytags":          SecurityCaseOperationActionapplyTags,
		"contentexport":      SecurityCaseOperationActioncontentExport,
		"converttopdf":       SecurityCaseOperationActionconvertToPdf,
		"estimatestatistics": SecurityCaseOperationActionestimateStatistics,
		"holdupdate":         SecurityCaseOperationActionholdUpdate,
		"index":              SecurityCaseOperationActionindex,
		"purgedata":          SecurityCaseOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseOperationAction(input)
	return &out, nil
}
