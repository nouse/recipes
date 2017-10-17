var directives = angular.module('guthub.directives', ['ngTouch']);

directives.directive('butterbar', ['$rootScope', function($rootScope) {
  return {
    restrict: 'A',
    link: function(scope, element, attrs) {
      element.addClass('hide');

      $rootScope.$on('$routeChangeStart', function() {
        element.removeClass('hide');
      });

      $rootScope.$on('$routeChangeSuccess', function() {
        element.addClass('hide');
      });

    }
  };
}]);
