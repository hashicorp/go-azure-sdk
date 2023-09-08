package meonenotesectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionPageClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionPageClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionPageClient: %+v", err)
	}

	return &MeOnenoteSectionPageClient{
		Client: client,
	}, nil
}
