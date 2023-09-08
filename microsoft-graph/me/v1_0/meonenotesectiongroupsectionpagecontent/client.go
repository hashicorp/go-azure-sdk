package meonenotesectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionPageContentClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
