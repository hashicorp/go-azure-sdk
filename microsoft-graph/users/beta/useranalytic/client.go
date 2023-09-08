package useranalytic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAnalyticClient struct {
	Client *msgraph.Client
}

func NewUserAnalyticClientWithBaseURI(api sdkEnv.Api) (*UserAnalyticClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useranalytic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAnalyticClient: %+v", err)
	}

	return &UserAnalyticClient{
		Client: client,
	}, nil
}
