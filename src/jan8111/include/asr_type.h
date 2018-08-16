/* =======================================================================================
*  Project          Auto Speech Recognition (ASR)
*  (c) Copyright    2014-2017
*  Company          Shanghai UVoice Technology CO., LTD
*                   All rights reserved
*  Secrecy Level    STRICTLY CONFIDENTIAL
* --------------------------------------------------------------------------------------*/
/**
 *  @internal
 *  @file asr_type.h
 *
 *  It includes general purpose definitions for all ASR submodules e.g. MACRO, TYPE
 *  definitions and etc.
 */
/*======================================================================================*/
/** @addtogroup DEF
*  @{*/
#ifndef _UNIVOICE_ASR_TYPE_H_
#define _UNIVOICE_ASR_TYPE_H_

////////////////////////////////////////////////////////////////////////////////
// type definition

typedef signed char         Int8;
typedef unsigned char       UInt8;
typedef signed short        Int16;
typedef unsigned short      UInt16;
typedef signed int          Int32;
typedef unsigned int        UInt32;
typedef unsigned short      WideChar;
typedef float               Float32;
typedef double              Float64;
#if defined WIN32
typedef __int64             Int64;
typedef unsigned __int64    UInt64;
#elif defined __GNUC__
typedef signed long long    Int64;
typedef unsigned long long  UInt64;
#else
#   error "Firstly, please define the 64bit integer"
#endif

#ifndef HANDLE
typedef void* HANDLE;
#endif

typedef enum
{
    RECOGNIZER_NOSPEECH = 3,
    RECOGNIZER_BOS = 2,
    RECOGNIZER_RESULT = 1,
    RECOGNIZER_OK = 0,
    RECOGNIZER_ERROR = -1,
    RECOGNIZER_BAD_STATE = -2,
    RECOGNIZER_BAD_POINTER = -3,
    RECOGNIZER_BAD_DATA = -4,
    RECOGNIZER_NO_MEMORY = -5,
    RECOGNIZER_NO_EXIST = -6,
    RECOGNIZER_EXIST = -7,
    RECOGNIZER_BAD_TYPE = -8,
    RECOGNIZER_SESSION_RUNNING = -9,
    RECOGNIZER_INIT_ERROR = -10,
    RECOGNIZER_THREAD_ERROR = -11,
    RECOGNIZER_INPUT_SIZE_ERROR = -12,
    RECOGNIZER_DECODER_IN_USE = -13,
    RECOGNIZER_OUT_OF_DATE = -14,
    RECOGNIZER_NO_LICENSEACCREDIT = -15,
} RecogRetCode;

typedef enum
{
    RecogResTypeNoResult = 0,
    RecogResTypePartial  = 1,
    RecogResTypeFull     = 2,
    RecogResTypeNBest    = 3,
    RecogResTypeLattice  = 4,
} RecogResType;

typedef enum
{
    kWfstUnknown  = 0,    //unused
    kWfst         = 1,    //for decoder and rescore
    kWfstCompress = 2,    //only for rescore
    kBnf          = 3,    //only for slot
    kListLine     = 4,    //only for slot
    kListJson     = 5,    //only for slot
//    kListXml      = 6,    //only for slot
	kListJsonMem  = 7,    //only for slot
} WfstType;

typedef struct Univoice_Acoustic_Param
{
    Int32 cpu_batch_size;
    Int32 sq_snr_estimate;              // 1: turn on, else: turn off, default: 0
    Int32 sq_clipping_dectect;          // 1: turn on, else: turn off, default: 0
} UnivoiceAcousticParam;

typedef struct Univoice_Decoder_Param
{
    Float32 lmScale;          // language model scale
    Float32 amScale;          // acoustic model scale
    Int32   maxTSN;
    Int32   minTSN;
    Float32 beamWidth;
    Float32 wordBeam;
    Float32 wordPenalty;
    Float32 loopPenalty;
    Float32 transPenalty;
    Int32   nBest;
    Int32   epsTransLimit;
    Int32   genPartialResult;    // if generate partial result during decodeing
    Int32   genPartialFrame;
    Int32   debug;
    Float32 blankSkip;
} UnivoiceDecoderParam;

typedef struct Univoice_Vad_Param
{
    Int32 vad_type;                     // 0: no vad
                                        // 1: vad_mode_1 (advanced, more accurate but more calculations), 
                                        // 2: vad_mode_2 (legacy)
                                        
    Float32 vad_threshold;              // default: 0.5
    Int32 vad_ms_lead_sil_timeout;      // default: 5000 ms(only for vad_mode_1)
    Int32 vad_ms_beg_acc_sph;           // default: 100 ms for vad_mode_1; 200 ms for vad_mode_2
    Int32 vad_ms_beg_acc_sph_lead_sil;  // default: 600 ms(only for vad_mode_1) 
    Int32 vad_ms_end_cont_sie;          // default: 1000 ms for vad_mode_1; 500 ms for vad_mode_2
} UnivoiceVadParam;

typedef enum E_Vad_Event_Tag
{
    E_Vad_Event_Null = 0,                   // Null event
    E_Vad_Event_Lst,                        // Leading silence timeout means no speech
    E_Vad_Event_Bos,                        // begin of speech
    E_Vad_Event_Eos,                        // end of speech
    E_Vad_Event_Over,                       // already have E_Vad_Event_Lst or E_Vad_Event_Eos, push wav continuously, you will get E_Vad_Event_Over
} E_Vad_Event;

typedef struct Univoice_License_Param
{
    UInt8 * p_ServerIpAddr;                 // the ip address of license server
    UInt32 n_ServerIpAddrLen;               // the length of server ip address
    UInt16 n_ServerPort;                    // the port of license server
    UInt8 * p_LocalIpAddr;                  // local ip address
    UInt32 n_LocalIpAddrLen;                // the length of local ip address
    Int32 i_BusinessType;                   // business type
} UnivoiceLicenseParam;


// log level
enum Log_Priority { LOG_INFO = 1, LOG_WARNING, LOG_ERROR, LOG_USER };

#ifdef __cplusplus
extern "C" {
#endif
typedef void (*PrintFunc)(const char*, int);
#ifdef __cplusplus
}  // extern "C"
#endif

typedef struct UnivoiceSessionCallback_S
{
    // Save feature stream format data
    Int32(* onFeatureData)(void* pInst, UInt8* data, UInt32 size);

    // Save acoustic stream format data
    Int32(* onAcousticData)(void* pInst, UInt8* data, UInt32 size);

    // Asr Vad Event
    Int32(* onRecogVadEvent)(void* pInst, E_Vad_Event vadEventType);
} UnivoiceSessionCallback;

#endif
/**@}*/
