package meplannerplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanClient: %+v", err)
	}

	return &MePlannerPlanClient{
		Client: client,
	}, nil
}
