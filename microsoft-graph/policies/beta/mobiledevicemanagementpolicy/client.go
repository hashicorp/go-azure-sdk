package mobiledevicemanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileDeviceManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewMobileDeviceManagementPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileDeviceManagementPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobiledevicemanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileDeviceManagementPolicyClient: %+v", err)
	}

	return &MobileDeviceManagementPolicyClient{
		Client: client,
	}, nil
}
