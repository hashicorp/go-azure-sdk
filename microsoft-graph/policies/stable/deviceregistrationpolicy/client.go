package deviceregistrationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationPolicyClient struct {
	Client *msgraph.Client
}

func NewDeviceRegistrationPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceRegistrationPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deviceregistrationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceRegistrationPolicyClient: %+v", err)
	}

	return &DeviceRegistrationPolicyClient{
		Client: client,
	}, nil
}
