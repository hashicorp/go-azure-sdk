package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionOverwriteConfiguration struct {
	ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`
	ODataType           *string                             `json:"@odata.type,omitempty"`
}
