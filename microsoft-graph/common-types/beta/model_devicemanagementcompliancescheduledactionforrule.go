package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplianceScheduledActionForRule struct {
	Id                            *string                                 `json:"id,omitempty"`
	ODataType                     *string                                 `json:"@odata.type,omitempty"`
	RuleName                      *string                                 `json:"ruleName,omitempty"`
	ScheduledActionConfigurations *[]DeviceManagementComplianceActionItem `json:"scheduledActionConfigurations,omitempty"`
}
