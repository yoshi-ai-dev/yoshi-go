# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalListResponse">GoalListResponse</a>

Methods:

- <code title="get /goals">client.Goals.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalListParams">GoalListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#GoalListResponse">GoalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

Methods:

- <code title="get /investments">client.Investments.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentListParams">InvestmentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#InvestmentListResponse">InvestmentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Income

Response Types:

- <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#IncomeGetResponse">IncomeGetResponse</a>

Methods:

- <code title="get /income">client.Income.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#IncomeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#IncomeGetResponse">IncomeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
- <code title="get /paper-trading/accounts/{accountId}/trades">client.PaperTrading.Accounts.Trades.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeListParams">PaperTradingAccountTradeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go">yoshi</a>.<a href="https://pkg.go.dev/github.com/yoshi-ai-dev/yoshi-go#PaperTradingAccountTradeListResponse">PaperTradingAccountTradeListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
