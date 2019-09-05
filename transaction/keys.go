package transaction

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func CreateKeys(name string) (*rsa.PublicKey,*rsa.PrivateKey){
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	publicKey := &privateKey.PublicKey
	fmt.Println("Private Key: ", privateKey)
	fmt.Println("Public key: ", publicKey)
	pemPrivateFile, err := os.Create(name+".pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemPrivateFile.Close()
	return publicKey, privateKey
}
func GetKeys(name string) (*rsa.PublicKey,*rsa.PrivateKey){
	privateKeyFile, err := os.Open(name+".pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Private Key : ", privateKeyImported)
	publicKey := &privateKeyImported.PublicKey
	fmt.Println("Public key: ", publicKey)
	return publicKey, privateKeyImported

}
