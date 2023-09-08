package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabel struct {
	ApplicableTo                *SensitivityLabelApplicableTo    `json:"applicableTo,omitempty"`
	ApplicationMode             *SensitivityLabelApplicationMode `json:"applicationMode,omitempty"`
	AssignedPolicies            *[]LabelPolicy                   `json:"assignedPolicies,omitempty"`
	AutoLabeling                *AutoLabeling                    `json:"autoLabeling,omitempty"`
	Description                 *string                          `json:"description,omitempty"`
	DisplayName                 *string                          `json:"displayName,omitempty"`
	Id                          *string                          `json:"id,omitempty"`
	IsDefault                   *bool                            `json:"isDefault,omitempty"`
	IsEndpointProtectionEnabled *bool                            `json:"isEndpointProtectionEnabled,omitempty"`
	LabelActions                *[]LabelActionBase               `json:"labelActions,omitempty"`
	Name                        *string                          `json:"name,omitempty"`
	ODataType                   *string                          `json:"@odata.type,omitempty"`
	Priority                    *int64                           `json:"priority,omitempty"`
	Sublabels                   *[]SensitivityLabel              `json:"sublabels,omitempty"`
	ToolTip                     *string                          `json:"toolTip,omitempty"`
}
