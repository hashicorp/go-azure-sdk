package userinsightshared

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightSharedClient struct {
	Client *msgraph.Client
}

func NewUserInsightSharedClientWithBaseURI(api sdkEnv.Api) (*UserInsightSharedClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinsightshared", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightSharedClient: %+v", err)
	}

	return &UserInsightSharedClient{
		Client: client,
	}, nil
}
