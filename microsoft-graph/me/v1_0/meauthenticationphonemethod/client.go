package meauthenticationphonemethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationPhoneMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationPhoneMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationPhoneMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationphonemethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationPhoneMethodClient: %+v", err)
	}

	return &MeAuthenticationPhoneMethodClient{
		Client: client,
	}, nil
}
