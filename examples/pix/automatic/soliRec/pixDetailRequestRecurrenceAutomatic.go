package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idSolicRec = "SC090893562025090205c9cfd85fc"

	res, err := efi.PixDetailRequestRecurrenceAutomatic(idSolicRec)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
