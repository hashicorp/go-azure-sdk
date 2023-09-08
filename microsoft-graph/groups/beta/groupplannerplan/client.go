package groupplannerplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanClient: %+v", err)
	}

	return &GroupPlannerPlanClient{
		Client: client,
	}, nil
}
