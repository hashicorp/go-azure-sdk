package usermanagedappregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserManagedAppRegistrationClient struct {
	Client *msgraph.Client
}

func NewUserManagedAppRegistrationClientWithBaseURI(api sdkEnv.Api) (*UserManagedAppRegistrationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermanagedappregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserManagedAppRegistrationClient: %+v", err)
	}

	return &UserManagedAppRegistrationClient{
		Client: client,
	}, nil
}
