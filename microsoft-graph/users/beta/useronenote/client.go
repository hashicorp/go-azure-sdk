package useronenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteClient: %+v", err)
	}

	return &UserOnenoteClient{
		Client: client,
	}, nil
}
