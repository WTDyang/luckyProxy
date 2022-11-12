package userRequest

import (
	"luckyProxy/client"
	"luckyProxy/common/logx"
	"luckyProxy/common/protocal"
)

// Handle user request
//
// write data to internet service
func Handle(c *client.Client, userRequest protocal.UserRequest) {
	internet, b := c.GetInternetService(userRequest.ConnId)
	if !b {
		logx.Error().Msg("internet not found, discard message")
		return
	}

	// write data to internet
	internet.Write(userRequest.Data)
}
