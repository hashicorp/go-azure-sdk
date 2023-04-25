package guestconfigurationconnectedvmwarevsphereassignmentsreports

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient struct {
	Client *resourcemanager.Client
}

func NewGuestConfigurationConnectedVMwarevSphereAssignmentsReportsClientWithBaseURI(api environments.Api) (*GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "guestconfigurationconnectedvmwarevsphereassignmentsreports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient: %+v", err)
	}

	return &GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient{
		Client: client,
	}, nil
}
