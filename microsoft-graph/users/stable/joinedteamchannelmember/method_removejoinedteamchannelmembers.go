package joinedteamchannelmember

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

type RemoveJoinedTeamChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ActionResultPart
}

type RemoveJoinedTeamChannelMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ActionResultPart
}

type RemoveJoinedTeamChannelMembersOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveJoinedTeamChannelMembersOperationOptions() RemoveJoinedTeamChannelMembersOperationOptions {
	return RemoveJoinedTeamChannelMembersOperationOptions{}
}

func (o RemoveJoinedTeamChannelMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveJoinedTeamChannelMembersOperationOptions) ToOData() *odata.Query {
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

func (o RemoveJoinedTeamChannelMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveJoinedTeamChannelMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveJoinedTeamChannelMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveJoinedTeamChannelMembers - Invoke action remove. Remove multiple members from a team in a single request. The
// response provides details about which memberships could and couldn't be removed.
func (c JoinedTeamChannelMemberClient) RemoveJoinedTeamChannelMembers(ctx context.Context, id stable.UserIdJoinedTeamIdChannelId, input RemoveJoinedTeamChannelMembersRequest, options RemoveJoinedTeamChannelMembersOperationOptions) (result RemoveJoinedTeamChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveJoinedTeamChannelMembersCustomPager{},
		Path:          fmt.Sprintf("%s/members/remove", id.ID()),
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

	temp := make([]stable.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// RemoveJoinedTeamChannelMembersComplete retrieves all the results into a single object
func (c JoinedTeamChannelMemberClient) RemoveJoinedTeamChannelMembersComplete(ctx context.Context, id stable.UserIdJoinedTeamIdChannelId, input RemoveJoinedTeamChannelMembersRequest, options RemoveJoinedTeamChannelMembersOperationOptions) (RemoveJoinedTeamChannelMembersCompleteResult, error) {
	return c.RemoveJoinedTeamChannelMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// RemoveJoinedTeamChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelMemberClient) RemoveJoinedTeamChannelMembersCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamIdChannelId, input RemoveJoinedTeamChannelMembersRequest, options RemoveJoinedTeamChannelMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result RemoveJoinedTeamChannelMembersCompleteResult, err error) {
	items := make([]stable.ActionResultPart, 0)

	resp, err := c.RemoveJoinedTeamChannelMembers(ctx, id, input, options)
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

	result = RemoveJoinedTeamChannelMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
