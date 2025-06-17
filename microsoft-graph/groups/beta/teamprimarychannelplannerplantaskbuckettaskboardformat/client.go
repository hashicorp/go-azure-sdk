package teamprimarychannelplannerplantaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelPlannerPlanTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelPlannerPlanTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelplannerplantaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelPlannerPlanTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &TeamPrimaryChannelPlannerPlanTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
