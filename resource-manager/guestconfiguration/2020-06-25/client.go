package v2020_06_25

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignmenthcrpreports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignmentreports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationconnectedvmwarevsphereassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationconnectedvmwarevsphereassignmentsreports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/guestconfiguration/2020-06-25/guestconfigurationhcrpassignments"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GuestConfigurationAssignmentHCRPReports                    *guestconfigurationassignmenthcrpreports.GuestConfigurationAssignmentHCRPReportsClient
	GuestConfigurationAssignmentReports                        *guestconfigurationassignmentreports.GuestConfigurationAssignmentReportsClient
	GuestConfigurationAssignments                              *guestconfigurationassignments.GuestConfigurationAssignmentsClient
	GuestConfigurationConnectedVMwarevSphereAssignments        *guestconfigurationconnectedvmwarevsphereassignments.GuestConfigurationConnectedVMwarevSphereAssignmentsClient
	GuestConfigurationConnectedVMwarevSphereAssignmentsReports *guestconfigurationconnectedvmwarevsphereassignmentsreports.GuestConfigurationConnectedVMwarevSphereAssignmentsReportsClient
	GuestConfigurationHCRPAssignments                          *guestconfigurationhcrpassignments.GuestConfigurationHCRPAssignmentsClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	guestConfigurationAssignmentHCRPReportsClient, err := guestconfigurationassignmenthcrpreports.NewGuestConfigurationAssignmentHCRPReportsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationAssignmentHCRPReports client: %+v", err)
	}
	configureFunc(guestConfigurationAssignmentHCRPReportsClient.Client)

	guestConfigurationAssignmentReportsClient, err := guestconfigurationassignmentreports.NewGuestConfigurationAssignmentReportsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationAssignmentReports client: %+v", err)
	}
	configureFunc(guestConfigurationAssignmentReportsClient.Client)

	guestConfigurationAssignmentsClient, err := guestconfigurationassignments.NewGuestConfigurationAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationAssignments client: %+v", err)
	}
	configureFunc(guestConfigurationAssignmentsClient.Client)

	guestConfigurationConnectedVMwarevSphereAssignmentsClient, err := guestconfigurationconnectedvmwarevsphereassignments.NewGuestConfigurationConnectedVMwarevSphereAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationConnectedVMwarevSphereAssignments client: %+v", err)
	}
	configureFunc(guestConfigurationConnectedVMwarevSphereAssignmentsClient.Client)

	guestConfigurationConnectedVMwarevSphereAssignmentsReportsClient, err := guestconfigurationconnectedvmwarevsphereassignmentsreports.NewGuestConfigurationConnectedVMwarevSphereAssignmentsReportsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationConnectedVMwarevSphereAssignmentsReports client: %+v", err)
	}
	configureFunc(guestConfigurationConnectedVMwarevSphereAssignmentsReportsClient.Client)

	guestConfigurationHCRPAssignmentsClient, err := guestconfigurationhcrpassignments.NewGuestConfigurationHCRPAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GuestConfigurationHCRPAssignments client: %+v", err)
	}
	configureFunc(guestConfigurationHCRPAssignmentsClient.Client)

	return &Client{
		GuestConfigurationAssignmentHCRPReports:                    guestConfigurationAssignmentHCRPReportsClient,
		GuestConfigurationAssignmentReports:                        guestConfigurationAssignmentReportsClient,
		GuestConfigurationAssignments:                              guestConfigurationAssignmentsClient,
		GuestConfigurationConnectedVMwarevSphereAssignments:        guestConfigurationConnectedVMwarevSphereAssignmentsClient,
		GuestConfigurationConnectedVMwarevSphereAssignmentsReports: guestConfigurationConnectedVMwarevSphereAssignmentsReportsClient,
		GuestConfigurationHCRPAssignments:                          guestConfigurationHCRPAssignmentsClient,
	}, nil
}
