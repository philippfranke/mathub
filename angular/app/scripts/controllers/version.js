'use strict';
angular.module('angularApp')
  .controller('VersionCtrl', function ($scope, api, $location, sharedProperties, userManagement) {
  	//user:
  	if(userManagement.getLoggedIn() === false){
  	  userManagement.retrieve();
  	}
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();

  	$scope.tex = 'Loading ...';

  	$scope.mark = 1;


  	$scope.displayVersion = function(number){
  		if($scope.mode === 'ass'){
  			api.getAssignmentVersion($scope.uni,$scope.lect,$scope.assi,number)
  			.success(function(data){
  				if(number === 1){
  					$scope.tex = '';
  				}else{
  					$scope.tex = data.tex.replace('\\\\','\\');
  				}
  			});
  			$scope.mark = number;
  		}else{
        api.getSolutionVersion($scope.sol,$scope.userId,number)
        .success(function(data){
          if(number === 1){
            $scope.tex = '';
          }else{
            $scope.tex = data.tex.replace('\\\\','\\');
          }
        });
        $scope.mark = number;
  		}
  	};

    $scope.revertVersion = function (number){
      if($scope.mode === 'ass'){
        api.revertAssignmentVersion($scope.uni,$scope.lect,$scope.assi,number)
        .success(function(){
          displayAssignment();
        });
      }else{
        api.revertSolutionVersion($scope.sol,$scope.userId,number)
        .success(function(){
          displaySolution();
        });
      }
    };


  	function displayAssignment(){
  		api.getAssignmentVersions($scope.uni,$scope.lect,$scope.assi)
  		.success(function(data){
  			$scope.data = data;
  			$scope.displayVersion(1);
  		});
  	}

  	function displaySolution(){
      api.getSolutionVersions($scope.sol,$scope.userId)
      .success(function(data){
        $scope.data = data;
        $scope.displayVersion(1);
      });
  	}


	function getShared(){
		$scope.uni = sharedProperties.getUni();
		$scope.lect = sharedProperties.getLect();
		$scope.assi = sharedProperties.getAssi();
    	$scope.sol = sharedProperties.getSol();
		if($scope.assi === 0||$scope.lect ===0|| $scope.uni ===0 || $scope.sol ===0){
			$location.path('/search');
		}else{
		      if($scope.sol === -1){
		        displayAssignment();
		        $scope.mode = 'ass';
		      }else{
		        displaySolution();
		        $scope.mode = 'sol';
		      }
		}
	}
	getShared();
  });