package driveroot

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

type ListDriveRootInvitesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Permission
}

type ListDriveRootInvitesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Permission
}

type ListDriveRootInvitesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListDriveRootInvitesOperationOptions() ListDriveRootInvitesOperationOptions {
	return ListDriveRootInvitesOperationOptions{}
}

func (o ListDriveRootInvitesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootInvitesOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootInvitesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootInvitesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootInvitesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootInvites - Invoke action invite. Sends a sharing invitation for a driveItem. A sharing invitation
// provides permissions to the recipients and optionally sends them an email with a sharing link.
func (c DriveRootClient) ListDriveRootInvites(ctx context.Context, id stable.MeDriveId, input ListDriveRootInvitesRequest, options ListDriveRootInvitesOperationOptions) (result ListDriveRootInvitesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveRootInvitesCustomPager{},
		Path:          fmt.Sprintf("%s/root/invite", id.ID()),
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

// ListDriveRootInvitesComplete retrieves all the results into a single object
func (c DriveRootClient) ListDriveRootInvitesComplete(ctx context.Context, id stable.MeDriveId, input ListDriveRootInvitesRequest, options ListDriveRootInvitesOperationOptions) (ListDriveRootInvitesCompleteResult, error) {
	return c.ListDriveRootInvitesCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveRootInvitesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootClient) ListDriveRootInvitesCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveId, input ListDriveRootInvitesRequest, options ListDriveRootInvitesOperationOptions, predicate PermissionOperationPredicate) (result ListDriveRootInvitesCompleteResult, err error) {
	items := make([]stable.Permission, 0)

	resp, err := c.ListDriveRootInvites(ctx, id, input, options)
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

	result = ListDriveRootInvitesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
