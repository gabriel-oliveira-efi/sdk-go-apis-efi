package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idSolicRec = "SC09089356202509035192e32902e"

	body := map[string]interface{}{
		"status": "CANCELADA",
	}

	res, err := efi.PixUpdateRequestRecurrenceAutomatic(idSolicRec, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
