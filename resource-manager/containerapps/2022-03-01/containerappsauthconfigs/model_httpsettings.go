package containerappsauthconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HttpSettings struct {
	ForwardProxy *ForwardProxy       `json:"forwardProxy,omitempty"`
	RequireHttps *bool               `json:"requireHttps,omitempty"`
	Routes       *HttpSettingsRoutes `json:"routes,omitempty"`
}
