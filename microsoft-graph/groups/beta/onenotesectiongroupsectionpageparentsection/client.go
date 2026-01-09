package onenotesectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
