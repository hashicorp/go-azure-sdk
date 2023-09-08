package meonenotesectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionPageContentClient: %+v", err)
	}

	return &MeOnenoteSectionPageContentClient{
		Client: client,
	}, nil
}
