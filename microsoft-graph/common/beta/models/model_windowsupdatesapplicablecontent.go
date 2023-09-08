package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesApplicableContent struct {
	CatalogEntry   *WindowsUpdatesCatalogEntry                   `json:"catalogEntry,omitempty"`
	MatchedDevices *[]WindowsUpdatesApplicableContentDeviceMatch `json:"matchedDevices,omitempty"`
	ODataType      *string                                       `json:"@odata.type,omitempty"`
}
