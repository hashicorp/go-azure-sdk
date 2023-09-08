package useractivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityClient struct {
	Client *msgraph.Client
}

func NewUserActivityClientWithBaseURI(api sdkEnv.Api) (*UserActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useractivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserActivityClient: %+v", err)
	}

	return &UserActivityClient{
		Client: client,
	}, nil
}
