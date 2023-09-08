package userinsightused

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightUsedClient struct {
	Client *msgraph.Client
}

func NewUserInsightUsedClientWithBaseURI(api sdkEnv.Api) (*UserInsightUsedClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinsightused", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightUsedClient: %+v", err)
	}

	return &UserInsightUsedClient{
		Client: client,
	}, nil
}
