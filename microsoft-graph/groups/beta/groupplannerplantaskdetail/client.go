package groupplannerplantaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanTaskDetailClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanTaskDetailClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplantaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanTaskDetailClient: %+v", err)
	}

	return &GroupPlannerPlanTaskDetailClient{
		Client: client,
	}, nil
}
