package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryEstimateStatisticsOperationAction string

const (
	EdiscoveryEstimateStatisticsOperationActionaddToReviewSet     EdiscoveryEstimateStatisticsOperationAction = "AddToReviewSet"
	EdiscoveryEstimateStatisticsOperationActionapplyTags          EdiscoveryEstimateStatisticsOperationAction = "ApplyTags"
	EdiscoveryEstimateStatisticsOperationActioncontentExport      EdiscoveryEstimateStatisticsOperationAction = "ContentExport"
	EdiscoveryEstimateStatisticsOperationActionconvertToPdf       EdiscoveryEstimateStatisticsOperationAction = "ConvertToPdf"
	EdiscoveryEstimateStatisticsOperationActionestimateStatistics EdiscoveryEstimateStatisticsOperationAction = "EstimateStatistics"
	EdiscoveryEstimateStatisticsOperationActionholdUpdate         EdiscoveryEstimateStatisticsOperationAction = "HoldUpdate"
	EdiscoveryEstimateStatisticsOperationActionindex              EdiscoveryEstimateStatisticsOperationAction = "Index"
	EdiscoveryEstimateStatisticsOperationActionpurgeData          EdiscoveryEstimateStatisticsOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryEstimateStatisticsOperationAction() []string {
	return []string{
		string(EdiscoveryEstimateStatisticsOperationActionaddToReviewSet),
		string(EdiscoveryEstimateStatisticsOperationActionapplyTags),
		string(EdiscoveryEstimateStatisticsOperationActioncontentExport),
		string(EdiscoveryEstimateStatisticsOperationActionconvertToPdf),
		string(EdiscoveryEstimateStatisticsOperationActionestimateStatistics),
		string(EdiscoveryEstimateStatisticsOperationActionholdUpdate),
		string(EdiscoveryEstimateStatisticsOperationActionindex),
		string(EdiscoveryEstimateStatisticsOperationActionpurgeData),
	}
}

func parseEdiscoveryEstimateStatisticsOperationAction(input string) (*EdiscoveryEstimateStatisticsOperationAction, error) {
	vals := map[string]EdiscoveryEstimateStatisticsOperationAction{
		"addtoreviewset":     EdiscoveryEstimateStatisticsOperationActionaddToReviewSet,
		"applytags":          EdiscoveryEstimateStatisticsOperationActionapplyTags,
		"contentexport":      EdiscoveryEstimateStatisticsOperationActioncontentExport,
		"converttopdf":       EdiscoveryEstimateStatisticsOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryEstimateStatisticsOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryEstimateStatisticsOperationActionholdUpdate,
		"index":              EdiscoveryEstimateStatisticsOperationActionindex,
		"purgedata":          EdiscoveryEstimateStatisticsOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryEstimateStatisticsOperationAction(input)
	return &out, nil
}
