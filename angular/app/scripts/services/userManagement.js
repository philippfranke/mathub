'use strict';
angular.module('angularApp')
    .service('userManagement',['$cookieStore', function ($cookieStore) {
        var isLoggedIn = false;
        var userName = '';
        var userId = 0;
        var userMail = '';

        return {
            store: function(){
                $cookieStore.put('isLoggedIn',isLoggedIn);
                $cookieStore.put('userName',userName);
                $cookieStore.put('userId',userId);
                $cookieStore.put('userMail',userMail);
            },
            retrieve: function(){
                isLoggedIn = $cookieStore.get('isLoggedIn');
                userName = $cookieStore.get('userName');
                userId = $cookieStore.get('userId');
                userMail = $cookieStore.get('userMail');
            },
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
    }]);