package usercreatedobject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCreatedObjectClient struct {
	Client *msgraph.Client
}

func NewUserCreatedObjectClientWithBaseURI(api sdkEnv.Api) (*UserCreatedObjectClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercreatedobject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCreatedObjectClient: %+v", err)
	}

	return &UserCreatedObjectClient{
		Client: client,
	}, nil
}
