// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.8.3;

contract Capsule {

    struct messageInfo{
        address Creator;
        uint256 CreateDate;          
        uint256 UnblockDate;            
        bool Status;             
        string Plaintext;         
        string Ciphertext; 
        bytes32 Hash;
        }
    

    mapping(address => uint256) private _nonce; 
    
    mapping(address => mapping(uint256 => bytes32)) private _index;

    mapping(bytes32 => messageInfo) private _mess;
    
    event Push(address who, uint256 index);
    event Unblock(address who, uint256 index);
    
    string public version;


    constructor () payable {
        
        version = "1.0.0";

    }
    

    function pushMessage( string memory ciphertext, uint256 time, bytes32 hash) public {
        require(time > (block.timestamp + 5 minutes),"Too close");
        
        
        messageInfo memory m_ ;
        
        bytes32 timeHash_ = sha256(abi.encodePacked(block.timestamp, msg.sender));
        bytes32 id_ = keccak256(abi.encodePacked(hash, timeHash_));
        
        require(_mess[id_].CreateDate == 0,"Existed");
        
        m_.Creator = msg.sender;
        m_.CreateDate = block.timestamp;
        m_.UnblockDate = time;
        m_.Status = false;
        m_.Ciphertext = ciphertext;
        m_.Plaintext= "";
        m_.Hash = hash;
        
        uint256 nonce = _nonce[msg.sender];
        
        _index[msg.sender][nonce] = id_;
        _nonce[msg.sender] = nonce + 1;
        
        _mess[id_] = m_;
        
        emit Push(msg.sender, nonce);
        
    }
    
    function UnblockMessage(uint256 index, string memory plaintext) public{
        require(_mess[_index[msg.sender][index]].UnblockDate < block.timestamp,"Not yet");
        require(_mess[_index[msg.sender][index]].Status == false, "Already");
        
        bytes32 hash = keccak256(abi.encodePacked(plaintext,msg.sender));
        require(_mess[_index[msg.sender][index]].Hash == hash, "Mismatch");
        
        
        _mess[_index[msg.sender][index]].Plaintext = plaintext;
        _mess[_index[msg.sender][index]].Status = true;
        
        emit Unblock(msg.sender, index);
    }
    
    function getMessageByIndex(address blogger, uint256 index) public view returns (string memory , uint256, uint256 , string memory, bytes32){
        
        return (_mess[_index[blogger][index]].Plaintext,
        _mess[_index[blogger][index]].CreateDate, 
        _mess[_index[blogger][index]].UnblockDate,
        _mess[_index[blogger][index]].Ciphertext,
        _mess[_index[blogger][index]].Hash);
        
    }
    
    function getMessageByHash(bytes32 id) public view returns (string memory , uint256, uint256 , string memory, bytes32){
        
        return (_mess[id].Plaintext,
        _mess[id].CreateDate, 
        _mess[id].UnblockDate,
        _mess[id].Ciphertext,
        _mess[id].Hash);
        
    }
    
}
