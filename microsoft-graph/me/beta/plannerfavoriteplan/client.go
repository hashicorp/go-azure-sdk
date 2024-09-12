package plannerfavoriteplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerFavoritePlanClient struct {
	Client *msgraph.Client
}

func NewPlannerFavoritePlanClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerFavoritePlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerfavoriteplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerFavoritePlanClient: %+v", err)
	}

	return &PlannerFavoritePlanClient{
		Client: client,
	}, nil
}
