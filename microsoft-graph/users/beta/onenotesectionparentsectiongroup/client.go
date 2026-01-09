package onenotesectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionParentSectionGroupClient: %+v", err)
	}

	return &OnenoteSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
