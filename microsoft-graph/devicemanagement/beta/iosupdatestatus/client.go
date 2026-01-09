package iosupdatestatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateStatusClient struct {
	Client *msgraph.Client
}

func NewIosUpdateStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*IosUpdateStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "iosupdatestatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IosUpdateStatusClient: %+v", err)
	}

	return &IosUpdateStatusClient{
		Client: client,
	}, nil
}
