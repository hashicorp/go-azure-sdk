package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionSettings struct {
	ContainerAppAuthEncryptionSecretName *string `json:"containerAppAuthEncryptionSecretName,omitempty"`
	ContainerAppAuthSigningSecretName    *string `json:"containerAppAuthSigningSecretName,omitempty"`
}
