package plannerplandetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanDetailClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplandetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanDetailClient: %+v", err)
	}

	return &PlannerPlanDetailClient{
		Client: client,
	}, nil
}
