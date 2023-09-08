package meonenotepageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenotePageParentSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenotePageParentSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenotePageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotepageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenotePageParentSectionClient: %+v", err)
	}

	return &MeOnenotePageParentSectionClient{
		Client: client,
	}, nil
}
