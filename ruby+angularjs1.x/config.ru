require 'sequel'

DB = Sequel.connect ENV['DATABASE_URL']
DB.extension :pg_json

use Rack::Static, :urls => ["/css", "/js", "/fonts", "/diff.json"], :root => "public"

require_relative 'app'
run App.app
