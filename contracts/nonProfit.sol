pragma solidity ^0.4.22;


contract nonProfit {
  enum Status { Available, Completed }

  uint constant ETHER_TO_WEI = 1000000000000000000;
  uint public totalRequests;

  event RequestCreated(uint);

  mapping (uint => Request) requests;

  struct Request {
    address vendorId;
    address donorId;
    address nonprofitId;
    uint amount;
    Status status;
  }

  constructor() public {
		totalRequests = 0;
	}

  function createRequest(address vendorId, address nonProfitId) public returns (uint) {
    totalRequests++;
    uint current = totalRequests - 1;
    requests[current].vendorId = vendorId;
    requests[current].nonprofitId = nonProfitId;
    emit RequestCreated(current);
		return current;
  }

  function setAmount(uint id, uint amount) public {
    requests[id].amount = amount;
  }

  function completeRequest(uint id, address donorId) public payable {
    requests[id].status = Status.Completed;
    requests[id].donorId = donorId;
    requests[id].vendorId.transfer(ETHER_TO_WEI * requests[id].amount);
  }

  function getRequestsByNonProfitId(address Id) public constant returns (uint[], address[], address[], uint[], Status[]) {
    uint[] memory ids = new uint[](totalRequests);
    address[] memory vendorIds = new address[](totalRequests);
		address[] memory donorIds = new address[](totalRequests);
    uint[] memory amount = new uint[](totalRequests);
    Status[] memory status = new Status[](totalRequests);
    uint num = 0;

		for(uint i = 0; i<totalRequests; i++){
          if (requests[i].nonprofitId == Id) {
            ids[num] = i;
            vendorIds[num] = requests[i].vendorId;
  				  donorIds[num] = requests[i].donorId;
            amount[num] = requests[i].amount;
            status[num] = requests[i].status;
            num++;
          }

    }
		return (ids, vendorIds, donorIds, amount, status);
  }

  function getRequestsByDonorId(address Id) public constant returns (uint[], address[], address[], uint[], Status[]) {
    uint[] memory ids = new uint[](totalRequests);
    address[] memory vendorIds = new address[](totalRequests);
		address[] memory nonprofitIds = new address[](totalRequests);
    uint[] memory amount = new uint[](totalRequests);
    Status[] memory status = new Status[](totalRequests);
    uint num = 0;

		for(uint i = 0; i<totalRequests; i++){
          if (requests[i].donorId == Id) {
            ids[num] = i;
            vendorIds[num] = requests[i].vendorId;
  				  nonprofitIds[num] = requests[i].nonprofitId;
            amount[num] = requests[i].amount;
            status[num] = requests[i].status;
            num++;
          }

    }
		return (ids, vendorIds, nonprofitIds, amount, status);
  }

  function getRequestsByVendorId(address Id) public constant returns (uint[], address[], address[], uint[], Status[]) {
    uint[] memory ids = new uint[](totalRequests);
    address[] memory nonprofitIds = new address[](totalRequests);
    address[] memory donorIds = new address[](totalRequests);
    uint[] memory amount = new uint[](totalRequests);
    Status[] memory status = new Status[](totalRequests);
    uint num = 0;

    for(uint i = 0; i<totalRequests; i++){
          if (requests[i].vendorId == Id) {
            ids[num] = i;
            nonprofitIds[num] = requests[i].nonprofitId;
            donorIds[num] = requests[i].donorId;
            amount[num] = requests[i].amount;
            status[num] = requests[i].status;
            num++;
          }

    }
    return (ids, nonprofitIds, donorIds, amount, status);
  }

  function getAvailableRequestsByNonProfitId(address Id) public constant returns (uint[], address[], address[], uint[]) {
    uint[] memory ids = new uint[](totalRequests);
    address[] memory vendorIds = new address[](totalRequests);
		address[] memory donorIds = new address[](totalRequests);
    uint[] memory amount = new uint[](totalRequests);
    uint num = 0;

		for(uint i = 0; i<totalRequests; i++){
          if (requests[i].nonprofitId == Id ||
            requests[i].status == Status.Available) {
            ids[num] = i;
            vendorIds[num] = requests[i].vendorId;
  				  donorIds[num] = requests[i].donorId;
            amount[num] = requests[i].amount;
            num++;
          }

    }
		return (ids, vendorIds, donorIds, amount);
  }

  function getClosedRequestsByNonProfitId(address Id) public constant returns (uint[], address[], address[], uint[]) {
    uint[] memory ids = new uint[](totalRequests);
    address[] memory vendorIds = new address[](totalRequests);
    address[] memory donorIds = new address[](totalRequests);
    uint[] memory amount = new uint[](totalRequests);
    uint num = 0;

    for(uint i = 0; i<totalRequests; i++){
          if (requests[i].nonprofitId == Id ||
            requests[i].status == Status.Completed) {
            ids[num] = i;
            vendorIds[num] = requests[i].vendorId;
            donorIds[num] = requests[i].donorId;
            amount[num] = requests[i].amount;
            num++;
          }

    }
    return (ids, vendorIds, donorIds, amount);
  }

}
