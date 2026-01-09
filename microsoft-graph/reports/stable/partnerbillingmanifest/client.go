package partnerbillingmanifest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerBillingManifestClient struct {
	Client *msgraph.Client
}

func NewPartnerBillingManifestClientWithBaseURI(sdkApi sdkEnv.Api) (*PartnerBillingManifestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "partnerbillingmanifest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PartnerBillingManifestClient: %+v", err)
	}

	return &PartnerBillingManifestClient{
		Client: client,
	}, nil
}
