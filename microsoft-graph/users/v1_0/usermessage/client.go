package usermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMessageClient struct {
	Client *msgraph.Client
}

func NewUserMessageClientWithBaseURI(api sdkEnv.Api) (*UserMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMessageClient: %+v", err)
	}

	return &UserMessageClient{
		Client: client,
	}, nil
}
