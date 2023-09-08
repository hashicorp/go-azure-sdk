package meplannerrecentplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerRecentPlanClient struct {
	Client *msgraph.Client
}

func NewMePlannerRecentPlanClientWithBaseURI(api sdkEnv.Api) (*MePlannerRecentPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerrecentplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerRecentPlanClient: %+v", err)
	}

	return &MePlannerRecentPlanClient{
		Client: client,
	}, nil
}
