/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type AlgoBlock struct {
	// hash to which this block belongs
	GenesisHash string `json:"genesisHash,omitempty"`
	// ID to which this block belongs
	GenesisId string `json:"genesisId,omitempty"`
	// Previous block hash
	PreviousBlockHash string `json:"previousBlockHash,omitempty"`
	// rewards
	Rewards *interface{} `json:"rewards,omitempty"`
	// Current round on which this block was appended to the chain
	Round float64 `json:"round,omitempty"`
	// Sortition seed.
	Seed string `json:"seed,omitempty"`
	// Block creation timestamp in seconds since eposh
	Timestamp float64 `json:"timestamp,omitempty"`
	// Array of transactions
	Txns []AlgoTx `json:"txns,omitempty"`
	// TransactionsRoot authenticates the set of transactions appearing in the block.
	Txn string `json:"txn,omitempty"`
	// TxnCounter counts the number of transations committed in the ledger
	Txnc float64 `json:"txnc,omitempty"`
	// upgrade state
	UpgradeState *interface{} `json:"upgradeState,omitempty"`
	// upgrade vote
	UpgradeVote *interface{} `json:"upgradeVote,omitempty"`
}
