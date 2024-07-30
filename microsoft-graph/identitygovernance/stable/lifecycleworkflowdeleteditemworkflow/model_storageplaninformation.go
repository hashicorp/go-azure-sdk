package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StoragePlanInformation struct {
	ODataType        *string `json:"@odata.type,omitempty"`
	UpgradeAvailable *bool   `json:"upgradeAvailable,omitempty"`
}
