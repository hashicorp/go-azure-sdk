// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_09_09

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/applicationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/desktop"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/msiximage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/msixpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/scalingplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/scalingplanpooledschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/sessionhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/startmenuitem"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/usersession"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/workspace"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Application               *application.ApplicationClient
	ApplicationGroup          *applicationgroup.ApplicationGroupClient
	Desktop                   *desktop.DesktopClient
	HostPool                  *hostpool.HostPoolClient
	MSIXPackage               *msixpackage.MSIXPackageClient
	MsixImage                 *msiximage.MsixImageClient
	ScalingPlan               *scalingplan.ScalingPlanClient
	ScalingPlanPooledSchedule *scalingplanpooledschedule.ScalingPlanPooledScheduleClient
	SessionHost               *sessionhost.SessionHostClient
	StartMenuItem             *startmenuitem.StartMenuItemClient
	UserSession               *usersession.UserSessionClient
	Workspace                 *workspace.WorkspaceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applicationClient := application.NewApplicationClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationClient.Client)

	applicationGroupClient := applicationgroup.NewApplicationGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationGroupClient.Client)

	desktopClient := desktop.NewDesktopClientWithBaseURI(endpoint)
	configureAuthFunc(&desktopClient.Client)

	hostPoolClient := hostpool.NewHostPoolClientWithBaseURI(endpoint)
	configureAuthFunc(&hostPoolClient.Client)

	mSIXPackageClient := msixpackage.NewMSIXPackageClientWithBaseURI(endpoint)
	configureAuthFunc(&mSIXPackageClient.Client)

	msixImageClient := msiximage.NewMsixImageClientWithBaseURI(endpoint)
	configureAuthFunc(&msixImageClient.Client)

	scalingPlanClient := scalingplan.NewScalingPlanClientWithBaseURI(endpoint)
	configureAuthFunc(&scalingPlanClient.Client)

	scalingPlanPooledScheduleClient := scalingplanpooledschedule.NewScalingPlanPooledScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&scalingPlanPooledScheduleClient.Client)

	sessionHostClient := sessionhost.NewSessionHostClientWithBaseURI(endpoint)
	configureAuthFunc(&sessionHostClient.Client)

	startMenuItemClient := startmenuitem.NewStartMenuItemClientWithBaseURI(endpoint)
	configureAuthFunc(&startMenuItemClient.Client)

	userSessionClient := usersession.NewUserSessionClientWithBaseURI(endpoint)
	configureAuthFunc(&userSessionClient.Client)

	workspaceClient := workspace.NewWorkspaceClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceClient.Client)

	return Client{
		Application:               &applicationClient,
		ApplicationGroup:          &applicationGroupClient,
		Desktop:                   &desktopClient,
		HostPool:                  &hostPoolClient,
		MSIXPackage:               &mSIXPackageClient,
		MsixImage:                 &msixImageClient,
		ScalingPlan:               &scalingPlanClient,
		ScalingPlanPooledSchedule: &scalingPlanPooledScheduleClient,
		SessionHost:               &sessionHostClient,
		StartMenuItem:             &startMenuItemClient,
		UserSession:               &userSessionClient,
		Workspace:                 &workspaceClient,
	}
}
