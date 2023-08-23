package realusermetrics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RealUserMetricsClient struct {
	Client *resourcemanager.Client
}

func NewRealUserMetricsClientWithBaseURI(sdkApi sdkEnv.Api) (*RealUserMetricsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "realusermetrics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RealUserMetricsClient: %+v", err)
	}

	return &RealUserMetricsClient{
		Client: client,
	}, nil
}
