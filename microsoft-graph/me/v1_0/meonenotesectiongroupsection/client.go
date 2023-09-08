package meonenotesectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionClient{
		Client: client,
	}, nil
}
