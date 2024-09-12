package entitlementmanagementaccesspackageassignmentaccesspolicycustomextensionstagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionStageSettingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionStageSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionStageSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageassignmentaccesspolicycustomextensionstagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionStageSettingClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageAssignmentAccessPolicyCustomExtensionStageSettingClient{
		Client: client,
	}, nil
}
