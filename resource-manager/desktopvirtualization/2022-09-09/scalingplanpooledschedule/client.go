package scalingplanpooledschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingPlanPooledScheduleClient struct {
	Client *resourcemanager.Client
}

func NewScalingPlanPooledScheduleClientWithBaseURI(api sdkEnv.Api) (*ScalingPlanPooledScheduleClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "scalingplanpooledschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScalingPlanPooledScheduleClient: %+v", err)
	}

	return &ScalingPlanPooledScheduleClient{
		Client: client,
	}, nil
}
