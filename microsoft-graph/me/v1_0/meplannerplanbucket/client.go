package meplannerplanbucket

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanBucketClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanBucketClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanBucketClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplanbucket", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanBucketClient: %+v", err)
	}

	return &MePlannerPlanBucketClient{
		Client: client,
	}, nil
}
