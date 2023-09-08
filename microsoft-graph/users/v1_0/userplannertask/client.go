package userplannertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerTaskClient struct {
	Client *msgraph.Client
}

func NewUserPlannerTaskClientWithBaseURI(api sdkEnv.Api) (*UserPlannerTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerTaskClient: %+v", err)
	}

	return &UserPlannerTaskClient{
		Client: client,
	}, nil
}
