package userplannerplantask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanTaskClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanTaskClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplantask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanTaskClient: %+v", err)
	}

	return &UserPlannerPlanTaskClient{
		Client: client,
	}, nil
}
