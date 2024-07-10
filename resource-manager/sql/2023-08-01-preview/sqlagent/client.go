package sqlagent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlAgentClient struct {
	Client *resourcemanager.Client
}

func NewSqlAgentClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlAgentClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlagent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlAgentClient: %+v", err)
	}

	return &SqlAgentClient{
		Client: client,
	}, nil
}
