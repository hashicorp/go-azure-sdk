package subscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionClient struct {
	Client *msgraph.Client
}

func NewSubscriptionClientWithBaseURI(sdkApi sdkEnv.Api) (*SubscriptionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "subscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubscriptionClient: %+v", err)
	}

	return &SubscriptionClient{
		Client: client,
	}, nil
}
