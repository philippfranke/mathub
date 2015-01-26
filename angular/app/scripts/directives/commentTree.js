'use strict';
angular.module('angularApp')
.directive('tree', function(RecursionHelper) {
    return {
        restrict: 'E',
        scope: {
            family: '=',
            userid:'=',
            prevtext: '=',
            submitreply: '='
        },
        templateUrl:'partials/commentPartial.html',
        compile: function(element) {
            // Use the compile function from the RecursionHelper,
            // And return the linking function(s) which it returns
            return RecursionHelper.compile(element);
        }
    };
});