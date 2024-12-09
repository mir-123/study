package controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// const version = "1.0.13"

// region encrypt

// 加密
// 密钥长度应为 128，192，256
func EncryptAES(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src = PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 填充字符串（末尾）
func PaddingText(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	return newPaddingStr
}
