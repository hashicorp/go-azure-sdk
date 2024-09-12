package entitlementmanagementassignmentpolicyaccesspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyAccessPackageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyAccessPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyAccessPackageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicyaccesspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyAccessPackageClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyAccessPackageClient{
		Client: client,
	}, nil
}
