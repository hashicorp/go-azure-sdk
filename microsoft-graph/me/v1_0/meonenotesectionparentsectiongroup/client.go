package meonenotesectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionParentSectionGroupClient: %+v", err)
	}

	return &MeOnenoteSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
