package entitlementmanagementassignmentpolicycustomextensionstagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyCustomExtensionStageSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicycustomextensionstagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient{
		Client: client,
	}, nil
}
