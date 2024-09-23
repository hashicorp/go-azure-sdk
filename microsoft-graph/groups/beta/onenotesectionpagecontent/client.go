package onenotesectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionPageContentClient: %+v", err)
	}

	return &OnenoteSectionPageContentClient{
		Client: client,
	}, nil
}
