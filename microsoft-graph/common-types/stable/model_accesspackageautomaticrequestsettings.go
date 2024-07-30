package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAutomaticRequestSettings struct {
	GracePeriodBeforeAccessRemoval             *string `json:"gracePeriodBeforeAccessRemoval,omitempty"`
	ODataType                                  *string `json:"@odata.type,omitempty"`
	RemoveAccessWhenTargetLeavesAllowedTargets *bool   `json:"removeAccessWhenTargetLeavesAllowedTargets,omitempty"`
	RequestAccessForAllowedTargets             *bool   `json:"requestAccessForAllowedTargets,omitempty"`
}
