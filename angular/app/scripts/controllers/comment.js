'use strict';
angular.module('angularApp')
  .controller('CommentCtrl', function ($scope,$filter, api, sharedProperties, $location, userManagement) {
  	//user
  	if(userManagement.getLoggedIn() === false){
  	  userManagement.retrieve();
  	}
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();
    $scope.userid = $scope.userId;

  	//control vars
  	$scope.type = '';
  	function showAssignment(uniID, lectureID, AssiID){
  		api.getAssignment(uniID, lectureID, AssiID)
			.success(function(data){
				$scope.data = data;
				var blobtex = data.tex.replace('\\\\','\\');
				$scope.tex = blobtex.split('\n');

				for(var i = 0; i<$scope.tex.length;i++){
					if($scope.tex[i] === ''){
						$scope.tex.splice(i,1);
					}
				}
			});
  	}

    function showSolution(solID){
      api.getSolution($scope.userId,solID)
	      .success(function(data){
	        $scope.data = data;
	        var blobtex = data.tex.replace('\\\\','\\');
	        $scope.tex = blobtex.split('\n');
	        for(var i = 0; i<$scope.tex.length;i++){
	        	if($scope.tex[i] === ''){
	        		$scope.tex.splice(i,1);
	        	}
	        }
	      });
    }

    $scope.submitreply = function(line,parent,content,edit){
      if(!edit){
        var e = {
          'ref_type' : $scope.type,
          'ref_id' : $scope.id,
          'ref_version' : 1, //change when versions are implemented
          'ref_line' : line,
          'parent_id' : parent,
          'user_id' : $scope.userId,
          'timestamp' : new Date().toJSON(),
          'text' : content
        };
        if(content !== ''){
          api.createComment(e)
            .success(function(){
              draw();
            });
        }
      }else{
        if(line !== -1){

          var i = {
            'ref_type' : $scope.type,
            'ref_id' : $scope.id,
            'text' : content
          };
          if(content !== ''){
            api.updateComment(i,parent)
              .success(function(){
                draw();
              });
          }
        }else{
          api.deleteComment(parent)
            .success(function(){
              draw();
            });
        }
      }
    };

    function getComments(type,id){
    	api.getComments(type,id)
    	.success(function(data){
    	  $scope.comments = data;
    	});
    }

    //executing
    function draw(){
      if(sharedProperties.getSol() === -1){
        $scope.type = 'assignment';
        $scope.id = sharedProperties.getAssi();
        getComments($scope.type,sharedProperties.getAssi());
        showAssignment(sharedProperties.getUni(),sharedProperties.getLect(),sharedProperties.getAssi());
      }else{
        if(sharedProperties.getSol() ===0){
          $location.path('/search');
        }
        $scope.type = 'solution';
        $scope.id = sharedProperties.getSol();
        getComments($scope.type,sharedProperties.getSol());
        showSolution(sharedProperties.getSol());
      }
    }

    draw();

  });