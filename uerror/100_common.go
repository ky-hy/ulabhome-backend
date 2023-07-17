package uerror

import "golang.org/x/xerrors"

var (
	ErrException = ErrDef{
		Code:   100,
		RpsMsg: "予期せぬエラーが発生しました。",
		Err:    xerrors.New("予期せぬエラーが発生しました。"),
	}
)
