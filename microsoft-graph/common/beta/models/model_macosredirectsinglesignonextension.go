package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSRedirectSingleSignOnExtension struct {
	Configurations      *[]KeyTypedValuePair `json:"configurations,omitempty"`
	ExtensionIdentifier *string              `json:"extensionIdentifier,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
	TeamIdentifier      *string              `json:"teamIdentifier,omitempty"`
	UrlPrefixes         *[]string            `json:"urlPrefixes,omitempty"`
}
