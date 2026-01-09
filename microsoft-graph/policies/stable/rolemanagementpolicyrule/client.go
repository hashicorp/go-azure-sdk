package rolemanagementpolicyrule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyRuleClient struct {
	Client *msgraph.Client
}

func NewRoleManagementPolicyRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementPolicyRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementpolicyrule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementPolicyRuleClient: %+v", err)
	}

	return &RoleManagementPolicyRuleClient{
		Client: client,
	}, nil
}
