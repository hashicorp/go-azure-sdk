package tdecertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TdeCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewTdeCertificatesClientWithBaseURI(api sdkEnv.Api) (*TdeCertificatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "tdecertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TdeCertificatesClient: %+v", err)
	}

	return &TdeCertificatesClient{
		Client: client,
	}, nil
}
