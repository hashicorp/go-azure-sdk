package scalingplanpooledschedule

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingPlanPooledScheduleClient struct {
	Client  autorest.Client
	baseUri string
}

func NewScalingPlanPooledScheduleClientWithBaseURI(endpoint string) ScalingPlanPooledScheduleClient {
	return ScalingPlanPooledScheduleClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
