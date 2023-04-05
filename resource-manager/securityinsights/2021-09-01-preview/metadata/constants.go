package metadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Kind string

const (
	KindAnalyticsRule         Kind = "AnalyticsRule"
	KindAnalyticsRuleTemplate Kind = "AnalyticsRuleTemplate"
	KindDataConnector         Kind = "DataConnector"
	KindDataType              Kind = "DataType"
	KindHuntingQuery          Kind = "HuntingQuery"
	KindInvestigationQuery    Kind = "InvestigationQuery"
	KindParser                Kind = "Parser"
	KindPlaybook              Kind = "Playbook"
	KindPlaybookTemplate      Kind = "PlaybookTemplate"
	KindSolution              Kind = "Solution"
	KindWatchlist             Kind = "Watchlist"
	KindWatchlistTemplate     Kind = "WatchlistTemplate"
	KindWorkbook              Kind = "Workbook"
	KindWorkbookTemplate      Kind = "WorkbookTemplate"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindAnalyticsRule),
		string(KindAnalyticsRuleTemplate),
		string(KindDataConnector),
		string(KindDataType),
		string(KindHuntingQuery),
		string(KindInvestigationQuery),
		string(KindParser),
		string(KindPlaybook),
		string(KindPlaybookTemplate),
		string(KindSolution),
		string(KindWatchlist),
		string(KindWatchlistTemplate),
		string(KindWorkbook),
		string(KindWorkbookTemplate),
	}
}

type Operator string

const (
	OperatorAND Operator = "AND"
	OperatorOR  Operator = "OR"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(OperatorAND),
		string(OperatorOR),
	}
}

type SourceKind string

const (
	SourceKindCommunity        SourceKind = "Community"
	SourceKindLocalWorkspace   SourceKind = "LocalWorkspace"
	SourceKindSolution         SourceKind = "Solution"
	SourceKindSourceRepository SourceKind = "SourceRepository"
)

func PossibleValuesForSourceKind() []string {
	return []string{
		string(SourceKindCommunity),
		string(SourceKindLocalWorkspace),
		string(SourceKindSolution),
		string(SourceKindSourceRepository),
	}
}

type SupportTier string

const (
	SupportTierCommunity SupportTier = "Community"
	SupportTierMicrosoft SupportTier = "Microsoft"
	SupportTierPartner   SupportTier = "Partner"
)

func PossibleValuesForSupportTier() []string {
	return []string{
		string(SupportTierCommunity),
		string(SupportTierMicrosoft),
		string(SupportTierPartner),
	}
}
