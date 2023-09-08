package medevicememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDeviceMemberOfClient struct {
	Client *msgraph.Client
}

func NewMeDeviceMemberOfClientWithBaseURI(api sdkEnv.Api) (*MeDeviceMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medevicememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDeviceMemberOfClient: %+v", err)
	}

	return &MeDeviceMemberOfClient{
		Client: client,
	}, nil
}
