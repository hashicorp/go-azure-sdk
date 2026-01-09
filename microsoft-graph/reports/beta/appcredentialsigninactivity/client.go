package appcredentialsigninactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivityClient struct {
	Client *msgraph.Client
}

func NewAppCredentialSignInActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*AppCredentialSignInActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appcredentialsigninactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppCredentialSignInActivityClient: %+v", err)
	}

	return &AppCredentialSignInActivityClient{
		Client: client,
	}, nil
}
