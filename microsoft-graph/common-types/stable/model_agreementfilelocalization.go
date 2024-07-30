package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementFileLocalization struct {
	CreatedDateTime *string                 `json:"createdDateTime,omitempty"`
	DisplayName     *string                 `json:"displayName,omitempty"`
	FileData        *AgreementFileData      `json:"fileData,omitempty"`
	FileName        *string                 `json:"fileName,omitempty"`
	Id              *string                 `json:"id,omitempty"`
	IsDefault       *bool                   `json:"isDefault,omitempty"`
	IsMajorVersion  *bool                   `json:"isMajorVersion,omitempty"`
	Language        *string                 `json:"language,omitempty"`
	ODataType       *string                 `json:"@odata.type,omitempty"`
	Versions        *[]AgreementFileVersion `json:"versions,omitempty"`
}
