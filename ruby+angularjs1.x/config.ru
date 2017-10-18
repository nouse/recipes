require 'sequel'

DB = Sequel.connect ENV['DATABASE_URL']
DB.extension :pg_json

require_relative 'app'
run App.freeze.app

require 'slim'
require 'refrigerator'
Refrigerator.freeze_core(:except=>[(Object.superclass || Object).name])
