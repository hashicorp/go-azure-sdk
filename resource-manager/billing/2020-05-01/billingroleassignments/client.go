package billingroleassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewBillingRoleAssignmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingRoleAssignmentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingroleassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingRoleAssignmentsClient: %+v", err)
	}

	return &BillingRoleAssignmentsClient{
		Client: client,
	}, nil
}
