package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesCurrentExportData struct {
	ClientMachineName                  *string `json:"clientMachineName,omitempty"`
	ODataType                          *string `json:"@odata.type,omitempty"`
	PendingObjectsAddition             *int64  `json:"pendingObjectsAddition,omitempty"`
	PendingObjectsDeletion             *int64  `json:"pendingObjectsDeletion,omitempty"`
	PendingObjectsUpdate               *int64  `json:"pendingObjectsUpdate,omitempty"`
	ServiceAccount                     *string `json:"serviceAccount,omitempty"`
	SuccessfulLinksProvisioningCount   *int64  `json:"successfulLinksProvisioningCount,omitempty"`
	SuccessfulObjectsProvisioningCount *int64  `json:"successfulObjectsProvisioningCount,omitempty"`
	TotalConnectorSpaceObjects         *int64  `json:"totalConnectorSpaceObjects,omitempty"`
}
