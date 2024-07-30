package privilegedacces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccesClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccesClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccesClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedacces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccesClient: %+v", err)
	}

	return &PrivilegedAccesClient{
		Client: client,
	}, nil
}
