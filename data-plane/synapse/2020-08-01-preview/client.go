package v2020_08_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2020-08-01-preview/checkprincipalaccess"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2020-08-01-preview/roleassignments"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2020-08-01-preview/synapserbacscopes"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2020-08-01-preview/synapseroledefinitions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	CheckPrincipalAccess   *checkprincipalaccess.CheckPrincipalAccessClient
	RoleAssignments        *roleassignments.RoleAssignmentsClient
	SynapseRbacScopes      *synapserbacscopes.SynapseRbacScopesClient
	SynapseRoleDefinitions *synapseroledefinitions.SynapseRoleDefinitionsClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	checkPrincipalAccessClient, err := checkprincipalaccess.NewCheckPrincipalAccessClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building CheckPrincipalAccess client: %+v", err)
	}
	configureFunc(checkPrincipalAccessClient.Client)

	roleAssignmentsClient, err := roleassignments.NewRoleAssignmentsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignments client: %+v", err)
	}
	configureFunc(roleAssignmentsClient.Client)

	synapseRbacScopesClient, err := synapserbacscopes.NewSynapseRbacScopesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building SynapseRbacScopes client: %+v", err)
	}
	configureFunc(synapseRbacScopesClient.Client)

	synapseRoleDefinitionsClient, err := synapseroledefinitions.NewSynapseRoleDefinitionsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building SynapseRoleDefinitions client: %+v", err)
	}
	configureFunc(synapseRoleDefinitionsClient.Client)

	return &Client{
		CheckPrincipalAccess:   checkPrincipalAccessClient,
		RoleAssignments:        roleAssignmentsClient,
		SynapseRbacScopes:      synapseRbacScopesClient,
		SynapseRoleDefinitions: synapseRoleDefinitionsClient,
	}, nil
}
