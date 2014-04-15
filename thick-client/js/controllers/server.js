function ServerCtrl($scope, $http) {
  $scope.$on('refresh', function() {
    $http({url: 'http://localhost:3000/server/' + $scope.server.ip})
    .success(function(data) {
      $scope.server.location = data.location;
      $scope.server.ping = data.ping;
    });
  });
};
