package userdevicecommandresponsepayload

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDeviceCommandResponsepayloadClient struct {
	Client *msgraph.Client
}

func NewUserDeviceCommandResponsepayloadClientWithBaseURI(api sdkEnv.Api) (*UserDeviceCommandResponsepayloadClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdevicecommandresponsepayload", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDeviceCommandResponsepayloadClient: %+v", err)
	}

	return &UserDeviceCommandResponsepayloadClient{
		Client: client,
	}, nil
}
