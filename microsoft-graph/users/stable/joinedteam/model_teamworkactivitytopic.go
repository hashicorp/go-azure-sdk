package joinedteam

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkActivityTopic struct {
	ODataType *string                      `json:"@odata.type,omitempty"`
	Source    *TeamworkActivityTopicSource `json:"source,omitempty"`
	Value     *string                      `json:"value,omitempty"`
	WebUrl    *string                      `json:"webUrl,omitempty"`
}
