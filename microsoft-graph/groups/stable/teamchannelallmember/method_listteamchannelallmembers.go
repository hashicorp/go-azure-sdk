package teamchannelallmember

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

type ListTeamChannelAllMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationMember
}

type ListTeamChannelAllMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationMember
}

type ListTeamChannelAllMembersOperationOptions struct {
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

func DefaultListTeamChannelAllMembersOperationOptions() ListTeamChannelAllMembersOperationOptions {
	return ListTeamChannelAllMembersOperationOptions{}
}

func (o ListTeamChannelAllMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamChannelAllMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamChannelAllMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamChannelAllMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelAllMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelAllMembers - Get allMembers from groups. A collection of membership records associated with the
// channel, including both direct and indirect members of shared channels.
func (c TeamChannelAllMemberClient) ListTeamChannelAllMembers(ctx context.Context, id stable.GroupIdTeamChannelId, options ListTeamChannelAllMembersOperationOptions) (result ListTeamChannelAllMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamChannelAllMembersCustomPager{},
		Path:          fmt.Sprintf("%s/allMembers", id.ID()),
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

	temp := make([]stable.ConversationMember, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalConversationMemberImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ConversationMember (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTeamChannelAllMembersComplete retrieves all the results into a single object
func (c TeamChannelAllMemberClient) ListTeamChannelAllMembersComplete(ctx context.Context, id stable.GroupIdTeamChannelId, options ListTeamChannelAllMembersOperationOptions) (ListTeamChannelAllMembersCompleteResult, error) {
	return c.ListTeamChannelAllMembersCompleteMatchingPredicate(ctx, id, options, ConversationMemberOperationPredicate{})
}

// ListTeamChannelAllMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelAllMemberClient) ListTeamChannelAllMembersCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdTeamChannelId, options ListTeamChannelAllMembersOperationOptions, predicate ConversationMemberOperationPredicate) (result ListTeamChannelAllMembersCompleteResult, err error) {
	items := make([]stable.ConversationMember, 0)

	resp, err := c.ListTeamChannelAllMembers(ctx, id, options)
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

	result = ListTeamChannelAllMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
