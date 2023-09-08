package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationalUrls struct {
	AppSignUpUrl                 *string `json:"appSignUpUrl,omitempty"`
	ODataType                    *string `json:"@odata.type,omitempty"`
	SingleSignOnDocumentationUrl *string `json:"singleSignOnDocumentationUrl,omitempty"`
}
