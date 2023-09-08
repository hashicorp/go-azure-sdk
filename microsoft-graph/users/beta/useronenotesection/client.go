package useronenotesection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionClient: %+v", err)
	}

	return &UserOnenoteSectionClient{
		Client: client,
	}, nil
}
