package entitlementmanagementaccesspackageassignmentapprovalstage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentApprovalStageClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentApprovalStageClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentApprovalStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentapprovalstage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentApprovalStageClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentApprovalStageClient{
		Client: client,
	}, nil
}
