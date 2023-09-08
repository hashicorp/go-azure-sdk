package userinsightsharedlastsharedmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightSharedLastSharedMethodClient struct {
	Client *msgraph.Client
}

func NewUserInsightSharedLastSharedMethodClientWithBaseURI(api sdkEnv.Api) (*UserInsightSharedLastSharedMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinsightsharedlastsharedmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightSharedLastSharedMethodClient: %+v", err)
	}

	return &UserInsightSharedLastSharedMethodClient{
		Client: client,
	}, nil
}
