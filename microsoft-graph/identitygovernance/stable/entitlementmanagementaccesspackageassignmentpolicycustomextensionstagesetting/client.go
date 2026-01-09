package entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient{
		Client: client,
	}, nil
}
