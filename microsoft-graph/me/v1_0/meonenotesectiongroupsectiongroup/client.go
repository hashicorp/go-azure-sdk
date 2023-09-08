package meonenotesectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionGroupClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
