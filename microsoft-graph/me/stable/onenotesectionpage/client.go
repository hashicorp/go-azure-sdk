package onenotesectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionPageClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionPageClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionPageClient: %+v", err)
	}

	return &OnenoteSectionPageClient{
		Client: client,
	}, nil
}
