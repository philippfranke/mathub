'use strict';
/*jshint camelcase: false */
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
    //what is shown right now in results
  	$scope.showLecture = false;
  	$scope.showAssignment = false;
    $scope.showSolution = false;
    //what is open to be added
    $scope.addUniversity = false;
    $scope.addLecture = false;
    $scope.addAssignment = false;
    //data transmitted to frontend
  	$scope.data = 0;
    //ids of all stages
  	$scope.uni = 0;
    $scope.lect = 0;
    $scope.ass = 0; 

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
        $scope.lastAssignments = data;
			});
  	}
    function getSolutions(assignmentID){
      api.getSolutions($scope.userId)
      .success(function(data){
        var correctData = [];
        for(var i = 0; i < data.length; i++){
          if(data[i].assignment_id === assignmentID){
            correctData.push(data[i]);
          }
        }
        $scope.data = correctData;
      });
    }

  	function fillUnis(){
  		getAllUnis();

      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
      $scope.addSolution = false;

  		$scope.showLecture = false;
  		$scope.showAssignment = false;
      $scope.showSolution = false;
  	}

  	function fillLectures(unid){
  		$scope.uni = unid;
  		getLectures(unid);

      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
      $scope.addSolution = false;

  		$scope.showLecture = true;
  		$scope.showAssignment = false;
      $scope.showSolution = false;
  	}

  	function fillAssignments(unid,lecid){
      $scope.uni = unid;
      $scope.lect = lecid;
  		getAssignments(unid,lecid);

      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
      $scope.addSolution = false;

  		$scope.showLecture = true;
  		$scope.showAssignment = true;
      $scope.showSolution = false;
  	}

    $scope.fillSolutions = function(assid){
      $scope.ass = assid;
      getSolutions(assid);

      $scope.addUniversity = false;
      $scope.addLecture = false;
      $scope.addAssignment = false;
      $scope.addSolution = false;

      $scope.showLecture = true;
      $scope.showAssignment = true;
      $scope.showSolution = true;
    };

  	$scope.Handler = function(id,unid){
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
  	};

    $scope.edit = function(contextID){
      sharedProperties.setUniEdit($scope.uni);
      sharedProperties.setLectEdit($scope.lect);
      if($scope.showAssignment && !$scope.showSolution){
        sharedProperties.setAssiEdit(contextID);
        sharedProperties.setSolEdit(-1);
      }else{
        sharedProperties.setAssiEdit($scope.ass);
        sharedProperties.setSolEdit(contextID);
      }
      $location.path('/edit');
    };

    $scope.view = function(contextID){
      console.log(contextID);
    };


    $scope.addItem = function(complete){
      var sendData = {};
      //uni add
      if(!$scope.showAssignment && !$scope.showLecture && !$scope.showSolution){
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
      if(!$scope.showAssignment && $scope.showLecture && !$scope.showSolution){
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
      if($scope.showAssignment && $scope.showLecture && !$scope.showSolution){
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

      if($scope.showAssignment && $scope.showLecture && $scope.showSolution){
        if(complete){
          var tex = '';
          for(var i = 0; i < $scope.lastAssignments.length; i++){
            if($scope.lastAssignments[i].id === $scope.ass){
              tex = $scope.lastAssignments[i].tex;
            }
          }
          sendData = {
            'assignment_id':$scope.ass,
            'tex':tex
          };
          api.createSolution($scope.userId,sendData)
          .success(function(){
            $scope.addSolution = false;
            $scope.fillSolutions($scope.ass);
          });
        }else{
          $scope.addSolution = true;
        }
      }
    };

  	fillUnis();
  });
