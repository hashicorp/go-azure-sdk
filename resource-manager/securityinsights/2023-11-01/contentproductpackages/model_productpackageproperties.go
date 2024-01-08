package contentproductpackages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductPackageProperties struct {
	Author                   *MetadataAuthor       `json:"author,omitempty"`
	Categories               *MetadataCategories   `json:"categories,omitempty"`
	ContentId                string                `json:"contentId"`
	ContentKind              PackageKind           `json:"contentKind"`
	ContentProductId         *string               `json:"contentProductId,omitempty"`
	ContentSchemaVersion     *string               `json:"contentSchemaVersion,omitempty"`
	Dependencies             *MetadataDependencies `json:"dependencies,omitempty"`
	Description              *string               `json:"description,omitempty"`
	DisplayName              string                `json:"displayName"`
	FirstPublishDate         *string               `json:"firstPublishDate,omitempty"`
	Icon                     *string               `json:"icon,omitempty"`
	InstalledVersion         *string               `json:"installedVersion,omitempty"`
	IsDeprecated             *Flag                 `json:"isDeprecated,omitempty"`
	IsFeatured               *Flag                 `json:"isFeatured,omitempty"`
	IsNew                    *Flag                 `json:"isNew,omitempty"`
	IsPreview                *Flag                 `json:"isPreview,omitempty"`
	LastPublishDate          *string               `json:"lastPublishDate,omitempty"`
	MetadataResourceId       *string               `json:"metadataResourceId,omitempty"`
	PackagedContent          *interface{}          `json:"packagedContent,omitempty"`
	Providers                *[]string             `json:"providers,omitempty"`
	PublisherDisplayName     *string               `json:"publisherDisplayName,omitempty"`
	Source                   *MetadataSource       `json:"source,omitempty"`
	Support                  *MetadataSupport      `json:"support,omitempty"`
	ThreatAnalysisTactics    *[]string             `json:"threatAnalysisTactics,omitempty"`
	ThreatAnalysisTechniques *[]string             `json:"threatAnalysisTechniques,omitempty"`
	Version                  string                `json:"version"`
}
