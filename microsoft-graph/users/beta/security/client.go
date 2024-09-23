package security

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityClient struct {
	Client *msgraph.Client
}

func NewSecurityClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "security", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityClient: %+v", err)
	}

	return &SecurityClient{
		Client: client,
	}, nil
}
