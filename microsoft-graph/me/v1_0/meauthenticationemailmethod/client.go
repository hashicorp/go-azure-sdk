package meauthenticationemailmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationEmailMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationEmailMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationEmailMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationemailmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationEmailMethodClient: %+v", err)
	}

	return &MeAuthenticationEmailMethodClient{
		Client: client,
	}, nil
}
