package uerror

import "golang.org/x/xerrors"

var (
	ErrException = ErrDef{
		Code:   100,
		RpsMsg: "予期せぬエラーが発生しました。",
		Err:    xerrors.New("予期せぬエラーが発生しました。"),
	}
	ErrPanic = ErrDef{
		Code:   101,
		RpsMsg: "予期せぬエラーが発生しました。",
		Err:    xerrors.New("パニックが発生しました。"),
	}
)
