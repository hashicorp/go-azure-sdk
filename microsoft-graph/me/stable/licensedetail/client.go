package licensedetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseDetailClient struct {
	Client *msgraph.Client
}

func NewLicenseDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*LicenseDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "licensedetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LicenseDetailClient: %+v", err)
	}

	return &LicenseDetailClient{
		Client: client,
	}, nil
}
