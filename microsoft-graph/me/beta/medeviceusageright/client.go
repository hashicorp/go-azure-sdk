package medeviceusageright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceUsageRightClient struct {
	Client *msgraph.Client
}

func NewMeDeviceUsageRightClientWithBaseURI(api sdkEnv.Api) (*MeDeviceUsageRightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medeviceusageright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceUsageRightClient: %+v", err)
	}

	return &MeDeviceUsageRightClient{
		Client: client,
	}, nil
}
