'use strict';
angular.module('angularApp')
  .controller('CommentCtrl', function ($scope, api, sharedProperties, $location, userManagement) {
  	//user
  	if(userManagement.getLoggedIn() === false){
  	  userManagement.retrieve();
  	}
  	$scope.loggedIn = userManagement.getLoggedIn();
  	$scope.userName = userManagement.getUserName();
  	$scope.userId = userManagement.getUserId();

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

    function getComments(){
    	/*api.getComments(type,id)
    	.success(function(data){
    	  $scope.comments = data;
    	});*/
    $scope.comments = [
        {
            'comment': {
                'id': 2,
                'ref_type': 'assignment',
                'ref_id': 1,
                'ref_version': 0,
                'ref_line': 0,
                'parent_id': 0,
                'user_id': 1,
                'username': 'bricktop',
                'timestamp': '0001-01-01T00:00:00Z',
                'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
            },
            'children': [
                {
                    'comment': {
                        'id': 5,
                        'ref_type': 'assignment',
                        'ref_id': 1,
                        'ref_version': 0,
                        'ref_line': 0,
                        'parent_id': 2,
                        'user_id': 1,
                        'username': 'bricktop',
                        'timestamp': '0001-01-01T00:00:00Z',
                        'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
                    },
                    'children': [
                {
                    'comment': {
                        'id': 5,
                        'ref_type': 'assignment',
                        'ref_id': 1,
                        'ref_version': 0,
                        'ref_line': 0,
                        'parent_id': 2,
                        'user_id': 1,
                        'username': 'bricktop',
                        'timestamp': '0001-01-01T00:00:00Z',
                        'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
                    }
                },
                {
                    'comment': {
                        'id': 6,
                        'ref_type': 'assignment',
                        'ref_id': 1,
                        'ref_version': 0,
                        'ref_line': 0,
                        'parent_id': 2,
                        'user_id': 1,
                        'username': 'bricktop',
                        'timestamp': '0001-01-01T00:00:00Z',
                        'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
                    }
                },
                {
                    'comment': {
                        'id': 7,
                        'ref_type': 'assignment',
                        'ref_id': 1,
                        'ref_version': 0,
                        'ref_line': 0,
                        'parent_id': 2,
                        'user_id': 1,
                        'username': 'bricktop',
                        'timestamp': '0001-01-01T00:00:00Z',
                        'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
                    }
                }
            ]
                },
                {
                    'comment': {
                        'id': 6,
                        'ref_type': 'assignment',
                        'ref_id': 1,
                        'ref_version': 0,
                        'ref_line': 0,
                        'parent_id': 2,
                        'user_id': 1,
                        'username': 'bricktop',
                        'timestamp': '0001-01-01T00:00:00Z',
                        'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
                    }
                },
                {
                    'comment': {
                        'id': 7,
                        'ref_type': 'assignment',
                        'ref_id': 1,
                        'ref_version': 0,
                        'ref_line': 0,
                        'parent_id': 2,
                        'user_id': 1,
                        'username': 'bricktop',
                        'timestamp': '0001-01-01T00:00:00Z',
                        'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
                    }
                }
            ]
        },
        {
          'comment': {
              'id': 2,
              'ref_type': 'assignment',
              'ref_id': 1,
              'ref_version': 0,
              'ref_line': 2,
              'parent_id': 0,
              'user_id': 1,
              'username': 'bricktop',
              'timestamp': '0001-01-01T00:00:00Z',
              'text': 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et'
          }
        }
    ];
    }

    //executing
    if(sharedProperties.getSol() === -1){
  		$scope.type = 'assignment';
  		getComments($scope.type,sharedProperties.getAssi());
  		showAssignment(sharedProperties.getUni(),sharedProperties.getLect(),sharedProperties.getAssi());
  	}else{
  		if(sharedProperties.getSol() ===0){
  			$location.path('/search');
  		}
  		$scope.type = 'solution';
  		getComments($scope.type,sharedProperties.getSol());
  		showSolution(sharedProperties.getSol());
  	}


  });