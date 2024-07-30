package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingLabel struct {
	ApplicationMode             *MatchingLabelApplicationMode `json:"applicationMode,omitempty"`
	Description                 *string                       `json:"description,omitempty"`
	DisplayName                 *string                       `json:"displayName,omitempty"`
	Id                          *string                       `json:"id,omitempty"`
	IsEndpointProtectionEnabled *bool                         `json:"isEndpointProtectionEnabled,omitempty"`
	LabelActions                *[]LabelActionBase            `json:"labelActions,omitempty"`
	Name                        *string                       `json:"name,omitempty"`
	ODataType                   *string                       `json:"@odata.type,omitempty"`
	PolicyTip                   *string                       `json:"policyTip,omitempty"`
	Priority                    *int64                        `json:"priority,omitempty"`
	ToolTip                     *string                       `json:"toolTip,omitempty"`
}
