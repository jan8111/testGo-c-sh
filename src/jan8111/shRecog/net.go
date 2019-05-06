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
	var bodySlc = make([]byte, 160)

	//sampleRate:=req.Header.Get("sampleRate")
	//fmt.Printf("sampleRate: %s", sampleRate)

	sessionId:=recogStart(contextId11[0])
	for {
		bodyLen, readErr := req.Body.Read(bodySlc)
		if readErr != nil {
			fmt.Println("read body end",readErr)
			break
		} else {
			//fmt.Println("resumeFile.bodyLen:",bodyLen)
			resumeFile(sessionId,bodySlc,bodyLen)
		}
	}
	resultStr:=recogEnd(sessionId)
	fmt.Println("result:", resultStr)
	io.WriteString(w, resultStr)
}

func StartHttpServer() {
	contextId11 = initEngine()
	if contextId11[0]!=nil {
		fmt.Println("initEngine success! ")
	}

	http.HandleFunc("/shCluster-worker/recog", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
