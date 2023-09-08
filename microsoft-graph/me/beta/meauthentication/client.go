package meauthentication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthentication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationClient: %+v", err)
	}

	return &MeAuthenticationClient{
		Client: client,
	}, nil
}
