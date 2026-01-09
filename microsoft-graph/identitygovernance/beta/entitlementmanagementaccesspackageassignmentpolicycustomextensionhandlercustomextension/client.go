package entitlementmanagementaccesspackageassignmentpolicycustomextensionhandlercustomextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentpolicycustomextensionhandlercustomextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionHandlerCustomExtensionClient{
		Client: client,
	}, nil
}
