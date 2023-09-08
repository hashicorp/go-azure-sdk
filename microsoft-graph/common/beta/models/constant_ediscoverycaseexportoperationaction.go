package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseExportOperationAction string

const (
	EdiscoveryCaseExportOperationActionaddToReviewSet     EdiscoveryCaseExportOperationAction = "AddToReviewSet"
	EdiscoveryCaseExportOperationActionapplyTags          EdiscoveryCaseExportOperationAction = "ApplyTags"
	EdiscoveryCaseExportOperationActioncontentExport      EdiscoveryCaseExportOperationAction = "ContentExport"
	EdiscoveryCaseExportOperationActionconvertToPdf       EdiscoveryCaseExportOperationAction = "ConvertToPdf"
	EdiscoveryCaseExportOperationActionestimateStatistics EdiscoveryCaseExportOperationAction = "EstimateStatistics"
	EdiscoveryCaseExportOperationActionholdUpdate         EdiscoveryCaseExportOperationAction = "HoldUpdate"
	EdiscoveryCaseExportOperationActionindex              EdiscoveryCaseExportOperationAction = "Index"
	EdiscoveryCaseExportOperationActionpurgeData          EdiscoveryCaseExportOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryCaseExportOperationAction() []string {
	return []string{
		string(EdiscoveryCaseExportOperationActionaddToReviewSet),
		string(EdiscoveryCaseExportOperationActionapplyTags),
		string(EdiscoveryCaseExportOperationActioncontentExport),
		string(EdiscoveryCaseExportOperationActionconvertToPdf),
		string(EdiscoveryCaseExportOperationActionestimateStatistics),
		string(EdiscoveryCaseExportOperationActionholdUpdate),
		string(EdiscoveryCaseExportOperationActionindex),
		string(EdiscoveryCaseExportOperationActionpurgeData),
	}
}

func parseEdiscoveryCaseExportOperationAction(input string) (*EdiscoveryCaseExportOperationAction, error) {
	vals := map[string]EdiscoveryCaseExportOperationAction{
		"addtoreviewset":     EdiscoveryCaseExportOperationActionaddToReviewSet,
		"applytags":          EdiscoveryCaseExportOperationActionapplyTags,
		"contentexport":      EdiscoveryCaseExportOperationActioncontentExport,
		"converttopdf":       EdiscoveryCaseExportOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryCaseExportOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryCaseExportOperationActionholdUpdate,
		"index":              EdiscoveryCaseExportOperationActionindex,
		"purgedata":          EdiscoveryCaseExportOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseExportOperationAction(input)
	return &out, nil
}
