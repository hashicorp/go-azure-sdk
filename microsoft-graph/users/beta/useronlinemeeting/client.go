package useronlinemeeting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeeting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingClient: %+v", err)
	}

	return &UserOnlineMeetingClient{
		Client: client,
	}, nil
}
