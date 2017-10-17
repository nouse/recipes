require 'roda'
require 'sequel'
require 'rack/protection'

class Recipe < Sequel::Model
end

class App < Roda
  use Rack::Session::Cookie, :secret => "28165294156053dfe44aa83a54004a1edccc3a5baa3bbb23e3a8aaef54004a30"
  plugin :render, engine: "slim"
  plugin :json
  plugin :all_verbs
  plugin :halt
  plugin :csrf, :header => "X-XSRF-TOKEN"
  plugin :static_path_info
  plugin :cookies
  use Rack::Protection

  route do |r|
    unless request.cookies['XSRF-TOKEN']
      response.set_cookie "XSRF-TOKEN", csrf_token
    end

    r.root do
      view "index"
    end

    r.get "new" do
      view "index"
    end

    r.on "recipes" do
      r.is do
        r.get do
          Recipe.select(:id, :title).map(&:values)
        end

        r.post do
          if recipe = Recipe.create(JSON.load(request.body))
            recipe.values
          else
            recipe.errors
            response.status = 500
          end
        end
      end

      r.is ":id" do |id|
        recipe = Recipe[id: id]
        unless recipe
          halt(404)
        end
        r.get do
          recipe.values
        end

        r.put do
          recipe.update_fields(JSON.load(request.body), %w[title description instructions ingredients])
          recipe.values
        end

        r.delete do
          recipe.destroy.values
        end
      end
    end

    r.get "view/:id" do |id|
      view "index"
    end

    r.get "edit/:id" do |id|
      view "index"
    end

    r.on "views" do
      r.get "recipeForm.html" do
        render "recipeForm"
      end

      r.get "list.html" do
        render "list"
      end

      r.get "viewRecipe.html" do
        render "view"
      end
    end

  end
end
