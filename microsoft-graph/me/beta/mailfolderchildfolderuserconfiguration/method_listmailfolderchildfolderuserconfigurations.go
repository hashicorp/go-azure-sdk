package mailfolderchildfolderuserconfiguration

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

type ListMailFolderChildFolderUserConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserConfiguration
}

type ListMailFolderChildFolderUserConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserConfiguration
}

type ListMailFolderChildFolderUserConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFolderUserConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolderUserConfigurations ...
func (c MailFolderChildFolderUserConfigurationClient) ListMailFolderChildFolderUserConfigurations(ctx context.Context, id MeMailFolderIdChildFolderId) (result ListMailFolderChildFolderUserConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderChildFolderUserConfigurationsCustomPager{},
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

// ListMailFolderChildFolderUserConfigurationsComplete retrieves all the results into a single object
func (c MailFolderChildFolderUserConfigurationClient) ListMailFolderChildFolderUserConfigurationsComplete(ctx context.Context, id MeMailFolderIdChildFolderId) (ListMailFolderChildFolderUserConfigurationsCompleteResult, error) {
	return c.ListMailFolderChildFolderUserConfigurationsCompleteMatchingPredicate(ctx, id, UserConfigurationOperationPredicate{})
}

// ListMailFolderChildFolderUserConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderUserConfigurationClient) ListMailFolderChildFolderUserConfigurationsCompleteMatchingPredicate(ctx context.Context, id MeMailFolderIdChildFolderId, predicate UserConfigurationOperationPredicate) (result ListMailFolderChildFolderUserConfigurationsCompleteResult, err error) {
	items := make([]beta.UserConfiguration, 0)

	resp, err := c.ListMailFolderChildFolderUserConfigurations(ctx, id)
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

	result = ListMailFolderChildFolderUserConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
