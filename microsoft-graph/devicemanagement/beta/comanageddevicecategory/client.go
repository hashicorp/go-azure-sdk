package comanageddevicecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagedDeviceCategoryClient struct {
	Client *msgraph.Client
}

func NewComanagedDeviceCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*ComanagedDeviceCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "comanageddevicecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComanagedDeviceCategoryClient: %+v", err)
	}

	return &ComanagedDeviceCategoryClient{
		Client: client,
	}, nil
}
