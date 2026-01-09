package plannerall

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerAllClient struct {
	Client *msgraph.Client
}

func NewPlannerAllClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerAllClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerall", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerAllClient: %+v", err)
	}

	return &PlannerAllClient{
		Client: client,
	}, nil
}
