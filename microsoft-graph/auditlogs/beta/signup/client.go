package signup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignUpClient struct {
	Client *msgraph.Client
}

func NewSignUpClientWithBaseURI(sdkApi sdkEnv.Api) (*SignUpClient, error) {
	client, err := msgraph.NewClient(sdkApi, "signup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SignUpClient: %+v", err)
	}

	return &SignUpClient{
		Client: client,
	}, nil
}
