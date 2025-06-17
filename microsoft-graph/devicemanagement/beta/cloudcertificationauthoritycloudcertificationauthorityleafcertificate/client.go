package cloudcertificationauthoritycloudcertificationauthorityleafcertificate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient struct {
	Client *msgraph.Client
}

func NewCloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudcertificationauthoritycloudcertificationauthorityleafcertificate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient: %+v", err)
	}

	return &CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient{
		Client: client,
	}, nil
}
