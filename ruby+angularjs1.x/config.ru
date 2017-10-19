require 'sequel'

DB = Sequel.connect ENV['DATABASE_URL']
DB.extension :pg_json

require_relative 'app'
use Rack::CommonLogger
run App.freeze.app

require 'refrigerator'
Refrigerator.freeze_core
