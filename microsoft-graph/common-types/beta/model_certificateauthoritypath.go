package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityPath struct {
	CertificateBasedApplicationConfigurations *[]CertificateBasedApplicationConfiguration `json:"certificateBasedApplicationConfigurations,omitempty"`
	Id                                        *string                                     `json:"id,omitempty"`
	ODataType                                 *string                                     `json:"@odata.type,omitempty"`
}
