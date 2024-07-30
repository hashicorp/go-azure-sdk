package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Endpoint struct {
	Capability         *string `json:"capability,omitempty"`
	DeletedDateTime    *string `json:"deletedDateTime,omitempty"`
	Id                 *string `json:"id,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	ProviderId         *string `json:"providerId,omitempty"`
	ProviderName       *string `json:"providerName,omitempty"`
	ProviderResourceId *string `json:"providerResourceId,omitempty"`
	Uri                *string `json:"uri,omitempty"`
}
