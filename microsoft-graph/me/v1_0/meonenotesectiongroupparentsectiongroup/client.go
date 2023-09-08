package meonenotesectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &MeOnenoteSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
