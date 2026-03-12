package hostnamebindings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostNameBindingsClient struct {
	Client *resourcemanager.Client
}

func NewHostNameBindingsClientWithBaseURI(sdkApi sdkEnv.Api) (*HostNameBindingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hostnamebindings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HostNameBindingsClient: %+v", err)
	}

	return &HostNameBindingsClient{
		Client: client,
	}, nil
}
