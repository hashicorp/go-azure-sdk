package useroutlooktask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskClient: %+v", err)
	}

	return &UserOutlookTaskClient{
		Client: client,
	}, nil
}
