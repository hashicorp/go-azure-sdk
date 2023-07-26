package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTEdgeAgentInfo struct {
	ImageName       string                     `json:"imageName"`
	ImageRepository *ImageRepositoryCredential `json:"imageRepository,omitempty"`
	Tag             string                     `json:"tag"`
}
