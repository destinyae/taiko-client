package txlistdecoder

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type BlobFetcher struct {
	rpc *rpc.Client
}

func NewBlobTxListFetcher(rpc *rpc.Client) *BlobFetcher {
	return &BlobFetcher{rpc}
}

func (d *BlobFetcher) Fetch(
	ctx context.Context,
	tx *types.Transaction,
	meta *bindings.TaikoDataBlockMetadata,
) ([]byte, error) {
	if !meta.BlobUsed {
		return nil, errBlobUnused
	}

	sidecars, err := d.rpc.GetBlobs(ctx, new(big.Int).SetUint64(meta.L1Height+1))
	if err != nil {
		return nil, err
	}

	log.Info("Fetch sidecars", "sidecars", sidecars)

	// for _, sidecar := range sidecars {
	// 	log.Info("Found sidecar", "KzgCommitment", sidecar.KzgCommitment, "blobHash", common.Bytes2Hex(meta.BlobHash[:]))

	// 	if sidecar.KzgCommitment == common.Bytes2Hex(meta.BlobHash[:]) {
	// 		return common.Hex2Bytes(sidecar.Blob), nil
	// 	}
	// }

	return common.Hex2Bytes(sidecars[0].Blob), nil

	return nil, errSidecarNotFound
}