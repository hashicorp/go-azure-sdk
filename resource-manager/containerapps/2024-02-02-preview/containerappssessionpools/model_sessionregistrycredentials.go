package containerappssessionpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionRegistryCredentials struct {
	PasswordSecretRef *string `json:"passwordSecretRef,omitempty"`
	RegistryServer    *string `json:"registryServer,omitempty"`
	Username          *string `json:"username,omitempty"`
}
