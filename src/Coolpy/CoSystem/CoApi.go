package CoSystem

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var CpVersion = "5.0.2.0"

func VersionGet(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, `{"ok":%d,"data":"%v"}`, 1, CpVersion)
}
