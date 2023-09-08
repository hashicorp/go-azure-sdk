package usercontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactExtensionClient struct {
	Client *msgraph.Client
}

func NewUserContactExtensionClientWithBaseURI(api sdkEnv.Api) (*UserContactExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactExtensionClient: %+v", err)
	}

	return &UserContactExtensionClient{
		Client: client,
	}, nil
}
