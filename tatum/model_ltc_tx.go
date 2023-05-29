/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type LtcTx struct {
	// Transaction hash.
	Hash string `json:"hash,omitempty"`
	// Witness hash in case of a SegWit transaction.
	WitnessHash string `json:"witnessHash,omitempty"`
	// Fee paid for this transaction, in LTC.
	Fee string `json:"fee,omitempty"`
	Rate string `json:"rate,omitempty"`
	Ps float64 `json:"ps,omitempty"`
	// Height of the block this transaction belongs to.
	BlockNumber float64 `json:"blockNumber,omitempty"`
	// Hash of the block this transaction belongs to.
	Block string `json:"block,omitempty"`
	// Time of the transaction.
	Ts float64 `json:"ts,omitempty"`
	// Index of the transaction in the block.
	Index float64 `json:"index,omitempty"`
	// Index of the transaction.
	Version float64 `json:"version,omitempty"`
	Flag float64 `json:"flag,omitempty"`
	// List of transactions, from which assets are being sent.
	Inputs []LtcTxInputs `json:"inputs,omitempty"`
	// List of recipient addresses and amounts to send to each of them.
	Outputs []LtcTxOutputs `json:"outputs,omitempty"`
	// Block this transaction was included in.
	Locktime float64 `json:"locktime,omitempty"`
}