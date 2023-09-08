package identityapiconnector

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateIdentityApiConnectorByIdUploadClientCertificateRequest struct {
	Password    *string `json:"password,omitempty"`
	Pkcs12Value *string `json:"pkcs12Value,omitempty"`
}
