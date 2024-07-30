package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtectionPolicySetting struct {
	DefaultLabelId                   *string `json:"defaultLabelId,omitempty"`
	Id                               *string `json:"id,omitempty"`
	IsDowngradeJustificationRequired *bool   `json:"isDowngradeJustificationRequired,omitempty"`
	IsMandatory                      *bool   `json:"isMandatory,omitempty"`
	MoreInfoUrl                      *string `json:"moreInfoUrl,omitempty"`
	ODataType                        *string `json:"@odata.type,omitempty"`
}
