'use strict';
angular.module('angularApp')
  .controller('CommentCtrl', function ($scope, api, sharedProperties, $location, userManagement) {
  	//user
  	if(userManagement.getLoggedIn() === false){
  	  userManagement.retrieve();
  	}
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();

  	//control vars
  	$scope.type = '';
  	function showAssignment(uniID, lectureID, AssiID){
  		api.getAssignment(uniID, lectureID, AssiID)
			.success(function(data){
				$scope.data = data;
				var blobtex = data.tex.replace('\\\\','\\');
				$scope.tex = blobtex.split('\n');
				for(var i = 0; i<$scope.tex.length;i++){
					if($scope.tex[i] === ''){
						$scope.tex.splice(i,1);
					}
				}
			});
  	}

    function showSolution(solID){
      api.getSolution($scope.userId,solID)
	      .success(function(data){
	        $scope.data = data;
	        var blobtex = data.tex.replace('\\\\','\\');
	        $scope.tex = blobtex.split('\n');
	        for(var i = 0; i<$scope.tex.length;i++){
	        	if($scope.tex[i] === '\\n'){
	        		$scope.tex.splice(i,1);
	        	}
	        }
	      });
    }

    function getComments(type,id){
    	api.getComments(type,id)
    	.success(function(data){
    	  $scope.comments = data;
    	});
    }

    //executing
    if(sharedProperties.getSol() === -1){
  		$scope.type = 'assignment';
  		getComments($scope.type,sharedProperties.getAssi());
  		showAssignment(sharedProperties.getUni(),sharedProperties.getLect(),sharedProperties.getAssi());
  	}else{
  		if(sharedProperties.getSol() ===0){
  			$location.path('/search');
  		}
  		$scope.type = 'solution';
  		getComments($scope.type,sharedProperties.getSol());
  		showSolution(sharedProperties.getSol);
  	}


  });