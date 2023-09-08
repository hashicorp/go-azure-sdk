package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveType struct {
	ClassificationMethod *SensitiveTypeClassificationMethod `json:"classificationMethod,omitempty"`
	Description          *string                            `json:"description,omitempty"`
	Id                   *string                            `json:"id,omitempty"`
	Name                 *string                            `json:"name,omitempty"`
	ODataType            *string                            `json:"@odata.type,omitempty"`
	PublisherName        *string                            `json:"publisherName,omitempty"`
	RulePackageId        *string                            `json:"rulePackageId,omitempty"`
	RulePackageType      *string                            `json:"rulePackageType,omitempty"`
	Scope                *SensitiveTypeScope                `json:"scope,omitempty"`
	SensitiveTypeSource  *SensitiveTypeSensitiveTypeSource  `json:"sensitiveTypeSource,omitempty"`
	State                *string                            `json:"state,omitempty"`
}
