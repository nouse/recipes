var services = angular.module('guthub.services', ['restangular']);

services.config(['$compileProvider', function ($compileProvider) {
  $compileProvider.debugInfoEnabled(false);
}]);

services.factory('RecipeService', ['Restangular', function(Restangular){
  var _recipeService = Restangular.all('recipes');

  return {
    getRecipes: function() {
      return _recipeService.getList();
    },
    getRecipe: function(id) {
      return _recipeService.get(id);
    },
    post: function(recipe) {
      return _recipeService.post(recipe);
    }
  };
}]);
