package testmodel

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/status-im/keycard-go/hexutils"

	"log"
	"testing"
)


func TestCrypto(t *testing.T)  {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	origData := []byte("Genesis Message") 
	dataHash := crypto.Keccak256Hash(origData).Bytes()
	hash := crypto.Keccak256Hash(dataHash, hexutils.HexToBytes("96216849c49358B10257cb55b28eA603c874b05E"))
	log.Println(hash.Hex()) // 0x4dc1102a2a66f9b1b3d07da46f1eaf617a79216f700458935989e3a7a11213d2

  // bytes32 hash = keccak256(abi.encodePacked(plaintext,msg.sender));
	// hash := crypto.Keccak256Hash(origData, hexutils.HexToBytes("96216849c49358B10257cb55b28eA603c874b05E"))
	// log.Println(hash.Hex()) // 0x5b053cba2bf2de145ec144ee7629418c8fe409a9b482a465048176dda42e5815

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	r := make([]byte, 32)

	copy(r, signature[0:32])

	log.Println(hexutil.Encode(signature)) // 0xb295c2fc32aebb8d71e007d9f2ffec40a1560907b3e1cb00fbd56624614dcd1237d4401fcfa0ab389b633c2eab8c6917ef929e831077073ff0abb48dd8919e0e01
	log.Println(hexutil.Encode(r)) // 0xb295c2fc32aebb8d71e007d9f2ffec40a1560907b3e1cb00fbd56624614dcd12

	key:= r // Encrypted key
	log.Println("Plaintext：", string(origData))

	log.Println("AES CBC model")
	encrypted := AesEncryptCBC(origData, key)
	log.Println("ciphertext(hex)：", hex.EncodeToString(encrypted)) // 42743acc8b782129f88555f33412ca0a
	log.Println("ciphertext(base64)：", base64.StdEncoding.EncodeToString(encrypted)) // QnQ6zIt4ISn4hVXzNBLKCg==
	decrypted := AesDecryptCBC(encrypted, key)
	log.Println("Decryption result：", string(decrypted)) //Genesis Message
}


func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {

	// k len :16, 24 or 32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted = make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted = make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
