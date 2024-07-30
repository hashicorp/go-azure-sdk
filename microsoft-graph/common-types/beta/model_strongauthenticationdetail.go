package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StrongAuthenticationDetail struct {
	EncryptedPinHashHistory *string `json:"encryptedPinHashHistory,omitempty"`
	Id                      *string `json:"id,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	ProofupTime             *int64  `json:"proofupTime,omitempty"`
}
