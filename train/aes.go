package train

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// AES-128 Iv向量获取
func GetAesIv(key string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	iv := make([]byte, block.BlockSize())
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	return hex.EncodeToString(iv), nil
}

// AES-128 CBC 加密
func Aes128CbcEncrypt(key string, iv string, data string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	ivBytes, err := hex.DecodeString(iv)
	if err != nil {
		return "", err
	}
	dataBytes := pad([]byte(data))
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCEncrypter(block, ivBytes)
	encrypted := make([]byte, len(dataBytes))
	blockMode.CryptBlocks(encrypted, dataBytes)

	return hex.EncodeToString(encrypted), nil
}

// AES-128 CBC 解密
func Aes128CbcDecrypt(key string, iv string, data string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	ivBytes, err := hex.DecodeString(iv)
	if err != nil {
		return "", err
	}
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	decrypted := make([]byte, len(dataBytes))
	blockMode.CryptBlocks(decrypted, dataBytes)

	decrypted = unpad(decrypted)
	return string(decrypted), nil
}

func pad(data []byte) []byte {
	// 添加PKCS#7填充
	blockSize := aes.BlockSize
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unpad(data []byte) []byte {
	// 移除PKCS#7填充
	length := len(data)
	unpadding := int(data[length-1])
	return data[:length-unpadding]
}

// AES - ECB 加密
func Aes128EcbEncrypt(orig string, key string) []byte {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := NewECBEncrypter(block)
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return cryted
}

func AesDecrypt(crytedByte []byte, key string) (string, error) {
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	if len(orig)%blockSize != 0 {
		return "", errors.New(fmt.Sprintf("fail to Notice, CryptBlocks"))
	}

	// 加密模式
	blockMode := NewECBDecrypter(block)

	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	var err error
	orig, err = PKCS7UnPadding(orig)
	if err != nil {
		return "", err
	}
	return string(orig), nil
}

// 补码
// AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length < 1 {
		return []byte{}, errors.New(fmt.Sprintf("fail to Notice, EcpmAesDecrypt1"))
	}
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte{}, errors.New(fmt.Sprintf("fail to Notice, EcpmAesDecrypt2"))
	}
	return origData[:(length - unpadding)], nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// returns a BlockMode which encrypts in electronic code book
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// returns a BlockMode which decrypts in electronic code book
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
