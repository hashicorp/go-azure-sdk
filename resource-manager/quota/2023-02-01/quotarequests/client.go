package quotarequests

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaRequestsClient struct {
	Client *resourcemanager.Client
}

func NewQuotaRequestsClientWithBaseURI(sdkApi sdkEnv.Api) (*QuotaRequestsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "quotarequests", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QuotaRequestsClient: %+v", err)
	}

	return &QuotaRequestsClient{
		Client: client,
	}, nil
}
