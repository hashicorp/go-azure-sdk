package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateIssuerHintsConfiguration struct {
	ODataType *string                                       `json:"@odata.type,omitempty"`
	State     *X509CertificateIssuerHintsConfigurationState `json:"state,omitempty"`
}
