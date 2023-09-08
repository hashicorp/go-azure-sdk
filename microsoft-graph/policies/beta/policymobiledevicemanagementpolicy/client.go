package policymobiledevicemanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyMobileDeviceManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyMobileDeviceManagementPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyMobileDeviceManagementPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policymobiledevicemanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyMobileDeviceManagementPolicyClient: %+v", err)
	}

	return &PolicyMobileDeviceManagementPolicyClient{
		Client: client,
	}, nil
}
