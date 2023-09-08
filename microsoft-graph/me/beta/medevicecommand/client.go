package medevicecommand

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceCommandClient struct {
	Client *msgraph.Client
}

func NewMeDeviceCommandClientWithBaseURI(api sdkEnv.Api) (*MeDeviceCommandClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medevicecommand", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceCommandClient: %+v", err)
	}

	return &MeDeviceCommandClient{
		Client: client,
	}, nil
}
