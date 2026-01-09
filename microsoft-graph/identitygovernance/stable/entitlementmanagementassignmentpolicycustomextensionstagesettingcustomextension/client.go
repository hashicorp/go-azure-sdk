package entitlementmanagementassignmentpolicycustomextensionstagesettingcustomextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicycustomextensionstagesettingcustomextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient{
		Client: client,
	}, nil
}
