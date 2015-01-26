'use strict';
angular.module('angularApp')
.directive('comeditor', function() {
    return {
        restrict: 'E',
        scope: {
            line: '=',
            parent: '='
        },
        templateUrl:'partials/commentEditor.html'
    };
});