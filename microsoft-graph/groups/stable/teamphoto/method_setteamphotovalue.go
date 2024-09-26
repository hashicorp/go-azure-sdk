package teamphoto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetTeamPhotoValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetTeamPhotoValueOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetTeamPhotoValueOperationOptions() SetTeamPhotoValueOperationOptions {
	return SetTeamPhotoValueOperationOptions{}
}

func (o SetTeamPhotoValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetTeamPhotoValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetTeamPhotoValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetTeamPhotoValue - Update media content for the navigation property photo in groups. The profile photo for the team.
func (c TeamPhotoClient) SetTeamPhotoValue(ctx context.Context, id stable.GroupId, input []byte, options SetTeamPhotoValueOperationOptions) (result SetTeamPhotoValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/photo/$value", id.ID()),
		RetryFunc:     options.RetryFunc,
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
