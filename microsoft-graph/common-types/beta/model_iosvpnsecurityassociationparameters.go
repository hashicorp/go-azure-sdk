package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParameters struct {
	LifetimeInMinutes           *int64                                                          `json:"lifetimeInMinutes,omitempty"`
	ODataType                   *string                                                         `json:"@odata.type,omitempty"`
	SecurityDiffieHellmanGroup  *int64                                                          `json:"securityDiffieHellmanGroup,omitempty"`
	SecurityEncryptionAlgorithm *IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm `json:"securityEncryptionAlgorithm,omitempty"`
	SecurityIntegrityAlgorithm  *IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm  `json:"securityIntegrityAlgorithm,omitempty"`
}
