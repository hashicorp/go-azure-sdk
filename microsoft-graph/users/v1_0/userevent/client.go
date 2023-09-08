package userevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventClient struct {
	Client *msgraph.Client
}

func NewUserEventClientWithBaseURI(api sdkEnv.Api) (*UserEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventClient: %+v", err)
	}

	return &UserEventClient{
		Client: client,
	}, nil
}
