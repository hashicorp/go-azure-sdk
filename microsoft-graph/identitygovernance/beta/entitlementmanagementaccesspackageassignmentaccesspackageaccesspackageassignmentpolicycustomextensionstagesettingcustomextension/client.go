package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicycustomextensionstagesettingcustomextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicycustomextensionstagesettingcustomextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient{
		Client: client,
	}, nil
}
