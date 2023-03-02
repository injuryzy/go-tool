package ursa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 使用公钥进行加密
func RSA_encrypter(path string, msg []byte) []byte {
	//首先从文件中提取公钥
	fp, _ := os.Open(path)
	defer fp.Close()
	//测量文件长度以便于保存
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	//下面的操作是与创建秘钥保存时相反的
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码,得到一个interface类型的pub
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//加密操作,需要将接口类型的pub进行类型断言得到公钥类型
	cipherText, _ := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), msg)
	return cipherText
}

// 使用私钥进行解密
func RSA_decrypter(path string, cipherText []byte) []byte {
	//同加密时，先将私钥从文件中取出，进行二次解码
	fp, _ := os.Open(path)
	defer fp.Close()
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	block, _ := pem.Decode(buf)
	PrivateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//二次解码完毕，调用解密函数
	afterDecrypter, _ := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, cipherText)
	return afterDecrypter
}

//func main() {
//	//加密
//	data := []byte("hello world")
//	encrypt := RSA_Encrypt(data, "public.pem")
//	fmt.Println(string(encrypt))
//
//	// 解密
//	decrypt := RSA_Decrypt(encrypt, "private.pem")
//	fmt.Println(string(decrypt))
//}
