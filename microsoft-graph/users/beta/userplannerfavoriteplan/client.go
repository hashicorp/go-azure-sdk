package userplannerfavoriteplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerFavoritePlanClient struct {
	Client *msgraph.Client
}

func NewUserPlannerFavoritePlanClientWithBaseURI(api sdkEnv.Api) (*UserPlannerFavoritePlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerfavoriteplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerFavoritePlanClient: %+v", err)
	}

	return &UserPlannerFavoritePlanClient{
		Client: client,
	}, nil
}
