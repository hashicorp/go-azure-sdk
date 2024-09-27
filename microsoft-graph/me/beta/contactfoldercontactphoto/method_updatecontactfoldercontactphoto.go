package contactfoldercontactphoto

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

type UpdateContactFolderContactPhotoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateContactFolderContactPhotoOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateContactFolderContactPhotoOperationOptions() UpdateContactFolderContactPhotoOperationOptions {
	return UpdateContactFolderContactPhotoOperationOptions{}
}

func (o UpdateContactFolderContactPhotoOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateContactFolderContactPhotoOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateContactFolderContactPhotoOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateContactFolderContactPhoto - Update the navigation property photo in me
func (c ContactFolderContactPhotoClient) UpdateContactFolderContactPhoto(ctx context.Context, id beta.MeContactFolderIdContactId, input beta.ProfilePhoto, options UpdateContactFolderContactPhotoOperationOptions) (result UpdateContactFolderContactPhotoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/photo", id.ID()),
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
