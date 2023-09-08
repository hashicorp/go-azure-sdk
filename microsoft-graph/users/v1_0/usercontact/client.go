package usercontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactClient struct {
	Client *msgraph.Client
}

func NewUserContactClientWithBaseURI(api sdkEnv.Api) (*UserContactClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactClient: %+v", err)
	}

	return &UserContactClient{
		Client: client,
	}, nil
}
