package devicesecuritygroups

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceSecurityGroupsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDeviceSecurityGroupsClientWithBaseURI(endpoint string) DeviceSecurityGroupsClient {
	return DeviceSecurityGroupsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
