package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CryptographySuite struct {
	AuthenticationTransformConstants *CryptographySuiteAuthenticationTransformConstants `json:"authenticationTransformConstants,omitempty"`
	CipherTransformConstants         *CryptographySuiteCipherTransformConstants         `json:"cipherTransformConstants,omitempty"`
	DhGroup                          *CryptographySuiteDhGroup                          `json:"dhGroup,omitempty"`
	EncryptionMethod                 *CryptographySuiteEncryptionMethod                 `json:"encryptionMethod,omitempty"`
	IntegrityCheckMethod             *CryptographySuiteIntegrityCheckMethod             `json:"integrityCheckMethod,omitempty"`
	ODataType                        *string                                            `json:"@odata.type,omitempty"`
	PfsGroup                         *CryptographySuitePfsGroup                         `json:"pfsGroup,omitempty"`
}
