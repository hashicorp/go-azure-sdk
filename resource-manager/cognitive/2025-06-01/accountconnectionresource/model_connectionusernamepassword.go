package accountconnectionresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionUsernamePassword struct {
	Password      *string `json:"password,omitempty"`
	SecurityToken *string `json:"securityToken,omitempty"`
	Username      *string `json:"username,omitempty"`
}
