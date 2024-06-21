package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppAuthTokenProperties struct {
	Expires *string `json:"expires,omitempty"`
	Token   *string `json:"token,omitempty"`
}
