package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryAddToReviewSetOperationAction string

const (
	SecurityEdiscoveryAddToReviewSetOperationActionaddToReviewSet     SecurityEdiscoveryAddToReviewSetOperationAction = "AddToReviewSet"
	SecurityEdiscoveryAddToReviewSetOperationActionapplyTags          SecurityEdiscoveryAddToReviewSetOperationAction = "ApplyTags"
	SecurityEdiscoveryAddToReviewSetOperationActioncontentExport      SecurityEdiscoveryAddToReviewSetOperationAction = "ContentExport"
	SecurityEdiscoveryAddToReviewSetOperationActionconvertToPdf       SecurityEdiscoveryAddToReviewSetOperationAction = "ConvertToPdf"
	SecurityEdiscoveryAddToReviewSetOperationActionestimateStatistics SecurityEdiscoveryAddToReviewSetOperationAction = "EstimateStatistics"
	SecurityEdiscoveryAddToReviewSetOperationActionholdUpdate         SecurityEdiscoveryAddToReviewSetOperationAction = "HoldUpdate"
	SecurityEdiscoveryAddToReviewSetOperationActionindex              SecurityEdiscoveryAddToReviewSetOperationAction = "Index"
	SecurityEdiscoveryAddToReviewSetOperationActionpurgeData          SecurityEdiscoveryAddToReviewSetOperationAction = "PurgeData"
)

func PossibleValuesForSecurityEdiscoveryAddToReviewSetOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryAddToReviewSetOperationActionaddToReviewSet),
		string(SecurityEdiscoveryAddToReviewSetOperationActionapplyTags),
		string(SecurityEdiscoveryAddToReviewSetOperationActioncontentExport),
		string(SecurityEdiscoveryAddToReviewSetOperationActionconvertToPdf),
		string(SecurityEdiscoveryAddToReviewSetOperationActionestimateStatistics),
		string(SecurityEdiscoveryAddToReviewSetOperationActionholdUpdate),
		string(SecurityEdiscoveryAddToReviewSetOperationActionindex),
		string(SecurityEdiscoveryAddToReviewSetOperationActionpurgeData),
	}
}

func parseSecurityEdiscoveryAddToReviewSetOperationAction(input string) (*SecurityEdiscoveryAddToReviewSetOperationAction, error) {
	vals := map[string]SecurityEdiscoveryAddToReviewSetOperationAction{
		"addtoreviewset":     SecurityEdiscoveryAddToReviewSetOperationActionaddToReviewSet,
		"applytags":          SecurityEdiscoveryAddToReviewSetOperationActionapplyTags,
		"contentexport":      SecurityEdiscoveryAddToReviewSetOperationActioncontentExport,
		"converttopdf":       SecurityEdiscoveryAddToReviewSetOperationActionconvertToPdf,
		"estimatestatistics": SecurityEdiscoveryAddToReviewSetOperationActionestimateStatistics,
		"holdupdate":         SecurityEdiscoveryAddToReviewSetOperationActionholdUpdate,
		"index":              SecurityEdiscoveryAddToReviewSetOperationActionindex,
		"purgedata":          SecurityEdiscoveryAddToReviewSetOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryAddToReviewSetOperationAction(input)
	return &out, nil
}
