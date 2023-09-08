package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryHoldOperationAction string

const (
	SecurityEdiscoveryHoldOperationActionaddToReviewSet     SecurityEdiscoveryHoldOperationAction = "AddToReviewSet"
	SecurityEdiscoveryHoldOperationActionapplyTags          SecurityEdiscoveryHoldOperationAction = "ApplyTags"
	SecurityEdiscoveryHoldOperationActioncontentExport      SecurityEdiscoveryHoldOperationAction = "ContentExport"
	SecurityEdiscoveryHoldOperationActionconvertToPdf       SecurityEdiscoveryHoldOperationAction = "ConvertToPdf"
	SecurityEdiscoveryHoldOperationActionestimateStatistics SecurityEdiscoveryHoldOperationAction = "EstimateStatistics"
	SecurityEdiscoveryHoldOperationActionholdUpdate         SecurityEdiscoveryHoldOperationAction = "HoldUpdate"
	SecurityEdiscoveryHoldOperationActionindex              SecurityEdiscoveryHoldOperationAction = "Index"
	SecurityEdiscoveryHoldOperationActionpurgeData          SecurityEdiscoveryHoldOperationAction = "PurgeData"
)

func PossibleValuesForSecurityEdiscoveryHoldOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryHoldOperationActionaddToReviewSet),
		string(SecurityEdiscoveryHoldOperationActionapplyTags),
		string(SecurityEdiscoveryHoldOperationActioncontentExport),
		string(SecurityEdiscoveryHoldOperationActionconvertToPdf),
		string(SecurityEdiscoveryHoldOperationActionestimateStatistics),
		string(SecurityEdiscoveryHoldOperationActionholdUpdate),
		string(SecurityEdiscoveryHoldOperationActionindex),
		string(SecurityEdiscoveryHoldOperationActionpurgeData),
	}
}

func parseSecurityEdiscoveryHoldOperationAction(input string) (*SecurityEdiscoveryHoldOperationAction, error) {
	vals := map[string]SecurityEdiscoveryHoldOperationAction{
		"addtoreviewset":     SecurityEdiscoveryHoldOperationActionaddToReviewSet,
		"applytags":          SecurityEdiscoveryHoldOperationActionapplyTags,
		"contentexport":      SecurityEdiscoveryHoldOperationActioncontentExport,
		"converttopdf":       SecurityEdiscoveryHoldOperationActionconvertToPdf,
		"estimatestatistics": SecurityEdiscoveryHoldOperationActionestimateStatistics,
		"holdupdate":         SecurityEdiscoveryHoldOperationActionholdUpdate,
		"index":              SecurityEdiscoveryHoldOperationActionindex,
		"purgedata":          SecurityEdiscoveryHoldOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryHoldOperationAction(input)
	return &out, nil
}
