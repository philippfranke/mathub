'use strict';
angular.module('angularApp')
  .controller('EditCtrl', function ($scope, api, sharedProperties, $location) {

  	$scope.editorOptions = {
  		lineNumbers: true,
  		matchBrackets: true,
  		mode: 'text/x-stex'
  	};

    $scope.assignment = '';

  	$scope.tex = 'Loading ...';

    $scope.saveTex = function(){
      var resultTex = {
        'tex': $scope.tex
      };
      resultTex.tex = resultTex.tex.replace('\\','\\\\');
      api.updateAssignment(sharedProperties.getUniEdit(),sharedProperties.getLectEdit(),sharedProperties.getAssiEdit(),resultTex).success(function(){
        $scope.saved =true;
      });

    };



  	function getAssignment(uniID, lectureID, AssiID){
  		api.getAssignment(uniID, lectureID, AssiID)
			.success(function(data){
				$scope.data = data;
				$scope.tex = data.tex.replace('\\\\','\\');
			});
  	}

  	function getUniversity(uniID){
  		api.getUnis(uniID)
  			.success(function(data){
  				$scope.university = data.name;
  			});
  	}

  	function getLectures(uniID,lectureID){
  		api.getLectures(uniID,lectureID)
  			.success(function(data){
  				$scope.lecture = data.name;
  			});
  	}

  	function displayAssignment (uniID, lectureID, AssiID){
  		getUniversity(uniID);
  		getLectures(uniID,lectureID);
  		getAssignment(uniID, lectureID, AssiID);
  	}

  	function getShared(){
  		var uni = sharedProperties.getUniEdit();
  		var lect = sharedProperties.getLectEdit();
  		var assi = sharedProperties.getAssiEdit();
  		if(assi === 0||lect ===0|| uni ===0 ){
  			$location.path('/search');
  		}else{
  			displayAssignment(uni,lect,assi);
  		}
  	}
  	getShared();
  });