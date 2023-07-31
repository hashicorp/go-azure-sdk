package rules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RulesClient struct {
	Client *resourcemanager.Client
}

func NewRulesClientWithBaseURI(api sdkEnv.Api) (*RulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "rules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RulesClient: %+v", err)
	}

	return &RulesClient{
		Client: client,
	}, nil
}
