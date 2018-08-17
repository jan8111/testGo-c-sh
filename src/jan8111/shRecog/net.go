package shRecog

import (
	"net/http"
	"log"
	"fmt"
	"io"
	"unsafe"
)

var contextId11 []unsafe.Pointer

func HelloServer(w http.ResponseWriter, req *http.Request) {
	var bodySlc = make([]byte, 3200)

	sessionId:=recogStart(contextId11[0])
	for {
		bodyLen, readErr := req.Body.Read(bodySlc)
		if readErr != nil {
			fmt.Println("read body end",readErr)
			break
		} else {
			resumeFile(sessionId,bodySlc,bodyLen)
		}
	}
	resultStr:=recogEnd(sessionId)
	fmt.Println("result:", resultStr)
	io.WriteString(w, resultStr)
}

func StartHttpServer() {
	contextId11 = initEngine()
	fmt.Println("initEngine end. contextId11: ",contextId11[0])

	http.HandleFunc("/shRecBase/recog", HelloServer)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
