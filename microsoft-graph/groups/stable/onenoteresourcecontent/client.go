package onenoteresourcecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteResourceContentClient struct {
	Client *msgraph.Client
}

func NewOnenoteResourceContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteResourceContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenoteresourcecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteResourceContentClient: %+v", err)
	}

	return &OnenoteResourceContentClient{
		Client: client,
	}, nil
}
