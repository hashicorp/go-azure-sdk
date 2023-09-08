package meonenotesectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
