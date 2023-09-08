package groupplanner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplanner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerClient: %+v", err)
	}

	return &GroupPlannerClient{
		Client: client,
	}, nil
}
