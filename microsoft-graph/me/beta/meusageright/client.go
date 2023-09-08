package meusageright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeUsageRightClient struct {
	Client *msgraph.Client
}

func NewMeUsageRightClientWithBaseURI(api sdkEnv.Api) (*MeUsageRightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meusageright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeUsageRightClient: %+v", err)
	}

	return &MeUsageRightClient{
		Client: client,
	}, nil
}
