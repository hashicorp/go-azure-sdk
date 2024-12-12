package activesessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyVaultCredentialsProperties struct {
	PasswordKeyVaultSecretUri string `json:"passwordKeyVaultSecretUri"`
	UsernameKeyVaultSecretUri string `json:"usernameKeyVaultSecretUri"`
}
