package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleManagedIdentityProvider struct {
	CertificateData *string `json:"certificateData,omitempty"`
	DeveloperId     *string `json:"developerId,omitempty"`
	DisplayName     *string `json:"displayName,omitempty"`
	Id              *string `json:"id,omitempty"`
	KeyId           *string `json:"keyId,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	ServiceId       *string `json:"serviceId,omitempty"`
}
