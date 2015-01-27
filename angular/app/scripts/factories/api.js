'use strict';
angular.module('angularApp')
    .factory('api', ['$http','userManagement', function($http,userManagement) {

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
        return $http.post(urlBase + '/unis', uni, {headers:  {'User': userManagement.getUserId()}});
    };

    api.updateUni = function (id, updatedUni) {
        return $http.patch(urlBase + '/unis/' + id, updatedUni , {headers:  {'User': userManagement.getUserId()}});
    };

    api.deleteUni = function (id) {
        return $http.delete(urlBase + '/unis/' + id,{headers:  {'User': userManagement.getUserId()}});
    };

    //lecture stuff

    api.getAllLectures = function (uniID) {
        return $http.get(urlBase + '/unis/'+uniID+'/lectures');
    };

    api.getLectures = function (uniID,lectureID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID);
    };

    api.createLecture = function (uniID, lecture) {
        return $http.post(urlBase + '/unis/'+ uniID + '/lectures', lecture,{headers:  {'User': userManagement.getUserId()}});
    };

    api.updateLecture = function (uniID,lectureID, updatedLecture) {
        return $http.patch(urlBase + '/unis/' + uniID + '/lectures/' + lectureID, updatedLecture,{headers:  {'User': userManagement.getUserId()}});
    };

    api.deleteLecture = function (uniID,lectureID) {
        return $http.delete(urlBase + '/unis/' + uniID + '/lectures/' + lectureID,{headers:  {'User': userManagement.getUserId()}});
    };

    //assignment stuff
    api.getAssignments = function (uniID,lectureID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments');
    };

    api.getAssignment = function (uniID,lectureID,assignmentID) {
        return $http.get(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments/' + assignmentID);
    };

    api.createAssignment = function(uniID,lectureID,data){
        return $http.post(urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments',data,{headers:  {'User': userManagement.getUserId()}});
    };

    api.updateAssignment = function (uniID,lectureID,assignmentID,tex){
        var url = urlBase + '/unis/' + uniID + '/lectures/' + lectureID + '/assignments/' + assignmentID;
        return $http.patch(url,tex,{headers:  {'User': userManagement.getUserId()}});
    };

    //solution stuff
    api.getSolutions = function (assi){
        return $http.get(urlBase + '/assignments/' + assi + '/solutions');
    };

    api.getSolution =  function (userID,solID){
        return $http.get(urlBase + '/users/' + userID + '/solutions/'+solID);
    };

    api.createSolution = function (userID,data){
        return $http.post(urlBase + '/users/' + userID + '/solutions',data,{headers:  {'User': userManagement.getUserId()}});
    };

    api.updateSolution =  function (userID,solID,tex){
        return $http.patch(urlBase + '/users/' + userID + '/solutions/'+solID,tex,{headers:  {'User': userManagement.getUserId()}});
    };

    //comment stuff

    api.getComments = function (reftype, refId){
        return $http.get(urlBase + '/comments/' + reftype+'/'+refId,{headers:  {'User': userManagement.getUserId()}});
    };

    api.createComment = function (data){
        return $http.post(urlBase + '/comments',data,{headers:  {'User': userManagement.getUserId()}});
    };

    api.updateComment = function (data,commentId){
        return $http.patch(urlBase + '/comments/' + commentId ,data,{headers:  {'User': userManagement.getUserId()}});
    };

    api.deleteComment = function (commentId){
        return $http.delete(urlBase + '/comments/'+commentId,{headers:  {'User': userManagement.getUserId()}});
    };

    //search stuff
    api.search = function (query){
        return $http.get(urlBase + '/search?query='+query);
    };

    //version related stuff

    api.getAssignmentVersions = function(uni,lect,assi){
        return $http.get(urlBase + '/unis/'+uni+'/lectures/'+lect+'/assignments/'+assi+'/versions');
    };

    api.getSolutionVersions = function(sol,user){
        return $http.get(urlBase + '/users/'+user+'/solutions/'+sol+'/versions');
    };

    api.getAssignmentVersion = function(uni,lect,assi,version){
        return $http.get(urlBase + '/unis/'+uni+'/lectures/'+lect+'/assignments/'+assi+'/versions/'+version);
    };

    api.getSolutionVersion = function(sol,user,version){
        return $http.get(urlBase + '/users/'+user+'/solutions/'+sol+'/versions/'+version);
    };

    api.revertAssignmentVersion = function (uni,lect,assi,version){
        return $http.patch(urlBase +'/unis/'+uni+'/lectures/'+lect+'/assignments/'+assi+'/versions/'+version);
    };

    api.revertSolutionVersion = function(sol,user,version){
        return $http.patch(urlBase + '/users/'+user+'/solutions/'+sol+'/versions/'+version);
    };

    return api;
}]).config(['$httpProvider', function($httpProvider) {
        $httpProvider.defaults.useXDomain = true;
        delete $httpProvider.defaults.headers.common['X-Requested-With'];
        $httpProvider.defaults.headers.patch = {'Content-Type': 'application/json;charset=utf-8'};
    }
]);