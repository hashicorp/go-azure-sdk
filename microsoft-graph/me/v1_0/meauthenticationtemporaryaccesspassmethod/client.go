package meauthenticationtemporaryaccesspassmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationTemporaryAccessPassMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationTemporaryAccessPassMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationTemporaryAccessPassMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationtemporaryaccesspassmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationTemporaryAccessPassMethodClient: %+v", err)
	}

	return &MeAuthenticationTemporaryAccessPassMethodClient{
		Client: client,
	}, nil
}
