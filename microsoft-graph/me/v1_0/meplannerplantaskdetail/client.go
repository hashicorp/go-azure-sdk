package meplannerplantaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanTaskDetailClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanTaskDetailClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplantaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanTaskDetailClient: %+v", err)
	}

	return &MePlannerPlanTaskDetailClient{
		Client: client,
	}, nil
}
