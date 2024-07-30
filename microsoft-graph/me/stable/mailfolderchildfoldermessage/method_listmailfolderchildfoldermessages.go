package mailfolderchildfoldermessage

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

type ListMailFolderChildFolderMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Message
}

type ListMailFolderChildFolderMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Message
}

type ListMailFolderChildFolderMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFolderMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolderMessages ...
func (c MailFolderChildFolderMessageClient) ListMailFolderChildFolderMessages(ctx context.Context, id MeMailFolderIdChildFolderId) (result ListMailFolderChildFolderMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMailFolderChildFolderMessagesCustomPager{},
		Path:       fmt.Sprintf("%s/messages", id.ID()),
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
		Values *[]stable.Message `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMailFolderChildFolderMessagesComplete retrieves all the results into a single object
func (c MailFolderChildFolderMessageClient) ListMailFolderChildFolderMessagesComplete(ctx context.Context, id MeMailFolderIdChildFolderId) (ListMailFolderChildFolderMessagesCompleteResult, error) {
	return c.ListMailFolderChildFolderMessagesCompleteMatchingPredicate(ctx, id, MessageOperationPredicate{})
}

// ListMailFolderChildFolderMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderMessageClient) ListMailFolderChildFolderMessagesCompleteMatchingPredicate(ctx context.Context, id MeMailFolderIdChildFolderId, predicate MessageOperationPredicate) (result ListMailFolderChildFolderMessagesCompleteResult, err error) {
	items := make([]stable.Message, 0)

	resp, err := c.ListMailFolderChildFolderMessages(ctx, id)
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

	result = ListMailFolderChildFolderMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
