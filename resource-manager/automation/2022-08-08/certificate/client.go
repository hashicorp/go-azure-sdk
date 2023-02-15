package certificate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateClient struct {
	Client *resourcemanager.Client
}

func NewCertificateClientWithBaseURI(api environments.Api) (*CertificateClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "certificate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateClient: %+v", err)
	}

	return &CertificateClient{
		Client: client,
	}, nil
}
