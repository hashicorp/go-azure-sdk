package meauthenticationfido2method

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationFido2MethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationFido2MethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationFido2MethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationfido2method", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationFido2MethodClient: %+v", err)
	}

	return &MeAuthenticationFido2MethodClient{
		Client: client,
	}, nil
}
