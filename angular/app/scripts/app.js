'use strict';

/**
 * @ngdoc overview
 * @name angularApp
 * @description
 * # angularApp
 *
 * Main module of the application.
 */
angular
  .module('angularApp', [
    'ngAnimate',
    'ngCookies',
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
    'ui.codemirror'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl'
      })
      .when('/search', {
        templateUrl: 'views/search.html',
        controller: 'SearchCtrl'
      })
      .when('/edit', {
        templateUrl: 'views/edit.html',
        controller: 'EditCtrl'
      })
      .when('/login', {
        templateUrl: 'views/login.html',
        controller: 'LoginCtrl'
      })
      .when('/profile', {
        templateUrl: 'views/profile.html',
        controller: 'ProfileCtrl'
      })
      .when('/comment', {
        templateUrl: 'views/comment.html',
        controller: 'CommentCtrl'
      })
      .when('/query', {
        templateUrl: 'views/query.html',
        controller: 'QueryCtrl'
      })
      .when('/version', {
        templateUrl: 'views/version.html',
        controller: 'VersionCtrl'
      })
      .when('/blog', {
        templateUrl: 'views/blog.html',
        controller: 'BlogCtrl'
      })
      .when('/contact', {
        templateUrl: 'views/contact.html',
        controller: 'ContactCtrl'
      })
      .when('/settings', {
        templateUrl: 'views/settings.html',
        controller: 'SettingsCtrl'
      })
      .when('/inbox', {
        templateUrl: 'views/inbox.html',
        controller: 'InboxCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });


