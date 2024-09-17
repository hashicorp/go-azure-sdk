package servicefabrics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceFabricsClient struct {
	Client *resourcemanager.Client
}

func NewServiceFabricsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServiceFabricsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "servicefabrics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceFabricsClient: %+v", err)
	}

	return &ServiceFabricsClient{
		Client: client,
	}, nil
}
