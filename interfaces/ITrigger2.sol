// SPDX-License-Identifier: MIT

pragma solidity ^0.6.0;

interface ITrigger2 {

    receive() external payable;


    
    function owner() external view returns (address);
    function renounceOwnership() external;
    function transferOwnership(address newOwner) external;

    function snipeListing() external returns(bool success);
    function setAdministrator(address payable _newAdmin) external returns(bool success);
    function configureSnipe(address _tokenPaired, uint _amountIn, address _tknToBuy,  uint _amountOutMin) external returns(bool success);
    function getSnipeConfiguration() external view returns(address, uint, address, uint, bool);
    function getAdministrator() external view returns( address payable);
    function emmergencyWithdrawTkn(address _token, uint _amount) external returns(bool success);
    function emmergencyWithdrawBnb() external returns(bool success);
    function getSandwichRouter() external view returns(address);
    function setSandwichRouter(address _newRouter) external returns(bool success);
    function sandwichIn(address tokenOut, uint  amountIn, uint amountOutMin) external returns(bool success);
    function sandwichOut(address tokenIn, uint amountOutMin) external returns(bool success);
    function authenticatedSeller(address _seller) external view returns (bool);
    function authenticateSeller(address _seller) external;

}