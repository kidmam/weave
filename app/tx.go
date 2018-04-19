package app

import (
	"github.com/confio/weave"
	"github.com/confio/weave/errors"
	"github.com/confio/weave/x/cash"
	"github.com/confio/weave/x/sigs"

	"github.com/iov-one/bcp-demo/x/hashlock"
)

//-------------------------------
// copied from weave/app verbatim
//
// any cleaner way to extend a tx with functionality?

// TxDecoder creates a Tx and unmarshals bytes into it
func TxDecoder(bz []byte) (weave.Tx, error) {
	tx := new(Tx)
	err := tx.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// make sure tx fulfills all interfaces
var _ weave.Tx = (*Tx)(nil)
var _ cash.FeeTx = (*Tx)(nil)
var _ sigs.SignedTx = (*Tx)(nil)
var _ hashlock.HashKeyTx = (*Tx)(nil)

// GetMsg switches over all types defined in the protobuf file
func (tx *Tx) GetMsg() (weave.Msg, error) {
	sum := tx.GetSum()
	if sum == nil {
		return nil, errors.ErrDecoding()
	}

	// make sure to cover all messages defined in protobuf
	switch t := sum.(type) {
	case *Tx_SendMsg:
		return t.SendMsg, nil
	case *Tx_SetNameMsg:
		return t.SetNameMsg, nil
	case *Tx_NewTokenMsg:
		return t.NewTokenMsg, nil
	case *Tx_CreateEscrowMsg:
		return t.CreateEscrowMsg, nil
	case *Tx_ReleaseEscrowMsg:
		return t.ReleaseEscrowMsg, nil
	case *Tx_ReturnEscrowMsg:
		return t.ReturnEscrowMsg, nil
	case *Tx_UpdateEscrowMsg:
		return t.UpdateEscrowMsg, nil
	}

	// we must have covered it above
	panic(sum)
	// return nil, errors.ErrUnknownTxType(nil) // alpe????
}

// GetSignBytes returns the bytes to sign...
func (tx *Tx) GetSignBytes() ([]byte, error) {
	// temporarily unset the signatures, as the sign bytes
	// should only come from the data itself, not previous signatures
	sigs := tx.Signatures
	tx.Signatures = nil

	bz, err := tx.Marshal()

	// reset the signatures after calculating the bytes
	tx.Signatures = sigs
	return bz, err
}
