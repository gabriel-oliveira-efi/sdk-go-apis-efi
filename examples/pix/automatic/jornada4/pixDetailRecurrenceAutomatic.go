package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idRec = "RN09089356202509040648092afd6"

	res, err := efi.PixDetailRecurrenceAutomatic(idRec)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
