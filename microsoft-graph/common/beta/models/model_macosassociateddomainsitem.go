package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAssociatedDomainsItem struct {
	ApplicationIdentifier  *string   `json:"applicationIdentifier,omitempty"`
	DirectDownloadsEnabled *bool     `json:"directDownloadsEnabled,omitempty"`
	Domains                *[]string `json:"domains,omitempty"`
	ODataType              *string   `json:"@odata.type,omitempty"`
}
