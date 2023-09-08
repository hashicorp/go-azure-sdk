package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppHostedMediaConfig struct {
	Blob      *string `json:"blob,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
