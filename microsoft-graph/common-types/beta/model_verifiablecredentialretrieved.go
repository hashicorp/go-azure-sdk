package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialRetrieved struct {
	ExpiryDateTime *string `json:"expiryDateTime,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
