package v2022_03_03

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	communityGalleriesClient, err := communitygalleries.NewCommunityGalleriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunityGalleries client: %+v", err)
	}
	configureFunc(communityGalleriesClient.Client)

	communityGalleryImageVersionsClient, err := communitygalleryimageversions.NewCommunityGalleryImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunityGalleryImageVersions client: %+v", err)
	}
	configureFunc(communityGalleryImageVersionsClient.Client)

	communityGalleryImagesClient, err := communitygalleryimages.NewCommunityGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunityGalleryImages client: %+v", err)
	}
	configureFunc(communityGalleryImagesClient.Client)

	galleriesClient, err := galleries.NewGalleriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Galleries client: %+v", err)
	}
	configureFunc(galleriesClient.Client)

	galleryApplicationVersionsClient, err := galleryapplicationversions.NewGalleryApplicationVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryApplicationVersions client: %+v", err)
	}
	configureFunc(galleryApplicationVersionsClient.Client)

	galleryApplicationsClient, err := galleryapplications.NewGalleryApplicationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryApplications client: %+v", err)
	}
	configureFunc(galleryApplicationsClient.Client)

	galleryImageVersionsClient, err := galleryimageversions.NewGalleryImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImageVersions client: %+v", err)
	}
	configureFunc(galleryImageVersionsClient.Client)

	galleryImagesClient, err := galleryimages.NewGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImages client: %+v", err)
	}
	configureFunc(galleryImagesClient.Client)

	gallerySharingUpdateClient, err := gallerysharingupdate.NewGallerySharingUpdateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GallerySharingUpdate client: %+v", err)
	}
	configureFunc(gallerySharingUpdateClient.Client)

	sharedGalleriesClient, err := sharedgalleries.NewSharedGalleriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedGalleries client: %+v", err)
	}
	configureFunc(sharedGalleriesClient.Client)

	sharedGalleryImageVersionsClient, err := sharedgalleryimageversions.NewSharedGalleryImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedGalleryImageVersions client: %+v", err)
	}
	configureFunc(sharedGalleryImageVersionsClient.Client)

	sharedGalleryImagesClient, err := sharedgalleryimages.NewSharedGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedGalleryImages client: %+v", err)
	}
	configureFunc(sharedGalleryImagesClient.Client)

	return &Client{
		CommunityGalleries:            communityGalleriesClient,
		CommunityGalleryImageVersions: communityGalleryImageVersionsClient,
		CommunityGalleryImages:        communityGalleryImagesClient,
		Galleries:                     galleriesClient,
		GalleryApplicationVersions:    galleryApplicationVersionsClient,
		GalleryApplications:           galleryApplicationsClient,
		GalleryImageVersions:          galleryImageVersionsClient,
		GalleryImages:                 galleryImagesClient,
		GallerySharingUpdate:          gallerySharingUpdateClient,
		SharedGalleries:               sharedGalleriesClient,
		SharedGalleryImageVersions:    sharedGalleryImageVersionsClient,
		SharedGalleryImages:           sharedGalleryImagesClient,
	}, nil
}
