'use strict';
angular.module('angularApp')
.directive('comeditor', function() {
    return {
        restrict: 'E',
        scope: {
            line: '=',
            submitreply: '=',
            userId:'=',
            edit: '=',
            reply:'=',
            prevtext: '=',
            parent: '='
        },
        templateUrl:'partials/commentEditor.html'
    };
});