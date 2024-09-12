package plannermydaytask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerMyDayTaskClient struct {
	Client *msgraph.Client
}

func NewPlannerMyDayTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerMyDayTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannermydaytask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerMyDayTaskClient: %+v", err)
	}

	return &PlannerMyDayTaskClient{
		Client: client,
	}, nil
}
