package policyappmanagementpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAppManagementPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyAppManagementPolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyAppManagementPolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyappmanagementpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAppManagementPolicyAppliesToClient: %+v", err)
	}

	return &PolicyAppManagementPolicyAppliesToClient{
		Client: client,
	}, nil
}
