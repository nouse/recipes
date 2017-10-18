require 'roda'
require 'sequel'
require 'rack/protection'

class Recipe < Sequel::Model
end

class App < Roda
  use Rack::Session::Cookie, :secret => "28165294156053dfe44aa83a54004a1edccc3a5baa3bbb23e3a8aaef54004a30"
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

      r.is Integer do |id|
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

  end
end
