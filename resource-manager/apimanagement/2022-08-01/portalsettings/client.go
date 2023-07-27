package portalsettings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalSettingsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPortalSettingsClientWithBaseURI(endpoint string) PortalSettingsClient {
	return PortalSettingsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
