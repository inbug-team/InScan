/*
负责人员：InBug Team
创建时间：2021/3/31
程序用途：可执行程序生成器动态传递
*/
package generate_exec

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
)

/*
1.将参数序列化成字符串
2.将字符串加密
3.获取模板文件长度
4.将模板文件字节存入可执行程序
5.将加密字符串追加到可执行程序末尾
6.将模板文件长度转化为4字节存到末尾
*/

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func IntToBytes(i int) []byte {
	buf := make([]byte, 4)
	if runtime.GOOS == "windows" {
		binary.LittleEndian.PutUint32(buf, uint32(i))
	} else {
		binary.BigEndian.PutUint32(buf, uint32(i))
	}
	return buf
}

func BytesToInt(buf []byte) int {
	return int(binary.BigEndian.Uint32(buf))
}

func GetExecPath() string {
	path, err := exec.LookPath(os.Args[0])
	CheckError(err)
	return path
}

// 补码
// AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// AES CBC 加密
func AesEncrypt(orig string, key string) string {
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
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryData := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryData, origData)
	return base64.StdEncoding.EncodeToString(cryData)
}

// AES CBC 解码
func AesDecrypt(cry string, key string) string {
	// 转成字节数组
	cryByte, _ := base64.StdEncoding.DecodeString(cry)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryByte))
	// 解密
	blockMode.CryptBlocks(orig, cryByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

func AppendExec(data interface{}, tplPath, execPath string) {
	// 1.读取模板内容
	dataByte, err := json.Marshal(data)
	CheckError(err)
	tplFile, err := os.OpenFile(tplPath, os.O_RDONLY, 0666)
	CheckError(err)
	defer tplFile.Close()
	fileInfo, err := tplFile.Stat()
	CheckError(err)
	size := fileInfo.Size()
	tplBuf := make([]byte, size)
	tplFile.Read(tplBuf)

	// 2.写入新的可执行程序
	execFile, err := os.Create(execPath)
	CheckError(err)
	defer execFile.Close()
	execFile.Write(tplBuf)
	execFile.WriteString(AesEncrypt(string(dataByte), "1234567890123456"))
	execFile.Write(IntToBytes(int(size)))
}
