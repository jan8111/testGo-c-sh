package shRecog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func buildConfig() SHConfigs{
	bb, _ := ioutil.ReadFile(os.Args[1])
	var configs SHConfigs
	err := json.Unmarshal(bb, &configs)
	if err != nil {
		fmt.Println(err.Error())
	}
	return configs
}

type Acoustic struct {
	Name string
	Path string
	Type string
	Device int
}
type Decoder struct {
	Name,Path,Type string
}

type SHContext struct {
	_context_ptr int
	AcousticName string
	ContextCode string
	BizMode int
	UnivoiceAcousticParam UnivoiceAcousticParam
	ContextDecoders []ContextDecoder
}

type UnivoiceAcousticParam struct {
	Cpu_batch_size int
	Sq_snr_estimate int
	Sq_clipping_dectect int
}

type UnivoiceLicenseParam struct {
	P_ServerIpAddr string
	N_ServerPort int
	I_BusinessType int
	P_LocalIpAddr string
}

type ContextDecoder struct {
	DecoderName string
	BSlot bool
	Rescore string
}

type SHConfigs struct {
	UnivoiceLicenseParam UnivoiceLicenseParam
	Acoustics  []Acoustic
	Decoders   []Decoder
	ShContexts []SHContext
}
