// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x47, /* |.......E| */
				0x4c, 0x6f, 0x73, 0x20, 0x41, 0x6e, 0x67, 0x65, /* |The Time| */
				0x6c, 0x65, 0x73, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |s 03/Jan| */
				0x73, 0x20, 0x31, 0x36, 0x2f, 0x53, 0x65, 0x70, /* |/2009 Ch| */
				0x2f, 0x32, 0x30, 0x32, 0x32, 0x20, 0x4c, 0x69, /* |ancellor| */
				0x62, 0x61, 0x6e, 0x6f, 0x2c, 0x20, 0x44, 0x65, /* | on brin| */
				0x70, 0x6f, 0x73, 0x69, 0x74, 0x61, 0x6e, 0x74, /* |k of sec|*/
				0x65, 0x73, 0x20, 0x69, 0x72, 0x72, 0x75, 0x6d, /* |ond bail| */
				0x70, 0x65, 0x6e, 0x20, 0x65, 0x6e, 0x20, 0x35, /* |out for |*/
				0x20, 0x62, 0x61, 0x6e, 0x63, 0x6f, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x35, 0x64, 0x05, 0x37, 0x70, 0x7b, /* |A.g....U| */
				0x79, 0x7e, 0x90, 0x08, 0x19, 0x96, 0x68, 0xf1, /* |H'.g..q0| */
				0xb8, 0xe8, 0xcc, 0x07, 0x97, 0x37, 0x67 ,0x6e,
				0x4c, 0x6a, 0x62, 0xe9, 0x39, 0x84, 0xe9, 0x70,
				0x56, 0xae, 0x45, 0xb7, 0x95, 0x4d, 0xea, 0x30,
				0xe3, 0x36, 0x9a, 0x4f, 0xeb, 0x4b, 0x35, 0xa8,
				0x19, 0xa4, 0xf0, 0xa4, 0x6c, 0x8f, 0x74, 0x1b,
				0xa9, 0xcd, 0xd3, 0x80, 0xf5, 0x74, 0x4c, 0xdc,
				0x68, 0xc7, 0xc5, 0xac,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf8, 0x6e, 0xd3, 0xf5, 0xcc, 0xcf, 0xf2, 0x7d, 
	0x9c, 0xa0, 0x62, 0x5e, 0xe6, 0x76, 0x18, 0xfc, 
	0xec, 0xf1, 0x55, 0x21, 0x07, 0xa8, 0x19, 0x63, 
	0x70, 0x9d, 0x48, 0xc6, 0x00, 0x00, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x9c, 0x25, 0x7e, 0x56, 0x5a, 0x12, 0x95, 0xde,
	0x81, 0xef, 0x4b, 0xb3, 0x32, 0xd3, 0xa1, 0x21,
	0x7f, 0x34, 0xba, 0x57, 0x73, 0x6c, 0x3b, 0xec,
	0x29, 0xde, 0x6b, 0x4a, 0x00, 0x18, 0x20, 0x89,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(0x632F3591, 0), // 2009-01-03 18:15:05 +0000 UTC
		Bits:       0x1d00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x8D823EF1,               // 2083236893
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf1, 0x6e, 0xd3, 0xf5, 0xcc, 0xcf, 0xf2, 0x7d, 
	0x9c, 0xa0, 0x62, 0x5e, 0xe6, 0x76, 0x18, 0xfc, 
	0xec, 0xf1, 0x55, 0x21, 0x07, 0xa8, 0x19, 0x63, 
	0x70, 0x9d, 0x48, 0xc6, 0x00, 0x00, 0x00, 0x00,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1664038289, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2374123249,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf0, 0x6e, 0xd3, 0xf5, 0xcc, 0xcf, 0xf2, 0x7d, 
	0x9c, 0xa0, 0x62, 0x5e, 0xe6, 0x76, 0x18, 0xfc, 
	0xec, 0xf1, 0x55, 0x21, 0x07, 0xa8, 0x19, 0x63, 
	0x70, 0x9d, 0x48, 0xc6, 0x00, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1664038289, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x8D823EF1,                // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf2, 0x6e, 0xd3, 0xf5, 0xcc, 0xcf, 0xf2, 0x7d, 
	0x9c, 0xa0, 0x62, 0x5e, 0xe6, 0x76, 0x18, 0xfc, 
	0xec, 0xf1, 0x55, 0x21, 0x07, 0xa8, 0x19, 0x63, 
	0x70, 0x9d, 0x48, 0xc6, 0x00, 0x00, 0x00, 0x00,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1664038289, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x1d00ffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      0x8D823EF1,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// sigNetGenesisHash is the hash of the first block in the block chain for the
// signet test network.
var sigNetGenesisHash = chainhash.Hash{
	0xf3, 0x6e, 0xd3, 0xf5, 0xcc, 0xcf, 0xf2, 0x7d, 
	0x9c, 0xa0, 0x62, 0x5e, 0xe6, 0x76, 0x18, 0xfc, 
	0xec, 0xf1, 0x55, 0x21, 0x07, 0xa8, 0x19, 0x63, 
	0x70, 0x9d, 0x48, 0xc6, 0x00, 0x00, 0x00, 0x00,
}

// sigNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the signet test network. It is the same as the merkle root for
// the main network.
var sigNetGenesisMerkleRoot = genesisMerkleRoot

// sigNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the signet test network.
var sigNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: sigNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1664038289, 0), // 2020-09-01 00:00:00 +0000 UTC
		Bits:       0x1d00ffff,               // 503543726 [00000377ae000000000000000000000000000000000000000000000000000000]
		Nonce:      2374123249,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
