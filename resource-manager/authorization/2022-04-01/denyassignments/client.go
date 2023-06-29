package denyassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DenyAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewDenyAssignmentsClientWithBaseURI(api environments.Api) (*DenyAssignmentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "denyassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DenyAssignmentsClient: %+v", err)
	}

	return &DenyAssignmentsClient{
		Client: client,
	}, nil
}
