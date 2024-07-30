package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageRight struct {
	CatalogId         *string          `json:"catalogId,omitempty"`
	Id                *string          `json:"id,omitempty"`
	ODataType         *string          `json:"@odata.type,omitempty"`
	ServiceIdentifier *string          `json:"serviceIdentifier,omitempty"`
	State             *UsageRightState `json:"state,omitempty"`
}
