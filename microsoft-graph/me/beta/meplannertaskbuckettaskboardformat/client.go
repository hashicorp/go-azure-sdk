package meplannertaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewMePlannerTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*MePlannerTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannertaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &MePlannerTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
