package main

/*
#cgo CFLAGS : -I../include
#cgo LDFLAGS: -L/home/shhan/tw_model/deps -lrecognizer

#include "asr_api.h"
#include "asr_type.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	ret :=C.recognizer_getVersion()
	str1:=C.GoString(ret)
	fmt.Println("recognizer_getVersion: ", str1)

	workpath1:=C.CString("/home/shhan/tw_model/")
	ret2 :=C.recognizer_setWorkPath(workpath1)
	C.free(unsafe.Pointer(workpath1))
	fmt.Println("recognizer_setWorkPath: ",ret2)

	name1:=C.CString("ctc")
	type1:=C.CString("dnn")
	path1:=C.CString("/home/shhan/tw_model/model/final_0808_online.bin")
	ret4:=C.recognizer_addAcoustic(name1,type1,path1,0)
	fmt.Println("recognizer_addAcoustic: ",ret4)


	name2:=C.CString("first-path")
	type2:=C.CString("wfst")
	path2:=C.CString("/home/shhan/tw_model/model/xiaoi_hotword_test_0326_5.65e-9_ctc.dat")
	ret5:=C.recognizer_addDecoder(name2,type2,path2)
	fmt.Println("recognizer_addDecoder: ",ret5)

	//ptr1 :=uintptr(0)
	//base1:= unsafe.Pointer(ptr1)
	var base1 unsafe.Pointer
	ret3:=C.recognizer_createContext(&base1)
	fmt.Println("recognizer_createContext: ",ret3)

	ret6:=C.recognizer_setContextAcoustic(base1,name1)
	fmt.Println("recognizer_setContextAcoustic: ",ret6)

	ret8:=C.recognizer_attachContextDecoder(base1,name2,false)
	fmt.Println("recognizer_attachContextDecoder: ",ret8)

	param1 := new (C.UnivoiceAcousticParam)
	(*param1).cpu_batch_size= 40
	(*param1).sq_snr_estimate= 0
	(*param1).sq_clipping_dectect= 0
	ret7:=C.recognizer_setContextAcousticParam(base1,param1)
	fmt.Println("recognizer_setContextAcousticParam: ",ret7)

}