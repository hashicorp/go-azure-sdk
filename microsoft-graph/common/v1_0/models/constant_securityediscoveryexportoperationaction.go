package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperationAction string

const (
	SecurityEdiscoveryExportOperationActionaddToReviewSet     SecurityEdiscoveryExportOperationAction = "AddToReviewSet"
	SecurityEdiscoveryExportOperationActionapplyTags          SecurityEdiscoveryExportOperationAction = "ApplyTags"
	SecurityEdiscoveryExportOperationActioncontentExport      SecurityEdiscoveryExportOperationAction = "ContentExport"
	SecurityEdiscoveryExportOperationActionconvertToPdf       SecurityEdiscoveryExportOperationAction = "ConvertToPdf"
	SecurityEdiscoveryExportOperationActionestimateStatistics SecurityEdiscoveryExportOperationAction = "EstimateStatistics"
	SecurityEdiscoveryExportOperationActionholdUpdate         SecurityEdiscoveryExportOperationAction = "HoldUpdate"
	SecurityEdiscoveryExportOperationActionindex              SecurityEdiscoveryExportOperationAction = "Index"
	SecurityEdiscoveryExportOperationActionpurgeData          SecurityEdiscoveryExportOperationAction = "PurgeData"
)

func PossibleValuesForSecurityEdiscoveryExportOperationAction() []string {
	return []string{
		string(SecurityEdiscoveryExportOperationActionaddToReviewSet),
		string(SecurityEdiscoveryExportOperationActionapplyTags),
		string(SecurityEdiscoveryExportOperationActioncontentExport),
		string(SecurityEdiscoveryExportOperationActionconvertToPdf),
		string(SecurityEdiscoveryExportOperationActionestimateStatistics),
		string(SecurityEdiscoveryExportOperationActionholdUpdate),
		string(SecurityEdiscoveryExportOperationActionindex),
		string(SecurityEdiscoveryExportOperationActionpurgeData),
	}
}

func parseSecurityEdiscoveryExportOperationAction(input string) (*SecurityEdiscoveryExportOperationAction, error) {
	vals := map[string]SecurityEdiscoveryExportOperationAction{
		"addtoreviewset":     SecurityEdiscoveryExportOperationActionaddToReviewSet,
		"applytags":          SecurityEdiscoveryExportOperationActionapplyTags,
		"contentexport":      SecurityEdiscoveryExportOperationActioncontentExport,
		"converttopdf":       SecurityEdiscoveryExportOperationActionconvertToPdf,
		"estimatestatistics": SecurityEdiscoveryExportOperationActionestimateStatistics,
		"holdupdate":         SecurityEdiscoveryExportOperationActionholdUpdate,
		"index":              SecurityEdiscoveryExportOperationActionindex,
		"purgedata":          SecurityEdiscoveryExportOperationActionpurgeData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryExportOperationAction(input)
	return &out, nil
}
