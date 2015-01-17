'use strict';
/*globals LaTeX2HTML5 */
angular.module('angularApp')
  .directive('latex', function () {
  	function link(scope, element) {
  		scope.$watch('data', function() {
  		        var TEX = new LaTeX2HTML5.TeX({
  		            tagName: 'section',
  		            className: 'latex-container',
  		            latex: scope.data
  		        });
  		        TEX.render();
  		        element.html(TEX.$el);
  		     });

  	}
  	return {
	restrict: 'E',
	    scope: {
	      data: '='
	    },
  		link: link
  	};
  });
