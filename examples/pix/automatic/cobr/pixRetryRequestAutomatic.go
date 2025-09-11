package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const txid = "0000000000000000000000000001"
	const data = "2021-03-05T22:01:35Z"

	res, err := efi.PixRetryRequestAutomatic(txid, data)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
