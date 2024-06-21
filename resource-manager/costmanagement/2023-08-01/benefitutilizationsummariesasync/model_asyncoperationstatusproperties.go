package benefitutilizationsummariesasync

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AsyncOperationStatusProperties struct {
	ReportUrl          *BenefitUtilizationSummaryReportSchema `json:"reportUrl,omitempty"`
	SecondaryReportUrl *BenefitUtilizationSummaryReportSchema `json:"secondaryReportUrl,omitempty"`
	ValidUntil         *string                                `json:"validUntil,omitempty"`
}
