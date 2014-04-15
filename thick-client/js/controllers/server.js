function ServerCtrl($scope, $http) {
  $scope.$on('refresh', function() {
    $http({url: 'http://localhost:3000/server/' + $scope.server.ip})
    .success(function(data) {
      $scope.server.name = data.name;
      $scope.server.status = data.status;
      $scope.server.ping = data.ping;
      $scope.server.location = data.location;
    });
  });
};
