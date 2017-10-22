var app = angular.module('guthub', ['guthub.directives', 'guthub.services', 'ngRoute']);

app.config(['$compileProvider', function ($compileProvider) {
  $compileProvider.debugInfoEnabled(false);
}]);

app.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
  $routeProvider.
    when('/', {
      controller: 'ListCtrl',
      resolve: {
        recipes: function(RecipeService) {
          return RecipeService.getRecipes();
        }
      },
      templateUrl: "/views/list.html"
    }).when('/edit/:recipeId', {
      controller: 'EditCtrl',
      resolve: {
        recipe: function(RecipeService, $route) {
          return RecipeService.getRecipe($route.current.params.recipeId);
        }
      },
      templateUrl: "/views/recipeForm.html"
    }).when('/view/:recipeId', {
      controller: 'ViewCtrl',
      resolve: {
        recipe: function(RecipeService, $route) {
          return RecipeService.getRecipe($route.current.params.recipeId);
        }
      },
      templateUrl: "/views/viewRecipe.html"
    }).when('/new', {
      controller: 'NewCtrl',
      templateUrl: "/views/recipeForm.html"
    }).otherwise({redirecTo: '/'});

  $locationProvider.html5Mode(true);
}]);

app.controller('NewCtrl', ['$scope', '$location', 'RecipeService', function($scope, $location, RecipeService) {
  $scope.recipe = {ingredients: []};

  $scope.save = function() {
    RecipeService.post($scope.recipe).then(function(recipe) {
      $location.path('/view/' + recipe.id);
    });
  };
}]);

app.controller('IngredientsCtrl', ['$scope', function($scope) {
  $scope.addIngredient = function() {
    $scope.recipe.ingredients.push({});
  };

  $scope.removeIngredient = function(index) {
    $scope.recipe.ingredients.splice(index, 1);
  };
}]);

app.controller('ListCtrl', ['$scope', 'recipes', function($scope, recipes) {
  $scope.recipes = recipes;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'recipe', function($scope, $location, recipe) {
  $scope.recipe = recipe;

  $scope.edit = function() {
    $location.path('/edit/' + recipe.id);
  };
}]);

app.controller('EditCtrl', ['$scope', '$location', 'recipe', function($scope, $location, recipe) {
  $scope.recipe = recipe;

  $scope.save = function() {
    $scope.recipe.put().then(function(recipe) {
      $location.path('/view/' + recipe.id);
    });
  };

  $scope.remove = function() {
    $scope.recipe.remove().then(function() {
      $location.path('/');
    });
  };
}]);

app.controller('navbarCtrl', ['$scope', '$rootScope', function($scope, $rootScope) {
  $scope.toggle = false;
  $rootScope.$on('$routeChangeStart', function() {
    $scope.toggle = false;
  });
}]);
