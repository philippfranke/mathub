'use strict';
angular.module('angularApp')
    .factory('api', ['$http', function($http) {

    var urlBase = 'http://localhost:8080';
    var api = {};

    api.getAllUnis = function () {
        return $http.get(urlBase + '/unis');
    };

    api.getUnis = function (id) {
        return $http.get(urlBase + '/unis/' + id);
    };

    api.createUni = function (uni) {
        return $http.post(urlBase + '/unis', uni);
    };

    api.updateUni = function (id, updatedUni) {
        return $http.patch(urlBase + '/unis/' + id, updatedUni);
    };

    api.deleteUni = function (id) {
        return $http.delete(urlBase + '/unis/' + id);
    };

    api.getAllLectures = function (uniID) {
        return $http.get(urlBase + '/unis/'+uniID+'/lectures');
    };

    api.getLectures = function (uniID,lectureID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID);
    };

    api.createLecture = function (uniID, lecture) {
        return $http.post(urlBase + '/unis/'+ uniID + '/lectures', lecture);
    };

    api.updateLecture = function (uniID,lectureID, updatedLecture) {
        return $http.patch(urlBase + '/unis/' + uniID + '/lectures/' + lectureID, updatedLecture);
    };

    api.deleteLecture = function (uniID,lectureID) {
        return $http.delete(urlBase + '/unis/' + uniID + '/lectures/' + lectureID);
    };

    api.getAssignments = function (uniID,lectureID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments');
    };


    return api;
}]).config(['$httpProvider', function($httpProvider) {
        $httpProvider.defaults.useXDomain = true;
        delete $httpProvider.defaults.headers.common['X-Requested-With'];
    }
]);