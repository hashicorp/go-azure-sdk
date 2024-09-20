package onenotesection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionClient: %+v", err)
	}

	return &OnenoteSectionClient{
		Client: client,
	}, nil
}
