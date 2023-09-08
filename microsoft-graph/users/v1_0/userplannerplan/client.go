package userplannerplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanClient: %+v", err)
	}

	return &UserPlannerPlanClient{
		Client: client,
	}, nil
}
