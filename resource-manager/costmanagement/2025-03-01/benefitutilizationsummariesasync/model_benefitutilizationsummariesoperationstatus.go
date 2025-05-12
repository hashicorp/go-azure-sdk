package benefitutilizationsummariesasync

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummariesOperationStatus struct {
	Input      *BenefitUtilizationSummariesRequest `json:"input,omitempty"`
	Properties *AsyncOperationStatusProperties     `json:"properties,omitempty"`
	Status     *OperationStatusType                `json:"status,omitempty"`
}
