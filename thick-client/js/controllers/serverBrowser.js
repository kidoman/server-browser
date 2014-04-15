var _ = require('lodash');
var gui = require('nw.gui');

function ServerBrowserCtrl($scope, $http, $timeout) {
  $scope.servers = [];
  $scope.predicate = 'ping';
  $scope.reverse = false;

  $scope.sortOn = function(predicate) {
    if ($scope.predicate === predicate) {
      $scope.reverse = !$scope.reverse;      
    } else {
      $scope.reverse = false;
      $scope.predicate = predicate;
    }
  };

  $scope.refresh = function() {
    $http({url: 'http://localhost:3000/servers'})
    .success(function(data) {
      $scope.servers = _(data).map(function(ip) {
        return {ip: ip, name: '', status: '', location: '', ping: 999};
      }).value();
      $timeout(function() {
        $scope.$broadcast('refresh');
      });
    });
  };

  $scope.quit = function() {
    var win = gui.Window.get();
    win.close();
  };
};
