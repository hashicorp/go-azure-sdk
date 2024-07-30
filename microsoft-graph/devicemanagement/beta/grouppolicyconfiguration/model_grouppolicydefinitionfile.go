package grouppolicyconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionFile struct {
	Definitions          *[]GroupPolicyDefinition             `json:"definitions,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	DisplayName          *string                              `json:"displayName,omitempty"`
	FileName             *string                              `json:"fileName,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	LanguageCodes        *[]string                            `json:"languageCodes,omitempty"`
	LastModifiedDateTime *string                              `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	PolicyType           *GroupPolicyDefinitionFilePolicyType `json:"policyType,omitempty"`
	Revision             *string                              `json:"revision,omitempty"`
	TargetNamespace      *string                              `json:"targetNamespace,omitempty"`
	TargetPrefix         *string                              `json:"targetPrefix,omitempty"`
}
