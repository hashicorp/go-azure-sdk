package policyrolemanagementpolicyrule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRoleManagementPolicyRuleClient struct {
	Client *msgraph.Client
}

func NewPolicyRoleManagementPolicyRuleClientWithBaseURI(api sdkEnv.Api) (*PolicyRoleManagementPolicyRuleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyrolemanagementpolicyrule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyRoleManagementPolicyRuleClient: %+v", err)
	}

	return &PolicyRoleManagementPolicyRuleClient{
		Client: client,
	}, nil
}
