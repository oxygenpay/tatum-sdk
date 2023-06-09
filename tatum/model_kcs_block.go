/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type KcsBlock struct {
	// Difficulty for this block.
	Difficulty string `json:"difficulty,omitempty"`
	// The 'extra data' field of this block.
	ExtraData string `json:"extraData,omitempty"`
	// The maximum gas allowed in this block.
	GasLimit float64 `json:"gasLimit,omitempty"`
	// The total used gas by all transactions in this block.
	GasUsed float64 `json:"gasUsed,omitempty"`
	// Hash of the block. 'null' when its pending block.
	Hash string `json:"hash,omitempty"`
	// The bloom filter for the logs of the block. 'null' when its pending block.
	LogsBloom string `json:"logsBloom,omitempty"`
	// The address of the beneficiary to whom the mining rewards were given.
	Miner string `json:"miner,omitempty"`
	MixHash string `json:"mixHash,omitempty"`
	// Hash of the generated proof-of-work. 'null' when its pending block.
	Nonce string `json:"nonce,omitempty"`
	// The block number. 'null' when its pending block.
	Number float64 `json:"number,omitempty"`
	// Hash of the parent block.
	ParentHash string `json:"parentHash,omitempty"`
	ReceiptsRoot string `json:"receiptsRoot,omitempty"`
	// SHA3 of the uncles data in the block.
	Sha3Uncles string `json:"sha3Uncles,omitempty"`
	// The size of this block in bytes.
	Size float64 `json:"size,omitempty"`
	// The root of the final state trie of the block.
	StateRoot string `json:"stateRoot,omitempty"`
	// The unix timestamp for when the block was collated.
	Timestamp float64 `json:"timestamp,omitempty"`
	// Total difficulty of the chain until this block.
	TotalDifficulty string `json:"totalDifficulty,omitempty"`
	// Array of transactions.
	Transactions []KcsTx `json:"transactions,omitempty"`
	// The root of the transaction trie of the block.
	TransactionsRoot string `json:"transactionsRoot,omitempty"`
}
