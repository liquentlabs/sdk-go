package types

import (
	"encoding/json"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type LiquentExecMsg struct {
	ExecutionData ExecutionData `json:"liquent_exec"`
}

type ExecutionData struct {
	Origin string      `json:"origin"`
	Name   string      `json:"name"`
	Args   interface{} `json:"args"`
}

func NewLiquentExecMsg(origin sdk.AccAddress, data string) (*LiquentExecMsg, error) {
	var e ExecutionData
	if err := json.Unmarshal([]byte(data), &e); err != nil {
		return nil, errors.Wrap(err, data)
	}

	if e.Origin == "" && origin.Empty() {
		return nil, errors.Wrap(sdkerrors.ErrInvalidAddress, "origin address is empty")
	}

	// override e.Origin for safety
	e.Origin = origin.String()

	return &LiquentExecMsg{
		ExecutionData: e,
	}, nil
}
