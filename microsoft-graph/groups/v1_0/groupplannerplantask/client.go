package groupplannerplantask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanTaskClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanTaskClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplantask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanTaskClient: %+v", err)
	}

	return &GroupPlannerPlanTaskClient{
		Client: client,
	}, nil
}
