package medeviceregistereduser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceRegisteredUserClient struct {
	Client *msgraph.Client
}

func NewMeDeviceRegisteredUserClientWithBaseURI(api sdkEnv.Api) (*MeDeviceRegisteredUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medeviceregistereduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceRegisteredUserClient: %+v", err)
	}

	return &MeDeviceRegisteredUserClient{
		Client: client,
	}, nil
}
