package onenotesectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionPageContentClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
