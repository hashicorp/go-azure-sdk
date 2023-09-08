package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeAppleVppLicensesActionResult struct {
	ActionName          *string                                        `json:"actionName,omitempty"`
	ActionState         *RevokeAppleVppLicensesActionResultActionState `json:"actionState,omitempty"`
	FailedLicensesCount *int64                                         `json:"failedLicensesCount,omitempty"`
	LastUpdatedDateTime *string                                        `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                        `json:"@odata.type,omitempty"`
	StartDateTime       *string                                        `json:"startDateTime,omitempty"`
	TotalLicensesCount  *int64                                         `json:"totalLicensesCount,omitempty"`
}
