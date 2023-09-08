package meauthenticationpasswordmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationPasswordMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationPasswordMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationPasswordMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationpasswordmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationPasswordMethodClient: %+v", err)
	}

	return &MeAuthenticationPasswordMethodClient{
		Client: client,
	}, nil
}
