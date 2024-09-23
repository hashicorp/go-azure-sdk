package authenticationsigninpreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationSignInPreferenceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationSignInPreferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationSignInPreferenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationsigninpreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationSignInPreferenceClient: %+v", err)
	}

	return &AuthenticationSignInPreferenceClient{
		Client: client,
	}, nil
}
