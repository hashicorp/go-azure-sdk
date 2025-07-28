package v2024_04_05

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationconnectedvmwarevsphereassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurationhcrpassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2024-04-05/guestconfigurations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GuestConfigurationAssignments                       *guestconfigurationassignments.GuestConfigurationAssignmentsClient
	GuestConfigurationConnectedVMwarevSphereAssignments *guestconfigurationconnectedvmwarevsphereassignments.GuestConfigurationConnectedVMwarevSphereAssignmentsClient
	GuestConfigurationHCRPAssignments                   *guestconfigurationhcrpassignments.GuestConfigurationHCRPAssignmentsClient
	Guestconfigurations                                 *guestconfigurations.GuestconfigurationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	guestConfigurationAssignmentsClient, err := guestconfigurationassignments.NewGuestConfigurationAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationAssignments client: %+v", err)
	}
	configureFunc(guestConfigurationAssignmentsClient.Client)

	guestConfigurationConnectedVMwarevSphereAssignmentsClient, err := guestconfigurationconnectedvmwarevsphereassignments.NewGuestConfigurationConnectedVMwarevSphereAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationConnectedVMwarevSphereAssignments client: %+v", err)
	}
	configureFunc(guestConfigurationConnectedVMwarevSphereAssignmentsClient.Client)

	guestConfigurationHCRPAssignmentsClient, err := guestconfigurationhcrpassignments.NewGuestConfigurationHCRPAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationHCRPAssignments client: %+v", err)
	}
	configureFunc(guestConfigurationHCRPAssignmentsClient.Client)

	guestconfigurationsClient, err := guestconfigurations.NewGuestconfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Guestconfigurations client: %+v", err)
	}
	configureFunc(guestconfigurationsClient.Client)

	return &Client{
		GuestConfigurationAssignments:                       guestConfigurationAssignmentsClient,
		GuestConfigurationConnectedVMwarevSphereAssignments: guestConfigurationConnectedVMwarevSphereAssignmentsClient,
		GuestConfigurationHCRPAssignments:                   guestConfigurationHCRPAssignmentsClient,
		Guestconfigurations:                                 guestconfigurationsClient,
	}, nil
}
