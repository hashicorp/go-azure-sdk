package virtualendpointprovisioningpolicyassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointProvisioningPolicyAssignmentClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointProvisioningPolicyAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointProvisioningPolicyAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointprovisioningpolicyassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointProvisioningPolicyAssignmentClient: %+v", err)
	}

	return &VirtualEndpointProvisioningPolicyAssignmentClient{
		Client: client,
	}, nil
}
