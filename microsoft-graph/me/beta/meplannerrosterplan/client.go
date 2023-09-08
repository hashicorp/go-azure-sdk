package meplannerrosterplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerRosterPlanClient struct {
	Client *msgraph.Client
}

func NewMePlannerRosterPlanClientWithBaseURI(api sdkEnv.Api) (*MePlannerRosterPlanClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerrosterplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerRosterPlanClient: %+v", err)
	}

	return &MePlannerRosterPlanClient{
		Client: client,
	}, nil
}
