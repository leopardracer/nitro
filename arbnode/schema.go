// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package arbnode

var (
	messagePrefix                       []byte = []byte("m") // maps a message sequence number to a message
	blockHashInputFeedPrefix            []byte = []byte("b") // maps a message sequence number to a block hash received through the input feed
	blockMetadataInputFeedPrefix        []byte = []byte("t") // maps a message sequence number to a blockMetaData byte array received through the input feed
	missingBlockMetadataInputFeedPrefix []byte = []byte("x") // maps a message sequence number whose blockMetaData byte array is missing to nil
	messageResultPrefix                 []byte = []byte("r") // maps a message sequence number to a message result
	legacyDelayedMessagePrefix          []byte = []byte("d") // maps a delayed sequence number to an accumulator and a message as serialized on L1
	rlpDelayedMessagePrefix             []byte = []byte("e") // maps a delayed sequence number to an accumulator and an RLP encoded message
	parentChainBlockNumberPrefix        []byte = []byte("p") // maps a delayed sequence number to a parent chain block number
	sequencerBatchMetaPrefix            []byte = []byte("s") // maps a batch sequence number to BatchMetadata
	delayedSequencedPrefix              []byte = []byte("a") // maps a delayed message count to the first sequencer batch sequence number with this delayed count

	messageCountKey             []byte = []byte("_messageCount")                // contains the current message count
	lastPrunedMessageKey        []byte = []byte("_lastPrunedMessageKey")        // contains the last pruned message key
	lastPrunedDelayedMessageKey []byte = []byte("_lastPrunedDelayedMessageKey") // contains the last pruned RLP delayed message key
	delayedMessageCountKey      []byte = []byte("_delayedMessageCount")         // contains the current delayed message count
	sequencerBatchCountKey      []byte = []byte("_sequencerBatchCount")         // contains the current sequencer message count
	dbSchemaVersion             []byte = []byte("_schemaVersion")               // contains a uint64 representing the database schema version
)

const currentDbSchemaVersion uint64 = 1
