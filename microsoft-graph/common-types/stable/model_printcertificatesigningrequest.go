package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintCertificateSigningRequest struct {
	Content      *string `json:"content,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	TransportKey *string `json:"transportKey,omitempty"`
}
