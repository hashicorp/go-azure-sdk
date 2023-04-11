package autoprovisioningsettings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoProvisioningSettingsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAutoProvisioningSettingsClientWithBaseURI(endpoint string) AutoProvisioningSettingsClient {
	return AutoProvisioningSettingsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
