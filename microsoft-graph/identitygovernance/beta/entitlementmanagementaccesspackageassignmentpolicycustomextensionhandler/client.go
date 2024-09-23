package entitlementmanagementaccesspackageassignmentpolicycustomextensionhandler

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentpolicycustomextensionhandler", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerClient{
		Client: client,
	}, nil
}
