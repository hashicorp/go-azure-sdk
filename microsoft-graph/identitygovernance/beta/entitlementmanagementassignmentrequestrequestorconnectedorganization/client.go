package entitlementmanagementassignmentrequestrequestorconnectedorganization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentRequestRequestorConnectedOrganizationClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentRequestRequestorConnectedOrganizationClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentRequestRequestorConnectedOrganizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentrequestrequestorconnectedorganization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentRequestRequestorConnectedOrganizationClient: %+v", err)
	}

	return &EntitlementManagementAssignmentRequestRequestorConnectedOrganizationClient{
		Client: client,
	}, nil
}
