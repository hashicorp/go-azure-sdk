package driveitem

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

type ListDriveItemInvitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Permission
}

type ListDriveItemInvitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Permission
}

type ListDriveItemInvitesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListDriveItemInvitesOperationOptions() ListDriveItemInvitesOperationOptions {
	return ListDriveItemInvitesOperationOptions{}
}

func (o ListDriveItemInvitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemInvitesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDriveItemInvitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemInvitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemInvitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemInvites - Invoke action invite. Sends a sharing invitation for a driveItem. A sharing invitation
// provides permissions to the recipients and optionally sends them an email with a sharing link.
func (c DriveItemClient) ListDriveItemInvites(ctx context.Context, id stable.GroupIdDriveIdItemId, input ListDriveItemInvitesRequest, options ListDriveItemInvitesOperationOptions) (result ListDriveItemInvitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveItemInvitesCustomPager{},
		Path:          fmt.Sprintf("%s/invite", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveItemInvitesComplete retrieves all the results into a single object
func (c DriveItemClient) ListDriveItemInvitesComplete(ctx context.Context, id stable.GroupIdDriveIdItemId, input ListDriveItemInvitesRequest, options ListDriveItemInvitesOperationOptions) (ListDriveItemInvitesCompleteResult, error) {
	return c.ListDriveItemInvitesCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveItemInvitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemClient) ListDriveItemInvitesCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdDriveIdItemId, input ListDriveItemInvitesRequest, options ListDriveItemInvitesOperationOptions, predicate PermissionOperationPredicate) (result ListDriveItemInvitesCompleteResult, err error) {
	items := make([]stable.Permission, 0)

	resp, err := c.ListDriveItemInvites(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListDriveItemInvitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
