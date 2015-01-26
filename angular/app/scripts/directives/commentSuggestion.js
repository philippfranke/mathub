'use strict';
module.directive('comTree', function() {
    return {
        restrict: 'E',
        scope: {family: '='},
        template: 
            '<p>{{ family.name }}</p>'+
            '<ul>' + 
                '<li ng-repeat="child in family.children">' + 
                    '<tree family="child"></tree>' +
                '</li>' +
            '</ul>'
    };
});