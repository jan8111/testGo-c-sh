package main

/*
#cgo CFLAGS : -I../include
#cgo LDFLAGS: -L/home/shhan/tw_model/deps -lrecognizer

#include "asr_api.h"
#include <stdlib.h>

int print111(char* str)
{
  return strlen(str);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
	"os"
	"bufio"
	)

func main() {
	contextId1 := initEngine()
	resultStr := recog(contextId1)
	fmt.Println("result:", resultStr)
}

func initEngine() unsafe.Pointer {
	ret := C.recognizer_getVersion()
	str1 := C.GoString(ret)
	fmt.Println("recognizer_getVersion: ", str1)

	workpath1 := C.CString("/home/shhan/tw_model/")
	ret2 := C.recognizer_setWorkPath(workpath1)
	C.free(unsafe.Pointer(workpath1))
	fmt.Println("recognizer_setWorkPath: ", ret2)

	name1 := C.CString("ctc")
	type1 := C.CString("dnn")
	path1 := C.CString("/home/shhan/tw_model/model/final_0808_online.bin")
	ret4 := C.recognizer_addAcoustic(name1, type1, path1, 0)
	fmt.Println("recognizer_addAcoustic: ", ret4)

	name2 := C.CString("first-path")
	type2 := C.CString("wfst")
	path2 := C.CString("/home/shhan/tw_model/model/xiaoi_hotword_test_0326_5.65e-9_ctc.dat")
	ret5 := C.recognizer_addDecoder(name2, type2, path2)
	fmt.Println("recognizer_addDecoder: ", ret5)

	var contextId1 unsafe.Pointer
	ret3 := C.recognizer_createContext(&contextId1)
	fmt.Println("recognizer_createContext: ", ret3)

	ret6 := C.recognizer_setContextAcoustic(contextId1, name1)
	fmt.Println("recognizer_setContextAcoustic: ", ret6)

	ret8 := C.recognizer_attachContextDecoder(contextId1, name2, false)
	fmt.Println("recognizer_attachContextDecoder: ", ret8)

	param1 := new(C.UnivoiceAcousticParam)
	(*param1).cpu_batch_size = 40
	(*param1).sq_snr_estimate = 0
	(*param1).sq_clipping_dectect = 0
	ret7 := C.recognizer_setContextAcousticParam(contextId1, param1)
	fmt.Println("recognizer_setContextAcousticParam: ", ret7)

	return contextId1
}

func recog(contextId1 unsafe.Pointer) string {
	var sessionId unsafe.Pointer
	ret := C.recognizer_createSession(&sessionId, contextId1, 16000);
	fmt.Println("recognizer_createSession: ", ret)

	ret2 := C.recognizer_startSession(sessionId, 0);
	fmt.Println("recognizer_startSession: ", ret2)

	resumeFile(sessionId)

	ret3 := C.recognizer_stopSession(sessionId);
	fmt.Println("recognizer_stopSession: ", ret3)

	result1 := C.CString("")
	ret4 := C.recognizer_getSessionResStr(sessionId, &result1);
	fmt.Println("recognizer_getSessionResStr: ", ret4)

	defer C.recognizer_destroySession(sessionId)
	//fmt.Println("recognizer_destroySession: ",ret5)


	size111:=C.print111(result1)
	var resultA unsafe.Pointer =(unsafe.Pointer) (result1)
	resultbb:=C.GoBytes(resultA, C.int(size111))
	return string(resultbb)

	//return C.GoString(result1)
}

func resumeFile(sessionId unsafe.Pointer) {
	f, error := os.Open("000000.wav")
	defer f.Close()
	if error != nil {
		fmt.Print("000000.wav not exist")
		return
	}

	r := bufio.NewReader(f)
	b1 := make([]byte, 3200)
	for {
		lenl, err := r.Read(b1)
		if err != nil {
			break
		}
		b2 := (*C.uchar)(unsafe.Pointer(&b1[0]))
		ret33 := C.recognizer_resumeSession(sessionId, b2, C.uint(lenl));
		fmt.Println("recognizer_resumeSession: ", ret33)
	}
}
