package joinedteamchannelallmember

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveJoinedTeamChannelAllMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ActionResultPart
}

type RemoveJoinedTeamChannelAllMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ActionResultPart
}

type RemoveJoinedTeamChannelAllMembersOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveJoinedTeamChannelAllMembersOperationOptions() RemoveJoinedTeamChannelAllMembersOperationOptions {
	return RemoveJoinedTeamChannelAllMembersOperationOptions{}
}

func (o RemoveJoinedTeamChannelAllMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveJoinedTeamChannelAllMembersOperationOptions) ToOData() *odata.Query {
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

func (o RemoveJoinedTeamChannelAllMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveJoinedTeamChannelAllMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveJoinedTeamChannelAllMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveJoinedTeamChannelAllMembers - Invoke action remove. Remove multiple members from a team in a single request.
// The response provides details about which memberships could and couldn't be removed.
func (c JoinedTeamChannelAllMemberClient) RemoveJoinedTeamChannelAllMembers(ctx context.Context, id stable.MeJoinedTeamIdChannelId, input RemoveJoinedTeamChannelAllMembersRequest, options RemoveJoinedTeamChannelAllMembersOperationOptions) (result RemoveJoinedTeamChannelAllMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveJoinedTeamChannelAllMembersCustomPager{},
		Path:          fmt.Sprintf("%s/allMembers/remove", id.ID()),
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

// RemoveJoinedTeamChannelAllMembersComplete retrieves all the results into a single object
func (c JoinedTeamChannelAllMemberClient) RemoveJoinedTeamChannelAllMembersComplete(ctx context.Context, id stable.MeJoinedTeamIdChannelId, input RemoveJoinedTeamChannelAllMembersRequest, options RemoveJoinedTeamChannelAllMembersOperationOptions) (RemoveJoinedTeamChannelAllMembersCompleteResult, error) {
	return c.RemoveJoinedTeamChannelAllMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// RemoveJoinedTeamChannelAllMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelAllMemberClient) RemoveJoinedTeamChannelAllMembersCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamIdChannelId, input RemoveJoinedTeamChannelAllMembersRequest, options RemoveJoinedTeamChannelAllMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result RemoveJoinedTeamChannelAllMembersCompleteResult, err error) {
	items := make([]stable.ActionResultPart, 0)

	resp, err := c.RemoveJoinedTeamChannelAllMembers(ctx, id, input, options)
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

	result = RemoveJoinedTeamChannelAllMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
