package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteExtensionInfoProperties struct {
	Authors                    *[]string          `json:"authors,omitempty"`
	Comment                    *string            `json:"comment,omitempty"`
	Description                *string            `json:"description,omitempty"`
	DownloadCount              *int64             `json:"download_count,omitempty"`
	ExtensionId                *string            `json:"extension_id,omitempty"`
	ExtensionType              *SiteExtensionType `json:"extension_type,omitempty"`
	ExtensionUrl               *string            `json:"extension_url,omitempty"`
	FeedUrl                    *string            `json:"feed_url,omitempty"`
	IconUrl                    *string            `json:"icon_url,omitempty"`
	InstalledDateTime          *string            `json:"installed_date_time,omitempty"`
	InstallerCommandLineParams *string            `json:"installer_command_line_params,omitempty"`
	LicenseUrl                 *string            `json:"license_url,omitempty"`
	LocalIsLatestVersion       *bool              `json:"local_is_latest_version,omitempty"`
	LocalPath                  *string            `json:"local_path,omitempty"`
	ProjectUrl                 *string            `json:"project_url,omitempty"`
	ProvisioningState          *string            `json:"provisioningState,omitempty"`
	PublishedDateTime          *string            `json:"published_date_time,omitempty"`
	Summary                    *string            `json:"summary,omitempty"`
	Title                      *string            `json:"title,omitempty"`
	Version                    *string            `json:"version,omitempty"`
}
