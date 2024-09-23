package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicycustomextensionstagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicycustomextensionstagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyCustomExtensionStageSettingClient{
		Client: client,
	}, nil
}
