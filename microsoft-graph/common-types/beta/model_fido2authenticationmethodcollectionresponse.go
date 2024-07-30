package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodCollectionResponse struct {
	ODataCount    *int64                       `json:"@odata.count,omitempty"`
	ODataNextLink *string                      `json:"@odata.nextLink,omitempty"`
	Value         *[]Fido2AuthenticationMethod `json:"value,omitempty"`
}
