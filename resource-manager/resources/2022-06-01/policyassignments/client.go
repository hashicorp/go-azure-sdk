package policyassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewPolicyAssignmentsClientWithBaseURI(api environments.Api) (*PolicyAssignmentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "policyassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAssignmentsClient: %+v", err)
	}

	return &PolicyAssignmentsClient{
		Client: client,
	}, nil
}
