package medeviceregisteredowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceRegisteredOwnerClient struct {
	Client *msgraph.Client
}

func NewMeDeviceRegisteredOwnerClientWithBaseURI(api sdkEnv.Api) (*MeDeviceRegisteredOwnerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medeviceregisteredowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceRegisteredOwnerClient: %+v", err)
	}

	return &MeDeviceRegisteredOwnerClient{
		Client: client,
	}, nil
}
