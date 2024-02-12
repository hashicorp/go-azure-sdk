package managedcertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewManagedCertificatesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedCertificatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managedcertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedCertificatesClient: %+v", err)
	}

	return &ManagedCertificatesClient{
		Client: client,
	}, nil
}
