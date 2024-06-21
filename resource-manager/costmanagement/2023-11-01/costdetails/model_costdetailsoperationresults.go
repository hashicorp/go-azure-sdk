package costdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostDetailsOperationResults struct {
	Error     *ErrorDetails          `json:"error,omitempty"`
	Id        *string                `json:"id,omitempty"`
	Manifest  *ReportManifest        `json:"manifest,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Status    *CostDetailsStatusType `json:"status,omitempty"`
	Type      *string                `json:"type,omitempty"`
	ValidTill *string                `json:"validTill,omitempty"`
}
