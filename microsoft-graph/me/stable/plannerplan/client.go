package plannerplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanClient: %+v", err)
	}

	return &PlannerPlanClient{
		Client: client,
	}, nil
}
