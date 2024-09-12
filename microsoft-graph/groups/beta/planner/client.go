package planner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerClient struct {
	Client *msgraph.Client
}

func NewPlannerClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "planner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerClient: %+v", err)
	}

	return &PlannerClient{
		Client: client,
	}, nil
}
