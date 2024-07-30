package siteinformationprotectionsensitivitylabelsublabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredSensitiveType struct {
	ClassificationAttributes *[]ClassificationAttribute `json:"classificationAttributes,omitempty"`
	Confidence               *int64                     `json:"confidence,omitempty"`
	Count                    *int64                     `json:"count,omitempty"`
	Id                       *string                    `json:"id,omitempty"`
	ODataType                *string                    `json:"@odata.type,omitempty"`
}
