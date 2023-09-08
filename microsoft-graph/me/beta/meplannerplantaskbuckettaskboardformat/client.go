package meplannerplantaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplantaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &MePlannerPlanTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
