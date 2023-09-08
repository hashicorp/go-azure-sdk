package usermanageddevicelogcollectionrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedDeviceLogCollectionRequestClient struct {
	Client *msgraph.Client
}

func NewUserManagedDeviceLogCollectionRequestClientWithBaseURI(api sdkEnv.Api) (*UserManagedDeviceLogCollectionRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanageddevicelogcollectionrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedDeviceLogCollectionRequestClient: %+v", err)
	}

	return &UserManagedDeviceLogCollectionRequestClient{
		Client: client,
	}, nil
}
