package siteinformationprotectionbitlocker

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionBitlockerClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionBitlockerClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionBitlockerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionbitlocker", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionBitlockerClient: %+v", err)
	}

	return &SiteInformationProtectionBitlockerClient{
		Client: client,
	}, nil
}
