package shRecog

/*
#cgo CFLAGS : -I../include
#cgo LDFLAGS: -L/home/shhan/tw_model/deps -lrecognizer

#include "asr_api.h"
#include <stdlib.h>

*/
import "C"
import (
	"fmt"
	"unsafe"
	"os"
	"bufio"
	"log"
)

func Main() {
	if len(os.Args) <= 1 {
		log.Fatalf("usage: ./testSh-Engine ../resource/000000.wav")
	}

	contextId1 := initEngine()
	recog(contextId1)

}

func initEngine() []unsafe.Pointer {
	config := buildConfig()
	fmt.Printf("config: %v \n",config)

	ret := C.recognizer_getVersion()
	str1 := C.GoString(ret)
	fmt.Println("recognizer_getVersion: ", str1)

	workpath1 := C.CString(config.Work_path)
	defer C.free(unsafe.Pointer(workpath1))
	ret2 := C.recognizer_setWorkPath(workpath1)
	fmt.Println(config.Work_path, " recognizer_setWorkPath: ", ret2)

	for _, Acoustic1 := range config.Acoustics {
		name1 := C.CString(Acoustic1.Name)
		defer C.free(unsafe.Pointer(name1))
		type1 := C.CString(Acoustic1.Type)
		defer C.free(unsafe.Pointer(type1))
		path1 := C.CString(Acoustic1.Path)
		defer C.free(unsafe.Pointer(path1))
		ret4 := C.recognizer_addAcoustic(name1, type1, path1, 0)
		fmt.Println(Acoustic1, "recognizer_addAcoustic: ", ret4)
	}

	for _, Decoders1 := range config.Decoders {
		name2 := C.CString(Decoders1.Name)
		defer C.free(unsafe.Pointer(name2))
		type2 := C.CString(Decoders1.Type)
		defer C.free(unsafe.Pointer(type2))
		path2 := C.CString(Decoders1.Path)
		defer C.free(unsafe.Pointer(path2))
		ret5 := C.recognizer_addDecoder(name2, type2, path2)
		fmt.Println(Decoders1, "recognizer_addDecoder: ", ret5)
	}

	ret1 := make([]unsafe.Pointer, len(config.ShContexts))
	for i, ShContexts1 := range config.ShContexts {
		var contextId1 unsafe.Pointer
		ret3 := C.recognizer_createContext(&contextId1)
		fmt.Println("recognizer_createContext: ", ret3)

		AcousticName11 := C.CString(ShContexts1.AcousticName)
		defer C.free(unsafe.Pointer(AcousticName11))
		ret6 := C.recognizer_setContextAcoustic(contextId1,AcousticName11 )
		fmt.Println("recognizer_setContextAcoustic: ", ret6)

		param1 := new(C.UnivoiceAcousticParam)
		(*param1).cpu_batch_size = C.int(ShContexts1.UnivoiceAcousticParam.Cpu_batch_size)
		(*param1).sq_snr_estimate = C.int(ShContexts1.UnivoiceAcousticParam.Sq_snr_estimate)
		(*param1).sq_clipping_dectect = C.int(ShContexts1.UnivoiceAcousticParam.Sq_clipping_dectect)
		ret7 := C.recognizer_setContextAcousticParam(contextId1, param1)
		fmt.Println("recognizer_setContextAcousticParam: ", ret7)

		for _, decoder1 := range ShContexts1.ContextDecoders {
			DecoderName11 := C.CString(decoder1.DecoderName)
			defer C.free(unsafe.Pointer(DecoderName11))
			ret8 := C.recognizer_attachContextDecoder(contextId1,DecoderName11 , C._Bool(decoder1.BSlot))
			fmt.Println("recognizer_attachContextDecoder: ", ret8)
		}

		ret1[i] = contextId1
	}

	return ret1
}

func recog(contextIds []unsafe.Pointer) {
	for _, contextId1 := range contextIds {
		var sessionId unsafe.Pointer
		ret := C.recognizer_createSession(&sessionId, contextId1, 16000);
		defer C.recognizer_destroySession(sessionId)
		fmt.Println("recognizer_createSession: ", ret)

		ret2 := C.recognizer_startSession(sessionId, 0);
		fmt.Println("recognizer_startSession: ", ret2)

		resumeFile(sessionId)

		ret3 := C.recognizer_stopSession(sessionId);
		fmt.Println("recognizer_stopSession: ", ret3)

		var resultptr unsafe.Pointer
		result11 := (*C.char)(resultptr)
		ret4 := C.recognizer_getSessionResStr(sessionId, &result11)
		fmt.Println("recognizer_getSessionResStr: ", ret4)
		resultStr := C.GoString(result11)
		fmt.Println("result:", resultStr)
	}
}

func resumeFile(sessionId unsafe.Pointer) {
	var filename = os.Args[1]
	f, error := os.Open(filename)
	defer f.Close()
	if error != nil {
		fmt.Print(filename, " Open fail")
	}

	r := bufio.NewReader(f)
	b1 := make([]byte, 3200)
	for {
		lenl, err := r.Read(b1)
		if err != nil {
			break
		}
		b2 := (*C.uchar)(unsafe.Pointer(&b1[0]))
		//b2:=C.CBytes(b1)
		C.recognizer_resumeSession(sessionId, (*C.uchar)(b2), C.uint(lenl));
		//C.free(unsafe.Pointer(b2))
	}
}
