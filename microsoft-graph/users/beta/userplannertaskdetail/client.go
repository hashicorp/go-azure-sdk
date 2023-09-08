package userplannertaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerTaskDetailClient struct {
	Client *msgraph.Client
}

func NewUserPlannerTaskDetailClientWithBaseURI(api sdkEnv.Api) (*UserPlannerTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannertaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerTaskDetailClient: %+v", err)
	}

	return &UserPlannerTaskDetailClient{
		Client: client,
	}, nil
}
