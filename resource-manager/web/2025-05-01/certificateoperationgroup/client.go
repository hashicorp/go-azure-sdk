package certificateoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewCertificateOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*CertificateOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "certificateoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateOperationGroupClient: %+v", err)
	}

	return &CertificateOperationGroupClient{
		Client: client,
	}, nil
}
