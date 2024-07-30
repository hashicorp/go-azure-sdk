package mailfolderuserconfiguration

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

type ListMailFolderUserConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserConfiguration
}

type ListMailFolderUserConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserConfiguration
}

type ListMailFolderUserConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderUserConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderUserConfigurations ...
func (c MailFolderUserConfigurationClient) ListMailFolderUserConfigurations(ctx context.Context, id MeMailFolderId) (result ListMailFolderUserConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderUserConfigurationsCustomPager{},
		Path:       fmt.Sprintf("%s/userConfigurations", id.ID()),
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
		Values *[]beta.UserConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMailFolderUserConfigurationsComplete retrieves all the results into a single object
func (c MailFolderUserConfigurationClient) ListMailFolderUserConfigurationsComplete(ctx context.Context, id MeMailFolderId) (ListMailFolderUserConfigurationsCompleteResult, error) {
	return c.ListMailFolderUserConfigurationsCompleteMatchingPredicate(ctx, id, UserConfigurationOperationPredicate{})
}

// ListMailFolderUserConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderUserConfigurationClient) ListMailFolderUserConfigurationsCompleteMatchingPredicate(ctx context.Context, id MeMailFolderId, predicate UserConfigurationOperationPredicate) (result ListMailFolderUserConfigurationsCompleteResult, err error) {
	items := make([]beta.UserConfiguration, 0)

	resp, err := c.ListMailFolderUserConfigurations(ctx, id)
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

	result = ListMailFolderUserConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
