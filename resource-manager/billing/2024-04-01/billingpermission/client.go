package billingpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPermissionClient struct {
	Client *resourcemanager.Client
}

func NewBillingPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingPermissionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingPermissionClient: %+v", err)
	}

	return &BillingPermissionClient{
		Client: client,
	}, nil
}
