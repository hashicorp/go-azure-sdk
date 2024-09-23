package onenotesectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionGroupClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
