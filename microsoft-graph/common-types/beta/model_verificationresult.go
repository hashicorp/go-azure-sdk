package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerificationResult struct {
	ODataType      *string `json:"@odata.type,omitempty"`
	SignatureValid *bool   `json:"signatureValid,omitempty"`
}
