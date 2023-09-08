package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryEstimateOperationAction string

const (
	SecurityEdiscoveryEstimateOperationActionaddToReviewSet     SecurityEdiscoveryEstimateOperationAction = "AddToReviewSet"
	SecurityEdiscoveryEstimateOperationActionapplyTags          SecurityEdiscoveryEstimateOperationAction = "ApplyTags"
	SecurityEdiscoveryEstimateOperationActioncontentExport      SecurityEdiscoveryEstimateOperationAction = "ContentExport"
	SecurityEdiscoveryEstimateOperationActionconvertToPdf       SecurityEdiscoveryEstimateOperationAction = "ConvertToPdf"
	SecurityEdiscoveryEstimateOperationActionestimateStatistics SecurityEdiscoveryEstimateOperationAction = "EstimateStatistics"
	SecurityEdiscoveryEstimateOperationActionholdUpdate         SecurityEdiscoveryEstimateOperationAction = "HoldUpdate"
	SecurityEdiscoveryEstimateOperationActionindex              SecurityEdiscoveryEstimateOperationAction = "Index"
	SecurityEdiscoveryEstimateOperationActionpurgeData          SecurityEdiscoveryEstimateOperationAction = "PurgeData"
)

func PossibleValuesForSecurityEdiscoveryEstimateOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryEstimateOperationActionaddToReviewSet),
		string(SecurityEdiscoveryEstimateOperationActionapplyTags),
		string(SecurityEdiscoveryEstimateOperationActioncontentExport),
		string(SecurityEdiscoveryEstimateOperationActionconvertToPdf),
		string(SecurityEdiscoveryEstimateOperationActionestimateStatistics),
		string(SecurityEdiscoveryEstimateOperationActionholdUpdate),
		string(SecurityEdiscoveryEstimateOperationActionindex),
		string(SecurityEdiscoveryEstimateOperationActionpurgeData),
	}
}

func parseSecurityEdiscoveryEstimateOperationAction(input string) (*SecurityEdiscoveryEstimateOperationAction, error) {
	vals := map[string]SecurityEdiscoveryEstimateOperationAction{
		"addtoreviewset":     SecurityEdiscoveryEstimateOperationActionaddToReviewSet,
		"applytags":          SecurityEdiscoveryEstimateOperationActionapplyTags,
		"contentexport":      SecurityEdiscoveryEstimateOperationActioncontentExport,
		"converttopdf":       SecurityEdiscoveryEstimateOperationActionconvertToPdf,
		"estimatestatistics": SecurityEdiscoveryEstimateOperationActionestimateStatistics,
		"holdupdate":         SecurityEdiscoveryEstimateOperationActionholdUpdate,
		"index":              SecurityEdiscoveryEstimateOperationActionindex,
		"purgedata":          SecurityEdiscoveryEstimateOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryEstimateOperationAction(input)
	return &out, nil
}
