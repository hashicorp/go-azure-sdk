package entitlementmanagementassignmentrequestrequestor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentRequestRequestorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentRequestRequestorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentRequestRequestorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentrequestrequestor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentRequestRequestorClient: %+v", err)
	}

	return &EntitlementManagementAssignmentRequestRequestorClient{
		Client: client,
	}, nil
}
