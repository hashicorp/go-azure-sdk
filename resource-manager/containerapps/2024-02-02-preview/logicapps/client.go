package logicapps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogicAppsClient struct {
	Client *resourcemanager.Client
}

func NewLogicAppsClientWithBaseURI(sdkApi sdkEnv.Api) (*LogicAppsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "logicapps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LogicAppsClient: %+v", err)
	}

	return &LogicAppsClient{
		Client: client,
	}, nil
}
