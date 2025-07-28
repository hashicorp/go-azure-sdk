package dtls

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DTLSClient struct {
	Client *resourcemanager.Client
}

func NewDTLSClientWithBaseURI(sdkApi sdkEnv.Api) (*DTLSClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dtls", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DTLSClient: %+v", err)
	}

	return &DTLSClient{
		Client: client,
	}, nil
}
