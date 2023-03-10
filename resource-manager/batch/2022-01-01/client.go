package v2022_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/applicationpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/batchaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/batchmanagements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/certificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/location"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/pool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2022-01-01/privatelinkresource"
)

type Client struct {
	Application               *application.ApplicationClient
	ApplicationPackage        *applicationpackage.ApplicationPackageClient
	BatchAccount              *batchaccount.BatchAccountClient
	BatchManagements          *batchmanagements.BatchManagementsClient
	Certificate               *certificate.CertificateClient
	Location                  *location.LocationClient
	Pool                      *pool.PoolClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applicationClient := application.NewApplicationClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationClient.Client)

	applicationPackageClient := applicationpackage.NewApplicationPackageClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationPackageClient.Client)

	batchAccountClient := batchaccount.NewBatchAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&batchAccountClient.Client)

	batchManagementsClient := batchmanagements.NewBatchManagementsClientWithBaseURI(endpoint)
	configureAuthFunc(&batchManagementsClient.Client)

	certificateClient := certificate.NewCertificateClientWithBaseURI(endpoint)
	configureAuthFunc(&certificateClient.Client)

	locationClient := location.NewLocationClientWithBaseURI(endpoint)
	configureAuthFunc(&locationClient.Client)

	poolClient := pool.NewPoolClientWithBaseURI(endpoint)
	configureAuthFunc(&poolClient.Client)

	privateEndpointConnectionClient := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionClient.Client)

	privateLinkResourceClient := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourceClient.Client)

	return Client{
		Application:               &applicationClient,
		ApplicationPackage:        &applicationPackageClient,
		BatchAccount:              &batchAccountClient,
		BatchManagements:          &batchManagementsClient,
		Certificate:               &certificateClient,
		Location:                  &locationClient,
		Pool:                      &poolClient,
		PrivateEndpointConnection: &privateEndpointConnectionClient,
		PrivateLinkResource:       &privateLinkResourceClient,
	}
}
