package privilegedaccess

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccess", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessClient: %+v", err)
	}

	return &PrivilegedAccessClient{
		Client: client,
	}, nil
}
