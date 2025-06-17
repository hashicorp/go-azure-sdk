package cloudcertificationauthorityleafcertificate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityLeafCertificateClient struct {
	Client *msgraph.Client
}

func NewCloudCertificationAuthorityLeafCertificateClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudCertificationAuthorityLeafCertificateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudcertificationauthorityleafcertificate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudCertificationAuthorityLeafCertificateClient: %+v", err)
	}

	return &CloudCertificationAuthorityLeafCertificateClient{
		Client: client,
	}, nil
}
