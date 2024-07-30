package setting

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

type ListSettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectorySetting
}

type ListSettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectorySetting
}

type ListSettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSettings ...
func (c SettingClient) ListSettings(ctx context.Context, id GroupId) (result ListSettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSettingsCustomPager{},
		Path:       fmt.Sprintf("%s/settings", id.ID()),
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
		Values *[]beta.DirectorySetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSettingsComplete retrieves all the results into a single object
func (c SettingClient) ListSettingsComplete(ctx context.Context, id GroupId) (ListSettingsCompleteResult, error) {
	return c.ListSettingsCompleteMatchingPredicate(ctx, id, DirectorySettingOperationPredicate{})
}

// ListSettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SettingClient) ListSettingsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate DirectorySettingOperationPredicate) (result ListSettingsCompleteResult, err error) {
	items := make([]beta.DirectorySetting, 0)

	resp, err := c.ListSettings(ctx, id)
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

	result = ListSettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
