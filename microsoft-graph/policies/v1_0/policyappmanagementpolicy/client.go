package policyappmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyAppManagementPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyAppManagementPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyappmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAppManagementPolicyClient: %+v", err)
	}

	return &PolicyAppManagementPolicyClient{
		Client: client,
	}, nil
}
