package uerror

import "golang.org/x/xerrors"

var (
	ErrDBException = ErrDef{
		Code:   200,
		RpsMsg: "予期せぬエラーが発生しました。",
		Err:    xerrors.New("userテーブルで予期せぬエラーが発生しました。"),
	}
)
