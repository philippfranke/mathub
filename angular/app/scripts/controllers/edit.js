'use strict';
angular.module('angularApp')
  .controller('EditCtrl', function ($scope, api, sharedProperties, $location, userManagement) {

    //user
    if(userManagement.getLoggedIn() === false){
      userManagement.retrieve();
    }
    $scope.loggedIn = userManagement.getLoggedIn();
    $scope.userName = userManagement.getUserName();
    $scope.userId = userManagement.getUserId();

    
  	$scope.editorOptions = {
  		lineNumbers: true,
  		matchBrackets: true,
  		mode: 'text/x-stex'
  	};

    $scope.assignment = '';

    $scope.mode = '';

  	$scope.tex = 'Loading ...';

    $scope.saveTex = function(){
      var resultTex = {
        'tex': $scope.tex
      };
      resultTex.tex = resultTex.tex.replace('\\','\\\\');
      if($scope.mode === 'ass'){
        api.updateAssignment(sharedProperties.getUni(),sharedProperties.getLect(),sharedProperties.getAssi(),resultTex).success(function(){
          $scope.saved =true;
        });
      }else{
        api.updateSolution($scope.userId,sharedProperties.getSol(),resultTex).success(function(){
          $scope.saved =true;
        });
      }

    };

    $scope.showComments = function(){
      $location.path('/comment');
    };

    $scope.showVersions = function(){
      $location.path('/version');
    };

    $scope.view = function(){
      console.log('view');
    };


  	function showAssignment(uniID, lectureID, AssiID){
  		api.getAssignment(uniID, lectureID, AssiID)
			.success(function(data){
				$scope.data = data;
				$scope.tex = data.tex.replace('\\\\','\\');
			});
  	}

    function showSolution(solID){
      api.getSolution($scope.userId,solID)
      .success(function(data){
        $scope.data = data;
        $scope.tex = data.tex.replace('\\\\','\\');
      });
    }

    function getAssignment(uniID, lectureID, AssiID){
      api.getAssignment(uniID, lectureID, AssiID)
      .success(function(data){
        if(data.name !== ''){
          $scope.assignment = data.name;
        }else{
          $scope.assignment = data.id;
        }
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
  		showAssignment(uniID, lectureID, AssiID);
  	}

    function displaySolution(uniID, lectID, AssiID, SolID){
      getUniversity(uniID);
      getLectures(uniID,lectID);
      getAssignment(uniID, lectID, AssiID);
      showSolution(SolID);
    }

  	function getShared(){
  		var uni = sharedProperties.getUni();
  		var lect = sharedProperties.getLect();
  		var assi = sharedProperties.getAssi();
      var sol = sharedProperties.getSol();
  		if(assi === 0||lect ===0|| uni ===0 || sol ===0){
  			$location.path('/search');
  		}else{
        if(sol === -1){
          displayAssignment(uni,lect,assi);
          $scope.mode = 'ass';
        }else{
          displaySolution(uni,lect,assi,sol);
          $scope.mode = 'sol';
        }
  		}
  	}
  	getShared();
  });