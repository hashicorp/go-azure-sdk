package plannerrecentplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRecentPlanClient struct {
	Client *msgraph.Client
}

func NewPlannerRecentPlanClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerRecentPlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerrecentplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerRecentPlanClient: %+v", err)
	}

	return &PlannerRecentPlanClient{
		Client: client,
	}, nil
}
