package policyrolemanagementpolicyeffectiverule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyRoleManagementPolicyEffectiveRuleClient struct {
	Client *msgraph.Client
}

func NewPolicyRoleManagementPolicyEffectiveRuleClientWithBaseURI(api sdkEnv.Api) (*PolicyRoleManagementPolicyEffectiveRuleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyrolemanagementpolicyeffectiverule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyRoleManagementPolicyEffectiveRuleClient: %+v", err)
	}

	return &PolicyRoleManagementPolicyEffectiveRuleClient{
		Client: client,
	}, nil
}
