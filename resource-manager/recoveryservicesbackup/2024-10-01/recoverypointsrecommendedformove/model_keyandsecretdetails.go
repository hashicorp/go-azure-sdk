package recoverypointsrecommendedformove

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyAndSecretDetails struct {
	BekDetails          *BEKDetails `json:"bekDetails,omitempty"`
	EncryptionMechanism *string     `json:"encryptionMechanism,omitempty"`
	KekDetails          *KEKDetails `json:"kekDetails,omitempty"`
}
