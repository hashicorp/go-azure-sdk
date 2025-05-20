package cloudhsmclustergetbysubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterGetBySubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterGetBySubscriptionClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterGetBySubscriptionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclustergetbysubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterGetBySubscriptionClient: %+v", err)
	}

	return &CloudHSMClusterGetBySubscriptionClient{
		Client: client,
	}, nil
}
