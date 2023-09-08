package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialSingleSignOnExtension struct {
	Configurations      *[]KeyTypedValuePair `json:"configurations,omitempty"`
	Domains             *[]string            `json:"domains,omitempty"`
	ExtensionIdentifier *string              `json:"extensionIdentifier,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
	Realm               *string              `json:"realm,omitempty"`
	TeamIdentifier      *string              `json:"teamIdentifier,omitempty"`
}
