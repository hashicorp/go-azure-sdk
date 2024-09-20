package standardoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardOperationClient struct {
	Client *resourcemanager.Client
}

func NewStandardOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*StandardOperationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "standardoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StandardOperationClient: %+v", err)
	}

	return &StandardOperationClient{
		Client: client,
	}, nil
}
