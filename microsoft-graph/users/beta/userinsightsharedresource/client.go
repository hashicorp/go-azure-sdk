package userinsightsharedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightSharedResourceClient struct {
	Client *msgraph.Client
}

func NewUserInsightSharedResourceClientWithBaseURI(api sdkEnv.Api) (*UserInsightSharedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinsightsharedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightSharedResourceClient: %+v", err)
	}

	return &UserInsightSharedResourceClient{
		Client: client,
	}, nil
}
