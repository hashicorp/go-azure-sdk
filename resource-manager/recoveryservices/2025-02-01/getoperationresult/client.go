package getoperationresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOperationResultClient struct {
	Client *resourcemanager.Client
}

func NewGetOperationResultClientWithBaseURI(sdkApi sdkEnv.Api) (*GetOperationResultClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "getoperationresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GetOperationResultClient: %+v", err)
	}

	return &GetOperationResultClient{
		Client: client,
	}, nil
}
