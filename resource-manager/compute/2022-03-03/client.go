package v2022_03_03

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/communitygalleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/communitygalleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/communitygalleryimageversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleryapplications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleryapplicationversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/galleryimageversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/gallerysharingupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/sharedgalleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/sharedgalleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/sharedgalleryimageversions"
)

type Client struct {
	CommunityGalleries            *communitygalleries.CommunityGalleriesClient
	CommunityGalleryImageVersions *communitygalleryimageversions.CommunityGalleryImageVersionsClient
	CommunityGalleryImages        *communitygalleryimages.CommunityGalleryImagesClient
	Galleries                     *galleries.GalleriesClient
	GalleryApplicationVersions    *galleryapplicationversions.GalleryApplicationVersionsClient
	GalleryApplications           *galleryapplications.GalleryApplicationsClient
	GalleryImageVersions          *galleryimageversions.GalleryImageVersionsClient
	GalleryImages                 *galleryimages.GalleryImagesClient
	GallerySharingUpdate          *gallerysharingupdate.GallerySharingUpdateClient
	SharedGalleries               *sharedgalleries.SharedGalleriesClient
	SharedGalleryImageVersions    *sharedgalleryimageversions.SharedGalleryImageVersionsClient
	SharedGalleryImages           *sharedgalleryimages.SharedGalleryImagesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	communityGalleriesClient := communitygalleries.NewCommunityGalleriesClientWithBaseURI(endpoint)
	configureAuthFunc(&communityGalleriesClient.Client)

	communityGalleryImageVersionsClient := communitygalleryimageversions.NewCommunityGalleryImageVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&communityGalleryImageVersionsClient.Client)

	communityGalleryImagesClient := communitygalleryimages.NewCommunityGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&communityGalleryImagesClient.Client)

	galleriesClient := galleries.NewGalleriesClientWithBaseURI(endpoint)
	configureAuthFunc(&galleriesClient.Client)

	galleryApplicationVersionsClient := galleryapplicationversions.NewGalleryApplicationVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryApplicationVersionsClient.Client)

	galleryApplicationsClient := galleryapplications.NewGalleryApplicationsClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryApplicationsClient.Client)

	galleryImageVersionsClient := galleryimageversions.NewGalleryImageVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryImageVersionsClient.Client)

	galleryImagesClient := galleryimages.NewGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryImagesClient.Client)

	gallerySharingUpdateClient := gallerysharingupdate.NewGallerySharingUpdateClientWithBaseURI(endpoint)
	configureAuthFunc(&gallerySharingUpdateClient.Client)

	sharedGalleriesClient := sharedgalleries.NewSharedGalleriesClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedGalleriesClient.Client)

	sharedGalleryImageVersionsClient := sharedgalleryimageversions.NewSharedGalleryImageVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedGalleryImageVersionsClient.Client)

	sharedGalleryImagesClient := sharedgalleryimages.NewSharedGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedGalleryImagesClient.Client)

	return Client{
		CommunityGalleries:            &communityGalleriesClient,
		CommunityGalleryImageVersions: &communityGalleryImageVersionsClient,
		CommunityGalleryImages:        &communityGalleryImagesClient,
		Galleries:                     &galleriesClient,
		GalleryApplicationVersions:    &galleryApplicationVersionsClient,
		GalleryApplications:           &galleryApplicationsClient,
		GalleryImageVersions:          &galleryImageVersionsClient,
		GalleryImages:                 &galleryImagesClient,
		GallerySharingUpdate:          &gallerySharingUpdateClient,
		SharedGalleries:               &sharedGalleriesClient,
		SharedGalleryImageVersions:    &sharedGalleryImageVersionsClient,
		SharedGalleryImages:           &sharedGalleryImagesClient,
	}
}
