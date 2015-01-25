'use strict';
angular.module('angularApp')
  .controller('ProfileCtrl', function ($scope, $location, userManagement) {
  	//user:
  	if(userManagement.getLoggedIn() === false){
  	  userManagement.retrieve();
  	}
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();
  	$scope.userMail = userManagement.getUserMail();

  	$scope.logout = function(){
  		userManagement.setUserName('');
  		userManagement.setUserId(0);
  		userManagement.setUserMail('');
  		userManagement.setLoggedIn(false);
  		userManagement.store();
  		$location.path('#/');
  	};
  });