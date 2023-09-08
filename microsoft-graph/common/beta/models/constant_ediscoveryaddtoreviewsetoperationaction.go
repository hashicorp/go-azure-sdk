package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryAddToReviewSetOperationAction string

const (
	EdiscoveryAddToReviewSetOperationActionaddToReviewSet     EdiscoveryAddToReviewSetOperationAction = "AddToReviewSet"
	EdiscoveryAddToReviewSetOperationActionapplyTags          EdiscoveryAddToReviewSetOperationAction = "ApplyTags"
	EdiscoveryAddToReviewSetOperationActioncontentExport      EdiscoveryAddToReviewSetOperationAction = "ContentExport"
	EdiscoveryAddToReviewSetOperationActionconvertToPdf       EdiscoveryAddToReviewSetOperationAction = "ConvertToPdf"
	EdiscoveryAddToReviewSetOperationActionestimateStatistics EdiscoveryAddToReviewSetOperationAction = "EstimateStatistics"
	EdiscoveryAddToReviewSetOperationActionholdUpdate         EdiscoveryAddToReviewSetOperationAction = "HoldUpdate"
	EdiscoveryAddToReviewSetOperationActionindex              EdiscoveryAddToReviewSetOperationAction = "Index"
	EdiscoveryAddToReviewSetOperationActionpurgeData          EdiscoveryAddToReviewSetOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryAddToReviewSetOperationAction() []string {
	return []string{
		string(EdiscoveryAddToReviewSetOperationActionaddToReviewSet),
		string(EdiscoveryAddToReviewSetOperationActionapplyTags),
		string(EdiscoveryAddToReviewSetOperationActioncontentExport),
		string(EdiscoveryAddToReviewSetOperationActionconvertToPdf),
		string(EdiscoveryAddToReviewSetOperationActionestimateStatistics),
		string(EdiscoveryAddToReviewSetOperationActionholdUpdate),
		string(EdiscoveryAddToReviewSetOperationActionindex),
		string(EdiscoveryAddToReviewSetOperationActionpurgeData),
	}
}

func parseEdiscoveryAddToReviewSetOperationAction(input string) (*EdiscoveryAddToReviewSetOperationAction, error) {
	vals := map[string]EdiscoveryAddToReviewSetOperationAction{
		"addtoreviewset":     EdiscoveryAddToReviewSetOperationActionaddToReviewSet,
		"applytags":          EdiscoveryAddToReviewSetOperationActionapplyTags,
		"contentexport":      EdiscoveryAddToReviewSetOperationActioncontentExport,
		"converttopdf":       EdiscoveryAddToReviewSetOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryAddToReviewSetOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryAddToReviewSetOperationActionholdUpdate,
		"index":              EdiscoveryAddToReviewSetOperationActionindex,
		"purgedata":          EdiscoveryAddToReviewSetOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryAddToReviewSetOperationAction(input)
	return &out, nil
}
