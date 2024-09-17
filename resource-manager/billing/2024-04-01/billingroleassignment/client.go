package billingroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleAssignmentClient struct {
	Client *resourcemanager.Client
}

func NewBillingRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingRoleAssignmentClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "billingroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingRoleAssignmentClient: %+v", err)
	}

	return &BillingRoleAssignmentClient{
		Client: client,
	}, nil
}
