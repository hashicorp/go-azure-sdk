package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseOperationAction string

const (
	EdiscoveryCaseOperationActionaddToReviewSet     EdiscoveryCaseOperationAction = "AddToReviewSet"
	EdiscoveryCaseOperationActionapplyTags          EdiscoveryCaseOperationAction = "ApplyTags"
	EdiscoveryCaseOperationActioncontentExport      EdiscoveryCaseOperationAction = "ContentExport"
	EdiscoveryCaseOperationActionconvertToPdf       EdiscoveryCaseOperationAction = "ConvertToPdf"
	EdiscoveryCaseOperationActionestimateStatistics EdiscoveryCaseOperationAction = "EstimateStatistics"
	EdiscoveryCaseOperationActionholdUpdate         EdiscoveryCaseOperationAction = "HoldUpdate"
	EdiscoveryCaseOperationActionindex              EdiscoveryCaseOperationAction = "Index"
	EdiscoveryCaseOperationActionpurgeData          EdiscoveryCaseOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryCaseOperationAction() []string {
	return []string{
		string(EdiscoveryCaseOperationActionaddToReviewSet),
		string(EdiscoveryCaseOperationActionapplyTags),
		string(EdiscoveryCaseOperationActioncontentExport),
		string(EdiscoveryCaseOperationActionconvertToPdf),
		string(EdiscoveryCaseOperationActionestimateStatistics),
		string(EdiscoveryCaseOperationActionholdUpdate),
		string(EdiscoveryCaseOperationActionindex),
		string(EdiscoveryCaseOperationActionpurgeData),
	}
}

func parseEdiscoveryCaseOperationAction(input string) (*EdiscoveryCaseOperationAction, error) {
	vals := map[string]EdiscoveryCaseOperationAction{
		"addtoreviewset":     EdiscoveryCaseOperationActionaddToReviewSet,
		"applytags":          EdiscoveryCaseOperationActionapplyTags,
		"contentexport":      EdiscoveryCaseOperationActioncontentExport,
		"converttopdf":       EdiscoveryCaseOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryCaseOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryCaseOperationActionholdUpdate,
		"index":              EdiscoveryCaseOperationActionindex,
		"purgedata":          EdiscoveryCaseOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseOperationAction(input)
	return &out, nil
}
