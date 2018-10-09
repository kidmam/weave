package username_test

import (
	"fmt"
	"testing"

	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/nft/username"
	"github.com/stretchr/testify/assert"
)

func TestIssueTokenMsgValidate(t *testing.T) {
	var helpers x.TestHelpers
	_, alice := helpers.MakeKey()

	specs := []struct {
		msg      username.IssueTokenMsg
		expError bool
	}{
		{ // happy path email
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      []byte("alice@example.com"),
				Details: username.TokenDetails{},
			},
			expError: false,
		},
		{ // happy path twitter
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      []byte("@iov_official"),
				Details: username.TokenDetails{},
			},
			expError: false,
		},
		{ // happy path phone
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      []byte("+491234567890"),
				Details: username.TokenDetails{},
			},
			expError: false,
		},
		{ // other characters
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      []byte("+-,._@"),
				Details: username.TokenDetails{},
			},
			expError: false,
		},
		{ // owner missing
			msg: username.IssueTokenMsg{
				Id:      []byte("alice@example.com"),
				Details: username.TokenDetails{},
			},
			expError: true,
		},
		{ // owner wrong format
			msg: username.IssueTokenMsg{
				Owner:   []byte("not an address"),
				Id:      []byte("alice@example.com"),
				Details: username.TokenDetails{},
			},
			expError: true,
		},
		{ // id too short
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      []byte("foo"),
				Details: username.TokenDetails{},
			},
			expError: true,
		},
		{ // id too long
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      anyIDWithLength(65),
				Details: username.TokenDetails{},
			},
			expError: true,
		},
		{ // id with forbidden character *
			msg: username.IssueTokenMsg{
				Owner:   alice.Address(),
				Id:      []byte("foo*bar"),
				Details: username.TokenDetails{},
			},
			expError: true,
		},
		// TODO: Add checks for approvals
		// TODO: Add checks for TokenDetails
	}
	for i, spec := range specs {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			err := spec.msg.Validate()
			if spec.expError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddChainAddressMsgValidate(t *testing.T) {
	specs := []struct {
		msg      username.AddChainAddressMsg
		expError bool
	}{
		{ // happy path
			msg: username.AddChainAddressMsg{
				Id:      []byte("me@example.com"),
				ChainID: []byte("myChain"),
				Address: []byte("myChainAddress"),
			},
		}, { // address missing
			msg: username.AddChainAddressMsg{
				Id:      []byte("me@example.com"),
				ChainID: []byte("myChain"),
			},
			expError: true,
		}, { // id missing
			msg: username.AddChainAddressMsg{
				ChainID: []byte("myChain"),
				Address: []byte("myChainAddress"),
			},
			expError: true,
		},
		{ // chainID missing
			msg: username.AddChainAddressMsg{
				Id:      []byte("me@example.com"),
				Address: []byte("myChainAddress"),
			},
			expError: true,
		},
	}
	for i, spec := range specs {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			err := spec.msg.Validate()
			if spec.expError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestRemoveChainAddressMsgValidate(t *testing.T) {
	specs := []struct {
		msg      username.RemoveChainAddressMsg
		expError bool
	}{
		{ // happy path
			msg: username.RemoveChainAddressMsg{
				Id:      []byte("me@example.com"),
				ChainID: []byte("myChain"),
				Address: []byte("myChainAddress"),
			},
		}, { // address missing
			msg: username.RemoveChainAddressMsg{
				Id:      []byte("me@example.com"),
				ChainID: []byte("myChain"),
			},
			expError: true,
		}, { // id missing
			msg: username.RemoveChainAddressMsg{
				ChainID: []byte("myChain"),
				Address: []byte("myChainAddress"),
			},
			expError: true,
		},
		{ // chainID missing
			msg: username.RemoveChainAddressMsg{
				Id:      []byte("me@example.com"),
				Address: []byte("myChainAddress"),
			},
			expError: true,
		},
	}
	for i, spec := range specs {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			err := spec.msg.Validate()
			if spec.expError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func anyIDWithLength(n int) []byte {
	r := make([]byte, n)
	for i := 0; i < n; i++ {
		r[i] = byte('a')
	}
	return r
}
