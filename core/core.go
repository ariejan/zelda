package core

import (
	"fmt"
	"os"

	"github.com/imroc/req"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func validateStatusCode(resp *req.Resp, statusCodes []int) {
	statusCode := resp.Response().StatusCode

	for _, code := range statusCodes {
		if code == statusCode {
			return
		}
	}

	fmt.Printf("==> Unexpected HTTP Status %d\n", statusCode)
	os.Exit(1)
}
