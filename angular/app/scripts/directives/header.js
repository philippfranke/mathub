'use strict';
angular.module('angularApp')
  .directive('matheader', function () {
	  return {
	  	restrict: 'E',
	  	templateUrl:'partials/header.html'
	  };
  });
