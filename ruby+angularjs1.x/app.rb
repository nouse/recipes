require 'roda'
require 'sequel'
require 'rack/protection'

class Recipe < Sequel::Model
end

class App < Roda
  use Rack::Session::Cookie, :secret => "insecure"
  plugin :json
  plugin :all_verbs
  plugin :halt
  plugin :csrf, :header => "X-XSRF-TOKEN"
  plugin :cookies
  use Rack::Protection, origin_whitelist: ["http://localhost:9292", "http://127.0.0.1:9292"], without_session: true

  route do |r|
    unless request.cookies['XSRF-TOKEN']
      response.set_cookie "XSRF-TOKEN", csrf_token
    end

    r.on "recipes" do
      r.is do
        # GET /recipes
        r.get do
          Recipe.select(:id, :title).map(&:values)
        end

        # POST /recipes
        r.post do
          if recipe = Recipe.create(JSON.load(request.body))
            recipe.values
          else
            recipe.errors
            response.status = 500
          end
        end
      end

      r.is Integer do |id|
        recipe = Recipe[id: id]
        unless recipe
          r.halt(404)
        end

        # GET /recipes/:id
        r.get do
          recipe.values
        end

        # PUT /recipes/:id
        r.put do
          recipe.update_fields(JSON.load(request.body), %w[title description instructions ingredients])
          recipe.values
        end

        # DELETE /recipes/:id
        r.delete do
          recipe.destroy.values
        end
      end
    end

  end
end
