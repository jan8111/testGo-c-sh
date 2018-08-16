package shRecog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func buildConfig() SHConfigs{
	bb, _ := ioutil.ReadFile("../resource/sh-config.json")
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

type ContextDecoder struct {
	DecoderName string
	BSlot bool
	Rescore string
}

type SHConfigs struct {
	Work_path  string
	DicPath    string
	Acoustics  []Acoustic
	Decoders   []Decoder
	ShContexts []SHContext
}
