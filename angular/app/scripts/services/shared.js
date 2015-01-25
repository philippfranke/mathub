'use strict';
angular.module('angularApp')
    .service('sharedProperties', function () {
        var uniEdit = 0;
        var lectEdit = 0;
        var assiEdit = 0;
        var solEdit = 0;


        return {
            getUni: function () {
                return uniEdit;
            },
            setUni: function(value) {
                uniEdit = value;
            },
            getLect: function () {
                return lectEdit;
            },
            setLect: function(value) {
                lectEdit = value;
            },
            getAssi: function () {
                return assiEdit;
            },
            setAssi: function(value) {
                assiEdit = value;
            },
            getSol: function () {
                return solEdit;
            },
            setSol: function(value) {
                solEdit = value;
            }
        };
    });