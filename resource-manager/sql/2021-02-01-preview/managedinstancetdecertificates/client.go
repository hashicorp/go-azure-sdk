package managedinstancetdecertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceTdeCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstanceTdeCertificatesClientWithBaseURI(api environments.Api) (*ManagedInstanceTdeCertificatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedinstancetdecertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstanceTdeCertificatesClient: %+v", err)
	}

	return &ManagedInstanceTdeCertificatesClient{
		Client: client,
	}, nil
}
