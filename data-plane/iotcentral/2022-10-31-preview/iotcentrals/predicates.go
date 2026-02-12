package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiTokenOperationPredicate struct {
	Expiry *string
	Id     *string
	Token  *string
}

func (p ApiTokenOperationPredicate) Matches(input ApiToken) bool {

	if p.Expiry != nil && (input.Expiry == nil || *p.Expiry != *input.Expiry) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Token != nil && (input.Token == nil || *p.Token != *input.Token) {
		return false
	}

	return true
}

type DashboardOperationPredicate struct {
	DisplayName *string
	Etag        *string
	Favorite    *bool
	Id          *string
	Personal    *bool
}

func (p DashboardOperationPredicate) Matches(input Dashboard) bool {

	if p.DisplayName != nil && *p.DisplayName != input.DisplayName {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Favorite != nil && (input.Favorite == nil || *p.Favorite != *input.Favorite) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Personal != nil && (input.Personal == nil || *p.Personal != *input.Personal) {
		return false
	}

	return true
}

type DeploymentManifestOperationPredicate struct {
	Data        *interface{}
	DisplayName *string
	Etag        *string
	Id          *string
}

func (p DeploymentManifestOperationPredicate) Matches(input DeploymentManifest) bool {

	if p.Data != nil && *p.Data != input.Data {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}

type DestinationOperationPredicate struct {
}

func (p DestinationOperationPredicate) Matches(input Destination) bool {

	return true
}

type DeviceOperationPredicate struct {
	DisplayName *string
	Enabled     *bool
	Etag        *string
	Id          *string
	Provisioned *bool
	Simulated   *bool
	Template    *string
}

func (p DeviceOperationPredicate) Matches(input Device) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Provisioned != nil && (input.Provisioned == nil || *p.Provisioned != *input.Provisioned) {
		return false
	}

	if p.Simulated != nil && (input.Simulated == nil || *p.Simulated != *input.Simulated) {
		return false
	}

	if p.Template != nil && (input.Template == nil || *p.Template != *input.Template) {
		return false
	}

	return true
}

type DeviceCommandOperationPredicate struct {
	ConnectionTimeout *int64
	Id                *string
	Request           *interface{}
	Response          *interface{}
	ResponseCode      *int64
	ResponseTimeout   *int64
}

func (p DeviceCommandOperationPredicate) Matches(input DeviceCommand) bool {

	if p.ConnectionTimeout != nil && (input.ConnectionTimeout == nil || *p.ConnectionTimeout != *input.ConnectionTimeout) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Request != nil && (input.Request == nil || *p.Request != *input.Request) {
		return false
	}

	if p.Response != nil && (input.Response == nil || *p.Response != *input.Response) {
		return false
	}

	if p.ResponseCode != nil && (input.ResponseCode == nil || *p.ResponseCode != *input.ResponseCode) {
		return false
	}

	if p.ResponseTimeout != nil && (input.ResponseTimeout == nil || *p.ResponseTimeout != *input.ResponseTimeout) {
		return false
	}

	return true
}

type DeviceGroupOperationPredicate struct {
	Description *string
	DisplayName *string
	Etag        *string
	Filter      *string
	Id          *string
}

func (p DeviceGroupOperationPredicate) Matches(input DeviceGroup) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && *p.DisplayName != input.DisplayName {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Filter != nil && *p.Filter != input.Filter {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}

type DeviceRelationshipOperationPredicate struct {
	Id     *string
	Source *string
	Target *string
}

func (p DeviceRelationshipOperationPredicate) Matches(input DeviceRelationship) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Source != nil && (input.Source == nil || *p.Source != *input.Source) {
		return false
	}

	if p.Target != nil && (input.Target == nil || *p.Target != *input.Target) {
		return false
	}

	return true
}

type DeviceTemplateOperationPredicate struct {
	CapabilityModel *interface{}
	Description     *string
	DisplayName     *string
	Etag            *string
	Id              *string
}

func (p DeviceTemplateOperationPredicate) Matches(input DeviceTemplate) bool {

	if p.CapabilityModel != nil && *p.CapabilityModel != input.CapabilityModel {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}

type EnrollmentGroupOperationPredicate struct {
	DisplayName *string
	Enabled     *bool
	Etag        *string
	Id          *string
	IdScope     *string
}

func (p EnrollmentGroupOperationPredicate) Matches(input EnrollmentGroup) bool {

	if p.DisplayName != nil && *p.DisplayName != input.DisplayName {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IdScope != nil && (input.IdScope == nil || *p.IdScope != *input.IdScope) {
		return false
	}

	return true
}

type ExportOperationPredicate struct {
	DisplayName    *string
	Enabled        *bool
	Filter         *string
	Id             *string
	LastExportTime *string
	Status         *string
}

func (p ExportOperationPredicate) Matches(input Export) bool {

	if p.DisplayName != nil && *p.DisplayName != input.DisplayName {
		return false
	}

	if p.Enabled != nil && *p.Enabled != input.Enabled {
		return false
	}

	if p.Filter != nil && (input.Filter == nil || *p.Filter != *input.Filter) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastExportTime != nil && (input.LastExportTime == nil || *p.LastExportTime != *input.LastExportTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}

type JobOperationPredicate struct {
	Description    *string
	DisplayName    *string
	End            *string
	Group          *string
	Id             *string
	ScheduledJobId *string
	Start          *string
	Status         *string
}

func (p JobOperationPredicate) Matches(input Job) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.End != nil && (input.End == nil || *p.End != *input.End) {
		return false
	}

	if p.Group != nil && *p.Group != input.Group {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ScheduledJobId != nil && (input.ScheduledJobId == nil || *p.ScheduledJobId != *input.ScheduledJobId) {
		return false
	}

	if p.Start != nil && (input.Start == nil || *p.Start != *input.Start) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}

type JobDeviceStatusOperationPredicate struct {
	Id     *string
	Status *string
}

func (p JobDeviceStatusOperationPredicate) Matches(input JobDeviceStatus) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}

type OrganizationOperationPredicate struct {
	DisplayName *string
	Id          *string
	Parent      *string
}

func (p OrganizationOperationPredicate) Matches(input Organization) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Parent != nil && (input.Parent == nil || *p.Parent != *input.Parent) {
		return false
	}

	return true
}

type RoleOperationPredicate struct {
	DisplayName *string
	Id          *string
}

func (p RoleOperationPredicate) Matches(input Role) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}

type ScheduledJobOperationPredicate struct {
	Completed   *bool
	Description *string
	DisplayName *string
	Enabled     *bool
	Etag        *string
	Group       *string
	Id          *string
}

func (p ScheduledJobOperationPredicate) Matches(input ScheduledJob) bool {

	if p.Completed != nil && (input.Completed == nil || *p.Completed != *input.Completed) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Group != nil && *p.Group != input.Group {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}

type UserOperationPredicate struct {
}

func (p UserOperationPredicate) Matches(input User) bool {

	return true
}
