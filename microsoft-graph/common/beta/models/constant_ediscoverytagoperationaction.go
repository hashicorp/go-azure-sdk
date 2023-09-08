package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTagOperationAction string

const (
	EdiscoveryTagOperationActionaddToReviewSet     EdiscoveryTagOperationAction = "AddToReviewSet"
	EdiscoveryTagOperationActionapplyTags          EdiscoveryTagOperationAction = "ApplyTags"
	EdiscoveryTagOperationActioncontentExport      EdiscoveryTagOperationAction = "ContentExport"
	EdiscoveryTagOperationActionconvertToPdf       EdiscoveryTagOperationAction = "ConvertToPdf"
	EdiscoveryTagOperationActionestimateStatistics EdiscoveryTagOperationAction = "EstimateStatistics"
	EdiscoveryTagOperationActionholdUpdate         EdiscoveryTagOperationAction = "HoldUpdate"
	EdiscoveryTagOperationActionindex              EdiscoveryTagOperationAction = "Index"
	EdiscoveryTagOperationActionpurgeData          EdiscoveryTagOperationAction = "PurgeData"
)

func PossibleValuesForEdiscoveryTagOperationAction() []string {
	return []string{
		string(EdiscoveryTagOperationActionaddToReviewSet),
		string(EdiscoveryTagOperationActionapplyTags),
		string(EdiscoveryTagOperationActioncontentExport),
		string(EdiscoveryTagOperationActionconvertToPdf),
		string(EdiscoveryTagOperationActionestimateStatistics),
		string(EdiscoveryTagOperationActionholdUpdate),
		string(EdiscoveryTagOperationActionindex),
		string(EdiscoveryTagOperationActionpurgeData),
	}
}

func parseEdiscoveryTagOperationAction(input string) (*EdiscoveryTagOperationAction, error) {
	vals := map[string]EdiscoveryTagOperationAction{
		"addtoreviewset":     EdiscoveryTagOperationActionaddToReviewSet,
		"applytags":          EdiscoveryTagOperationActionapplyTags,
		"contentexport":      EdiscoveryTagOperationActioncontentExport,
		"converttopdf":       EdiscoveryTagOperationActionconvertToPdf,
		"estimatestatistics": EdiscoveryTagOperationActionestimateStatistics,
		"holdupdate":         EdiscoveryTagOperationActionholdUpdate,
		"index":              EdiscoveryTagOperationActionindex,
		"purgedata":          EdiscoveryTagOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryTagOperationAction(input)
	return &out, nil
}
