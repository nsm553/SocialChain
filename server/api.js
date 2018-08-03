const Web3 = require("web3");
const web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
var contract = require("../contracts/build/contracts/nonProfit.json");

const address = contract.networks[
    Object.keys(contract.networks)[0]
].address;

var nonProfitInstance = new web3.eth.Contract(contract.abi, address);
const maxGas = 150000;
const etherToWei = 1000000000000000000;

var exports = module.exports = {};

exports.completeRequest = function (id, donorId, amount) {
    var data = nonProfitInstance.methods.completeRequest(id, donorId)
      .send({from: donorId, value:(amount * etherToWei)})
      .on('receipt', function(receipt){
        console.log("Completed!")
    })
};

exports.createRequest = function (vendorId, nonProfitId, requesterId, amount) {
    var data = nonProfitInstance.methods.createRequest(vendorId, nonProfitId)
      .send({from: requesterId})
      .on('receipt', function(receipt){
        console.log("Request Created!")
        console.log(receipt)
        console.log(receipt['events']['RequestCreated']['returnValues'][0])
        nonProfitInstance.methods.setAmount(receipt['events']['RequestCreated']['returnValues'][0], amount).
          send({from: requesterId}).on('error', console.error)
      })
      .on('error', console.error);
};

exports.getRequestsByNPId = function (Id, requesterId) {
    nonProfitInstance.methods.getRequestsByNonProfitId(Id).call({from: requesterId}).then(function(result){
      var requests = [];
      var request = {}
      for (i=0; i<result[0].length; i++) {
          request = {}
          request['id'] = result['0'][i]
          request['vendor'] = result['1'][i]
          request['donor'] = result['2'][i]
          request['amount'] = result['3'][i]
          request['status'] = result['4'][i]
          requests.push(request)
      }
      console.log(requests);
    });
}

nonProfitId = '0x7178a5cc39d1ccd359383099ad746bb4d9144778';
vendorId = '0xea461c69a1876d3590ed2d1964761c44ff90289d';
requesterId = '0x2ad2c60524a6c30185898d16bceb67ec63ef4a6a';
donorId = '0x8fece4b1bfcefd306853d1f7759c73150cf80ff7';

web3.eth.getBalance(nonProfitId, function (error, result) {
    if (!error) {
        console.log("Non Profit");
        console.log(web3.utils.fromWei(result));
    } else {
        console.error(error);
    }
});

web3.eth.getBalance(vendorId, function (error, result) {
    if (!error) {
        console.log("Vendor");
        console.log(web3.utils.fromWei(result));
    } else {
        console.error(error);
    }
});

web3.eth.getBalance(donorId, function (error, result) {
    if (!error) {
        console.log("Donor");
        console.log(web3.utils.fromWei(result));
    } else {
        console.error(error);
    }
});

web3.eth.getBalance(requesterId, function (error, result) {
    if (!error) {
        console.log("Requester");
        console.log(web3.utils.fromWei(result));
    } else {
        console.error(error);
    }
});

//exports.createRequest(vendorId, nonProfitId, requesterId, 5);
//exports.getRequestsByNPId(nonProfitId, requesterId);
//exports.completeRequest(0, donorId, 5);
