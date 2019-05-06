/* =======================================================================================
*  Project          Auto Speech Recognition (ASR)
*  (c) Copyright    2014-2017
*  Company          Shanghai UVoice Technology CO., LTD
*                   All rights reserved
*  Secrecy Level    STRICTLY CONFIDENTIAL
* --------------------------------------------------------------------------------------*/
/**
 *  @internal
 *  @file asr_api.h
 *
 *  Prototypes of the Auto Speech Recognition (ASR) API functions.
 *
 *  This header file contains all prototypes of the API functions
 *  of the Auto Speech Recognition (ASR) module.
 */
/*======================================================================================*/
/** @addtogroup DEF
*  @{*/
#ifndef _UNIVOICE_ASR_API_H_
#define _UNIVOICE_ASR_API_H_

#include "asr_type.h"

#ifdef __cplusplus
extern "C"
{
#endif

/////////////////////////////////////////////////////////////////////////////////////////////////
//global api
RecogRetCode recognizer_setWorkPath(const char* path);
RecogRetCode recognizer_setDicPath(const char* path);
const char* recognizer_getVersion();
void recognizer_setOutputFanti();
void recognizer_setPrintFunc(PrintFunc pFunc);
RecogRetCode recognizer_getModel(const char* type, unsigned char **model, unsigned int *modelLen);
RecogRetCode recognizer_init(const UnivoiceLicenseParam* param);
RecogRetCode recognizer_unInit();
RecogRetCode recognizer_setActivateHandle(void* pv_activate_handle);

//model api
RecogRetCode recognizer_addAcoustic(const char* name, const char* type, const char* path, int device);
RecogRetCode recognizer_addAcousticMem(const char* name, const char* type, const unsigned char* gram, const int gramLen);
RecogRetCode recognizer_delAcoustic(const char* name);
RecogRetCode recognizer_getAcousticInfo(const char* name, AcousticInfo acoustic_info, void* value);

//ret: 1-yes 0-no
Int32 recognizer_hasAcoustic(const char* name);
const char* recognizer_getAcousticVersion(const char* name);
//type:"wfst","wfst-compress","bnf","list_line","list_json","list_xml","list_json_mem"
//slot name format "grammar#slot" or "#slot"
RecogRetCode recognizer_addDecoder(const char* name, const char* type, const char* path);
RecogRetCode recognizer_addDecoderMem(const char* name, const char* type, const unsigned char* gram, const int gramLen);
RecogRetCode recognizer_delDecoder(const char* name);
RecogRetCode recognizer_renameDecoder(const char* old_name, const char* new_name);
RecogRetCode recognizer_aliasDecoder(const char* src_name, const char* alias_name);
int recognizer_getDecoderStateNum(const char* name);
//ret: 1-yes 0-no
Int32 recognizer_hasDecoder(const char* name);
const char* recognizer_getDecoderVersion(const char* name);

RecogRetCode recognizer_addPuncModel(const char* name, const char* type, const char* path);
RecogRetCode recognizer_delPuncModel();

/////////////////////////////////////////////////////////////////////////////////////////////////
//context api
RecogRetCode recognizer_createContext(void** base);
void recognizer_destroyContext(void* base);
RecogRetCode recognizer_setContextAcoustic(void* base, const char* acoustic_name);
RecogRetCode recognizer_attachContextDecoder(void* base, const char* decoder_name, _Bool bSlot, float weight);
RecogRetCode recognizer_detachContextDecoder(void* base, const char* decoder_name, _Bool bSlot);
RecogRetCode recognizer_setContextRescore(void* base, const char* decoder_name, const char* rescore_name);
RecogRetCode recognizer_setContextAcousticParam(void* base, const UnivoiceAcousticParam* param);
RecogRetCode recognizer_getContextAcousticParam(void* base, const UnivoiceAcousticParam* param);
RecogRetCode recognizer_setContextDecoderParam(void* base, const char* decoder_name, const UnivoiceDecoderParam* param);
RecogRetCode recognizer_getContextDecoderParam(void* base, const char* decoder_name, const UnivoiceDecoderParam* param);
RecogRetCode recognizer_setContextVadParam(void* base, const UnivoiceVadParam* param);
RecogRetCode recognizer_getContextVadParam(void* base, const UnivoiceVadParam* param);

/////////////////////////////////////////////////////////////////////////////////////////////////
//session api
RecogRetCode recognizer_createSession(void** session, void* base);
void recognizer_destroySession(void* session);
RecogRetCode recognizer_setSessionCallBack(void* session, UnivoiceSessionCallback* callback, void* pInst);
RecogRetCode recognizer_startSession(void* session, const Int32 nThreadIndex);
//Be careful! Should only input 10ms data(320 bytes for 16KHz, or 160 bytes for 8Khz) !!!
RecogRetCode recognizer_resumeSession(void* session, const UInt8* data, UInt32 len);
RecogRetCode recognizer_resumeSessionWithFeature(void* session, const UInt8* data, UInt32 len);
RecogRetCode recognizer_resumeSessionWithAcoustic(void* session, const UInt8* data, UInt32 len);
RecogRetCode recognizer_stopSession(void* session, _Bool bFullSlot);
RecogRetCode recognizer_getSessionResStr(void* session, const char** result);
RecogRetCode recognizer_getSessionResJson(void* session, const char** result);
RecogRetCode recognizer_getSessionResPhone(void* session, const char** result);
RecogRetCode recognizer_getSessionAnalyzeResStr(void* session, const char** result);
RecogRetCode recognizer_getSessionAnalyzeResJson(void* session, const char** result);
RecogRetCode recognizer_getSessionLat(void* session, const char** latticeOut);
RecogRetCode recognizer_getSessionLatJson(void* session, const char** latticeOut);
RecogRetCode recognizer_getSessionNBest2(void* session, const char** * nbestStr, int* num);
RecogRetCode recognizer_getSessionNBestJson(void* session, const char** nbestStr);
RecogRetCode recognizer_abortSession(void* session);
RecogRetCode recognizer_getSessionLmScore(void* session, const char *pInJson, const char** result);

/////////////////////////////////////////////////////////////////////////////////////////////////
//lexion
//word should be utf-8 format, phone can use pingyin or phoneset
//only available when use FstMaker lib
void recognizer_initFst();
void recognizer_destroyFst();
const char* recognizer_getFstMakerVersion();
void recognizer_setLexicon(char* word, char* phone);
const char* recognizer_getLexicon(char* word);
void recognizer_clearLexicon();
_Bool recognizer_getFstBreakFlag();
void recognizer_setFstBreakFlag(_Bool bFlag);
#ifdef __cplusplus
}
#endif

#endif
/**@}*/
