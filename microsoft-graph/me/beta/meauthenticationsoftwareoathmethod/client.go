package meauthenticationsoftwareoathmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationSoftwareOathMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationSoftwareOathMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationSoftwareOathMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationsoftwareoathmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationSoftwareOathMethodClient: %+v", err)
	}

	return &MeAuthenticationSoftwareOathMethodClient{
		Client: client,
	}, nil
}
