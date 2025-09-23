package groupquotasubscriptionids

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaSubscriptionIdsClient struct {
	Client *resourcemanager.Client
}

func NewGroupQuotaSubscriptionIdsClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupQuotaSubscriptionIdsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "groupquotasubscriptionids", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupQuotaSubscriptionIdsClient: %+v", err)
	}

	return &GroupQuotaSubscriptionIdsClient{
		Client: client,
	}, nil
}
