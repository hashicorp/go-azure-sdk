package meplannerplandetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanDetailClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanDetailClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplandetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanDetailClient: %+v", err)
	}

	return &MePlannerPlanDetailClient{
		Client: client,
	}, nil
}
