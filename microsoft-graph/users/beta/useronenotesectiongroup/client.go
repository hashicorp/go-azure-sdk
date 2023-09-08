package useronenotesectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupClient: %+v", err)
	}

	return &UserOnenoteSectionGroupClient{
		Client: client,
	}, nil
}
