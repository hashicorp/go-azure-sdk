package groupsiteonenotesectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotesectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteSectionPageClient: %+v", err)
	}

	return &GroupSiteOnenoteSectionPageClient{
		Client: client,
	}, nil
}
