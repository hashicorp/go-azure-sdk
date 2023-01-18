package rolemanagementpolicyassignments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyAssignmentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRoleManagementPolicyAssignmentsClientWithBaseURI(endpoint string) RoleManagementPolicyAssignmentsClient {
	return RoleManagementPolicyAssignmentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
