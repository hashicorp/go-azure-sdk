package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppContentFile struct {
	AzureStorageUri                   *string                          `json:"azureStorageUri,omitempty"`
	AzureStorageUriExpirationDateTime *string                          `json:"azureStorageUriExpirationDateTime,omitempty"`
	CreatedDateTime                   *string                          `json:"createdDateTime,omitempty"`
	Id                                *string                          `json:"id,omitempty"`
	IsCommitted                       *bool                            `json:"isCommitted,omitempty"`
	IsDependency                      *bool                            `json:"isDependency,omitempty"`
	IsFrameworkFile                   *bool                            `json:"isFrameworkFile,omitempty"`
	Manifest                          *string                          `json:"manifest,omitempty"`
	Name                              *string                          `json:"name,omitempty"`
	ODataType                         *string                          `json:"@odata.type,omitempty"`
	Size                              *int64                           `json:"size,omitempty"`
	SizeEncrypted                     *int64                           `json:"sizeEncrypted,omitempty"`
	UploadState                       *MobileAppContentFileUploadState `json:"uploadState,omitempty"`
}
