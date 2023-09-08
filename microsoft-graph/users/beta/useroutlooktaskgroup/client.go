package useroutlooktaskgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskGroupClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskGroupClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskGroupClient: %+v", err)
	}

	return &UserOutlookTaskGroupClient{
		Client: client,
	}, nil
}
