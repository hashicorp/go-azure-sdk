package invitedusersponsor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInvitedUserSponsorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListInvitedUserSponsorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListInvitedUserSponsorsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListInvitedUserSponsorsOperationOptions() ListInvitedUserSponsorsOperationOptions {
	return ListInvitedUserSponsorsOperationOptions{}
}

func (o ListInvitedUserSponsorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListInvitedUserSponsorsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListInvitedUserSponsorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListInvitedUserSponsorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInvitedUserSponsorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInvitedUserSponsors - Get invitedUserSponsors from invitations. The users or groups who are sponsors of the
// invited user. Sponsors are users and groups that are responsible for guest users' privileges in the tenant and for
// keeping the guest users' information and access up to date.
func (c InvitedUserSponsorClient) ListInvitedUserSponsors(ctx context.Context, options ListInvitedUserSponsorsOperationOptions) (result ListInvitedUserSponsorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListInvitedUserSponsorsCustomPager{},
		Path:          "/invitations/invitedUserSponsors",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListInvitedUserSponsorsComplete retrieves all the results into a single object
func (c InvitedUserSponsorClient) ListInvitedUserSponsorsComplete(ctx context.Context, options ListInvitedUserSponsorsOperationOptions) (ListInvitedUserSponsorsCompleteResult, error) {
	return c.ListInvitedUserSponsorsCompleteMatchingPredicate(ctx, options, DirectoryObjectOperationPredicate{})
}

// ListInvitedUserSponsorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InvitedUserSponsorClient) ListInvitedUserSponsorsCompleteMatchingPredicate(ctx context.Context, options ListInvitedUserSponsorsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListInvitedUserSponsorsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListInvitedUserSponsors(ctx, options)
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

	result = ListInvitedUserSponsorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
