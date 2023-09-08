package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryPurgeDataOperationAction string

const (
	EdiscoveryPurgeDataOperationActionaddToReviewSet     EdiscoveryPurgeDataOperationAction = "AddToReviewSet"
	EdiscoveryPurgeDataOperationActionapplyTags          EdiscoveryPurgeDataOperationAction = "ApplyTags"
	EdiscoveryPurgeDataOperationActioncontentExport      EdiscoveryPurgeDataOperationAction = "ContentExport"
	EdiscoveryPurgeDataOperationActionconvertToPdf       EdiscoveryPurgeDataOperationAction = "ConvertToPdf"
	EdiscoveryPurgeDataOperationActionestimateStatistics EdiscoveryPurgeDataOperationAction = "EstimateStatistics"
	EdiscoveryPurgeDataOperationActionholdUpdate         EdiscoveryPurgeDataOperationAction = "HoldUpdate"
	EdiscoveryPurgeDataOperationActionindex              EdiscoveryPurgeDataOperationAction = "Index"
	EdiscoveryPurgeDataOperationActionpurgeData          EdiscoveryPurgeDataOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryPurgeDataOperationAction() []string {
	return []string{
		string(EdiscoveryPurgeDataOperationActionaddToReviewSet),
		string(EdiscoveryPurgeDataOperationActionapplyTags),
		string(EdiscoveryPurgeDataOperationActioncontentExport),
		string(EdiscoveryPurgeDataOperationActionconvertToPdf),
		string(EdiscoveryPurgeDataOperationActionestimateStatistics),
		string(EdiscoveryPurgeDataOperationActionholdUpdate),
		string(EdiscoveryPurgeDataOperationActionindex),
		string(EdiscoveryPurgeDataOperationActionpurgeData),
	}
}

func parseEdiscoveryPurgeDataOperationAction(input string) (*EdiscoveryPurgeDataOperationAction, error) {
	vals := map[string]EdiscoveryPurgeDataOperationAction{
		"addtoreviewset":     EdiscoveryPurgeDataOperationActionaddToReviewSet,
		"applytags":          EdiscoveryPurgeDataOperationActionapplyTags,
		"contentexport":      EdiscoveryPurgeDataOperationActioncontentExport,
		"converttopdf":       EdiscoveryPurgeDataOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryPurgeDataOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryPurgeDataOperationActionholdUpdate,
		"index":              EdiscoveryPurgeDataOperationActionindex,
		"purgedata":          EdiscoveryPurgeDataOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryPurgeDataOperationAction(input)
	return &out, nil
}
