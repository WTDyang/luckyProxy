package handler

import (
	"luckyProxy/common/protocal"
	"luckyProxy/server/internal/api/ws/handler/internetResponse"
)

func Dispatch(burst protocal.Burst) {
	switch burst.Type {
	case protocal.IntranetResponseType:
		internetResponse.Handle(burst.IntranetResponse)
	}
}
