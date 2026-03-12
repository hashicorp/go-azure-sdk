package identifieroperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentifierOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewIdentifierOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*IdentifierOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "identifieroperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentifierOperationGroupClient: %+v", err)
	}

	return &IdentifierOperationGroupClient{
		Client: client,
	}, nil
}
