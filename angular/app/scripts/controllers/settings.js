'use strict';
angular.module('angularApp')
  .controller('SettingsCtrl', function ($scope,userManagement) {
  	//user
    if(userManagement.getLoggedIn() === false){
      userManagement.retrieve();
    }
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();

  });