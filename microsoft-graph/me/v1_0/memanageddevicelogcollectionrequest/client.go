package memanageddevicelogcollectionrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeManagedDeviceLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewMeManagedDeviceLogCollectionRequestClientWithBaseURI(api sdkEnv.Api) (*MeManagedDeviceLogCollectionRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memanageddevicelogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeManagedDeviceLogCollectionRequestClient: %+v", err)
	}

	return &MeManagedDeviceLogCollectionRequestClient{
		Client: client,
	}, nil
}
