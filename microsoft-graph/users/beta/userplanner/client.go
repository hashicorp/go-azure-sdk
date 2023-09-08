package userplanner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerClient struct {
	Client *msgraph.Client
}

func NewUserPlannerClientWithBaseURI(api sdkEnv.Api) (*UserPlannerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplanner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerClient: %+v", err)
	}

	return &UserPlannerClient{
		Client: client,
	}, nil
}
