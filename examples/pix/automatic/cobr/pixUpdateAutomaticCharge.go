package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const txid = "00000000000000000000000001"

	body := map[string]interface{}{
		"status": "CANCELADA",
	}

	res, err := efi.PixUpdateAutomaticCharge(txid, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
