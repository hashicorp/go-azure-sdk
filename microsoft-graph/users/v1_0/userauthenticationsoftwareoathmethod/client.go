package userauthenticationsoftwareoathmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationSoftwareOathMethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationSoftwareOathMethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationSoftwareOathMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationsoftwareoathmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationSoftwareOathMethodClient: %+v", err)
	}

	return &UserAuthenticationSoftwareOathMethodClient{
		Client: client,
	}, nil
}
