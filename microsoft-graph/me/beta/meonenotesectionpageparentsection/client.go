package meonenotesectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionPageParentSectionClient: %+v", err)
	}

	return &MeOnenoteSectionPageParentSectionClient{
		Client: client,
	}, nil
}
