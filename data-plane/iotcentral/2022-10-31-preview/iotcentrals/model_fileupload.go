package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileUpload struct {
	Account          *string          `json:"account,omitempty"`
	ConnectionString string           `json:"connectionString"`
	Container        string           `json:"container"`
	Etag             *string          `json:"etag,omitempty"`
	ReadAccess       *bool            `json:"readAccess,omitempty"`
	SasTtl           *string          `json:"sasTtl,omitempty"`
	State            *FileUploadState `json:"state,omitempty"`
}
