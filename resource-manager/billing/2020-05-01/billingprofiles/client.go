package billingprofiles

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfilesClient struct {
	Client *resourcemanager.Client
}

func NewBillingProfilesClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingProfilesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingprofiles", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingProfilesClient: %+v", err)
	}

	return &BillingProfilesClient{
		Client: client,
	}, nil
}
