package userownedobject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOwnedObjectClient struct {
	Client *msgraph.Client
}

func NewUserOwnedObjectClientWithBaseURI(api sdkEnv.Api) (*UserOwnedObjectClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userownedobject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOwnedObjectClient: %+v", err)
	}

	return &UserOwnedObjectClient{
		Client: client,
	}, nil
}
