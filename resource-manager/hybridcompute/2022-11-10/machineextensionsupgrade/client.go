package machineextensionsupgrade

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineExtensionsUpgradeClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMachineExtensionsUpgradeClientWithBaseURI(endpoint string) MachineExtensionsUpgradeClient {
	return MachineExtensionsUpgradeClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
