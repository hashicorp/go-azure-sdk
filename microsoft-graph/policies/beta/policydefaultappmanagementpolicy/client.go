package policydefaultappmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDefaultAppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyDefaultAppManagementPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyDefaultAppManagementPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policydefaultappmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyDefaultAppManagementPolicyClient: %+v", err)
	}

	return &PolicyDefaultAppManagementPolicyClient{
		Client: client,
	}, nil
}
