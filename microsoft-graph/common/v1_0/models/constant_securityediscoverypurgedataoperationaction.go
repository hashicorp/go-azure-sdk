package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryPurgeDataOperationAction string

const (
	SecurityEdiscoveryPurgeDataOperationActionaddToReviewSet     SecurityEdiscoveryPurgeDataOperationAction = "AddToReviewSet"
	SecurityEdiscoveryPurgeDataOperationActionapplyTags          SecurityEdiscoveryPurgeDataOperationAction = "ApplyTags"
	SecurityEdiscoveryPurgeDataOperationActioncontentExport      SecurityEdiscoveryPurgeDataOperationAction = "ContentExport"
	SecurityEdiscoveryPurgeDataOperationActionconvertToPdf       SecurityEdiscoveryPurgeDataOperationAction = "ConvertToPdf"
	SecurityEdiscoveryPurgeDataOperationActionestimateStatistics SecurityEdiscoveryPurgeDataOperationAction = "EstimateStatistics"
	SecurityEdiscoveryPurgeDataOperationActionholdUpdate         SecurityEdiscoveryPurgeDataOperationAction = "HoldUpdate"
	SecurityEdiscoveryPurgeDataOperationActionindex              SecurityEdiscoveryPurgeDataOperationAction = "Index"
	SecurityEdiscoveryPurgeDataOperationActionpurgeData          SecurityEdiscoveryPurgeDataOperationAction = "PurgeData"
)

func PossibleValuesForSecurityEdiscoveryPurgeDataOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryPurgeDataOperationActionaddToReviewSet),
		string(SecurityEdiscoveryPurgeDataOperationActionapplyTags),
		string(SecurityEdiscoveryPurgeDataOperationActioncontentExport),
		string(SecurityEdiscoveryPurgeDataOperationActionconvertToPdf),
		string(SecurityEdiscoveryPurgeDataOperationActionestimateStatistics),
		string(SecurityEdiscoveryPurgeDataOperationActionholdUpdate),
		string(SecurityEdiscoveryPurgeDataOperationActionindex),
		string(SecurityEdiscoveryPurgeDataOperationActionpurgeData),
	}
}

func parseSecurityEdiscoveryPurgeDataOperationAction(input string) (*SecurityEdiscoveryPurgeDataOperationAction, error) {
	vals := map[string]SecurityEdiscoveryPurgeDataOperationAction{
		"addtoreviewset":     SecurityEdiscoveryPurgeDataOperationActionaddToReviewSet,
		"applytags":          SecurityEdiscoveryPurgeDataOperationActionapplyTags,
		"contentexport":      SecurityEdiscoveryPurgeDataOperationActioncontentExport,
		"converttopdf":       SecurityEdiscoveryPurgeDataOperationActionconvertToPdf,
		"estimatestatistics": SecurityEdiscoveryPurgeDataOperationActionestimateStatistics,
		"holdupdate":         SecurityEdiscoveryPurgeDataOperationActionholdUpdate,
		"index":              SecurityEdiscoveryPurgeDataOperationActionindex,
		"purgedata":          SecurityEdiscoveryPurgeDataOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryPurgeDataOperationAction(input)
	return &out, nil
}
