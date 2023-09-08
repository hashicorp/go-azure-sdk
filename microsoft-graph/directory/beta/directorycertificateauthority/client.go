package directorycertificateauthority

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryCertificateAuthorityClient struct {
	Client *msgraph.Client
}

func NewDirectoryCertificateAuthorityClientWithBaseURI(api sdkEnv.Api) (*DirectoryCertificateAuthorityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directorycertificateauthority", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryCertificateAuthorityClient: %+v", err)
	}

	return &DirectoryCertificateAuthorityClient{
		Client: client,
	}, nil
}
