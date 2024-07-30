package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethod struct {
	AaGuid                  *string                                    `json:"aaGuid,omitempty"`
	AttestationCertificates *[]string                                  `json:"attestationCertificates,omitempty"`
	AttestationLevel        *Fido2AuthenticationMethodAttestationLevel `json:"attestationLevel,omitempty"`
	CreatedDateTime         *string                                    `json:"createdDateTime,omitempty"`
	DisplayName             *string                                    `json:"displayName,omitempty"`
	Id                      *string                                    `json:"id,omitempty"`
	Model                   *string                                    `json:"model,omitempty"`
	ODataType               *string                                    `json:"@odata.type,omitempty"`
}
