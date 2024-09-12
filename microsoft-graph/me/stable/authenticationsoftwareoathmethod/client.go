package authenticationsoftwareoathmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationSoftwareOathMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationSoftwareOathMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationSoftwareOathMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationsoftwareoathmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationSoftwareOathMethodClient: %+v", err)
	}

	return &AuthenticationSoftwareOathMethodClient{
		Client: client,
	}, nil
}
