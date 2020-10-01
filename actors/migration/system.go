package migration

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	system0 "github.com/filecoin-project/specs-actors/actors/builtin/system"
	cid "github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"

	system2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/system"
)

type systemMigrator struct {
}

func (m *systemMigrator) MigrateState(ctx context.Context, store cbor.IpldStore, head cid.Cid, _ MigrationInfo) (cid.Cid, abi.TokenAmount, error) {
	// No change
	var _ = system2.State(system0.State{})
	return head, big.Zero(), nil
}
