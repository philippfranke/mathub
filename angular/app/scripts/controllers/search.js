'use strict';
angular.module('angularApp')
  .controller('SearchCtrl', function ($scope, api, $location, sharedProperties) {
  	$scope.showLecture = false;
  	$scope.showAssignment = false;
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
        console.log(data);
				$scope.data = data;
			});
  	}

  	function fillUnis(){
  		getAllUnis();
  		$scope.showLecture = false;
  		$scope.showAssignment = false;
  	}

  	function fillLectures(unid){
  		$scope.uni = unid;
  		getLectures(unid);
  		$scope.showLecture = true;
  		$scope.showAssignment = false;
  	}

  	function fillAssignments(unid,lecid){
  		getAssignments(unid,lecid);
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

  	fillUnis();
  });
