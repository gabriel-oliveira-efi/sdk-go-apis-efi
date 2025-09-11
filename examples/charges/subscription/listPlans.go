package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	res, err := efi.ListPlans(20, 0) // limit e offset

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
