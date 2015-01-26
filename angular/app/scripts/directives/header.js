'use strict';
angular.module('angularApp')
  .directive('matheader',['$location','sharedProperties', function ($location,sharedProperties) {
  	function link(scope) {
  		scope.redirectquery = function(){
  			sharedProperties.setQuery(scope.querydata);
  			if($location.url() === '/query'){
  				scope.query = scope.querydata;
  			}else{
  				$location.path('/query');
  			}
  		};

  	}
	  return {
	  	restrict: 'E',
	  	link: link,
	  	templateUrl:'partials/header.html'
	  };
  }]);
