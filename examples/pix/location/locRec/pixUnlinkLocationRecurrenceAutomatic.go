package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const id = "1"

	res, err := efi.PixUnlinkLocationRecurrenceAutomatic(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
