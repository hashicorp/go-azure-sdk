package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnTokenIssuanceStartCustomExtensionHandler struct {
	Configuration   *CustomExtensionOverwriteConfiguration `json:"configuration,omitempty"`
	CustomExtension *OnTokenIssuanceStartCustomExtension   `json:"customExtension,omitempty"`
	ODataType       *string                                `json:"@odata.type,omitempty"`
}
