package userplannerrecentplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerRecentPlanClient struct {
	Client *msgraph.Client
}

func NewUserPlannerRecentPlanClientWithBaseURI(api sdkEnv.Api) (*UserPlannerRecentPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerrecentplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerRecentPlanClient: %+v", err)
	}

	return &UserPlannerRecentPlanClient{
		Client: client,
	}, nil
}
