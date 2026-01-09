package certificateauthority

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityClient struct {
	Client *msgraph.Client
}

func NewCertificateAuthorityClientWithBaseURI(sdkApi sdkEnv.Api) (*CertificateAuthorityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "certificateauthority", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CertificateAuthorityClient: %+v", err)
	}

	return &CertificateAuthorityClient{
		Client: client,
	}, nil
}
