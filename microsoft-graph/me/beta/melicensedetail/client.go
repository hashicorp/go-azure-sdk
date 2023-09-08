package melicensedetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeLicenseDetailClient struct {
	Client *msgraph.Client
}

func NewMeLicenseDetailClientWithBaseURI(api sdkEnv.Api) (*MeLicenseDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "melicensedetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeLicenseDetailClient: %+v", err)
	}

	return &MeLicenseDetailClient{
		Client: client,
	}, nil
}
