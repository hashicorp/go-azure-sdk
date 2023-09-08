package meauthenticationwindowshelloforbusinessmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationWindowsHelloForBusinessMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationWindowsHelloForBusinessMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationWindowsHelloForBusinessMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationwindowshelloforbusinessmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationWindowsHelloForBusinessMethodClient: %+v", err)
	}

	return &MeAuthenticationWindowsHelloForBusinessMethodClient{
		Client: client,
	}, nil
}
