package productsubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductSubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewProductSubscriptionClientWithBaseURI(api sdkEnv.Api) (*ProductSubscriptionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "productsubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProductSubscriptionClient: %+v", err)
	}

	return &ProductSubscriptionClient{
		Client: client,
	}, nil
}
