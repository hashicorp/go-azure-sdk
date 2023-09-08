package medevicetransitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceTransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewMeDeviceTransitiveMemberOfClientWithBaseURI(api sdkEnv.Api) (*MeDeviceTransitiveMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medevicetransitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceTransitiveMemberOfClient: %+v", err)
	}

	return &MeDeviceTransitiveMemberOfClient{
		Client: client,
	}, nil
}
