'use strict';
/*jshint camelcase: false */
angular.module('angularApp')
  .controller('LoginCtrl', function ($scope, api, $location, userManagement) {
    //user:
    if(userManagement.getLoggedIn() === false){
      userManagement.retrieve();
    }
    $scope.loggedIn = userManagement.getLoggedIn();
    $scope.userName = userManagement.getUserName();
    $scope.userId = userManagement.getUserId();

    
  	$scope.error = '';
  	$scope.lemail = '';
  	$scope.remail = '';
  	$scope.lpassword_hash = '';
  	$scope.rpassword_hash = '';

  	$scope.rusername = '';

  	$scope.loginAction = function(){

  		$scope.error = '';

  		if($scope.lemail !== '' && $scope.lpassword_hash !== ''){
  			var sendData = {
  				'email': $scope.lemail,
  				'password_hash': $scope.lpassword_hash
  			};
	  		api.loginUser(sendData)
				.success(function(data){
					userManagement.setUserName(data.name);
					userManagement.setUserId(data.id);
					userManagement.setUserMail(data.email);
					userManagement.setLoggedIn(true);
          userManagement.store();
					$location.path('#/');
				})
				.error(function(){
					$scope.error = 'Credentials not correct.';
				});
  		}else{
  			$scope.error = 'Credentials not complete.';
  		}
  	};
  	$scope.registerAction = function(){

  		$scope.error = '';

  		if($scope.remail !== '' && $scope.rpassword_hash !== '' && $scope.rusername !== ''){
  			var sendData = {
  				'name' : $scope.rusername,
  				'email': $scope.remail,
  				'password_hash': $scope.rpassword_hash
  			};
	  		api.createUser(sendData)
				.success(function(data){
					userManagement.setUserName(data.name);
					userManagement.setUserId(data.id);
					userManagement.setUserMail(data.email);
					userManagement.setLoggedIn(true);
          userManagement.store();
					$location.path('#/');
				})
				.error(function(){
					$scope.error = 'An error has occured.';
				});
  		}else{
  			$scope.error = 'Credentials not complete.';
  		}
  	};
  });