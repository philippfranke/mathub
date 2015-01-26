'use strict';
module.directive('comTree', function(RecursionHelper) {
    return {
        restrict: 'E',
        scope: {family: '='},
        template: 
            '<p>{{ family.name }}</p>'+
            '<ul>' + 
                '<li ng-repeat="child in family.children">' + 
                    '<tree family="child"></tree>' +
                '</li>' +
            '</ul>',
        compile: function(element) {
            // Use the compile function from the RecursionHelper,
            // And return the linking function(s) which it returns
            return RecursionHelper.compile(element);
        }
    };
});