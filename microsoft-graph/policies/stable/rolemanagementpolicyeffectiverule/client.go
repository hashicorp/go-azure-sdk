package rolemanagementpolicyeffectiverule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyEffectiveRuleClient struct {
	Client *msgraph.Client
}

func NewRoleManagementPolicyEffectiveRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementPolicyEffectiveRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementpolicyeffectiverule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementPolicyEffectiveRuleClient: %+v", err)
	}

	return &RoleManagementPolicyEffectiveRuleClient{
		Client: client,
	}, nil
}
