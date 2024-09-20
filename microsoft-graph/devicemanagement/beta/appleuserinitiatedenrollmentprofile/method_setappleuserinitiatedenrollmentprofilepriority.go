package appleuserinitiatedenrollmentprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetAppleUserInitiatedEnrollmentProfilePriorityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions() SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions {
	return SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions{}
}

func (o SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetAppleUserInitiatedEnrollmentProfilePriority - Invoke action setPriority
func (c AppleUserInitiatedEnrollmentProfileClient) SetAppleUserInitiatedEnrollmentProfilePriority(ctx context.Context, id beta.DeviceManagementAppleUserInitiatedEnrollmentProfileId, input SetAppleUserInitiatedEnrollmentProfilePriorityRequest, options SetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions) (result SetAppleUserInitiatedEnrollmentProfilePriorityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/setPriority", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
