package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseHoldOperationAction string

const (
	EdiscoveryCaseHoldOperationActionaddToReviewSet     EdiscoveryCaseHoldOperationAction = "AddToReviewSet"
	EdiscoveryCaseHoldOperationActionapplyTags          EdiscoveryCaseHoldOperationAction = "ApplyTags"
	EdiscoveryCaseHoldOperationActioncontentExport      EdiscoveryCaseHoldOperationAction = "ContentExport"
	EdiscoveryCaseHoldOperationActionconvertToPdf       EdiscoveryCaseHoldOperationAction = "ConvertToPdf"
	EdiscoveryCaseHoldOperationActionestimateStatistics EdiscoveryCaseHoldOperationAction = "EstimateStatistics"
	EdiscoveryCaseHoldOperationActionholdUpdate         EdiscoveryCaseHoldOperationAction = "HoldUpdate"
	EdiscoveryCaseHoldOperationActionindex              EdiscoveryCaseHoldOperationAction = "Index"
	EdiscoveryCaseHoldOperationActionpurgeData          EdiscoveryCaseHoldOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryCaseHoldOperationAction() []string {
	return []string{
		string(EdiscoveryCaseHoldOperationActionaddToReviewSet),
		string(EdiscoveryCaseHoldOperationActionapplyTags),
		string(EdiscoveryCaseHoldOperationActioncontentExport),
		string(EdiscoveryCaseHoldOperationActionconvertToPdf),
		string(EdiscoveryCaseHoldOperationActionestimateStatistics),
		string(EdiscoveryCaseHoldOperationActionholdUpdate),
		string(EdiscoveryCaseHoldOperationActionindex),
		string(EdiscoveryCaseHoldOperationActionpurgeData),
	}
}

func parseEdiscoveryCaseHoldOperationAction(input string) (*EdiscoveryCaseHoldOperationAction, error) {
	vals := map[string]EdiscoveryCaseHoldOperationAction{
		"addtoreviewset":     EdiscoveryCaseHoldOperationActionaddToReviewSet,
		"applytags":          EdiscoveryCaseHoldOperationActionapplyTags,
		"contentexport":      EdiscoveryCaseHoldOperationActioncontentExport,
		"converttopdf":       EdiscoveryCaseHoldOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryCaseHoldOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryCaseHoldOperationActionholdUpdate,
		"index":              EdiscoveryCaseHoldOperationActionindex,
		"purgedata":          EdiscoveryCaseHoldOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseHoldOperationAction(input)
	return &out, nil
}
