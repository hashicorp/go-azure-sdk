package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppRevokeLicensesActionResult struct {
	ActionFailureReason *IosVppAppRevokeLicensesActionResultActionFailureReason `json:"actionFailureReason,omitempty"`
	ActionName          *string                                                 `json:"actionName,omitempty"`
	ActionState         *IosVppAppRevokeLicensesActionResultActionState         `json:"actionState,omitempty"`
	FailedLicensesCount *int64                                                  `json:"failedLicensesCount,omitempty"`
	LastUpdatedDateTime *string                                                 `json:"lastUpdatedDateTime,omitempty"`
	ManagedDeviceId     *string                                                 `json:"managedDeviceId,omitempty"`
	ODataType           *string                                                 `json:"@odata.type,omitempty"`
	StartDateTime       *string                                                 `json:"startDateTime,omitempty"`
	TotalLicensesCount  *int64                                                  `json:"totalLicensesCount,omitempty"`
	UserId              *string                                                 `json:"userId,omitempty"`
}
