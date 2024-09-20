package linkers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkersClient struct {
	Client *resourcemanager.Client
}

func NewLinkersClientWithBaseURI(sdkApi sdkEnv.Api) (*LinkersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "linkers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LinkersClient: %+v", err)
	}

	return &LinkersClient{
		Client: client,
	}, nil
}
