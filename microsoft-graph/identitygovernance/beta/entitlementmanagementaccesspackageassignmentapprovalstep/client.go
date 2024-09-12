package entitlementmanagementaccesspackageassignmentapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentApprovalStepClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentApprovalStepClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentApprovalStepClient{
		Client: client,
	}, nil
}
