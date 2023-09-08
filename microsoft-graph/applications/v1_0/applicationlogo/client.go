package applicationlogo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationLogoClient struct {
	Client *msgraph.Client
}

func NewApplicationLogoClientWithBaseURI(api sdkEnv.Api) (*ApplicationLogoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationlogo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationLogoClient: %+v", err)
	}

	return &ApplicationLogoClient{
		Client: client,
	}, nil
}
