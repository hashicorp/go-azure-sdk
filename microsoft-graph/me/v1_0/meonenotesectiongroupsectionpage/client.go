package meonenotesectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionPageClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
