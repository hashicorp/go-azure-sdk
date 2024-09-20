package billingpermissions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPermissionsClient struct {
	Client *resourcemanager.Client
}

func NewBillingPermissionsClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingPermissionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "billingpermissions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingPermissionsClient: %+v", err)
	}

	return &BillingPermissionsClient{
		Client: client,
	}, nil
}
