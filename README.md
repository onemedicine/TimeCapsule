# TimeCapsule
Ethereum Time Capsules

Ropsten contract address:0xa90681A17030DC91E4402A9815CA0Ae49911F76e

onedapp: https://oneclickdapp.com/vienna-camera

## Process

> Data encryption and decryption off-chain;
> The demo uses the AES algorithm, and the key is from the user's R value after signing the Plaintext Hash and Address:
 

 **1. Send a message**
 
 Hash  = Keccak256Hash(Plaintext + user's address)
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
