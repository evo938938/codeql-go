// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import (
	"compress/zlib"
	"io"
)

func TaintStepTest_CompressZlibNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader656 := sourceCQL.(io.Reader)
	intoReadCloser414, _ := zlib.NewReader(fromReader656)
	return intoReadCloser414
}

func TaintStepTest_CompressZlibNewReaderDict_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader518 := sourceCQL.(io.Reader)
	intoReadCloser650, _ := zlib.NewReaderDict(fromReader518, nil)
	return intoReadCloser650
}

func TaintStepTest_CompressZlibNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter784 := sourceCQL.(*zlib.Writer)
	var intoWriter957 io.Writer
	intermediateCQL := zlib.NewWriter(intoWriter957)
	link(fromWriter784, intermediateCQL)
	return intoWriter957
}

func TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter520 := sourceCQL.(*zlib.Writer)
	var intoWriter443 io.Writer
	intermediateCQL, _ := zlib.NewWriterLevel(intoWriter443, 0)
	link(fromWriter520, intermediateCQL)
	return intoWriter443
}

func TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter127 := sourceCQL.(*zlib.Writer)
	var intoWriter483 io.Writer
	intermediateCQL, _ := zlib.NewWriterLevelDict(intoWriter483, 0, nil)
	link(fromWriter127, intermediateCQL)
	return intoWriter483
}

func TaintStepTest_CompressZlibWriterReset_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriter989 := sourceCQL.(zlib.Writer)
	var intoWriter982 io.Writer
	fromWriter989.Reset(intoWriter982)
	return intoWriter982
}

func TaintStepTest_CompressZlibWriterWrite_B0I0O0(sourceCQL interface{}) interface{} {
	fromByte417 := sourceCQL.([]byte)
	var intoWriter584 zlib.Writer
	intoWriter584.Write(fromByte417)
	return intoWriter584
}

func TaintStepTest_CompressZlibResetterReset_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader991 := sourceCQL.(io.Reader)
	var intoResetter881 zlib.Resetter
	intoResetter881.Reset(fromReader991, nil)
	return intoResetter881
}

func RunAllTaints_CompressZlib() {
	{
		source := newSource(0)
		out := TaintStepTest_CompressZlibNewReader_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_CompressZlibNewReaderDict_B0I0O0(source)
		sink(1, out)
	}
	{
		source := newSource(2)
		out := TaintStepTest_CompressZlibNewWriter_B0I0O0(source)
		sink(2, out)
	}
	{
		source := newSource(3)
		out := TaintStepTest_CompressZlibNewWriterLevel_B0I0O0(source)
		sink(3, out)
	}
	{
		source := newSource(4)
		out := TaintStepTest_CompressZlibNewWriterLevelDict_B0I0O0(source)
		sink(4, out)
	}
	{
		source := newSource(5)
		out := TaintStepTest_CompressZlibWriterReset_B0I0O0(source)
		sink(5, out)
	}
	{
		source := newSource(6)
		out := TaintStepTest_CompressZlibWriterWrite_B0I0O0(source)
		sink(6, out)
	}
	{
		source := newSource(7)
		out := TaintStepTest_CompressZlibResetterReset_B0I0O0(source)
		sink(7, out)
	}
}
