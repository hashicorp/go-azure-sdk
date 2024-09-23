package usageright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageRightClient struct {
	Client *msgraph.Client
}

func NewUsageRightClientWithBaseURI(sdkApi sdkEnv.Api) (*UsageRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "usageright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UsageRightClient: %+v", err)
	}

	return &UsageRightClient{
		Client: client,
	}, nil
}
