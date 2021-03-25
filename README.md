# TimeCapsule
Ethereum Time Capsules

Ropsten contract address:0x45681A97c22c661C4234490ED7d9F84348b39dD5

onedapp: https://oneclickdapp.com/pigment-parlor

## Process

> Data encryption and decryption off-chain;
> The demo uses the AES algorithm, and the key is from the user's R value after signing the Plaintext Hash and Address:
 

 **1. Send a message**
 
 DataHash=Keccak256Hash(Plaintext)
 Hash  = Keccak256Hash(DataHash + user's address)
 signature = Sign(Hash)
 AESkey = signature.r
 ciphertext = AesEncryptCBC(Plaintext,AESkey)

 CapsuleContract.pushMessage( ciphertext, now + 1 year , Hash) 
 with event Push()
 
 

 **2. Unblock message**

 ciphertext , Hash = CapsuleContract.getMessageByIndex()
 or  CapsuleContract.getMessageByHash()
 
 signature = Sign(Hash)
 AESkey = signature.r
 Plaintext = AesDecryptCBC(ciphertext , AESkey)
 CapsuleContract.UnblockMessage(index, Plaintext) 
 with  event Unblock()
 

 **3. View message**
 
  Plaintext  = CapsuleContract.getMessageByIndex()
 or  CapsuleContract.getMessageByHash()
 
 
 
## Interface
 pushMessage(string memory ciphertext, uint256 time, bytes32 hash)
 UnblockMessage(uint256 index, string memory plaintext)
 getMessageByHash(address blogger, uint256 index)
 getMessageByIndex(bytes32 id)
 version 
