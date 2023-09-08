package groupplannerplandetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanDetailClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanDetailClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplandetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanDetailClient: %+v", err)
	}

	return &GroupPlannerPlanDetailClient{
		Client: client,
	}, nil
}
