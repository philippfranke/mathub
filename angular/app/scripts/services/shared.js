'use strict';
angular.module('angularApp')
    .service('sharedProperties', function () {
        var uniEdit = 0;
        var lectEdit = 0;
        var assiEdit = 0;

        return {
            getUniEdit: function () {
                return uniEdit;
            },
            setUniEdit: function(value) {
                uniEdit = value;
            },
            getLectEdit: function () {
                return lectEdit;
            },
            setLectEdit: function(value) {
                lectEdit = value;
            },
            getAssiEdit: function () {
                return assiEdit;
            },
            setAssiEdit: function(value) {
                assiEdit = value;
            }
        };
    });