'use strict';
angular.module('angularApp')
  .controller('SearchCtrl', function ($scope, api) {
  	$scope.status = 0;
  	$scope.data = 0;

  	function getAllUnis(){
  		api.getAllUnis()
	  		.success(function(data){
	  			$scope.data = data;
	  		});
  	}

  	getAllUnis();
  });
