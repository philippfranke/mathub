'use strict';
angular.module('angularApp')
    .factory('api', ['$http', function($http) {

    var urlBase = 'http://192.168.59.103';
    var api = {};

    //user stuff
    api.loginUser = function (credentials) {
        return $http.post(urlBase + '/login',credentials);
    };

    api.getUser = function (userID) {
        return $http.get(urlBase + '/users/' + userID);
    };

    api.createUser = function (credentials) {
        return $http.post(urlBase + '/users',credentials);
    };


    //university stuff
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

    //lecture stuff

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

    //assignment stuff
    api.getAssignments = function (uniID,lectureID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments');
    };

    api.getAssignment = function (uniID,lectureID,assignmentID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments/' + assignmentID);
    };

    api.createAssignment = function(uniID,lectureID,data){
        return $http.post(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments',data);
    };

    api.updateAssignment = function (uniID,lectureID,assignmentID,tex){
        var url = urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments/' + assignmentID;
        return $http.patch(url,tex);
    };

    //solution stuff
    api.getSolutions = function (userID){
        return $http.get(urlBase + '/users/' + userID + '/solutions');
    };

    api.getSolution =  function (userID,solID){
        return $http.get(urlBase + '/users/' + userID + '/solutions/'+solID);
    };

    api.createSolution = function (userID,data){
        return $http.post(urlBase + '/users/' + userID + '/solutions',data);
    };

    api.updateSolution =  function (userID,solID,tex){
        return $http.patch(urlBase + '/users/' + userID + '/solutions/'+solID,tex);
    };

    //comment stuff

    api.getComments = function (reftype, refId){
        return $http.get(urlBase + '/comments/' + reftype+'/'+refId);
    };

    api.createComment = function (data){
        return $http.post(urlBase + '/comments',data);
    };

    return api;
}]).config(['$httpProvider', function($httpProvider) {
        $httpProvider.defaults.useXDomain = true;
        delete $httpProvider.defaults.headers.common['X-Requested-With'];
        $httpProvider.defaults.headers.patch = {'Content-Type': 'application/json;charset=utf-8'};
    }
]);