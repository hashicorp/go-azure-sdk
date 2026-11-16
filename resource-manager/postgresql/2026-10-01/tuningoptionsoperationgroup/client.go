package tuningoptionsoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TuningOptionsOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewTuningOptionsOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*TuningOptionsOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "tuningoptionsoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TuningOptionsOperationGroupClient: %+v", err)
	}

	return &TuningOptionsOperationGroupClient{
		Client: client,
	}, nil
}
