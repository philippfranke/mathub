'use strict';
angular.module('angularApp')
    .service('userManagement', function () {
        var isLoggedIn = false;
        var userName = '';
        var userId = 0;
        var userMail = '';

        return {
            getLoggedIn: function () {
                return isLoggedIn;
            },
            setLoggedIn: function(value) {
                isLoggedIn = value;
            },
            getUserName: function () {
                return userName;
            },
            setUserName: function(value) {
                userName = value;
            },
            getUserId: function () {
                return userId;
            },
            setUserId: function(value) {
                userId = value;
            },
            getUserMail: function () {
                return userMail;
            },
            setUserMail: function(value) {
                userMail = value;
            }
        };
    });