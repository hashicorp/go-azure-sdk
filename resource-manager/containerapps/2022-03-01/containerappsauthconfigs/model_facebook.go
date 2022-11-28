package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Facebook struct {
	Enabled         *bool            `json:"enabled,omitempty"`
	GraphApiVersion *string          `json:"graphApiVersion,omitempty"`
	Login           *LoginScopes     `json:"login"`
	Registration    *AppRegistration `json:"registration"`
}
