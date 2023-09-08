package userplannerplandetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanDetailClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanDetailClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplandetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanDetailClient: %+v", err)
	}

	return &UserPlannerPlanDetailClient{
		Client: client,
	}, nil
}
