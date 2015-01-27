'use strict';
/*jshint camelcase: false */
angular.module('angularApp')
  .controller('QueryCtrl', function ($scope, api, $location, sharedProperties, userManagement) {
  	//user:
  	if(userManagement.getLoggedIn() === false){
  	  userManagement.retrieve();
  	}
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();
  	$scope.query = sharedProperties.getQuery();

  	function getQueryResult(query){
		api.search(query)
		.success(function(data){
			$scope.data = data;
		});
  	}


  	$scope.edit = function(thing){
  	  sharedProperties.setUni(thing.university_id);
  	  sharedProperties.setLect(thing.lecture_id);
  	  if(thing.type === 'assignments'){
  	    sharedProperties.setAssi(thing.assignment_id);
  	    sharedProperties.setSol(-1);
  	  }else{
  	    sharedProperties.setAssi(thing.assignment_id);
  	    sharedProperties.setSol(thing.solution_id);
  	  }
  	  $location.path('/edit');
  	};


  	$scope.comments = function(thing){
		sharedProperties.setUni(thing.university_id);
		sharedProperties.setLect(thing.lecture_id);
		if(thing.type === 'assignments'){
		  sharedProperties.setAssi(thing.assignment_id);
		  sharedProperties.setSol(-1);
		}else{
		  sharedProperties.setAssi(thing.assignment_id);
		  sharedProperties.setSol(thing.solution_id);
		}
  	  $location.path('/comment');
  	};


  	$scope.$watch('query', function() {
  			getQueryResult($scope.query);
  	     });

  	getQueryResult($scope.query);
  });