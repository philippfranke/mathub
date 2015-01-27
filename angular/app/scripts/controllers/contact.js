'use strict';
angular.module('angularApp')
  .controller('ContactCtrl', function ($scope,userManagement) {
  	//user
    if(userManagement.getLoggedIn() === false){
      userManagement.retrieve();
    }
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();

  });