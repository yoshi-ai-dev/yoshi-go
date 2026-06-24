# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountListResponse">AccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountNewRealEstateResponse">AccountNewRealEstateResponse</a>

Methods:

- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/real-estate">client.Accounts.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountService.NewRealEstate">NewRealEstate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountNewRealEstateParams">AccountNewRealEstateParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountNewRealEstateResponse">AccountNewRealEstateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BalanceSeries

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountBalanceSeriesListResponse">AccountBalanceSeriesListResponse</a>

Methods:

- <code title="get /accounts/{id}/balances">client.Accounts.BalanceSeries.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountBalanceSeriesService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountBalanceSeriesListParams">AccountBalanceSeriesListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountBalanceSeriesListResponse">AccountBalanceSeriesListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionListResponse">TransactionListResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionChangesResponse">TransactionChangesResponse</a>

Methods:

- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionListResponse">TransactionListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/changes">client.Transactions.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionService.Changes">Changes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionChangesParams">TransactionChangesParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransactionChangesResponse">TransactionChangesResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardIdentityHints

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CardIdentityHintListResponse">CardIdentityHintListResponse</a>

Methods:

- <code title="get /card-identity-hints">client.CardIdentityHints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CardIdentityHintService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CardIdentityHintListParams">CardIdentityHintListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CardIdentityHintListResponse">CardIdentityHintListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scores

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#ScoreListResponse">ScoreListResponse</a>

Methods:

- <code title="get /scores">client.Scores.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#ScoreService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#ScoreListResponse">ScoreListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Goals

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalNewResponse">GoalNewResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalUpdateResponse">GoalUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalListResponse">GoalListResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalDeleteResponse">GoalDeleteResponse</a>

Methods:

- <code title="post /goals">client.Goals.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalNewParams">GoalNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalNewResponse">GoalNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /goals/{id}">client.Goals.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalUpdateParams">GoalUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalUpdateResponse">GoalUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /goals">client.Goals.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalListParams">GoalListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalListResponse">GoalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /goals/{id}">client.Goals.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalDeleteResponse">GoalDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Recurring

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#RecurringListResponse">RecurringListResponse</a>

Methods:

- <code title="get /recurring">client.Recurring.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#RecurringService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#RecurringListResponse">RecurringListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Benefits

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitExpiringResponse">BenefitExpiringResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitSummaryResponse">BenefitSummaryResponse</a>

Methods:

- <code title="get /benefits/expiring">client.Benefits.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitService.Expiring">Expiring</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitExpiringParams">BenefitExpiringParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitExpiringResponse">BenefitExpiringResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /benefits/summary">client.Benefits.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitService.Summary">Summary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenefitSummaryResponse">BenefitSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Investments

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentListResponse">InvestmentListResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentHoldingHistoryResponse">InvestmentHoldingHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentHoldingsResponse">InvestmentHoldingsResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentPerformanceResponse">InvestmentPerformanceResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentTransactionsResponse">InvestmentTransactionsResponse</a>

Methods:

- <code title="get /investments">client.Investments.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentListParams">InvestmentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentListResponse">InvestmentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/holdings/history">client.Investments.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentService.HoldingHistory">HoldingHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentHoldingHistoryParams">InvestmentHoldingHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentHoldingHistoryResponse">InvestmentHoldingHistoryResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/holdings">client.Investments.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentService.Holdings">Holdings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentHoldingsParams">InvestmentHoldingsParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentHoldingsResponse">InvestmentHoldingsResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/performance">client.Investments.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentService.Performance">Performance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentPerformanceParams">InvestmentPerformanceParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentPerformanceResponse">InvestmentPerformanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/transactions">client.Investments.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentService.Transactions">Transactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentTransactionsParams">InvestmentTransactionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentTransactionsResponse">InvestmentTransactionsResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Benchmarks

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenchmarkReplayResponse">BenchmarkReplayResponse</a>

Methods:

- <code title="post /benchmarks/replay">client.Benchmarks.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenchmarkService.Replay">Replay</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenchmarkReplayParams">BenchmarkReplayParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BenchmarkReplayResponse">BenchmarkReplayResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Trades

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TradeNewResponse">TradeNewResponse</a>

Methods:

- <code title="post /trades">client.Trades.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TradeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TradeNewParams">TradeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TradeNewResponse">TradeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transfers

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransferNewResponse">TransferNewResponse</a>

Methods:

- <code title="post /transfers">client.Transfers.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransferNewParams">TransferNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#TransferNewResponse">TransferNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Income

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#IncomeGetResponse">IncomeGetResponse</a>

Methods:

- <code title="get /income">client.Income.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#IncomeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#IncomeGetResponse">IncomeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Spending

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SpendingGetResponse">SpendingGetResponse</a>

Methods:

- <code title="get /spending">client.Spending.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SpendingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SpendingGetParams">SpendingGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SpendingGetResponse">SpendingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# NetWorth

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#NetWorthHistoryResponse">NetWorthHistoryResponse</a>

Methods:

- <code title="get /net-worth/history">client.NetWorth.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#NetWorthService.History">History</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#NetWorthHistoryParams">NetWorthHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#NetWorthHistoryResponse">NetWorthHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CreditDebt

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CreditDebtGetResponse">CreditDebtGetResponse</a>

