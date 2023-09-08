package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFile struct {
	Content                          *string                                      `json:"content,omitempty"`
	DefaultLanguageCode              *string                                      `json:"defaultLanguageCode,omitempty"`
	Definitions                      *[]GroupPolicyDefinition                     `json:"definitions,omitempty"`
	Description                      *string                                      `json:"description,omitempty"`
	DisplayName                      *string                                      `json:"displayName,omitempty"`
	FileName                         *string                                      `json:"fileName,omitempty"`
	GroupPolicyOperations            *[]GroupPolicyOperation                      `json:"groupPolicyOperations,omitempty"`
	GroupPolicyUploadedLanguageFiles *[]GroupPolicyUploadedLanguageFile           `json:"groupPolicyUploadedLanguageFiles,omitempty"`
	Id                               *string                                      `json:"id,omitempty"`
	LanguageCodes                    *[]string                                    `json:"languageCodes,omitempty"`
	LastModifiedDateTime             *string                                      `json:"lastModifiedDateTime,omitempty"`
	ODataType                        *string                                      `json:"@odata.type,omitempty"`
	PolicyType                       *GroupPolicyUploadedDefinitionFilePolicyType `json:"policyType,omitempty"`
	Revision                         *string                                      `json:"revision,omitempty"`
	Status                           *GroupPolicyUploadedDefinitionFileStatus     `json:"status,omitempty"`
	TargetNamespace                  *string                                      `json:"targetNamespace,omitempty"`
	TargetPrefix                     *string                                      `json:"targetPrefix,omitempty"`
	UploadDateTime                   *string                                      `json:"uploadDateTime,omitempty"`
}
