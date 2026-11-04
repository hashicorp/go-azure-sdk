package servertrustgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerTrustGroupsClient struct {
	Client *resourcemanager.Client
}

func NewServerTrustGroupsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerTrustGroupsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "servertrustgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerTrustGroupsClient: %+v", err)
	}

	return &ServerTrustGroupsClient{
		Client: client,
	}, nil
}
