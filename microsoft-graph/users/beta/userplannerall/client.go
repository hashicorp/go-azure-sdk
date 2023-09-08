package userplannerall

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerAllClient struct {
	Client *msgraph.Client
}

func NewUserPlannerAllClientWithBaseURI(api sdkEnv.Api) (*UserPlannerAllClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerall", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerAllClient: %+v", err)
	}

	return &UserPlannerAllClient{
		Client: client,
	}, nil
}
