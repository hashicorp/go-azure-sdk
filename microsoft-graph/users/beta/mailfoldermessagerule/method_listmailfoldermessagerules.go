package mailfoldermessagerule

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

type ListMailFolderMessageRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MessageRule
}

type ListMailFolderMessageRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MessageRule
}

type ListMailFolderMessageRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderMessageRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderMessageRules ...
func (c MailFolderMessageRuleClient) ListMailFolderMessageRules(ctx context.Context, id UserIdMailFolderId) (result ListMailFolderMessageRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderMessageRulesCustomPager{},
		Path:       fmt.Sprintf("%s/messageRules", id.ID()),
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
		Values *[]beta.MessageRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMailFolderMessageRulesComplete retrieves all the results into a single object
func (c MailFolderMessageRuleClient) ListMailFolderMessageRulesComplete(ctx context.Context, id UserIdMailFolderId) (ListMailFolderMessageRulesCompleteResult, error) {
	return c.ListMailFolderMessageRulesCompleteMatchingPredicate(ctx, id, MessageRuleOperationPredicate{})
}

// ListMailFolderMessageRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderMessageRuleClient) ListMailFolderMessageRulesCompleteMatchingPredicate(ctx context.Context, id UserIdMailFolderId, predicate MessageRuleOperationPredicate) (result ListMailFolderMessageRulesCompleteResult, err error) {
	items := make([]beta.MessageRule, 0)

	resp, err := c.ListMailFolderMessageRules(ctx, id)
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

	result = ListMailFolderMessageRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
