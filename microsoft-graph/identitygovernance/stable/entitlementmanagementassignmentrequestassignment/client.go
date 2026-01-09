package entitlementmanagementassignmentrequestassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentRequestAssignmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentRequestAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentRequestAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentrequestassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentRequestAssignmentClient: %+v", err)
	}

	return &EntitlementManagementAssignmentRequestAssignmentClient{
		Client: client,
	}, nil
}
