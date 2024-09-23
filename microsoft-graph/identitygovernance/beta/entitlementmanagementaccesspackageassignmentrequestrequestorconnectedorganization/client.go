package entitlementmanagementaccesspackageassignmentrequestrequestorconnectedorganization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentRequestRequestorConnectedOrganizationClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentRequestRequestorConnectedOrganizationClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentRequestRequestorConnectedOrganizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentrequestrequestorconnectedorganization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentRequestRequestorConnectedOrganizationClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentRequestRequestorConnectedOrganizationClient{
		Client: client,
	}, nil
}
