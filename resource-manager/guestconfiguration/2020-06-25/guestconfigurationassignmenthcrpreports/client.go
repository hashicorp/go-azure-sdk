package guestconfigurationassignmenthcrpreports

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentHCRPReportsClient struct {
	Client *resourcemanager.Client
}

func NewGuestConfigurationAssignmentHCRPReportsClientWithBaseURI(sdkApi sdkEnv.Api) (*GuestConfigurationAssignmentHCRPReportsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "guestconfigurationassignmenthcrpreports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GuestConfigurationAssignmentHCRPReportsClient: %+v", err)
	}

	return &GuestConfigurationAssignmentHCRPReportsClient{
		Client: client,
	}, nil
}
