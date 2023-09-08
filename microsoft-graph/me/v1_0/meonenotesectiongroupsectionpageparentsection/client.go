package meonenotesectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
