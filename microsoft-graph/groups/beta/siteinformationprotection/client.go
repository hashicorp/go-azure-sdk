package siteinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionClient: %+v", err)
	}

	return &SiteInformationProtectionClient{
		Client: client,
	}, nil
}
