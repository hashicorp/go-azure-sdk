package joinedteamphoto

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

type SetJoinedTeamPhotoValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetJoinedTeamPhotoValueOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetJoinedTeamPhotoValueOperationOptions() SetJoinedTeamPhotoValueOperationOptions {
	return SetJoinedTeamPhotoValueOperationOptions{}
}

func (o SetJoinedTeamPhotoValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetJoinedTeamPhotoValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetJoinedTeamPhotoValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetJoinedTeamPhotoValue - Update media content for the navigation property photo in users. The profile photo for the
// team.
func (c JoinedTeamPhotoClient) SetJoinedTeamPhotoValue(ctx context.Context, id stable.UserIdJoinedTeamId, input []byte, options SetJoinedTeamPhotoValueOperationOptions) (result SetJoinedTeamPhotoValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/photo/$value", id.ID()),
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
