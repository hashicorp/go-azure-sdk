package userinsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightClient struct {
	Client *msgraph.Client
}

func NewUserInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightClient: %+v", err)
	}

	return &UserInsightClient{
		Client: client,
	}, nil
}
