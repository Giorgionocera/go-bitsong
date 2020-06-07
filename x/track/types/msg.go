package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Content messages types and routes
const (
	TypeMsgTrackPublish  = "track_publish"
	TypeMsgTrackTokenize = "track_tokenize"
	TypeMsgTokenMint     = "token_mint"
)

var _ sdk.Msg = MsgTrackPublish{}

type MsgTrackPublish struct {
	TrackInfo []byte         `json:"track_info" yaml:"track_info"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgTrackPublish(info []byte, creator sdk.AccAddress) MsgTrackPublish {
	return MsgTrackPublish{
		TrackInfo: info,
		Creator:   creator,
	}
}

func (msg MsgTrackPublish) Route() string { return RouterKey }
func (msg MsgTrackPublish) Type() string  { return TypeMsgTrackPublish }

func (msg MsgTrackPublish) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTrackPublish) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgTrackPublish) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
	/*addrs := make([]sdk.AccAddress, len(msg.Dao))
	for i, de := range msg.Dao {
		addrs[i] = de.Address
	}

	return addrs*/
}

func (msg MsgTrackPublish) String() string {
	// TODO
	return fmt.Sprintf(`Msg Track Add
Title: %s`,
		msg.Creator,
	)
}

var _ sdk.Msg = MsgTrackTokenize{}

type MsgTrackTokenize struct {
	TrackID uint64         `json:"track_id" yaml:"track_id"`
	Denom   string         `json:"denom" yaml:"denom"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgTrackTokenize(trackID uint64, denom string, creator sdk.AccAddress) MsgTrackTokenize {
	return MsgTrackTokenize{
		TrackID: trackID,
		Denom:   denom,
		Creator: creator,
	}
}

func (msg MsgTrackTokenize) Route() string { return RouterKey }
func (msg MsgTrackTokenize) Type() string  { return TypeMsgTrackTokenize }

func (msg MsgTrackTokenize) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTrackTokenize) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgTrackTokenize) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

func (msg MsgTrackTokenize) String() string {
	// TODO
	return fmt.Sprintf(`Msg Track Tokenize
ID: %d`,
		msg.TrackID,
	)
}

var _ sdk.Msg = MsgTokenMint{}

type MsgTokenMint struct {
	TrackID   uint64         `json:"track_id" yaml:"track_id"`
	Amount    sdk.Coin       `json:"amount" yaml:"amount"`
	Recipient sdk.AccAddress `json:"reciptient" yaml:"recipient"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgTokenMint(trackID uint64, amount sdk.Coin, recipient sdk.AccAddress, creator sdk.AccAddress) MsgTokenMint {
	return MsgTokenMint{
		TrackID:   trackID,
		Amount:    amount,
		Recipient: recipient,
		Creator:   creator,
	}
}

func (msg MsgTokenMint) Route() string { return RouterKey }
func (msg MsgTokenMint) Type() string  { return TypeMsgTokenMint }

func (msg MsgTokenMint) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTokenMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgTokenMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

func (msg MsgTokenMint) String() string {
	// TODO
	return fmt.Sprintf(`Msg Token Mint
ID: %d`,
		msg.TrackID,
	)
}