package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGatewaySettingsParameters struct {
	RestAuthCredentialIsEnabled *bool   `json:"restAuthCredential.isEnabled,omitempty"`
	RestAuthCredentialPassword  *string `json:"restAuthCredential.password,omitempty"`
	RestAuthCredentialUsername  *string `json:"restAuthCredential.username,omitempty"`
}
