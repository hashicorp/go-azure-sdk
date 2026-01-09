package virtualendpointprovisioningpolicyassignmentassigneduserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointprovisioningpolicyassignmentassigneduserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient: %+v", err)
	}

	return &VirtualEndpointProvisioningPolicyAssignmentAssignedUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
