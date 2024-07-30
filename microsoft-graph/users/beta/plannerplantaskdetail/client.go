package plannerplantaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanTaskDetailClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplantaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanTaskDetailClient: %+v", err)
	}

	return &PlannerPlanTaskDetailClient{
		Client: client,
	}, nil
}
