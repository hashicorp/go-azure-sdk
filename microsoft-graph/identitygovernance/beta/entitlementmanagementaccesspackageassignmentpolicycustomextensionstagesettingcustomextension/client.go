package entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesettingcustomextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesettingcustomextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient{
		Client: client,
	}, nil
}
