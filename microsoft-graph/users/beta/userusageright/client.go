package userusageright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserUsageRightClient struct {
	Client *msgraph.Client
}

func NewUserUsageRightClientWithBaseURI(api sdkEnv.Api) (*UserUsageRightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userusageright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserUsageRightClient: %+v", err)
	}

	return &UserUsageRightClient{
		Client: client,
	}, nil
}
