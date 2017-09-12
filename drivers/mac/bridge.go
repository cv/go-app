package mac

import "C"
import (
	"github.com/murlokswarm/app/bridge"
)

func handleMacOSRequest(url string, p bridge.Payload, returnID string) (res bridge.Payload, err error) {
	panic("not implemented")
}

//export goRequest
func goRequest(url *C.char, payload *C.char) {
	driver.golang.Request(
		C.GoString(url),
		bridge.PayloadFromString(C.GoString(payload)),
	)
}

//export goRequestWithResult
// res should be free after each call of goRequestWithResult.
func goRequestWithResult(url *C.char, payload *C.char) (res *C.char) {
	pret := driver.golang.RequestWithResponse(
		C.GoString(url),
		bridge.PayloadFromString(C.GoString(payload)),
	)

	if pret != nil {
		res = C.CString(pret.String())
	}
	return
}
