package main

import (
	"bytes"
	"errors"
)

const (
	NOPADDING = iota
	PKCS5PADDING
)

//PKCS5补位
func PKCS5Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

//去除PKCS5补位
func PKCS5Unpadding(text []byte) []byte {
	length := len(text)
	padtext := int(text[length-1])
	return text[:(length - padtext)]
}

//补位方法
func Padding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("text length invalid")
		}
	case PKCS5PADDING:
		return PKCS5Padding(text, 8), nil
	default:
		return nil, errors.New("padding type error")
	}

	return text, nil
}

//去除补位方法
func UnPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("text length invalid")
		}
	case PKCS5PADDING:
		return PKCS5Unpadding(text), nil
	default:
		return nil, errors.New("padding type error.")
	}
	return text, nil
}
