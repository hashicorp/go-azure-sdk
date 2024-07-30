package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityPolicySettings struct {
	ApplicableTo                              *SensitivityPolicySettingsApplicableTo `json:"applicableTo,omitempty"`
	DowngradeSensitivityRequiresJustification *bool                                  `json:"downgradeSensitivityRequiresJustification,omitempty"`
	HelpWebUrl                                *string                                `json:"helpWebUrl,omitempty"`
	Id                                        *string                                `json:"id,omitempty"`
	IsMandatory                               *bool                                  `json:"isMandatory,omitempty"`
	ODataType                                 *string                                `json:"@odata.type,omitempty"`
}