Methods:

- <code title="get /credit-debt">client.CreditDebt.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CreditDebtService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#CreditDebtGetResponse">CreditDebtGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Automations

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AutomationListResponse">AutomationListResponse</a>

Methods:

- <code title="get /automations">client.Automations.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AutomationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AutomationListParams">AutomationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AutomationListResponse">AutomationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Briefs

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefGetResponse">BriefGetResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefListResponse">BriefListResponse</a>

Methods:

- <code title="get /briefs/{id}">client.Briefs.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefGetResponse">BriefGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /briefs">client.Briefs.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefListParams">BriefListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#BriefListResponse">BriefListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Securities

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityGetResponse">SecurityGetResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecuritySearchResponse">SecuritySearchResponse</a>

Methods:

- <code title="get /securities/{symbol}">client.Securities.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityGetResponse">SecurityGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /securities/search">client.Securities.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecuritySearchParams">SecuritySearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecuritySearchResponse">SecuritySearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Options

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityOptionChainResponse">SecurityOptionChainResponse</a>

Methods:

- <code title="get /securities/{symbol}/options/chain">client.Securities.Options.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityOptionService.Chain">Chain</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityOptionChainParams">SecurityOptionChainParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityOptionChainResponse">SecurityOptionChainResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PriceHistory

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityPriceHistoryListResponse">SecurityPriceHistoryListResponse</a>

Methods:

- <code title="get /securities/{symbol}/price-history">client.Securities.PriceHistory.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityPriceHistoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityPriceHistoryListParams">SecurityPriceHistoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#SecurityPriceHistoryListResponse">SecurityPriceHistoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Me

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeGetResponse">MeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeSummaryResponse">MeSummaryResponse</a>

Methods:

- <code title="get /me">client.Me.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeGetResponse">MeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /me/summary">client.Me.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeService.Summary">Summary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeSummaryParams">MeSummaryParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#MeSummaryResponse">MeSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PaperTrading

## Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountNewResponse">PaperTradingAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountListResponse">PaperTradingAccountListResponse</a>

Methods:

- <code title="post /paper-trading/accounts">client.PaperTrading.Accounts.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountNewParams">PaperTradingAccountNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountNewResponse">PaperTradingAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /paper-trading/accounts">client.PaperTrading.Accounts.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountListResponse">PaperTradingAccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Holdings

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountHoldingListResponse">PaperTradingAccountHoldingListResponse</a>

Methods:

- <code title="get /paper-trading/accounts/{accountId}/holdings">client.PaperTrading.Accounts.Holdings.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountHoldingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountHoldingListResponse">PaperTradingAccountHoldingListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Trades

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeNewResponse">PaperTradingAccountTradeNewResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeListResponse">PaperTradingAccountTradeListResponse</a>

Methods:

- <code title="post /paper-trading/accounts/{accountId}/trades">client.PaperTrading.Accounts.Trades.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeNewParams">PaperTradingAccountTradeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeNewResponse">PaperTradingAccountTradeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /paper-trading/accounts/{accountId}/trades">client.PaperTrading.Accounts.Trades.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeListParams">PaperTradingAccountTradeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeListResponse">PaperTradingAccountTradeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

## Endpoints

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointNewResponse">WebhookEndpointNewResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointGetResponse">WebhookEndpointGetResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointUpdateResponse">WebhookEndpointUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointListResponse">WebhookEndpointListResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointDeleteResponse">WebhookEndpointDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointRotateResponse">WebhookEndpointRotateResponse</a>
- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointTestResponse">WebhookEndpointTestResponse</a>

Methods:

- <code title="post /webhooks/endpoints">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointNewParams">WebhookEndpointNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointNewResponse">WebhookEndpointNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/endpoints/{id}">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointGetResponse">WebhookEndpointGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /webhooks/endpoints/{id}">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointUpdateParams">WebhookEndpointUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointUpdateResponse">WebhookEndpointUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/endpoints">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointListResponse">WebhookEndpointListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/endpoints/{id}">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointDeleteResponse">WebhookEndpointDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/endpoints/{id}/rotate">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointRotateResponse">WebhookEndpointRotateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/endpoints/{id}/test">client.Webhooks.Endpoints.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointService.Test">Test</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEndpointTestResponse">WebhookEndpointTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Deliveries

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookDeliveryListResponse">WebhookDeliveryListResponse</a>

Methods:

- <code title="get /webhooks/deliveries">client.Webhooks.Deliveries.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookDeliveryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookDeliveryListParams">WebhookDeliveryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookDeliveryListResponse">WebhookDeliveryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Events

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEventListResponse">WebhookEventListResponse</a>

Methods:

- <code title="get /webhooks/events">client.Webhooks.Events.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookEventListResponse">WebhookEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Portal

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookPortalGetResponse">WebhookPortalGetResponse</a>

Methods:

- <code title="get /webhooks/portal">client.Webhooks.Portal.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookPortalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#WebhookPortalGetResponse">WebhookPortalGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Approvals

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#ApprovalGetResponse">ApprovalGetResponse</a>

Methods:

- <code title="get /approvals/{threadId}">client.Approvals.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#ApprovalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#ApprovalGetResponse">ApprovalGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
