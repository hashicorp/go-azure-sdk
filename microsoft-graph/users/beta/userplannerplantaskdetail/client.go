package userplannerplantaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanTaskDetailClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanTaskDetailClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplantaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanTaskDetailClient: %+v", err)
	}

	return &UserPlannerPlanTaskDetailClient{
		Client: client,
	}, nil
}
