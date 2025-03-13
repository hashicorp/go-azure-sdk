package v2024_04_03

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/appattachpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/appattachpackageinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/application"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/applicationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/desktop"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/hostpool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/msiximage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/msixpackage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/privatelink"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/scalingplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/scalingplanpersonalschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/scalingplanpooledschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/sessionhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/startmenuitem"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/usersession"
	"github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-03/workspace"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppAttachPackage            *appattachpackage.AppAttachPackageClient
	AppAttachPackageInfo        *appattachpackageinfo.AppAttachPackageInfoClient
	Application                 *application.ApplicationClient
	ApplicationGroup            *applicationgroup.ApplicationGroupClient
	Desktop                     *desktop.DesktopClient
	HostPool                    *hostpool.HostPoolClient
	MSIXPackage                 *msixpackage.MSIXPackageClient
	MsixImage                   *msiximage.MsixImageClient
	PrivateLink                 *privatelink.PrivateLinkClient
	ScalingPlan                 *scalingplan.ScalingPlanClient
	ScalingPlanPersonalSchedule *scalingplanpersonalschedule.ScalingPlanPersonalScheduleClient
	ScalingPlanPooledSchedule   *scalingplanpooledschedule.ScalingPlanPooledScheduleClient
	SessionHost                 *sessionhost.SessionHostClient
	StartMenuItem               *startmenuitem.StartMenuItemClient
	UserSession                 *usersession.UserSessionClient
	Workspace                   *workspace.WorkspaceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appAttachPackageClient, err := appattachpackage.NewAppAttachPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppAttachPackage client: %+v", err)
	}
	configureFunc(appAttachPackageClient.Client)

	appAttachPackageInfoClient, err := appattachpackageinfo.NewAppAttachPackageInfoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppAttachPackageInfo client: %+v", err)
	}
	configureFunc(appAttachPackageInfoClient.Client)

	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	applicationGroupClient, err := applicationgroup.NewApplicationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGroup client: %+v", err)
	}
	configureFunc(applicationGroupClient.Client)

	desktopClient, err := desktop.NewDesktopClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Desktop client: %+v", err)
	}
	configureFunc(desktopClient.Client)

	hostPoolClient, err := hostpool.NewHostPoolClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HostPool client: %+v", err)
	}
	configureFunc(hostPoolClient.Client)

	mSIXPackageClient, err := msixpackage.NewMSIXPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MSIXPackage client: %+v", err)
	}
	configureFunc(mSIXPackageClient.Client)

	msixImageClient, err := msiximage.NewMsixImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MsixImage client: %+v", err)
	}
	configureFunc(msixImageClient.Client)

	privateLinkClient, err := privatelink.NewPrivateLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLink client: %+v", err)
	}
	configureFunc(privateLinkClient.Client)

	scalingPlanClient, err := scalingplan.NewScalingPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlan client: %+v", err)
	}
	configureFunc(scalingPlanClient.Client)

	scalingPlanPersonalScheduleClient, err := scalingplanpersonalschedule.NewScalingPlanPersonalScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlanPersonalSchedule client: %+v", err)
	}
	configureFunc(scalingPlanPersonalScheduleClient.Client)

	scalingPlanPooledScheduleClient, err := scalingplanpooledschedule.NewScalingPlanPooledScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlanPooledSchedule client: %+v", err)
	}
	configureFunc(scalingPlanPooledScheduleClient.Client)

	sessionHostClient, err := sessionhost.NewSessionHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SessionHost client: %+v", err)
	}
	configureFunc(sessionHostClient.Client)

	startMenuItemClient, err := startmenuitem.NewStartMenuItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StartMenuItem client: %+v", err)
	}
	configureFunc(startMenuItemClient.Client)

	userSessionClient, err := usersession.NewUserSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSession client: %+v", err)
	}
	configureFunc(userSessionClient.Client)

	workspaceClient, err := workspace.NewWorkspaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspace client: %+v", err)
	}
	configureFunc(workspaceClient.Client)

	return &Client{
		AppAttachPackage:            appAttachPackageClient,
		AppAttachPackageInfo:        appAttachPackageInfoClient,
		Application:                 applicationClient,
		ApplicationGroup:            applicationGroupClient,
		Desktop:                     desktopClient,
		HostPool:                    hostPoolClient,
		MSIXPackage:                 mSIXPackageClient,
		MsixImage:                   msixImageClient,
		PrivateLink:                 privateLinkClient,
		ScalingPlan:                 scalingPlanClient,
		ScalingPlanPersonalSchedule: scalingPlanPersonalScheduleClient,
		ScalingPlanPooledSchedule:   scalingPlanPooledScheduleClient,
		SessionHost:                 sessionHostClient,
		StartMenuItem:               startMenuItemClient,
		UserSession:                 userSessionClient,
		Workspace:                   workspaceClient,
	}, nil
}
