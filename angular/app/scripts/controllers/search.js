'use strict';
angular.module('angularApp')
  .controller('SearchCtrl', function ($scope, api, $location, sharedProperties, userManagement) {

    //user:
    if(userManagement.getLoggedIn() === false){
      userManagement.retrieve();
    }
    $scope.loggedIn = userManagement.getLoggedIn();
    $scope.userName = userManagement.getUserName();
    $scope.userId = userManagement.getUserId();

    //rest
  	$scope.showLecture = false;
  	$scope.showAssignment = false;
    $scope.addUniversity = false;
    $scope.addLecture = false;
    $scope.addAssignment = false;
  	$scope.data = 0;
  	$scope.uni = 0;

  	function getAllUnis(){
  		api.getAllUnis()
	  		.success(function(data){
	  			$scope.data = data;
	  		});
  	}
  	function getLectures(uniID){
  		api.getAllLectures(uniID)
			.success(function(data){
				$scope.data = data;
			});
  	}
  	function getAssignments(uniID, lectureID){
  		api.getAssignments(uniID, lectureID)
			.success(function(data){
				$scope.data = data;
			});
  	}

  	function fillUnis(){
  		getAllUnis();
      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
  		$scope.showLecture = false;
  		$scope.showAssignment = false;
  	}

  	function fillLectures(unid){
  		$scope.uni = unid;
  		getLectures(unid);
      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
  		$scope.showLecture = true;
  		$scope.showAssignment = false;
  	}

  	function fillAssignments(unid,lecid){
      $scope.uni = unid;
      $scope.lect = lecid;
  		getAssignments(unid,lecid);
      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
  		$scope.showLecture = true;
  		$scope.showAssignment = true;
  	}

  	$scope.Handler = function(id,unid,lectureid){
      if(!angular.isUndefined(lectureid)){
        sharedProperties.setLectEdit(lectureid);
        sharedProperties.setAssiEdit(id);
        $location.path('/edit');
      }else{
        if(angular.isUndefined(id)){
          fillUnis();
        }else{
          if(angular.isUndefined(unid)){
              fillLectures(id);
          }else{
            sharedProperties.setUniEdit(unid);
            fillAssignments(unid,id);
          }
        }

      }
  	};

    $scope.addItem = function(complete){
      var sendData = {};
      //uni add
      if(!$scope.showAssignment && !$scope.showLecture){
        if(complete){
          sendData = {
            'name':$scope.universityData
          };
          api.createUni(sendData)
          .success(function(){
            $scope.addUniversity = false;
            fillUnis();
          });
        }else{
          $scope.addUniversity = true;
        }
      }
      //lecture add
      if(!$scope.showAssignment && $scope.showLecture ){
        if(complete){
          sendData = {
            'name':$scope.lectureData
          };
          api.createLecture($scope.uni,sendData)
          .success(function(){
            $scope.addLecture = false;
            fillLectures($scope.uni);
          });
        }else{
          $scope.addLecture = true;
        }
      }
      //assignment add
      if($scope.showAssignment && $scope.showLecture ){
        if(complete){
          sendData = {
            'name':$scope.assignmentData,
            'tex':''
          };
          api.createAssignment($scope.uni,$scope.lect,sendData)
          .success(function(){
            $scope.addAssignment = false;
            fillAssignments($scope.uni,$scope.lect);
          });
        }else{
          $scope.addAssignment = true;
        }
      }
    };

  	fillUnis();
  });
