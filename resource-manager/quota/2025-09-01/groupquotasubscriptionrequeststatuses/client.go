package groupquotasubscriptionrequeststatuses

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaSubscriptionRequestStatusesClient struct {
	Client *resourcemanager.Client
}

func NewGroupQuotaSubscriptionRequestStatusesClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupQuotaSubscriptionRequestStatusesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "groupquotasubscriptionrequeststatuses", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupQuotaSubscriptionRequestStatusesClient: %+v", err)
	}

	return &GroupQuotaSubscriptionRequestStatusesClient{
		Client: client,
	}, nil
}
