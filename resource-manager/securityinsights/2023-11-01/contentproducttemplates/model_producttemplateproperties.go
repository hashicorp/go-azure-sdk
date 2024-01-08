package contentproducttemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductTemplateProperties struct {
	Author                   *MetadataAuthor       `json:"author,omitempty"`
	Categories               *MetadataCategories   `json:"categories,omitempty"`
	ContentId                string                `json:"contentId"`
	ContentKind              Kind                  `json:"contentKind"`
	ContentProductId         *string               `json:"contentProductId,omitempty"`
	ContentSchemaVersion     *string               `json:"contentSchemaVersion,omitempty"`
	CustomVersion            *string               `json:"customVersion,omitempty"`
	Dependencies             *MetadataDependencies `json:"dependencies,omitempty"`
	DisplayName              string                `json:"displayName"`
	FirstPublishDate         *string               `json:"firstPublishDate,omitempty"`
	Icon                     *string               `json:"icon,omitempty"`
	IsDeprecated             *Flag                 `json:"isDeprecated,omitempty"`
	LastPublishDate          *string               `json:"lastPublishDate,omitempty"`
	PackageId                *string               `json:"packageId,omitempty"`
	PackageKind              *PackageKind          `json:"packageKind,omitempty"`
	PackageName              *string               `json:"packageName,omitempty"`
	PackageVersion           *string               `json:"packageVersion,omitempty"`
	PackagedContent          *interface{}          `json:"packagedContent,omitempty"`
	PreviewImages            *[]string             `json:"previewImages,omitempty"`
	PreviewImagesDark        *[]string             `json:"previewImagesDark,omitempty"`
	Providers                *[]string             `json:"providers,omitempty"`
	Source                   MetadataSource        `json:"source"`
	Support                  *MetadataSupport      `json:"support,omitempty"`
	ThreatAnalysisTactics    *[]string             `json:"threatAnalysisTactics,omitempty"`
	ThreatAnalysisTechniques *[]string             `json:"threatAnalysisTechniques,omitempty"`
	Version                  string                `json:"version"`
}
