package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseIndexOperationAction string

const (
	EdiscoveryCaseIndexOperationActionaddToReviewSet     EdiscoveryCaseIndexOperationAction = "AddToReviewSet"
	EdiscoveryCaseIndexOperationActionapplyTags          EdiscoveryCaseIndexOperationAction = "ApplyTags"
	EdiscoveryCaseIndexOperationActioncontentExport      EdiscoveryCaseIndexOperationAction = "ContentExport"
	EdiscoveryCaseIndexOperationActionconvertToPdf       EdiscoveryCaseIndexOperationAction = "ConvertToPdf"
	EdiscoveryCaseIndexOperationActionestimateStatistics EdiscoveryCaseIndexOperationAction = "EstimateStatistics"
	EdiscoveryCaseIndexOperationActionholdUpdate         EdiscoveryCaseIndexOperationAction = "HoldUpdate"
	EdiscoveryCaseIndexOperationActionindex              EdiscoveryCaseIndexOperationAction = "Index"
	EdiscoveryCaseIndexOperationActionpurgeData          EdiscoveryCaseIndexOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryCaseIndexOperationAction() []string {
	return []string{
		string(EdiscoveryCaseIndexOperationActionaddToReviewSet),
		string(EdiscoveryCaseIndexOperationActionapplyTags),
		string(EdiscoveryCaseIndexOperationActioncontentExport),
		string(EdiscoveryCaseIndexOperationActionconvertToPdf),
		string(EdiscoveryCaseIndexOperationActionestimateStatistics),
		string(EdiscoveryCaseIndexOperationActionholdUpdate),
		string(EdiscoveryCaseIndexOperationActionindex),
		string(EdiscoveryCaseIndexOperationActionpurgeData),
	}
}

func parseEdiscoveryCaseIndexOperationAction(input string) (*EdiscoveryCaseIndexOperationAction, error) {
	vals := map[string]EdiscoveryCaseIndexOperationAction{
		"addtoreviewset":     EdiscoveryCaseIndexOperationActionaddToReviewSet,
		"applytags":          EdiscoveryCaseIndexOperationActionapplyTags,
		"contentexport":      EdiscoveryCaseIndexOperationActioncontentExport,
		"converttopdf":       EdiscoveryCaseIndexOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryCaseIndexOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryCaseIndexOperationActionholdUpdate,
		"index":              EdiscoveryCaseIndexOperationActionindex,
		"purgedata":          EdiscoveryCaseIndexOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseIndexOperationAction(input)
	return &out, nil
}
