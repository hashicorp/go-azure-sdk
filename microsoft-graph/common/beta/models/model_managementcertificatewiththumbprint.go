package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementCertificateWithThumbprint struct {
	Certificate *string `json:"certificate,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Thumbprint  *string `json:"thumbprint,omitempty"`
}
