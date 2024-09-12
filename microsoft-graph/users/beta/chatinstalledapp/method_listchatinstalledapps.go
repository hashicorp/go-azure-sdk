package chatinstalledapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListChatInstalledAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TeamsAppInstallation
}

type ListChatInstalledAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TeamsAppInstallation
}

type ListChatInstalledAppsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListChatInstalledAppsOperationOptions() ListChatInstalledAppsOperationOptions {
	return ListChatInstalledAppsOperationOptions{}
}

func (o ListChatInstalledAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListChatInstalledAppsOperationOptions) ToOData() *odata.Query {
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

func (o ListChatInstalledAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListChatInstalledAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatInstalledAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatInstalledApps - Get installedApps from users. A collection of all the apps in the chat. Nullable.
func (c ChatInstalledAppClient) ListChatInstalledApps(ctx context.Context, id beta.UserIdChatId, options ListChatInstalledAppsOperationOptions) (result ListChatInstalledAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListChatInstalledAppsCustomPager{},
		Path:          fmt.Sprintf("%s/installedApps", id.ID()),
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

	temp := make([]beta.TeamsAppInstallation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalTeamsAppInstallationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.TeamsAppInstallation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListChatInstalledAppsComplete retrieves all the results into a single object
func (c ChatInstalledAppClient) ListChatInstalledAppsComplete(ctx context.Context, id beta.UserIdChatId, options ListChatInstalledAppsOperationOptions) (ListChatInstalledAppsCompleteResult, error) {
	return c.ListChatInstalledAppsCompleteMatchingPredicate(ctx, id, options, TeamsAppInstallationOperationPredicate{})
}

// ListChatInstalledAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatInstalledAppClient) ListChatInstalledAppsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdChatId, options ListChatInstalledAppsOperationOptions, predicate TeamsAppInstallationOperationPredicate) (result ListChatInstalledAppsCompleteResult, err error) {
	items := make([]beta.TeamsAppInstallation, 0)

	resp, err := c.ListChatInstalledApps(ctx, id, options)
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

	result = ListChatInstalledAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
