package plannertaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannertaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &PlannerTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
