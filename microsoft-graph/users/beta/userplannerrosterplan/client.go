package userplannerrosterplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerRosterPlanClient struct {
	Client *msgraph.Client
}

func NewUserPlannerRosterPlanClientWithBaseURI(api sdkEnv.Api) (*UserPlannerRosterPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerrosterplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerRosterPlanClient: %+v", err)
	}

	return &UserPlannerRosterPlanClient{
		Client: client,
	}, nil
}
