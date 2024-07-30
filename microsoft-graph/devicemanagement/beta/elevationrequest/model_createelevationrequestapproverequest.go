package elevationrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateElevationRequestApproveRequest struct {
	ReviewerJustification *string `json:"reviewerJustification,omitempty"`
}
